package config

import (
	"net/http"
	"os"
	"strings"
	
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
	"https://iteung.ulbi.ac.id",
	"https://whatsauth.github.io",
	"https://rofinafiin.github.io",
	"https://gocroot.github.io/",
	"https://gocroot-baru.herokuapp.com/",
	"https://kimseokgis.github.io",
	"https://kimseokgis.advocata.me",
}

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Request-Headers, token, Access-Control-Allow-Origin, Authorization, Bearer, login",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}

func isAllowedOrigin(origin string) bool {
	for _, o := range origins {
		if o == origin {
			return true
		}
	}
	return false
}

func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {