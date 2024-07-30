package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)





func GetCorsConfig() cors.Config {
	origins := []string{
		"http://127.0.0.1:5500",
		"http://127.0.0.1:5501",
	}
	return cors.Config{
		AllowOrigins:     strings.Join(origins, ","),
		AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Request-Headers, token, Access-Control-Allow-Origin, Authorization, Bearer, login",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}
}

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
