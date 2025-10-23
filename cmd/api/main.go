package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/config"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/di"
	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"go.uber.org/zap"
)

// @title API Backend IA
// @version 1.0
// @description Backend API for AI interactions using Gemini models.
// @host 161.153.200.7:8080
// @schemes http https
func main() {

	cfg := config.LoadConfig()
	logger.InitWith(cfg.Environment, cfg.LogLevel)
	log := logger.Log

	log.Info("🚀 Starting AI Backend with Gemini + DI...")

	// Build container (inject logger instance)
	c := di.New(log)

	// Start HTTP server (non-blocking)
	if err := c.Start(); err != nil {
		log.Fatal("server start failed", zap.Error(err))
	}

	// OS signal handling for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block until signal

	log.Info("🔻 Shutdown signal received, stopping...")
	c.Stop(10 * time.Second)
}
