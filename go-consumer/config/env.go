package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvironment() {
	// loads environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
