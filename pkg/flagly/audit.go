package flagly

import "github.com/gin-gonic/gin"

type AuditService interface {
	TrackEvaluation(c *gin.Context, flag Flag, user User, environment string, result bool)
}
