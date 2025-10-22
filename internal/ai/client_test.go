package ai_test

import (
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/config"
)

func TestGeminiClientConnection(t *testing.T) {
	cfg := config.LoadConfig()

	client := ai.NewGeminiClient(cfg.GeminiAPIKey)
	err := client.TestConnection()

	if err != nil {
		t.Fatalf("❌ Gemini connection failed: %v", err)
	}

	t.Log("✅ Gemini client connection successful")
}
