package controllers

import (
	"net/http"

	"github.com/kimseokgis/backend-ai/helper"
)

// createHomeResponse creates the response for the Home handler.
// Returns a map with a welcome message.
func createHomeResponse() map[string]string {
	return map[string]string{
		"message": "Welcome to Makmur AI!",
	}
}

// writeJSONResponse writes a JSON response using the helper package.
// Takes an HTTP response writer, status code, and the data to be written.
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	helper.WriteJSON(w, statusCode, data)
}
