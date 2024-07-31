package helpers

import (
	"context"
	"os"

	"github.com/aiteung/atdb"
	"github.com/kimseokgis/backend-ai/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SetConnection() *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv("MONGOSTRING"),
		DBName:   "AI",
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func InsertUser(db *mongo.Database, user model.User) interface{} {
	collection := db.Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil
	}
	return result.InsertedID
}