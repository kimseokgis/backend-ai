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

// getDBInfo constructs the database info from environment variables.
// Returns the database information.
func getDBInfo() atdb.DBInfo {
	return atdb.DBInfo{
		DBString: os.Getenv("MONGOSTRING"),
		DBName:   "AI",
	}
}

// connectToMongoDB establishes a connection to MongoDB.
// Takes the database info and returns the database connection.
func connectToMongoDB(dbInfo atdb.DBInfo) *mongo.Database {
	return atdb.MongoConnect(dbInfo)
}

// insertUserToDB inserts a user into the specified MongoDB collection.
// Returns the inserted ID or an error if insertion fails.
func insertUserToDB(collection *mongo.Collection, user model.User) (interface{}, error) {
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// findUserInDB finds a user in the specified MongoDB collection by username.
// Returns the user model or an error if retrieval fails.
func findUserInDB(collection *mongo.Collection, username string) (*model.User, error) {
	filter := bson.M{"username": username}
	var storedUser model.User
	err := collection.FindOne(context.TODO(), filter).Decode(&storedUser)
	if err != nil {
		return nil, err
	}
	return &storedUser, nil
}

// hashPassword generates a hashed password using bcrypt.
// Returns the hashed password or an error if hashing fails.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}




