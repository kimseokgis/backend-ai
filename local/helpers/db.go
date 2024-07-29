package helpers

import (
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection() *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv("MONGOSTRING"),
		DBName:   "AI",
	}
	return atdb.MongoConnect(DBmongoinfo)
}
