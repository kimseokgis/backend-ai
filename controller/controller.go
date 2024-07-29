package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
)

func HomeMakmur(w http.ResponseWriter, r *http.Request) {
	Response := fmt.Sprintf("Makmur AI chooy %s", "8080")
	response, err := json.Marshal(Response)
	if err != nil {
		http.Error(w, "Internal server error: JSON marshaling failed", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	return
}

func NotFound(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Message = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, resp)
}

func Comment(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	comment := new(model.Comment)
	resp.Status = false
	conn := helper.SetConnection()
	err := json.NewDecoder(req.Body).Decode(comment)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
		helper.WriteJSON(respw, http.StatusNotAcceptable, resp)
	}
	insID := helper.InsertComment(conn, *comment)
	resp.Status = true
	resp.Message = fmt.Sprintf("Data berhasil diinsert %s", insID)
	helper.WriteJSON(respw, http.StatusOK, resp)
	return
}
