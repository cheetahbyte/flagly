//go:build !test

package main

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/cheetahbyte/flagly/apis"
	"github.com/cheetahbyte/flagly/internal/audit"
	"github.com/cheetahbyte/flagly/internal/storage"
	"github.com/cheetahbyte/flagly/pkg/flagly/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//go:embed all:dashboard/dist
var embeddedFS embed.FS

func main() {
	configFile := flag.String("config", "./flagly.yml", "Path to the configuration file")
	flag.Parse()

	logger := zap.Must(zap.NewDevelopment())
	if os.Getenv("GIN_MODE") == "release" {
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	store, err := storage.InitStorage(*configFile)
	if err != nil {
		sugar.Fatalf("Failed to initialize storage from config file '%s': %v", *configFile, err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.Use(middleware.ContextLogger(sugar))
	router.Use(middleware.ErrorHandlerMiddleware())

	auditService := audit.NewDefaultAuditService()

	apiGroup := router.Group("/api")
	apis.NewGeneralAPI(store).RegisterRoutes(apiGroup)
	apis.NewFlagAPI(store, auditService).RegisterRoutes(apiGroup)
	apis.NewEnvironmentAPI(store).RegisterRoutes(apiGroup)

	if os.Getenv("GIN_MODE") == "release" {
		distFS, err := fs.Sub(embeddedFS, "dashboard/dist")
		if err != nil {
			sugar.Fatalf("Failed to create sub-filesystem for embedded assets: %v", err)
		}
		fileServer := http.FileServer(http.FS(distFS))

		router.NoRoute(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
				return
			}

			filePath := strings.TrimPrefix(c.Request.URL.Path, "/")
			_, err := distFS.Open(filePath)

			if err != nil {
				indexBytes, readErr := fs.ReadFile(distFS, "index.html")
				if readErr != nil {
					sugar.Error("Could not read index.html from embedded fs", zap.Error(readErr))
					c.String(http.StatusInternalServerError, "index.html not found")
					return
				}
				c.Data(http.StatusOK, "text/html; charset=utf-8", indexBytes)
			} else {
				fileServer.ServeHTTP(c.Writer, c.Request)
			}
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	sugar.Infof("Server listening on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		sugar.Fatalf("Error starting server: %v", err)
	}
}
