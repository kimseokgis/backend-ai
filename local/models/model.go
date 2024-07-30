package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user in the system.
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`  // ID is the unique identifier for the user.
	Username     string             `json:"username"`       // Username is the user's username.
	Email        string             `json:"email"`          // Email is the user's email address.
	Password     string             `json:"password"`       // Password is the user's plain text password. This field should not be stored in the database.
	PasswordHash string             `json:"password_hash"`  // PasswordHash is the hashed version of the user's password.
}

