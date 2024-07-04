package controller

// import (
// 	"context"
// 	"crypto/rand"
// 	"encoding/hex"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"golang.org/x/crypto/argon2"
// )

//user
// func GetUserFromID(db *mongo.Database, col string, _id primitive.ObjectID) (*User, error) {
// 	cols := db.Collection(col)
// 	filter := bson.M{"_id": _id}

// 	user := new(Ticket)

// 	err := cols.FindOne(context.Background(), filter).Decode(user)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, fmt.Errorf("no data found for ID %s", _id.Hex())
// 		}
// 		return nil, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
// 	}

// 	return user, nil
// }

//login
// func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
// 	if insertedDoc.Email == "" || insertedDoc.Password == "" {
// 		return user, fmt.Errorf("Dimohon untuk melengkapi data")
// 	}
// 	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
// 		return user, fmt.Errorf("Email tidak valid")
// 	}
// 	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db)
// 	if err != nil {
// 		return
// 	}
// 	salt, err := hex.DecodeString(existsDoc.Salt)
// 	if err != nil {
// 		return user, fmt.Errorf("kesalahan server : salt")
// 	}
// 	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
// 	if hex.EncodeToString(hash) != existsDoc.Password {
// 		return user, fmt.Errorf("password salah")
// 	}
// 	return existsDoc, nil
// }

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/kimseokgis/backen-ai/model"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// func Login Register
func Register(Mongoenv, dbname string, r *http.Request) string {
	resp := new(model.Credential)
	userdata := new(model.User)
	resp.Status = false
	conn := MongoCreateConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		hash, err := HashPass(userdata.PasswordHash)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(conn, userdata.Username, userdata.Email, userdata.Password, hash)
		resp.Message = "Berhasil Input data"
	}
	response := ReturnStringStruct(resp)
	return response
}

// func Login User
func Login(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response model.Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
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
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Username atau Password Salah"
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
