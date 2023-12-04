package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
