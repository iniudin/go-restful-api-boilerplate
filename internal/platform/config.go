package platform

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_NAME     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_URL      string
	APP_NAME          string
	APP_ENV           string
	APP_KEY           string
	APP_DEBUG         string
	APP_HOST          string
	APP_PORT          string
}

func LoadConfig() error {
	return godotenv.Load()
}

func NewConfig() *Config {
	if err := LoadConfig(); err != nil {
		return nil
	}
	return &Config{
		DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
		DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
		DATABASE_USER:     os.Getenv("DATABASE_USER"),
		DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
		DATABASE_URL:      os.Getenv("DATABASE_URL"),
		APP_NAME:          os.Getenv("APP_NAME"),
		APP_ENV:           os.Getenv("APP_ENV"),
		APP_KEY:           os.Getenv("APP_KEY"),
		APP_DEBUG:         os.Getenv("APP_DEBUG"),
		APP_HOST:          os.Getenv("APP_HOST"),
		APP_PORT:          os.Getenv("APP_PORT"),
	}
}
