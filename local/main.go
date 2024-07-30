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


}
