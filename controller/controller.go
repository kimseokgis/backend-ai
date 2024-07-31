package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
)

// func response homemakmur
func HomeMakmur(w http.ResponseWriter, r *http.Request) {
	Response := fmt.Sprintf("Makmur AI chooy %s", "8080")
	response, err := json.Marshal(Response)
	if err != nil {
		http.Error(w, "Internal server error: JSON marshaling failed", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	return
}

// func notfound resp
func NotFound(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Message = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, resp)
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

	config.DB.Save(&user)
	return c.JSON(user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return helper.ErrorResponse(c, "User not found")
	}

	config.DB.Delete(&user)
	return c.SendStatus(fiber.StatusNoContent)
}