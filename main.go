package main

import (
	"flag"
	"log"
	"os"

	"github.com/cheetahbyte/flagly/apis"
	"github.com/cheetahbyte/flagly/internal/flagly"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	configFile := flag.String("config", "./flagly.yml", "Path to the configuration file")
	flag.Parse()
	router := gin.Default()

	logger := zap.Must(zap.NewDevelopment())
	if os.Getenv("GIN_MODE") == "release" {
		logger = zap.Must(zap.NewProduction())
	}

	defer logger.Sync()
	sugar := logger.Sugar()

	store, err := flagly.InitStorage(*configFile)
	if err != nil {
		sugar.Fatalf("Failed to read file: %v", err)
	}
	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.Use(flagly.ContextLogger(sugar))
	router.Use(flagly.ErrorHandlerMiddleware())

	flagApi := apis.NewFlagAPI(store)
	environmentApi := apis.NewEnvironmentAPI(store)
	generalApi := apis.NewGeneralAPI(store)
	generalApi.RegisterRoutes(router)
	flagApi.RegisterRoutes(router)
	environmentApi.RegisterRoutes(router)

	log.Println("Server listening on http://localhost:8080")
	if err := router.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
