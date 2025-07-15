package flagly

import "github.com/gin-gonic/gin"

type AuditService interface {
	TrackEvaluation(c *gin.Context, flag Flag, user User, environment string, result bool)
}

type DefaultAuditService struct{}

func NewDefaultAuditService() AuditService {
	return &DefaultAuditService{}
}

func (s *DefaultAuditService) TrackEvaluation(c *gin.Context, flag Flag, user User, environment string, result bool) {
	logger := GetLogger(c)

	logger.Infow("Flag evaluation completed",
		"flag_key", flag.Key,
		"user_id", user.ID,
		"environment", environment,
		"result", result,
	)
}
