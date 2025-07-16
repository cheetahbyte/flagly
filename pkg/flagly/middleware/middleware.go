package middleware

import (
	"log"
	"net/http"

	custom_errors "github.com/cheetahbyte/flagly/internal/error"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ContextLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			log.Printf("Request error: %v", err)
			c.Header("Content-Type", "application/problem+json")
			if apiErr, ok := err.(*custom_errors.APIError); ok {
				c.JSON(apiErr.Status, gin.H{
					"type":     apiErr.Type,
					"title":    apiErr.Title,
					"status":   apiErr.Status,
					"detail":   apiErr.Detail,
					"instance": c.Request.URL.Path,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"type":     "/errors/internal-server-error",
					"title":    "Internal Server Error",
					"status":   http.StatusInternalServerError,
					"detail":   "An unexpected error occurred. Please try again later.",
					"instance": c.Request.URL.Path,
					"code":     "INTERNAL_SERVER_ERROR",
				})
			}
			c.Abort()
		}
	}
}
