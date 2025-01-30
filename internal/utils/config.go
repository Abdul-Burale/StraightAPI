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

func GetMongoDatabaseName() string {
	return os.Getenv("DATABASE_NAME")
}

func GetMongoDatabaseCollection() string {
	return os.Getenv("COLLECTION_NAME")
}

// FireBase Methods

func GetFirebaseType() string {
	return os.Getenv("FIREBASE_TYPE")
}
func GetFirebaseProjectID() string {
	return os.Getenv("FIREBASE_PROJECT_ID")
}
func GetFirebasePrivateKeyID() string {
	return os.Getenv("FIREBASE_PRIVATE_KEY_ID")
}
func GetFirebasePrivateKey() string {
	return os.Getenv("FIREBASE_PRIVATE_KEY")
}
func GetFirebaseClientEmail() string {
	return os.Getenv("FIREBASE_CLIENT_EMAIL")
}
func GetFirebaseClientID() string {
	return os.Getenv("FIREBASE_CLIENT_ID")
}
func GetFirebaseAuth_URI() string {
	return os.Getenv("FIREBASE_AUTH_URI")
}
func GetFirebaseToken_URI() string {
	return os.Getenv("FIREBASE_TOKEN_URI")
}
func GetFirebaseAuthProviderX509Cert() string {
	return os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERTL_URL")
}
func GetFirebaseClientX509Cert() string {
	return os.Getenv("FIREBASE_CLIENT_X509_CERT_URL")
}
func GetFirebaseUniverseDomain() string {
	return os.Getenv("FIREBASE_UNIVERSE_DOMAIN")
}
