package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var FiberConfig = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Makmur",
	AppName:       "Makmur AI",
}
var (
	IPport       = os.Getenv("IP") + ":" + os.Getenv("PORT")
	PrivateKey   = os.Getenv("PRIVATEKEY")
	PublicKey    = os.Getenv("PUBLICKEY")
	Internalhost = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")
)

var origins = []string{
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
}
var CorsConfig = cors.Config{
	AllowOrigins:     strings.Join(origins, ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Request-Headers, token, Access-Control-Allow-Origin, Authorization, Bearer, login",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}

func IsAllowedOrigin(origin string) bool {
	for _, o := range origins {
		if o == origin {
			return true
		}
	}
	return false
}
