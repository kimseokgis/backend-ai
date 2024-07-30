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



func LoginUser(c *fiber.Ctx) error {
	var credentials model.User
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	conn := helper.SetConnection()
	defer conn.Client().Disconnect(context.TODO())

	// Retrieve user from database
	var storedUser model.User
	err := conn.Collection("users").FindOne(context.TODO(), bson.M{"username": credentials.Username}).Decode(&storedUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username atau Password Anda Salah",
		})
	}

	// Compare hash and password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(credentials.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username atau Password Anda Salah",
		})
	}

	// Generate token or handle successful login
	token, err := helper.EncodeWithUsername(storedUser.Username, config.PrivateKey)
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
