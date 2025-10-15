package ai

import "context"

type AIClient interface {
	Chat(ctx context.Context, message string) (string, error)
}
