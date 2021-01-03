package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represent app configuration
type Config struct {
	Port string
}

// GetConfig return configurations
func GetConfig() *Config {
	config := &Config{Port: getEnvVariable("PORT")}

	return config
}

func getEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
