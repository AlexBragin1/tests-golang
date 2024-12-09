package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	MongoDB  string
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return Config{
		MongoURI: os.Getenv("MONGO_URI"),
		MongoDB:  os.Getenv("MONGO_DB"),
	}, nil
}
