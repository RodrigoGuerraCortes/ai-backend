package mocks

import (
	"context"
)

type MockAIClient struct {
	MockChat func(ctx context.Context, message string) (string, error)
}

func (m *MockAIClient) Chat(ctx context.Context, message string) (string, error) {
	return m.MockChat(ctx, message)
}
