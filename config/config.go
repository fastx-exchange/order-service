package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
}
