package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

// loads env file
func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
