package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL string
    Port        string
}

func Load() *Config {

    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on system environment variables")
    }

    return &Config{
        DatabaseURL: mustGet("DATABASE_URL"),
        Port:        getOrDefault("PORT", "3000"),
    }
}

func mustGet(key string) string {
    val := os.Getenv(key)
    if val == "" {
        log.Fatalf("Required env var %s is not set", key)
    }
    return val
}

func getOrDefault(key, fallback string) string {
    val := os.Getenv(key)
    if val == "" {
        return fallback
    }
    return val
}