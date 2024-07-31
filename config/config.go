package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/helper"
	"os"
)

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,

	func ConnectDatabase() {
		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		
	if err != nil {
		panic("Failed to connect to the database!")
	}

DB = database
fmt.Println("Database connection established")
}