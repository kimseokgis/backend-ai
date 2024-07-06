package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/helper"
)

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Iteung",
	AppName:       "Message Router",
}
var IPport, netstring = helper.GetAddress()
