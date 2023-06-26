package server

import (
	"github.com/drhernandez/go-starter-project/src/logger"
	"github.com/labstack/echo/v4"
)

func mapUrls(router *echo.Echo, app *app) {
	router.Any("/", func(c echo.Context) error {
		logger.Info(c.Request().Context(), logger.WithTags("tag1", "value1", "tag2", 2), "este es un mensaje")
		return nil
	})
}
