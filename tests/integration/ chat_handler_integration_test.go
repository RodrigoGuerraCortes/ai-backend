//go:build integration
// +build integration

package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/router"
)

func TestChatIntegration(t *testing.T) {
	// Given real Gemini client
	geminiClient := ai.NewGeminiClient()

	// Build router
	r := router.NewRouter(geminiClient)

	// Build request
	body := map[string]string{"message": "Hola, ¿cómo estás?"}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/v1/chat", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["reply"] == "" {
		t.Fatalf("expected a reply from AI, got empty response")
	}

	t.Logf("✅ Integration chat test passed. AI replied: %s", response["reply"])
}
