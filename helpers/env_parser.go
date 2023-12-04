package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from the .env file.
func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

// GetEnvVariable retrieves the value of the specified environment variable.
func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
