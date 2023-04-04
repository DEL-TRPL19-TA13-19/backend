package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"ta13-svc/docs"
	_ "ta13-svc/docs"
	"ta13-svc/internal/factory"
	"ta13-svc/internal/usecase/alternative"
	"ta13-svc/internal/usecase/auth"
	"ta13-svc/internal/usecase/collection"
	"ta13-svc/internal/usecase/tps"
)

func Init(e *echo.Echo, f *factory.Factory) {

	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to API AHP Method TPS Locations")
		return c.String(http.StatusOK, message)
	})

	docs.SwaggerInfo.Host = os.Getenv("BASE_URL")

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	auth.NewHandler(f).Route(e.Group("/auth"))
	tps.NewHandler(f).Route(e.Group("/tps"))
	collection.NewHandler(f).Route(e.Group("/collection"))
	alternative.NewHandler(f).Route(e.Group("/alternative"))
}
