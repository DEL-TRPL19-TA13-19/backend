package middleware

import (
	"github.com/labstack/echo/v4"
)

func Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//TODO
		return next(c)
	}
}
