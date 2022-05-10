package server

import (
	"fmt"
	"github.com/drhernandez/go-starter-project/src/middlewares"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func StartServer(port int) error {
	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: uuid.NewString,
	}))
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/health")
		},
		Format: "${time_rfc3339} level:INFO [request_id:${id}] | method=${method}, uri=${uri}, status=${status} error=${error} latency=${latency_human}\n",
	}))
	router.Use(middlewares.AddRequestIDToContext())

	mapUrls(router, nil)

	return router.Start(fmt.Sprintf(":%d", port))
}
