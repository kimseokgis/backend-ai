package main

import (
	"github.com/kimseokgis/backend-ai/url"
	"log"
	"net/http"
	"time"
)

// Define the logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		
	config.ConnectDatabase()

	// Middleware
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":8080")
}