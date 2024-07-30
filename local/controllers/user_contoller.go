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
