package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/helper"
)

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Makmur",
	AppName:       "Makmur ai",
}
var IPport, netstring = helper.GetAddress()
