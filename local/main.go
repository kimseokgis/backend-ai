package main

import (
	"log"
	"os"

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

	// Determine port
	port := os.Getenv("PORT") // Try to get PORT from environment variable
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server is running on http://127.0.0.1:%s\n", port)
	log.Fatal(app.Listen(":" + port))
}
