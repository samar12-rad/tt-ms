package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
}
