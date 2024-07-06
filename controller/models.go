package controller

import (
	"github.com/kimseokgis/backend-ai/helper"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"net/http"
)

func PredictModels(w http.ResponseWriter, r *http.Request) {
	model, err := tf.LoadSavedModel("indobert_model", []string{"serve"}, nil)
	if err != nil {
		http.Error(w, "Failed to load model", http.StatusInternalServerError)
		return
	}
	defer model.Session.Close()

	helper.WriteJSON(w, http.StatusOK, model)
}
