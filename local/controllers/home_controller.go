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

