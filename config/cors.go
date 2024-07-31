package config

import (
	"github.com/gofiber/fiber/v2"
	"aidanwoods.dev/go-paseto"
	"time"
)

// GenerateToken creates a new PASETO token
func GenerateToken(userID string) (string, error) {
	now := time.Now()
	builder := paseto.NewToken().
		SetIssuedAt(now).
		SetExpiration(now.Add(24 * time.Hour)).
		SetSubject(userID)

	secret := []byte("YELLOW SUBMARINE, BLACK WIZARDRY")
	return builder.V2Encrypt(secret, nil)
}

// ErrorResponse sends an error response
func ErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": message,
	})
}