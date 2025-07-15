package audit

import (
	"github.com/cheetahbyte/flagly/internal/utils"
	"github.com/cheetahbyte/flagly/pkg/flagly"
	"github.com/gin-gonic/gin"
)

type AuditService interface {
	TrackEvaluation(c *gin.Context, flag flagly.Flag, user flagly.User, environment string, result bool)
}

type DefaultAuditService struct{}

func NewDefaultAuditService() AuditService {
	return &DefaultAuditService{}
}

func (s *DefaultAuditService) TrackEvaluation(c *gin.Context, flag flagly.Flag, user flagly.User, environment string, result bool) {
	logger := utils.GetLogger(c)

	logger.Infow("Flag evaluation completed",
		"flag_key", flag.Key,
		"user_id", user.ID,
		"environment", environment,
		"result", result,
	)
}
