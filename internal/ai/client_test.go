package ai_test

import (
	"os"
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/config"
)

func TestGeminiClientConnection(t *testing.T) {

	if os.Getenv("CI") == "true" {
		t.Skip("Skipping Gemini API test in CI environment")
	}

	cfg := config.LoadConfig()

	client := ai.NewGeminiClient(cfg.GeminiAPIKey)
	err := client.TestConnection()

	if err != nil {
		t.Fatalf("❌ Gemini connection failed: %v", err)
	}

	t.Log("✅ Gemini client connection successful")
}
