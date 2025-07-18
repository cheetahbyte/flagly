package evaluation

import (
	"fmt"

	"github.com/cheetahbyte/flagly/pkg/flagly"
	"github.com/spaolacci/murmur3"
)

type DefaultEvaluationService struct{}

func NewDefaultAuditService() flagly.EvaluationService {
	return &DefaultEvaluationService{}
}

func calculateRolloutBucket(identifier string) int {
	hash := murmur3.Sum32([]byte(identifier))
	return int(hash % 100)
}

func (s *DefaultEvaluationService) EvaluateFlag(flag flagly.Flag, user flagly.User, environment string) (bool, error) {
	env, ok := flag.Environments[environment]
	if !ok {
		return false, nil
	}
	fmt.Println(env)
	if env.Rollout.Percentage == 0 || env.Rollout.Percentage == 100 {
		return true, nil
	}

	// TODO: make this more dynamic
	var stickiness string
	switch env.Rollout.Stickiness {
	case "user_id":
		stickiness = user.ID
	default:
		stickiness = user.ID
	}

	// Rollout specific logic
	hashedPercentage := calculateRolloutBucket(stickiness)
	return env.Rollout.Percentage > hashedPercentage, nil
}
