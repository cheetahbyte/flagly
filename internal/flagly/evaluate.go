package flagly

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

type User struct {
	ID string
}

// calculates the rollout percentage bucket in which the user is located.
func calculateRolloutBucket(identifier string) int {
	hash := murmur3.Sum32([]byte(identifier))
	return int(hash % 100)
}

func EvaluateFlag(flag Flag, user User, environment string) bool {
	// check if environment is enabled
	env, ok := flag.Environments[environment]
	if !ok {
		return false
	}
	fmt.Println(env)
	if env.Rollout.Percentage == 0 || env.Rollout.Percentage == 100 {
		return true
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
	return env.Rollout.Percentage > hashedPercentage
}
