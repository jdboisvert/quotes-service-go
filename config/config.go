package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	ConnectionString string
}

func GetConfiguration() Configuration {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	configuration := Configuration{
		os.Getenv("CONNECTION_STRING"),
	}

	return configuration
}
