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
	"https://whatsauth.github.io",
	"https://rofinafiin.github.io",
	"https://gocroot.github.io/",
	"https://gocroot-baru.herokuapp.com/",
	"https://kimseokgis.github.io",
	"https://kimseokgis.advocata.me",
}

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),

// ErrorResponse sends an error response
func ErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": message,
	})
}