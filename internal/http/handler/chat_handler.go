package handler

import (
	"net/http"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/dto"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatService *service.ChatService
}

func NewChatHandler(gemini *ai.GeminiClient) *ChatHandler {
	return &ChatHandler{
		chatService: service.NewChatService(gemini),
	}
}

func (h *ChatHandler) Chat(c *gin.Context) {
	var req dto.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := h.chatService.Chat(c, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ChatResponse{Reply: response})
}
