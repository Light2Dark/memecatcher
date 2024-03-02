package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Light2Dark/memecatcher/internal"
	templates "github.com/Light2Dark/memecatcher/templates/fetchMeme"
	"github.com/labstack/echo/v4"
)

type Meme struct {
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
	numMemesRequested := c.FormValue("numMemes")

	memeAPIUrl := fmt.Sprintf("https://meme-api.com/gimme/%s", numMemesRequested)

	res, err := http.Get(memeAPIUrl)
	if err != nil {
		return err
	}

	var meme Meme
	err = json.NewDecoder(res.Body).Decode(&meme)
	if err != nil {
		return err
	}

	indexToTitle := make(map[int]string, len(meme.Memes))

	for i, m := range meme.Memes {
		indexToTitle[i] = strings.ToLower(m.Title)
	}

	prompt := `Given a hashmap where each sentence number corresponds to a meme, with the hashmap as follows: ` + fmt.Sprintf("%v", indexToTitle) + `. The user has entered the search term: ` + search + `. Your task is to return the sentence number of the best match to the search term. In the world of memes, words often carry slang meanings, so be creative in finding a match. Often there is no word matches, just make up a loose connection between the search term and the meme. If there is no match, return 0. Provide your answer in the format: "Return: sentence number, Explanation: explanation". Make your explanation brief but change it to something witty or sarcastic.

	Here are some examples of the hashmap and user searches and their expected return values:
	Hashmap: {0: "I am a cat", 1: "I am a dog", 2: "I am a human"}
	User search: "Doggo"
	Your response: Return: 1, Explanation: The title "doggo" indicates a dog

	Hashmap: {0: "Rolling in dough", 1: "I am a dog", 2: "I am a human"}
	User search: "Balling"
	Your response: Return: 0, Explanation: The title "balling" is a slang term for "rolling in dough"
	`

	resp, err := internal.ChatCompletion(app.openAiClient, prompt, 50)
	if err != nil {
		return err
	}

	re := regexp.MustCompile("[0-9]")
	memeNumString := re.FindString(resp)
	memeNumber, err := strconv.Atoi(memeNumString)
	if err != nil {
		fmt.Println("error converting meme number to int:", err)
		memeNumber = 0
	}

	highestImgQualityUrl := meme.Memes[memeNumber].Preview[len(meme.Memes[memeNumber].Preview)-1]
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
