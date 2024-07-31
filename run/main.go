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
func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register your handler
	mux.HandleFunc("/", url.Web)

	// Wrap the mux with the logging middleware
	loggedMux := loggingMiddleware(mux)

	// Start the server
	app.Listen(":8080")
}