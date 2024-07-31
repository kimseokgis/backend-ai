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

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register your handler
	mux.HandleFunc("/", url.Web)

	// Wrap the mux with the logging middleware
	loggedMux := loggingMiddleware(mux)

	// Start the server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
