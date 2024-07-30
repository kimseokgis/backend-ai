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

// checkPasswordHash compares a hashed password with a plain password.
// Returns true if the passwords match, false otherwise.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// SetConnection establishes a connection to the MongoDB database.
// Returns the database connection.
func SetConnection() *mongo.Database {
	dbInfo := getDBInfo()
	return connectToMongoDB(dbInfo)
}

// InsertUser inserts a user into the "users" collection in the database.
// Returns the inserted ID or nil if insertion fails.
func InsertUser(db *mongo.Database, user model.User) interface{} {
	collection := db.Collection("users")
	insertedID, err := insertUserToDB(collection, user)
	if err != nil {
		return nil
	}
	return insertedID
}

// ValidatePassword validates the user's password against the stored hash in the database.
// Returns true if the password is valid, false otherwise.
func ValidatePassword(conn *mongo.Database, user model.User) bool {
	collection := conn.Collection("users")
	storedUser, err := findUserInDB(collection, user.Username)
	if err != nil {
		return false
	}
	return checkPasswordHash(user.Password, storedUser.PasswordHash)
}