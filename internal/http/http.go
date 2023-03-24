package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	_ "ta13-svc/docs"
	"ta13-svc/internal/factory"
	"ta13-svc/internal/usecase/auth"
)

func Init(e *echo.Echo, f *factory.Factory) {

	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome")
		return c.String(http.StatusOK, message)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	auth.NewHandler(f).Route(e.Group("/auth"))
}
