package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) healthcheckHandler(c echo.Context) error {
	jsonStruct := struct {
		Status      string `json:"status"`
		Environment string `json:"environment"`
	}{
		Status:      "OK",
		Environment: "development",
	}
	return c.JSON(http.StatusOK, jsonStruct)
}
