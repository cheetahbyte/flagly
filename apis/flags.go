package apis

import (
	"errors"
	"fmt"

	"github.com/cheetahbyte/flagly/internal/flagly"
	"github.com/gin-gonic/gin"
)

type FlaglyContext struct {
	Environment string `json:"environment" binding:"required"`
}

func GetAllFlags(c *gin.Context) {
	logger := flagly.GetLogger(c)
	logger.Info("Fetching all flags")
	c.JSON(200, flagly.Store.Flags)
}

func GetFlag(c *gin.Context) {
	logger := flagly.GetLogger(c)
	flag_key := c.Param("flag")
	logger.Infow("Fetching a single flag",
		"flag_key", flag_key,
	)
	var selectedFlag *flagly.Flag
	for _, f := range flagly.Store.Flags {
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

type PostEvaluateFlagDTO struct {
	Flag        string      `json:"flag"`
	User        flagly.User `json:"user"`
	Environment string      `json:"environment"`
}

func PostEvaluateFlag(c *gin.Context) {
	logger := flagly.GetLogger(c)
	var data PostEvaluateFlagDTO
	if err := c.ShouldBind(&data); err != nil {
		logger.Info(err.Error())
		c.Error(err)
		return
	}
	var flag *flagly.Flag
	for _, f := range flagly.Store.Flags {
		logger.Info(f)
		if f.Key == data.Flag {
			flag = &f
		}
	}
	if flag == nil {
		c.Error(errors.New(("flag not found")))
		return
	}

	result := flagly.EvaluateFlag(*flag, data.User, data.Environment)

	if !result {
		c.JSON(200, gin.H{"enabled": false})
		return
	}
	c.JSON(200, gin.H{"enabled": true})

}
