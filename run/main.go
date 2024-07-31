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
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
		
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
	// Middleware
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":8080")
}