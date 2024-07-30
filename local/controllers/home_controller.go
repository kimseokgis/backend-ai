package controllers

import (
	"net/http"

	"github.com/kimseokgis/backend-ai/helper"
)

// createHomeResponse creates the response for the Home handler.


// writeJSONResponse writes a JSON response using the helper package.
// Takes an HTTP response writer, status code, and the data to be written.
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	helper.WriteJSON(w, statusCode, data)
}

// Home handles the home route.
// Creates a response with a welcome message and writes it as JSON.
func Home(w http.ResponseWriter, r *http.Request) {
	response := createHomeResponse()
	writeJSONResponse(w, http.StatusOK, response)
}