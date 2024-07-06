package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

var origins = []string{
	"https://auth.ulbi.ac.id",
	"https://sip.ulbi.ac.id",
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
	"https://euis.ulbi.ac.id",
	"https://home.ulbi.ac.id",
	"https://alpha.ulbi.ac.id",
	"https://dias.ulbi.ac.id",
	"https://iteung.ulbi.ac.id",
	"https://whatsauth.github.io",
	"https://rofinafiin.github.io",
	"https://gocroot.github.io/",
	"https://gocroot-baru.herokuapp.com/",
}

func main() {
	//go whatsauth.RunHub()
	site := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Makmur",
		AppName:       "Makmur ai",
	})
	site.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(origins[:], ","),
		AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Request-Headers, token, Access-Control-Allow-Origin, Authorization, Bearer",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	},
	))
	Web(site)
	log.Fatal(site.Listen("8080"))
}

func Web(page *fiber.App) {
}
