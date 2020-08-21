package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler: Home | Handler send to home page
func homeHandler(c echo.Context) error {
	Debugf("Publishing the home page")
	return c.Render(http.StatusUnauthorized, "home.html", "")
}

// Handler: Second | Handler send to second page
func secondHandler(c echo.Context) error {
	Debugf("Publishing the second page")
	return c.Render(http.StatusUnauthorized, "second.html", "")
}

// Handler: Api | Handler that serves the API data
func apiHandler(c echo.Context) error {
	Debugf("Publishing the api")
	return c.JSON(http.StatusOK, map[string]string{
		`key1`: `value1`,
		`key2`: `value2`,
		`key3`: `value3`,
		`keyN`: `valueN`,
	})
}