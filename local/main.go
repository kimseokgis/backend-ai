package local

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kimseokgis/backend-ai/local/config"
)

func main() {
	app := fiber.New(config.FiberConfig)

	// Middleware
	app.Use(logger.New())
}
