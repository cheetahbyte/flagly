package apis

import (
	"errors"

	"github.com/cheetahbyte/flagly/internal/storage"
	"github.com/gin-gonic/gin"
)

func GetAllEnvironments(c *gin.Context) {
	c.JSON(200, storage.Store.Environments)
}

func GetEnvironment(c *gin.Context) {
	env_name := c.Param("env")
	environments := storage.Store.Environments

	for _, env := range environments {
		if env == env_name {
			c.JSON(200, gin.H{"name": env})
			return
		}
	}
	c.Error(errors.New("environment not found"))
}
