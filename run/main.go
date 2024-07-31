package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/routes"
)

func main() {
	app := fiber.New()

	// Connect to the database
	config.ConnectDatabase()

	// Middleware
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":8080")
}