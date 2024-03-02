package main

import (
	"net/http"

	"github.com/Light2Dark/memecatcher/internal"
	templates "github.com/Light2Dark/memecatcher/templates/home"
	"github.com/labstack/echo/v4"
)

func (app *application) homeHandler(c echo.Context) error {
	cookie, err := internal.ReadCookie(c)
	if err == nil && cookie != nil {
		userID := cookie.Value
		_, err := internal.DoesUserExist(app.db, userID)
		if err != nil {
			return err
		}
		app.user.ID = userID
	}

	return Render(c, http.StatusOK, templates.Index())
}