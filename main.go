package backend_ai

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/routes"
)

func init() {
	functions.HTTP("makmur", url.Web)
}
