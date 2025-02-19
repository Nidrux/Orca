package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nidrux/orca/internal/config"
	"github.com/nidrux/orca/internal/database"
	"github.com/nidrux/orca/internal/routes"
)

func main() {

	config.InitConfig()
	config.InitWebServer()

	app := config.GetWebServer()
	cfg := config.GetConfig()

	database.InitDatabase()

	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	// First init the api routes. every other route will fallback to the react dist
	routes.ApiRoutes()
	app.Static("/", "./web/dist")
	app.Static("*", "../../web/dist/index.html")

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Port)))
}
