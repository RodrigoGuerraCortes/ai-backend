package router_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai/mocks"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/router"
	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestRouter_Healthz(t *testing.T) {

	gin.SetMode(gin.TestMode)

	logger.Log = zap.NewNop() // 👈 evita nil panic

	r := router.NewRouter(&mocks.MockAIClient{
		MockChat: func(_ context.Context, _ string) (string, error) { return "ok", nil },
	})

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
