package apis

import (
	"errors"
	"fmt"

	"github.com/cheetahbyte/flagly/internal"
	"github.com/gin-gonic/gin"
)

type FlaglyContext struct {
	Environment string `json:"environment" binding:"required"`
}

func GetAllFlags(c *gin.Context) {
	logger := internal.GetLogger(c)
	logger.Info("Fetching all flags")
	c.JSON(200, internal.Store.Flags)
}

func GetFlag(c *gin.Context) {
	logger := internal.GetLogger(c)
	flag_key := c.Param("flag")
	logger.Infow("Fetching a single flag",
		"flag_key", flag_key,
	)
	var selectedFlag *internal.Flag
	for _, f := range internal.Store.Flags {
		if f.Key == flag_key {
			selectedFlag = &f
			break
		}
	}
	if selectedFlag == nil {
		msg := fmt.Sprintf("flag '%s' not found", flag_key)
		logger.Warn(msg)
		c.Error(errors.New("flag not found"))
		return
	}
	c.JSON(200, selectedFlag)
}

func GetFlagEnabled(c *gin.Context) {
	logger := internal.GetLogger(c)
	flag_key := c.Param("flag")
	environment := c.Query("environment")
	if environment == "" {
		c.Error(errors.New(("no environment provided")))
		return
	}

	var flag *internal.Flag
	for _, f := range internal.Store.Flags {
		logger.Info(f)
		if f.Key == flag_key {
			flag = &f
		}
	}
	if flag == nil {
		c.Error(errors.New(("flag not found")))
		return
	}

	if !internal.CheckEnvironment(environment, flag.Conditions) {
		c.JSON(200, gin.H{"enabled": false})
		return
	}

	c.JSON(200, gin.H{
		"enabled": flag.Enabled,
	})
}
