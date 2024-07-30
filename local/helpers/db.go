package helpers

import (
	"context"
	"os"

	"github.com/aiteung/atdb"
	"github.com/kimseokgis/backend-ai/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// getDBInfo constructs the database info from environment variables.
// Returns the database information.
func getDBInfo() atdb.DBInfo {
	return atdb.DBInfo{
		DBString: os.Getenv("MONGOSTRING"),
		DBName:   "AI",
	}
}

// connectToMongoDB establishes a connection to MongoDB.
// Takes the database info and returns the database connection.
func connectToMongoDB(dbInfo atdb.DBInfo) *mongo.Database {
	return atdb.MongoConnect(dbInfo)
}





