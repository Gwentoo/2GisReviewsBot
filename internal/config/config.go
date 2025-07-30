package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	TelegramToken string `env:"TELEGRAM_TOKEN"`
	DBHost        string `env:"DB_HOST"`
	DBPort        string `env:"DB_PORT"`
	DBUser        string `env:"DB_USER"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBName        string `env:"DB_NAME"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		DBPort:        os.Getenv("DB_PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
	}
}
