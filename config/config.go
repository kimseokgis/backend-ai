package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

				Prefork:       true,
				CaseSensitive: true,
				StrictRouting: true,
				ServerHeader:  "Makmur",
				AppName:       "Makmur ai",
}
		var IPport, netstring = helper.GetAddress()

	var PrivateKey = os.Getenv("PRIVATEKEY")
	var PublicKey = os.Getenv("PUBLICKEY")