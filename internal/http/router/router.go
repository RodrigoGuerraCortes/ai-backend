package router

import (
	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/handler"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(aiClient ai.AIClient) *gin.Engine {

	r := gin.Default()

	// Apply middlewares
	r.Use(gin.Recovery()) // built-in panic recovery
	r.Use(middleware.RequestLogger())
	r.Use(middleware.ErrorHandler())

	api := r.Group("/api/v1")
	{
		chatHandler := handler.NewChatHandler(aiClient)
		api.POST("/chat", chatHandler.Chat)
	}

	return r
}
