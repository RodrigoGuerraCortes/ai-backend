package router

import (
	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	_ "github.com/RodrigoGuerraCortes/ai-backend/internal/http/docs"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/handler"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(aiClient ai.AIClient) *gin.Engine {

	r := gin.Default()

	// Apply middlewares
	r.Use(gin.Recovery()) // built-in panic recovery
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.ErrorHandler())

	api := r.Group("/api/v1")
	{
		chatHandler := handler.NewChatHandler(aiClient)
		api.POST("/chat", chatHandler.Chat)
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
