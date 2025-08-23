package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	SERVER_PORT string
}

func LoadConfig() *Config{

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		SERVER_PORT: os.Getenv("SERVER_PORT"),
	}
	return config
}
