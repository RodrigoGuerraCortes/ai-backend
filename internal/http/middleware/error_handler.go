package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler captures unhandled errors and panics
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   c.Errors[0].Error(),
			})
		}
	}
}
