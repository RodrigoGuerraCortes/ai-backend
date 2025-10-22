package middleware

import (
	"net/http"

	"github.com/RodrigoGuerraCortes/ai-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorHandler captures unhandled errors and panics
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Log.Error("panic recovered", zap.Any("error", rec))
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Internal server error",
				})
				c.Abort()
			}
		}()

		c.Next()

		// Handle explicit errors from handlers
		if len(c.Errors) > 0 {
			logger.Log.Error("handler error", zap.String("message", c.Errors[0].Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   c.Errors[0].Error(),
			})
		}
	}
}
