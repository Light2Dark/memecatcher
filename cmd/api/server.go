package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Light2Dark/memecatcher/internal"
	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type application struct {
	db           *sql.DB
	openAiClient *openai.Client
	user         *internal.User
}
type config struct {
	openaiKey string
	dsn       string
	port      string
}

func Render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Writer.WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()
	config, err := loadConfig()

	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	openaiClient, err := internal.CreateOpenAIClient(config.openaiKey)
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	db, err := internal.NewDBConn(config.dsn)
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

	e.Logger.Fatal(e.Start(":" + config.port))
}

func loadConfig() (*config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found, using default values")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if os.Getenv("OPENAI_API_KEY") == "" || os.Getenv("NEON_DSN") == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return &config{
		openaiKey: os.Getenv("OPENAI_API_KEY"),
		dsn:       os.Getenv("NEON_DSN"),
		port:      port,
	}, nil
}
