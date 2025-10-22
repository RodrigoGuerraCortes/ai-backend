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
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// shared helper for creating router
func setupRouter(t *testing.T) *gin.Engine {
	_ = godotenv.Load(".env")
	geminiClient := ai.NewGeminiClient()
	return router.NewRouter(geminiClient)
}

// ===========================
// ✅ SUCCESS TEST
// ===========================
func TestChatIntegration_Success(t *testing.T) {
	r := setupRouter(t)

	body := map[string]string{"message": "Hola, ¿cómo estás?"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/v1/chat", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	// Basic schema checks
	if response["success"] != true {
		t.Fatalf("expected success=true, got %v", response["success"])
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok || data["reply"] == "" {
		t.Fatalf("expected reply field in data, got %v", data)
	}

	t.Logf("✅ Chat success test passed. Reply: %v", data["reply"])
}

// ===========================
// ⚠️ INVALID REQUEST TEST
// ===========================
func TestChatIntegration_InvalidBody(t *testing.T) {
	r := setupRouter(t)

	req, _ := http.NewRequest("POST", "/api/v1/chat", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if response["success"] != false {
		t.Fatalf("expected success=false, got %v", response["success"])
	}

	if response["error"] == "" {
		t.Fatalf("expected error message, got %v", response)
	}

	t.Logf("✅ Chat invalid body test passed. Error: %v", response["error"])
}
