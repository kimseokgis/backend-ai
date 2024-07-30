package controllers

import (
	"net/http"

	"github.com/kimseokgis/backend-ai/helper"
)

// createHomeResponse creates the response for the Home handler.


// writeJSONResponse writes a JSON response using the helper package.


// Home handles the home route.
// Creates a response with a welcome message and writes it as JSON.
func Home(w http.ResponseWriter, r *http.Request) {
	response := createHomeResponse()
	writeJSONResponse(w, http.StatusOK, response)
}