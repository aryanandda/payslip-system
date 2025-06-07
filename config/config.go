package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string // e.g. "8080"
	DatabaseDSN string // e.g. postgres://user:pass@localhost:5432/db?sslmode=disable
	JWTSecret   string // for signing tokens
	Environment string // development | production | test
}

var (
	cfg  *Config
	once sync.Once
)

// Load reads from .env or environment and returns a singleton Config instance.
func Load() *Config {
	once.Do(func() {
		// Try to load from .env file
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found. Using environment variables.")
		}

		cfg = &Config{
			AppPort:     getEnv("PORT", "8080"),
			DatabaseDSN: getEnv("DATABASE_DSN", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
			JWTSecret:   getEnv("JWT_SECRET", "supersecretkey"),
			Environment: getEnv("ENV", "development"),
		}
	})

	return cfg
}

// getEnv returns env variable or fallback if not set.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
