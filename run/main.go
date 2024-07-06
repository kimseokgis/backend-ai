package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(config.IPport))
}
