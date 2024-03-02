package main

import (
	"database/sql"
	"os"

	"github.com/Light2Dark/memecatcher/internal"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type application struct {
	db           *sql.DB
	openAiClient *openai.Client
	user         *internal.User
}

func Render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Writer.WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()
	openaiClient, err := internal.CreateOpenAIClient()
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	db, err := internal.NewDBConn()
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		db:           db,
		openAiClient: openaiClient,
		user:         &internal.User{},
	}

	e.Static("/templates/css", "templates/css")
	e.Static("/templates/assets", "templates/assets")

	e.GET("/v1/healthcheck", app.healthcheckHandler)
	e.GET("/", app.homeHandler)
	e.POST("/fetchMeme", app.fetchMemeHandler)
	e.GET("/getAllMemes", app.getMemesHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
