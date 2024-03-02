package main

import (
	"fmt"
	"net/http"

	"github.com/Light2Dark/memecatcher/internal"
	templates "github.com/Light2Dark/memecatcher/templates/getAllMemes"
	"github.com/labstack/echo/v4"
)

func (app *application) getMemesHandler(c echo.Context) error {
	if app.user.ID == "" {
		return c.JSON(http.StatusUnauthorized, "user not authenticated")
	}

	memes, err := internal.GetAllMemes(app.db, app.user.ID, 6)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return Render(c, http.StatusOK, templates.Meme(memes))
}
