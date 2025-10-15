package service

import (
	"context"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
)

type ChatService struct {
	ai ai.AIClient
}

func NewChatService(aiClient ai.AIClient) *ChatService {
	return &ChatService{ai: aiClient}
}

func (s *ChatService) Chat(ctx context.Context, message string) (string, error) {
	return s.ai.Chat(ctx, message)
}
