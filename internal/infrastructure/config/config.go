package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBType      string
	DBDSN       string
	GRPCPort    string
	Environment string
	LogLevel    string
}

var PRODUCTION_ENV = "production"

func Load() Config {
	if os.Getenv("ENV") != PRODUCTION_ENV {
		if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
			log.Printf("No .env file found: %v", err)
		}
	}

	return Config{
		DBType:      getEnv("DB_TYPE", "postgres"),
		DBDSN:       getEnv("DB_DSN", "postgresql://user:password@localhost:5432/userdb?sslmode=disable"),
		GRPCPort:    getEnv("GRPC_PORT", "50051"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
