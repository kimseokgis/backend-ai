package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kimseokgis/backend-ai/local/config"
	"github.com/kimseokgis/backend-ai/local/controllers"
)

// main is the entry point of the application.
func main() {
	// Initialize Fiber app with custom configuration
	app := fiber.New(config.FiberConfig)

	// Middleware
	// Logger middleware logs HTTP requests
	app.Use(logger.New())

	// Routes
	// Root route to check if the server is running
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Makmur AI!")
	})

	// Route for user registration
	app.Post("/register", controllers.RegisterUser)

	// Route for user login
	app.Post("/login", controllers.LoginUser)

	// Determine port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}


}
