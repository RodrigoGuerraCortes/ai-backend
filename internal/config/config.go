package config

import (
	"os"

	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	GeminiAPIKey string
	Environment  string
	LogLevel     string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	cfg := &Config{
		Port:         getEnv("PORT", "8080"),
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),
		Environment:  getEnv("ENV", "development"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
	}

	if cfg.GeminiAPIKey == "" {
		logger.Log.Fatal("Missing GEMINI_API_KEY in environment variables")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
