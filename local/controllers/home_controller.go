package controllers

import (
	"net/http"

	"github.com/kimseokgis/backend-ai/helper"
)

func Home(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to Makmur AI!",
	}
	helper.WriteJSON(w, http.StatusOK, response)
}
