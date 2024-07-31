package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/helper"
	"os"
)

var Iteung = fiber.Config{
				Prefork:       true,
				CaseSensitive: true,
				StrictRouting: true,
				ServerHeader:  "Makmur",
				AppName:       "Makmur ai",
}
		var IPport, netstring = helper.GetAddress()

	var PrivateKey = os.Getenv("PRIVATEKEY")
	var PublicKey = os.Getenv("PUBLICKEY")