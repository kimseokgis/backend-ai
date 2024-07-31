package main

import (
	"github.com/kimseokgis/backend-ai/url"
	"log"
	"net/http"
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