package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	db "ta13-svc/database"
	"ta13-svc/database/migration"
	"ta13-svc/internal/factory"
	"ta13-svc/internal/http"
	"ta13-svc/internal/middleware"
	"ta13-svc/pkg/elasticsearch"
	"ta13-svc/pkg/env"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

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
	db.Init()
	migration.Init()
	elasticsearch.Init()
	var PORT = os.Getenv("PORT")

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)
	e.Logger.Fatal(e.Start(":" + PORT))
}
