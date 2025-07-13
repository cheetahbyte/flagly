package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cheetahbyte/flagly/apis"
	"github.com/cheetahbyte/flagly/internal"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ContextLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}

func main() {
	router := gin.Default()

	logger := zap.Must(zap.NewDevelopment())
	if os.Getenv("GIN_MODE") == "release" {
		logger = zap.Must(zap.NewProduction())
	}

	defer logger.Sync()
	sugar := logger.Sugar()

	internal.InitStorage()

	router.Use(ErrorHandler())
	router.Use(gin.Recovery())
	router.Use(ContextLogger(sugar))

	router.GET("/flags", apis.GetAllFlags)
	router.GET("/flags/:flag", apis.GetFlag)
	router.GET("/flags/:flag/enabled", apis.GetFlagEnabled)

	router.GET("/environments", apis.GetAllEnvironments)
	router.GET("/environments/:env", apis.GetEnvironment)

	log.Println("Server listening on http://localhost:8080")
	if err := router.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
