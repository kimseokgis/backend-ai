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

	if err != nil {
		panic("Failed to connect to the database!")
	}
	
DB = database
fmt.Println("Database connection established")
}