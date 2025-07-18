package evaluation

import (
	"strconv"
	"testing"

	"github.com/cheetahbyte/flagly/pkg/flagly"
)

func TestEvaluateFlag_UnknownEnvironment(t *testing.T) {
	service := &DefaultEvaluationService{}
	flag := flagly.Flag{
		Environments: map[string]flagly.Environment{},
	}
	user := flagly.User{ID: "123"}

	result, err := service.EvaluateFlag(flag, user, "production")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result {
		t.Errorf("expected false for unknown environment, got true")
	}
}

func TestEvaluateFlag_PercentageZero(t *testing.T) {
	service := &DefaultEvaluationService{}
	flag := flagly.Flag{
		Environments: map[string]flagly.Environment{
			"production": {
				Enabled: true,
				Rollout: flagly.Rollout{Percentage: 0, Stickiness: "user_id"},
			},
		},
	}
	user := flagly.User{ID: "user1"}

	result, err := service.EvaluateFlag(flag, user, "production")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result {
		t.Errorf("expected true when rollout is 0%%, got false")
	}
}

func TestEvaluateFlag_Percentage100(t *testing.T) {
	service := &DefaultEvaluationService{}
	flag := flagly.Flag{
		Environments: map[string]flagly.Environment{
			"production": {
				Enabled: true,
				Rollout: flagly.Rollout{Percentage: 100, Stickiness: "user_id"},
			},
		},
	}
	user := flagly.User{ID: "user1"}

	result, err := service.EvaluateFlag(flag, user, "production")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result {
		t.Errorf("expected true for 100%% rollout, got false")
	}
}

func TestEvaluateFlag_ConsistentRollout(t *testing.T) {
	service := &DefaultEvaluationService{}
	user := flagly.User{ID: "user42"}
	flag := flagly.Flag{
		Environments: map[string]flagly.Environment{
			"production": {
				Enabled: true,
				Rollout: flagly.Rollout{Percentage: 50, Stickiness: "user_id"},
			},
		},
	}

	result1, err1 := service.EvaluateFlag(flag, user, "production")
	result2, err2 := service.EvaluateFlag(flag, user, "production")

	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected error(s): %v, %v", err1, err2)
	}
	if result1 != result2 {
		t.Errorf("expected consistent result for same user, got %v and %v", result1, result2)
	}
}

func TestEvaluateFlag_WithinRollout(t *testing.T) {
	service := &DefaultEvaluationService{}

	// Find a user ID that hashes to a bucket < 10
	var userID string
	for i := 0; i < 10000; i++ {
		id := "user" + strconv.Itoa(i)
		if calculateRolloutBucket(id) < 10 {
			userID = id
			break
		}
	}
	if userID == "" {
		t.Fatalf("could not find user ID with bucket < 10")
	}

	user := flagly.User{ID: userID}
	flag := flagly.Flag{
		Environments: map[string]flagly.Environment{
			"production": {
				Enabled: true,
				Rollout: flagly.Rollout{Percentage: 10, Stickiness: "user_id"},
			},
		},
	}

	result, err := service.EvaluateFlag(flag, user, "production")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result {
		t.Errorf("expected true for user within rollout range")
	}
}
