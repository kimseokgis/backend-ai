package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var FiberConfig = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Makmur",
	AppName:       "Makmur AI",
}
var (
	IPport     = os.Getenv("IP") + ":" + os.Getenv("PORT")
	PrivateKey = os.Getenv("PRIVATEKEY")
)
