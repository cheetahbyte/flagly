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
	logger.Info("Attempting to fetch all flags")

	flags := flagly.Store.Flags
	logger.Infow("Successfully fetched all flags",
		"count", len(flags),
	)
	c.JSON(200, flags)
}

func GetFlag(c *gin.Context) {
	logger := flagly.GetLogger(c)
	flagKey := c.Param("flag")

	logger.Infow("Attempting to fetch a single flag",
		"flag_key", flagKey,
	)

	var selectedFlag *flagly.Flag
	for _, f := range flagly.Store.Flags {
		if f.Key == flagKey {
			selectedFlag = &f
			break
		}
	}

	if selectedFlag == nil {
		msg := fmt.Sprintf("flag '%s' not found", flagKey)
		logger.Warnw(msg,
			"flag_key", flagKey,
		)
		c.Error(errors.New("flag not found"))
		return
	}

	logger.Infow("Successfully found flag",
		"flag_key", flagKey,
	)
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
		logger.Errorw("Failed to bind request data",
			"error", err.Error(),
		)
		c.Error(err)
		return
	}

	logger.Infow("Attempting to evaluate a flag",
		"flag_key", data.Flag,
		"user_id", data.User.ID,
		"environment", data.Environment,
	)

	var flag *flagly.Flag
	for _, f := range flagly.Store.Flags {
		if f.Key == data.Flag {
			flag = &f
			break
		}
	}

	if flag == nil {
		logger.Warnw("Evaluation failed because flag was not found",
			"flag_key", data.Flag,
		)
		c.Error(errors.New("flag not found"))
		return
	}

	result := flagly.EvaluateFlag(*flag, data.User, data.Environment)

	logger.Infow("Flag evaluation completed",
		"flag_key", flag.Key,
		"user_id", data.User.ID,
		"environment", data.Environment,
		"result", result,
	)

	c.JSON(200, gin.H{"enabled": result})
}
