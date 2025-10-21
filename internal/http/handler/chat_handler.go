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

func NewChatHandler(ai ai.AIClient) *ChatHandler {
	return &ChatHandler{
		chatService: service.NewChatService(ai),
	}
}

func (h *ChatHandler) Chat(c *gin.Context) {
	var req dto.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	reply, err := h.chatService.Chat(c, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Data: gin.H{
			"reply": reply,
		},
	})
}
