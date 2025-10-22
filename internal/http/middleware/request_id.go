package middleware

import (
	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const RequestIDKey = "request_id"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get or create request ID
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Add to context and response
		c.Set(RequestIDKey, requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		// Log start of request
		logger.Log.Info("➡️ Incoming request",
			zap.String("request_id", requestID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		// Process next middleware/handler
		c.Next()

		// Log end of request
		logger.Log.Info("✅ Request completed",
			zap.String("request_id", requestID),
			zap.Int("status", c.Writer.Status()),
		)
	}
}
