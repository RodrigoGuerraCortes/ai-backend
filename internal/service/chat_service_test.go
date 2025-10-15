package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai/mocks"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/service"
)

func TestChatService_Success(t *testing.T) {
	mockAI := &mocks.MockAIClient{
		MockChat: func(ctx context.Context, message string) (string, error) {
			return "Hello from mock AI!", nil
		},
	}

	svc := service.NewChatService(mockAI)
	resp, err := svc.Chat(context.Background(), "Hello")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp != "Hello from mock AI!" {
		t.Errorf("expected mock response, got %s", resp)
	}
}

func TestChatService_Error(t *testing.T) {
	mockAI := &mocks.MockAIClient{
		MockChat: func(ctx context.Context, message string) (string, error) {
			return "", errors.New("AI error")
		},
	}

	svc := service.NewChatService(mockAI)
	_, err := svc.Chat(context.Background(), "Hello")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
