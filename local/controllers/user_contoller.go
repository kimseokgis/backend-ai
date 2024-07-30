package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/local/config"
	"github.com/kimseokgis/backend-ai/local/helpers"
	"github.com/kimseokgis/backend-ai/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// parseUser parses the request body into a User model.
// Returns the User model or an error if parsing fails.
func parseUser(c *fiber.Ctx) (*model.User, error) {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// hashPassword hashes the given password.
// Returns the hashed password or an error if hashing fails.
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// saveUser saves the user to the database.
// Returns an error if saving fails.
func saveUser(user model.User) error {
	conn := helper.SetConnection()
	defer conn.Client().Disconnect(context.TODO())
	return helpers.InsertUser(conn, user)
}

// findUserByUsername retrieves a user from the database by username.
// Returns the User model or an error if retrieval fails.
func findUserByUsername(username string) (*model.User, error) {
	conn := helper.SetConnection()
	defer conn.Client().Disconnect(context.TODO())

	var user model.User
	err := conn.Collection("users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// comparePasswords compares a hashed password with a plain password.
// Returns an error if the passwords do not match.
func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

