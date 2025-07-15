package apis

import (
	"net/http"

	"github.com/cheetahbyte/flagly/internal/version"
	"github.com/cheetahbyte/flagly/pkg/flagly"
	"github.com/gin-gonic/gin"
)

type GeneralAPI struct {
	store *flagly.Storage
}

func NewGeneralAPI(store *flagly.Storage) *GeneralAPI {
	return &GeneralAPI{store: store}
}

func (api *GeneralAPI) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/status", api.GetStatus)
	router.GET("/health", api.GetHealth)
}

func (g *GeneralAPI) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": version.Version,
	})
}

func (g *GeneralAPI) GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{"ok": 1})
}
