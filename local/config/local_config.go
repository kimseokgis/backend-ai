package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)







// IsAllowedOrigin checks if the origin is allowed
func IsAllowedOrigin(origin string) bool {
	origins := []string{
		"http://127.0.0.1:5500",
		"http://127.0.0.1:5501",
	}
	for _, o := range origins {
		if o == origin {
			return true
		}
	}
	return false
}
