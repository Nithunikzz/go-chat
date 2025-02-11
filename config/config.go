package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds application configurations
type Config struct {
	ServerPort  string
	DatabaseURL string
}

// LoadConfig reads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:secret@db:5432/chatapp?sslmode=disable"),
	}
}

// getEnv retrieves environment variables or sets a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
