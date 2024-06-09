package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/Light2Dark/memecatcher/internal"
	templates "github.com/Light2Dark/memecatcher/templates/fetchMeme"
	dropdown "github.com/Light2Dark/memecatcher/templates/home/dropdown"
	"github.com/labstack/echo/v4"
)

type MemeResponse struct {
	Count int `json:"count"`
	Memes []struct {
		PostLink  string   `json:"postLink"`
		Subreddit string   `json:"subreddit"`
		Title     string   `json:"title"`
		URL       string   `json:"url"`
		NSFW      bool     `json:"nsfw,omitempty"`
		Spoiler   bool     `json:"spoiler,omitempty"`
		Author    string   `json:"author,omitempty"`
		UPS       int      `json:"ups,omitempty"`
		Preview   []string `json:"preview,omitempty"`
	}
}

type Meme struct {
	Title                string
	HighestImgQualityUrl string
	Subreddit            string
}

func (app *application) fetchMemeHandler(c echo.Context) error {
	if app.user.ID == "" {
		newID := internal.GenerateUserID()
		err := internal.CreateUser(app.db, newID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = internal.WriteCookie(c, newID)
		if err != nil {
			return err
		}

		app.user.ID = newID
	}

	var subredditsChosen []string = make([]string, 0, len(dropdown.Subreddits))
	for _, subreddit := range dropdown.Subreddits {
		exists := c.FormValue(subreddit)
		if exists != "" {
			subredditsChosen = append(subredditsChosen, subreddit)
		}
	}

	var extraSubreddits = c.FormValue("extraSubreddits")
	if extraSubreddits != "" {
		subs := strings.Split(extraSubreddits, ",")
		for _, sub := range subs {
			subredditsChosen = append(subredditsChosen, strings.TrimSpace(sub))
		}
	}

	search := c.FormValue("search")
	numMemesRequested, err := strconv.Atoi(c.FormValue("numMemes"))
	numMemesRequested = numMemesRequested / 3 // Division factor
	nsfw := c.FormValue("nsfw")

	if err != nil {
		return err
	}

	includeNsfw := false
	if nsfw == "on" {
		includeNsfw = true
	}

	var wg sync.WaitGroup
	var ch = make(chan MemeResponse, len(subredditsChosen))

	for _, subreddit := range subredditsChosen {
		wg.Add(1)
		memeAPIUrl := fmt.Sprintf("https://meme-api.com/gimme/%s/%d", subreddit, numMemesRequested)

		go func() {
			defer wg.Done()
			res, err := http.Get(memeAPIUrl)
			if err != nil {
				c.Logger().Error("error fetching data from", memeAPIUrl, ":", err)
				return
			}
			defer res.Body.Close()

			var memeResponse MemeResponse
			err = json.NewDecoder(res.Body).Decode(&memeResponse)
			if err != nil {
				c.Logger().Error("error decoding meme response:", err)
				return
			}
			ch <- memeResponse
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var indexToMeme = make(map[int]Meme, len(subredditsChosen)*numMemesRequested)

	count := 0
	for memeRes := range ch {
		for _, meme := range memeRes.Memes {
			if !includeNsfw && meme.NSFW {
				continue
			}

			memeToAdd := Meme{
				Title:                meme.Title,
				Subreddit:            meme.Subreddit,
				HighestImgQualityUrl: meme.Preview[len(meme.Preview)-1],
			}
			indexToMeme[count] = memeToAdd

			count++
		}
	}

	var memePrompt strings.Builder
	for i, meme := range indexToMeme {
		memePrompt.WriteString(fmt.Sprintf("%d:[title: %s, subreddit: %s] ", i, meme.Title, meme.Subreddit))
	}
	c.Logger().Debug("memePrompt:", memePrompt.String())

	prompt := `Given a hashmap where each sentence number corresponds to a meme array. The first part of the array is the meme title while the second part of the array is the subreddit it's pulled from. Here is the hashmap: ` + memePrompt.String() + `. The user has entered the search term: ` + search + `. Your task is to return the sentence number of the best match to the search term. In the world of memes, words often carry slang meanings, so be creative in finding a match. Think very broadly in all subjects to find a match. When there is no match, return a random meme (it's number) and come up with a funny explanation about the person searching this. Provide your answer in the format: "Return: sentence number, Explanation: explanation". Make your explanation brief but change it to something witty or sarcastic, refer the user as "you" or "your".

	Here are some examples of the hashmap and user searches and their expected return values:
	Hashmap: 1:[title: 2meirl4meirl, subreddit: dog] 5:[title: noPrivilegesNoProblems, subreddit: ProgrammerHumor], 0:[title: The Freudian surface is much too slippery., subreddit: dankmemes]
	User search: "Doggo"
	Your response: Return: 1, Explanation: Perhaps this 2meirl meme from the dog land will make you feel better.

	Hashmap: 1: [title: Rolling in dough, subreddit: dollar], 2: [title: I am a dog, subreddit: dog]
	User search: "Balling"
	Your response: Return: 0, Explanation: You be "balling" by rolling in dough in these dollar streets.
	`

	resp, err := internal.ChatCompletion(app.openAiClient, prompt, 100)
	if err != nil {
		return err
	}

	// fmt.Println("memes", indexToMeme)
	// fmt.Println("response", resp)

	re := regexp.MustCompile(`\d+`)
	memeNumString := re.FindString(resp)
	memeNumber, err := strconv.Atoi(memeNumString)
	if err != nil {
		fmt.Println("error converting meme number to int:", err)
		memeNumber = 0
	}

	highestImgQualityUrl := indexToMeme[memeNumber].HighestImgQualityUrl
	memeExplanation := resp[strings.Index(resp, "Explanation: ")+len("Explanation: "):]
	memeExplanation = strings.Replace(memeExplanation, "hashmap", "memes", -1)

	// fmt.Println("memeChosen", memeNumber, "memeExplanation", memeExplanation)
	c.Logger().Debug("memeExplanation:", memeExplanation)

	err = internal.InsertMeme(app.db, app.user.ID, highestImgQualityUrl)
	if err != nil {
		fmt.Println("error inserting meme into db:", err)
		return err
	}

	return Render(c, http.StatusOK, templates.FetchMeme(highestImgQualityUrl, memeExplanation))
}
