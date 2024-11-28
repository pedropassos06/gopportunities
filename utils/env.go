package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// loads env file
func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
