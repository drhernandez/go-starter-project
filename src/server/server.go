package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func StartServer(port int) error {
	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.RequestID())
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/health")
		},
		Format: "request_id:${id} level=INFO method=${method}, uri=${uri}, status=${status} error=${error} latency=${latency_human}\n",
	}))

	mapUrls(router, nil)

	return router.Start(fmt.Sprintf(":%d", port))
}
