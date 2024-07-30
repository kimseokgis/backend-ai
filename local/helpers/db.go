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


// connectToMongoDB establishes a connection to MongoDB.


// insertUserToDB inserts a user into the specified MongoDB collection.


// findUserInDB finds a user in the specified MongoDB collection by username.


// hashPassword generates a hashed password using bcrypt.


// checkPasswordHash compares a hashed password with a plain password.


// SetConnection establishes a connection to the MongoDB database.


// InsertUser inserts a user into the "users" collection in the database.


// ValidatePassword validates the user's password against the stored hash in the database.


// HashPass hashes a password using bcrypt.


// CheckPasswordHash compares a hashed password with a plain password.
