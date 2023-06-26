package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
)

func AddRequestIDToContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = c.Response().Header().Get(echo.HeaderXRequestID)
			}
			newCtx := context.WithValue(c.Request().Context(), echo.HeaderXRequestID, requestID)
			c.SetRequest(c.Request().WithContext(newCtx))
			return next(c)
		}
	}
}
