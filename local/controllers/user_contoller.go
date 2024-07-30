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
