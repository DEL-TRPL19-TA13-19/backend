package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"ta13-svc/controller/user"
	_ "ta13-svc/docs"
	"ta13-svc/storage"
)

// @title API TA
// @version 1.0
// @description Dokumentasi API D4 TRPL 19 TA 13.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /
func main() {
	port := ":9000"

	storage.NewDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", user.GetHello)
	e.GET("/users", user.GetUsers)

	e.Logger.Fatal(e.Start(port))
}
