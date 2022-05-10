package server

import "github.com/labstack/echo/v4"

func mapUrls(router *echo.Echo, app *app) {
	router.Any("/", func(c echo.Context) error {
		return nil
	})
}
