package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
)

func ChatPredictUsingRegexp(w http.ResponseWriter, r *http.Request) {
	resp := new(model.Credential)
	req := new(model.Requests)
	chat := new(model.Chats)
	token := w.Header().Get("login")
	if token == "" {
		resp.Message = "token is empty"
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotAcceptable, resp)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
		helper.WriteJSON(w, http.StatusNotAcceptable, resp)
		return
	} else {
		decoder, err := helper.Decoder(config.PublicKey, token)
		if err != nil {
			resp.Message = err.Error()
			resp.Status = false
			helper.WriteJSON(w, http.StatusBadRequest, resp)
			return
		}
		db := helper.SetConnection()
		_, err = helper.FindUserByUsername(db, decoder.User)
		if err != nil {
			resp.Message = fmt.Sprintf("Data tidak ditemukan : %s\n", err.Error())
			resp.Status = false
			helper.WriteJSON(w, http.StatusNotFound, resp)
			return
		}

		reply, err := helper.QueriesDataRegexp(db, context.TODO(), req.Messages)
		if err != nil {
			resp.Message = fmt.Sprintf("error get Replies: %s\n", err.Error())
			resp.Status = false
			helper.WriteJSON(w, http.StatusNotFound, resp)
			return
		}
		chat.IdChats = reply.ID.Hex()
		chat.Message = reply.Question
		chat.Responses = reply.Answer
	}
	helper.WriteJSON(w, http.StatusOK, chat)
	return
}
