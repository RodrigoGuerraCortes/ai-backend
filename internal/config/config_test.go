package config_test

import (
	"os"
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/config"
)

func TestLoadConfig_EnvOverrides(t *testing.T) {
	os.Setenv("PORT", "9090")
	os.Setenv("GEMINI_API_KEY", "dummy")
	os.Setenv("ENV", "development")
	os.Setenv("LOG_LEVEL", "debug")
	defer os.Clearenv()

	cfg := config.LoadConfig()
	if cfg.Port != "9090" || cfg.Environment != "development" || cfg.LogLevel != "debug" {
		t.Fatalf("env not applied: %#v", cfg)
	}
}
