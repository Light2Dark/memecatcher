package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrCookieNotFound = echo.NewHTTPError(http.StatusNotFound, "Cookie not found")

func WriteCookie(c echo.Context, userID string) error {
	cookie := new(http.Cookie)
	cookie.Name = "memeUser"
	cookie.Value = userID
	c.SetCookie(cookie)
	return nil
}

func ReadCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie("memeUser")
	if err != nil {
		return nil, ErrCookieNotFound
	}
	return cookie, nil
}
