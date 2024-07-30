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


// hashPassword hashes the given password.


// saveUser saves the user to the database.

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

// generateToken generates a JWT token for the given username.
// Returns the token or an error if token generation fails.
func generateToken(username string) (string, error) {
	return helper.EncodeWithUsername(username, config.PrivateKey)
}

// RegisterUser handles user registration.
// Parses the request body, hashes the password, and saves the user to the database.
// Returns a success or error message as JSON.
func RegisterUser(c *fiber.Ctx) error {
	user, err := parseUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	user.PasswordHash = hash
	user.Password = "" // Clear plain password

	if err := saveUser(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error saving user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// LoginUser handles user login.
// Parses the request body, retrieves the user from the database,
// compares passwords, and generates a JWT token if successful.
// Returns a success message with the token or an error message as JSON.
func LoginUser(c *fiber.Ctx) error {
	credentials, err := parseUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	storedUser, err := findUserByUsername(credentials.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username atau Password Anda Salah",
		})
	}

	if err := comparePasswords(storedUser.PasswordHash, credentials.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username atau Password Anda Salah",
		})
	}

	token, err := generateToken(storedUser.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Selamat Datang Anda Berhasil Login",
		"token":   token,
	})
}
