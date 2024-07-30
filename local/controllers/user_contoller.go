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


// comparePasswords compares a hashed password with a plain password.


// generateToken generates a JWT token for the given username.


// RegisterUser handles user registration.


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
