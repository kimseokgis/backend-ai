package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kimseokgis/backend-ai/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"os"

	"github.com/aiteung/atdb"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection() *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv("MONGOSTRING"),
		DBName:   "AI",
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func InsertUserdata(MongoConn *mongo.Database, username, email, password, passwordhash string) (InsertedID interface{}) {
	req := new(model.User)
	req.Username = username
	req.Email = email
	req.Password = password
	req.PasswordHash = passwordhash
	return InsertOneDoc(MongoConn, "users", req)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func CreateNewUserRole(mongoconn *mongo.Database, collection string, userdata model.User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPass(userdata.PasswordHash)
	if err != nil {
		return err
	}
	userdata.PasswordHash = hashedPassword

	// Insert the admin data into the database
	return atdb.InsertOneDoc(mongoconn, collection, userdata)
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

func IsPasswordValid(mongoconn *mongo.Database, userdata model.User) bool {
	filter := bson.M{
		"$or": []bson.M{
			{"username": userdata.Username},
			{"email": userdata.Email},
		},
	}

	var res model.User
	err := mongoconn.Collection("users").FindOne(context.TODO(), filter).Decode(&res)

	if err == nil {
		return CheckPasswordHash(userdata.PasswordHash, res.PasswordHash)
	}
	return false
}
