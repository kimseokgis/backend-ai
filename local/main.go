package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kimseokgis/backend-ai/local/config"
	"github.com/kimseokgis/backend-ai/local/controllers"
)

func main() {
	app := fiber.New(config.FiberConfig)

	// Middleware
	app.Use(logger.New())

