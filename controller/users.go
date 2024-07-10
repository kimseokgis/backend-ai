package controller

import (
	"context"
	"encoding/json"
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"github.com/whatsauth/watoken"
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

func LoginUsers(w http.ResponseWriter, r *http.Request) {
	resp := new(model.Credential)
	var userdata model.User
	resp.Status = false
	conn := helper.SetConnection()
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		if helper.IsPasswordValid(conn, userdata) {
			resp.Status = true
			tokenstring, err := watoken.Encode(userdata.Username, config.PrivateKey)
			if err != nil {
				resp.Message = "Gagal Encode Token : " + err.Error()
				helper.WriteJSON(w, http.StatusBadRequest, resp)
			} else {
				resp.Message = "Selamat Datang Anda Berhasil Login"
				resp.Token = tokenstring
				helper.WriteJSON(w, http.StatusOK, resp)
			}
		} else {
			resp.Message = "Username atau Password Anda Salah"
			helper.WriteJSON(w, http.StatusBadRequest, resp)
		}
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    resp := new(model.Credential)
    resp.Status = false
    conn := helper.SetConnection()
    defer conn.Client().Disconnect(context.TODO())

    username := r.URL.Query().Get("username")
    if username == "" {
        resp.Message = "Username tidak boleh kosong"
        helper.WriteJSON(w, http.StatusBadRequest, resp)
        return
    }

    err := helper.FindUserByUsername(conn, username)
    if err != nil {
        resp.Message = "Pengguna tidak ditemukan: " + err.Error()
        helper.WriteJSON(w, http.StatusNotFound, resp)
        return
    }

    resp.Status = true
    resp.Message = "Pengguna ditemukan"
    helper.WriteJSON(w, http.StatusOK, resp)
}