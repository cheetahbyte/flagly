package apis

import (
	"net/http"

	custom_errors "github.com/cheetahbyte/flagly/internal/error"
	"github.com/cheetahbyte/flagly/pkg/flagly"
	"github.com/gin-gonic/gin"
)

type EnvironmentAPI struct {
	store *flagly.Storage
}

func NewEnvironmentAPI(store *flagly.Storage) *EnvironmentAPI {
	return &EnvironmentAPI{store: store}
}

func (api *EnvironmentAPI) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/environments", api.GetAllEnvironments)
	router.GET("/environments/:env", api.GetEnvironment)
}

func (api *EnvironmentAPI) GetAllEnvironments(c *gin.Context) {
	c.JSON(200, api.store.Environments)
}

func (api *EnvironmentAPI) GetEnvironment(c *gin.Context) {
	env_name := c.Param("env")
	environments := api.store.Environments

	for _, env := range environments {
		if env == env_name {
			c.JSON(200, gin.H{"name": env})
			return
		}
	}
	c.Error(custom_errors.NewAPIError(http.StatusNotFound,
		"/errors/environment-not-found",
		"Environment not found",
		"The requested environment was not found on the server."))
}
