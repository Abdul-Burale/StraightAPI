package utils

import (
	"log"
	"os"

	// godotenv is a package to load environment variables from .env
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
	// Load the .env file into environment variables
	err := godotenv.Load()
	if err != nil {
		// If there is an error loading the .env file, log the error and terminate app
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// GetMongoURI retives the MongoDB URI from the environment variables
func GetMongoURI() string {
	// Return the value of the environment variable MONGO_URI
	return os.Getenv("MONGO_URI")
}
