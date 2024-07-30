package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/local/config"
	"github.com/kimseokgis/backend-ai/local/helpers"
	"github.com/kimseokgis/backend-ai/model"
)

func RegisterUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	conn := helper.SetConnection()
	defer conn.Client().Disconnect(context.TODO())

	hash, err := helper.HashPass(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}
	user.PasswordHash = hash
	helpers.InsertUser(conn, user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var userdata model.User
	if err := c.BodyParser(&userdata); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing application/json: " + err.Error(),
		})
	}

	conn := helper.SetConnection()
	defer conn.Client().Disconnect(context.TODO())

	if helper.IsPasswordValid(conn, userdata) {
		tokenstring, err := helper.EncodeWithUsername(userdata.Username, config.PrivateKey)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Gagal Encode Token : " + err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Selamat Datang Anda Berhasil Login",
			"token":   tokenstring,
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Username atau Password Anda Salah",
	})
}
