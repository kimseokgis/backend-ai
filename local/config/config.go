package config

import "github.com/gofiber/fiber/v2"

var FiberConfig = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Makmur",
	AppName:       "Makmur AI",
}
