package ai_test

import (
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
)

func TestGeminiClientConnection(t *testing.T) {
	client := ai.NewGeminiClient()
	err := client.TestConnection()

	if err != nil {
		t.Fatalf("❌ Gemini connection failed: %v", err)
	}

	t.Log("✅ Gemini client connection successful")
}
