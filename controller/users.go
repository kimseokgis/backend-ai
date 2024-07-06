package controller

import (
	"context"
	"encoding/json"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
)

func RegisterUsers(w http.ResponseWriter, r *http.Request) {
	resp := new(model.Credential)
	userdata := new(model.User)
	resp.Status = false
	conn := helper.SetConnection()
	err := json.NewDecoder(r.Body).Decode(userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
		helper.WriteJSON(w, http.StatusNotAcceptable, resp)
	} else {
		resp.Status = true
		hash, err := helper.HashPass(userdata.PasswordHash)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
			helper.WriteJSON(w, http.StatusBadRequest, resp)
		}
		helper.InsertUserdata(conn, userdata.Username, userdata.Email, userdata.Password, hash)
		resp.Message = "Berhasil Registrasi Data"
		defer conn.Client().Disconnect(context.TODO())
	}
	helper.WriteJSON(w, http.StatusOK, resp)

}
