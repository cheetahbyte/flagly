package apis

import (
	"net/http"

	"github.com/cheetahbyte/flagly/internal/flagly"
	"github.com/gin-gonic/gin"
)

type GeneralAPI struct {
	store *flagly.Storage
}

func NewGeneralAPI(store *flagly.Storage) *GeneralAPI {
	return &GeneralAPI{store: store}
}

func (api *GeneralAPI) RegisterRoutes(router *gin.Engine) {
	router.GET("/api/status", api.GetStatus)
	router.GET("/api/health", api.GetHealth)
}

func (g *GeneralAPI) GetStatus(c *gin.Context) {
	version := "dirty"
	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
}

func (g *GeneralAPI) GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{"ok": 1})
}
