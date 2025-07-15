package apis

import (
	"fmt"
	"net/http"

	custom_errors "github.com/cheetahbyte/flagly/internal/error"
	"github.com/cheetahbyte/flagly/internal/evaluation"
	"github.com/cheetahbyte/flagly/internal/utils"
	"github.com/cheetahbyte/flagly/pkg/flagly"
	"github.com/gin-gonic/gin"
)

type FlagAPI struct {
	store        *flagly.Storage
	auditService flagly.AuditService
}

func NewFlagAPI(store *flagly.Storage, auditService flagly.AuditService) *FlagAPI {
	return &FlagAPI{store: store, auditService: auditService}
}

func (api *FlagAPI) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/flags", api.GetFlags)
	router.GET("/flags/:flag", api.GetFlag)
	router.POST("/flags/evaluate", api.PostEvaluateFlag)
}

func (api *FlagAPI) GetFlags(c *gin.Context) {
	c.JSON(http.StatusOK, api.store.Flags)
}

func (api *FlagAPI) GetFlag(c *gin.Context) {
	logger := utils.GetLogger(c)
	flagKey := c.Param("flag")

	logger.Infow("Attempting to fetch a single flag",
		"flag_key", flagKey,
	)

	var selectedFlag *flagly.Flag
	for _, f := range api.store.Flags {
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
		c.Error(custom_errors.NewAPIError(http.StatusNotFound,
			"/errors/flag-not-found",
			"Flag not found",
			"The requested flag was not found on the server."))
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

func (api *FlagAPI) PostEvaluateFlag(c *gin.Context) {
	logger := utils.GetLogger(c)
	var data PostEvaluateFlagDTO

	if err := c.ShouldBind(&data); err != nil {
		logger.Errorw("Failed to bind request data",
			"error", err.Error(),
		)
		c.Error(err)
		return
	}

	var flag *flagly.Flag
	for _, f := range api.store.Flags {
		if f.Key == data.Flag {
			flag = &f
			break
		}
	}

	if flag == nil {
		logger.Warnw("Evaluation failed because flag was not found",
			"flag_key", data.Flag,
		)
		c.Error(custom_errors.NewAPIError(http.StatusNotFound,
			"/errors/flag-not-found",
			"Flag not found",
			"The requested flag was not found on the server."))
		return
	}

	result := evaluation.EvaluateFlag(*flag, data.User, data.Environment)

	api.auditService.TrackEvaluation(c, *flag, data.User, data.Environment, result)

	c.JSON(200, gin.H{"enabled": result})
}
