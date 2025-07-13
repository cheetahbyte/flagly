package flagly

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetLogger(c *gin.Context) *zap.SugaredLogger {
	return c.MustGet("logger").(*zap.SugaredLogger)
}
