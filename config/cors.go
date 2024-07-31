package config

import (
	"net/http"
	"os"
	"strings"
	
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
	"https://iteung.ulbi.ac.id",
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