package main

import (
	"github.com/RodrigoGuerraCortes/ai-backend/internal/container"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/router"
	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.Sync()

	logger.Log.Info("🚀 Starting AI Backend with Gemini + DI...")

	c := container.BuildContainer()

	if err := c.AIClient.TestConnection(); err != nil {
		logger.Log.Fatal("❌ Gemini connection failed", zap.Error(err))
	}

	r := router.NewRouter(c.AIClient)
	port := ":" + c.Config.Port

	logger.Log.Info("✅ Server starting", zap.String("port", port))

	if err := r.Run(port); err != nil {
		logger.Log.Fatal("❌ Failed to start server", zap.Error(err))
	}
}
