package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	DBSSL   string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppPort: getEnv("APP_PORT", ":8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "go_user"),
		DBPass:  getEnv("DB_PASSWORD", "password123"),
		DBName:  getEnv("DB_NAME", "go_clean_api"),
		DBSSL:   getEnv("DB_SSLMODE", "disable"),
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultValue
}
