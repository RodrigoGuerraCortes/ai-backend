package container

import (
	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/config"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/handler"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/service"
	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"go.uber.org/zap"
)

type Container struct {
	Config      *config.Config
	AIClient    *ai.GeminiClient
	ChatService *service.ChatService
	ChatHandler *handler.ChatHandler
}

func BuildContainer() *Container {
	logger.Init()
	defer logger.Sync()

	cfg := config.LoadConfig()
	logger.Log.Info("✅ Configuration loaded successfully",
		zap.String("env", cfg.Environment),
		zap.String("port", cfg.Port),
	)

	aiClient := ai.NewGeminiClient()
	logger.Log.Info("✅ Gemini client created")

	chatService := service.NewChatService(aiClient)
	chatHandler := handler.NewChatHandler(aiClient)

	return &Container{
		Config:      cfg,
		AIClient:    aiClient,
		ChatService: chatService,
		ChatHandler: chatHandler,
	}
}
