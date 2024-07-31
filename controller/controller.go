package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// GetUsers returns all users
func GetUsers(c *fiber.Ctx) error {
	response, err := json.Marshal(Response)
	var users []model.User
	config.DB.Find(&users)
	return c.JSON(users)
	}

// GetUser returns a user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return helper.ErrorResponse(c, "User not found")
	}
	return c.JSON(user)
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	var user model.User
	helper.WriteJSON(respw, http.StatusNotFound, resp)

	if err := c.BodyParser(&user); err != nil {
		return helper.ErrorResponse(c, err.Error())
	}

	// Hash the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return helper.ErrorResponse(c, "Could not create user")
	}

	return c.JSON(user)
}

// UpdateUser updates a user by ID
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return helper.ErrorResponse(c, "User not found")
	}

	if err := c.BodyParser(&user); err != nil {
		return helper.ErrorResponse(c, err.Error())
	}
