package controller

import (
	"context"
	"fmt"
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
	"strings"
)

func ChatPredictUsingRegexp(w http.ResponseWriter, r *http.Request) {
	resp := new(model.Credential)
	chat := new(model.Chats)
	token := r.Header.Get("login")
	if token == "" {
		resp.Message = "token is empty"
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotAcceptable, resp)
		return
	}
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Missing key parameter", http.StatusBadRequest)
		return
	}
	key := keys[0]
	fmt.Println(key)
	decoder, err := helper.DecodeGetUser(config.PublicKey, token)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = false
		helper.WriteJSON(w, http.StatusBadRequest, resp)
		return
	}
	db := helper.SetConnection()
	fmt.Println(decoder)

	_, err = helper.FindUserByUsername(db, decoder)
	if err != nil {
		resp.Message = fmt.Sprintf("Data tidak ditemukan : %s\n"+
			"Username: %s\n", err.Error(), decoder)
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotFound, resp)
		return
	}
	if strings.Contains(key, "_") {
		cihuy := strings.Replace(key, "_", " ", 10)
		key = cihuy
		keysSlices := strings.Split(cihuy, " ")
		if len(keysSlices) >= 5 {
			key = keysSlices[len(keysSlices)-4] + " " + keysSlices[len(keysSlices)-3] + " " + keysSlices[len(keysSlices)-2] + " " + keysSlices[len(keysSlices)-1]
		} else if len(keysSlices) >= 3 {
			key = keysSlices[len(keysSlices)-2] + " " + keysSlices[len(keysSlices)-1]
		} else if len(keysSlices) >= 2 {
			key = keysSlices[len(keysSlices)-1]
		}
	}
	fmt.Printf("%+v\n", key)
	reply, err := helper.QueriesDataRegexp(db, context.TODO(), key)
	if err != nil {
		resp.Message = "Aduh aduh aduhhhaiii, aku ga ngerti nihh coba nanya yang lain dongg biar aku ngertiin kamu..."
		resp.Status = false
		chat.Responses = resp.Message
		helper.WriteJSON(w, http.StatusNotFound, resp)
		return
	}
	chat.IdChats = reply.ID.Hex()
	chat.Message = reply.Question
	chat.Responses = reply.Answer
	defer db.Client().Disconnect(context.Background())
	helper.WriteJSON(w, http.StatusOK, chat)
	return
}
