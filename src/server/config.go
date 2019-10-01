package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Struct to hold the server configuration
type config struct {
	EnvLoaded bool
	Port string
}

var Config config

func LoadEnv() {
	// Load environment variables from .env if it exists
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("No .env file, so just use the environment")
	}

	Config = config{
		EnvLoaded: true,
		Port: getEnv("PORT"),
	}
}

func getEnv(f string) string {
	return os.Getenv(f)
}
