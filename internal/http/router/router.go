package router

import (
	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(gemini *ai.GeminiClient) *gin.Engine {
	r := gin.Default()

	chatHandler := handler.NewChatHandler(gemini)
	api := r.Group("/api/v1")
	api.POST("/chat", chatHandler.Chat)

	return r
}
