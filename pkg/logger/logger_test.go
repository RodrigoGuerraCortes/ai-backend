package logger_test

import (
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/config"
	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
)

func TestInitDoesNotPanic(t *testing.T) {
	cfg := config.LoadConfig()
	logger.InitWith(cfg.Environment, cfg.LogLevel)
	defer logger.Sync()
	// Si llega acá, ok
}
