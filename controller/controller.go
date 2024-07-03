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