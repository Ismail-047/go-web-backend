package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port        string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on system environment variables")
	}

	return &Config{
		DatabaseUrl: mustGet("DATABASE_URL"),
		Port:        getOrDefault("PORT", "3000"),
	}
}

func mustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Required env %v is not set", val)
	}
	return val
}

func getOrDefault(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		val = fallback
	}
	return val
}