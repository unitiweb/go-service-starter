package env

import (
	"github.com/joho/godotenv"
	"os"
)

// Load environment variables
// This function should be used in the main.go file
func LoadEnv() {
	// Load environment variables from .env if it exists
	err := godotenv.Load()
	if err != nil {
		panic("No .env file, so just use the environment")
	}
}

// Get an environment variable. They must be loaded with "loadEnv()" first
func Get(f string) string {
	return os.Getenv(f)
}
