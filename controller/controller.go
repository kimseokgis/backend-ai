package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

// func Login Register
func Register(Mongoenv, dbname string, r *http.Request) string {
	resp := new(model.Credential)
	userdata := new(model.User)
	resp.Status = false
	conn := helper.MongoCreateConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		hash, err := HashPass(userdata.PasswordHash)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
		}
		helper.InsertUserdata(conn, userdata.Username, userdata.Email, userdata.Password, hash)
		resp.Message = "Berhasil Registrasi Data"
	}
	response := ReturnStringStruct(resp)
	return response
}

// func Login User
func Login(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response model.Credential
	Response.Status = false
	mconn := helper.SetConnection(dbname)
	var datauser model.User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang Anda Berhasil Login"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Username atau Password Anda Salah"
		}
	}

	return GCFReturnStruct(Response)
}

// return struct
func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func ReturnStringStruct(Data any) string {
	jsonee, _ := json.Marshal(Data)
	return string(jsonee)
}

// Password
func HashPass(passwordhash string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordhash), 14)
	return string(bytes), err
}

func CheckPasswordHash(passwordhash, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordhash))
	return err == nil
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata model.User) bool {
	filter := bson.M{
		"$or": []bson.M{
			{"username": userdata.Username},
			{"email": userdata.Email},
		},
	}

	var res model.User
	err := mongoconn.Collection(collection).FindOne(context.TODO(), filter).Decode(&res)

	if err == nil {
		return CheckPasswordHash(userdata.PasswordHash, res.PasswordHash)
	}
	return false
}
