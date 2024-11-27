package utils

import (
	"fmt"
	"log"
	"os"

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

// loads env variables
func LoadEnvVariable(key string) string {
	// Load .env file
	envVar := os.Getenv(key)
	if envVar == "" {
		fmt.Errorf("environment variable %s is not set", envVar)
	}
	return envVar
}
