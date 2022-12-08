package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/raystainerz/gourl/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.Resolve)
	app.Post("/api/v1", routes.Shorten)
	app.Post("/api/v1/urlinfo", routes.Urlinfo)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment file.")
	}

	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	app.Listen(":3000") 
}
