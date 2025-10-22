package config

import (
	"log"
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

	// Handle missing API key gracefully
	if cfg.GeminiAPIKey == "" {
		// If logger is not initialized (tests/CI), fallback to std log
		if logger.Log != nil {
			logger.Log.Warn("⚠️ GEMINI_API_KEY not found — using dummy key (likely in CI)")
		} else {
			log.Println("⚠️ GEMINI_API_KEY not found — using dummy key (likely in CI)")
		}

		// Assign dummy key to avoid nil pointer in tests
		cfg.GeminiAPIKey = "dummy-key-for-tests"
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
