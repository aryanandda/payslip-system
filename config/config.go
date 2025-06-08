package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	DatabaseDSN string
	JWTSecret   string
	Environment string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
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

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
