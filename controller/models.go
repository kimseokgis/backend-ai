package controller

import (
	"context"
	"fmt"
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
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
	fmt.Printf("%+v\n", key)
	reply, err := helper.QueriesDataRegexp(db, context.TODO(), key)
	if err != nil {
		resp.Message = fmt.Sprintf("error get Replies: %s\n", err.Error())
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotFound, resp)
		return
	}
	chat.IdChats = reply.ID.Hex()
	chat.Message = reply.Question
	chat.Responses = reply.Answer
	helper.WriteJSON(w, http.StatusOK, chat)
	return
}
