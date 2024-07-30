package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kimseokgis/backend-ai/local/config"
	"github.com/kimseokgis/backend-ai/local/controllers"
)

func main() {
	app := fiber.New(config.FiberConfig)

	// Middleware
	app.Use(logger.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Makmur AI!")
	})
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)

	// Start server
	port := config.IPport
	if port == "" {
		port = ":8080"
	}

	log.Fatal(app.Listen(port))
}
