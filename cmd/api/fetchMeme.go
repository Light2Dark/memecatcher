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

	var subreddits []string = []string{"memes", "dankmemes", "wholesomememes", "Animemes", "artmemes", "holesome", "2meirl4meirl", "wholesomememes", "shitposting"}

	var wg sync.WaitGroup
	var ch = make(chan MemeResponse, len(subreddits))
	var indexToMemes = make(map[int]Meme, len(subreddits)*numMemesRequested)

	for _, subreddit := range subreddits {
		wg.Add(1)
		memeAPIUrl := fmt.Sprintf("https://meme-api.com/gimme/%s/%d", subreddit, numMemesRequested)

		go func() {
			defer wg.Done()
			res, err := http.Get(memeAPIUrl)
			if err != nil {
				fmt.Println("error fetching data from", memeAPIUrl, ":", err)
				return
			}
			defer res.Body.Close()

			var memeResponse MemeResponse
			err = json.NewDecoder(res.Body).Decode(&memeResponse)
			if err != nil {
				fmt.Println("error decoding meme response:", err)
				return
			}
			ch <- memeResponse
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0
	for memeRes := range ch {
		for _, meme := range memeRes.Memes {
			if !includeNsfw && meme.NSFW {
				continue
			}

			memeToAdd := Meme{
				Title:                meme.Title,
				HighestImgQualityUrl: meme.Preview[len(meme.Preview)-1],
			}

			indexToMemes[count] = memeToAdd
			count++
		}
	}

	var indexToTitle = make(map[int]string, len(indexToMemes))
	for i, meme := range indexToMemes {
		indexToTitle[i] = meme.Title
	}

	prompt := `Given a hashmap where each sentence number corresponds to a meme, with the hashmap as follows: ` + fmt.Sprintf("%v", indexToTitle) + `. The user has entered the search term: ` + search + `. Your task is to return the sentence number of the best match to the search term. In the world of memes, words often carry slang meanings, so be creative in finding a match. Think very broadly in all subjects to find a match. When there is no match, return a random meme (it's number) and come up with a funny explanation about the person searching this. Provide your answer in the format: "Return: sentence number, Explanation: explanation". Make your explanation brief but change it to something witty or sarcastic.

	Here are some examples of the hashmap and user searches and their expected return values:
	Hashmap: {0: "I am a cat", 1: "I am a dog", 2: "I am a human"}
	User search: "Doggo"
	Your response: Return: 1, Explanation: The title "doggo" indicates a dog

	Hashmap: {0: "Rolling in dough", 1: "I am a dog", 2: "I am a human"}
	User search: "Balling"
	Your response: Return: 0, Explanation: The title "balling" is a slang term for "rolling in dough"
	`

	resp, err := internal.ChatCompletion(app.openAiClient, prompt, 100)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`\d+`)
	memeNumString := re.FindString(resp)
	memeNumber, err := strconv.Atoi(memeNumString)
	if err != nil {
		fmt.Println("error converting meme number to int:", err)
		memeNumber = 0
	}

	highestImgQualityUrl := indexToMemes[memeNumber].HighestImgQualityUrl
	memeExplanation := resp[strings.Index(resp, "Explanation: ")+len("Explanation: "):]
	memeExplanation = strings.Replace(memeExplanation, "hashmap", "memes", -1)

	fmt.Println("response:", resp, "memes:", indexToTitle)

	err = internal.InsertMeme(app.db, app.user.ID, highestImgQualityUrl)
	if err != nil {
		fmt.Println("error inserting meme into db:", err)
		return err
	}

	return Render(c, http.StatusOK, templates.FetchMeme(highestImgQualityUrl, memeExplanation))
}
