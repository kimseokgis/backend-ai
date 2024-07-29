package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kimseokgis/backend-ai/model"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
}
