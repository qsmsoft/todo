package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBSslMode  string

	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Printf("Error converting DB_PORT to int: %v", err)
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSPORT"),
		DBSslMode:  os.Getenv("DB_SSL_MODE"),

		Port: os.Getenv("SERVER_PORT"),
	}

	return config
}
