package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
	func ConnectDatabase() {
		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		var IPport, netstring = helper.GetAddress()

	var PrivateKey = os.Getenv("PRIVATEKEY")
	var PublicKey = os.Getenv("PUBLICKEY")