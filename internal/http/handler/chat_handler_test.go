package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai/mocks"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func TestChatHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAI := &mocks.MockAIClient{
		MockChat: func(_ context.Context, _ string) (string, error) {
			return "Hello from mock!", nil
		},
	}
	h := handler.NewChatHandler(mockAI)

	r := gin.New()
	r.POST("/api/v1/chat", h.Chat)

	body, _ := json.Marshal(map[string]string{"message": "hola"})
	req, _ := http.NewRequest("POST", "/api/v1/chat", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestChatHandler_InvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAI := &mocks.MockAIClient{ // no se usará
		MockChat: func(_ context.Context, _ string) (string, error) { return "", nil },
	}
	h := handler.NewChatHandler(mockAI)

	r := gin.New()
	r.POST("/api/v1/chat", h.Chat)

	req, _ := http.NewRequest("POST", "/api/v1/chat", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
