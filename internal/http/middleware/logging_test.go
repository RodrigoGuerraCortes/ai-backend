package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/middleware"
	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

// TestRequestLogger ensures middleware runs and calls c.Next()
func TestRequestLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// 🧠 Initialize zap before using middleware
	logger.Init()
	defer logger.Sync()

	// Track if handler executed
	handlerCalled := false

	// Fake route handler
	testHandler := func(c *gin.Context) {
		handlerCalled = true
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}

	// Setup router with our middleware
	r := gin.New()
	r.Use(middleware.RequestLogger())
	r.GET("/ping", testHandler)

	// Perform request
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ✅ Assertions
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	if !handlerCalled {
		t.Fatal("expected handler to be called, but it was not")
	}
}
