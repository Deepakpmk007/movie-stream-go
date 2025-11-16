package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

func InitDB() {
	// Load env
	_ = godotenv.Load(".env")

	MONGO_URI := os.Getenv("DATABASE_URL")
	DB_NAME := os.Getenv("DATABASE_NAME")

	if MONGO_URI == "" {
		log.Fatal("DATABASE_URL not set!")
	}
	if DB_NAME == "" {
		log.Fatal("DATABASE_NAME not set!")
	}

	// Connect
	client, err := mongo.Connect(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal("MongoDB connection failed → ", err)
	}

	Client = client
	Database = client.Database(DB_NAME)

	log.Println("MongoDB connected successfully!")
}

// Helper function to get any collection
func Collection(name string) *mongo.Collection {
	return Database.Collection(name)
}
