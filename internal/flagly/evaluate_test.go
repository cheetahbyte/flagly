package flagly

import (
	"strconv"
	"testing"
)

func TestEvaluateFlag_UnknownEnvironment(t *testing.T) {
	flag := Flag{
		Environments: map[string]Environment{},
	}
	user := User{ID: "123"}

	result := EvaluateFlag(flag, user, "production")
	if result {
		t.Errorf("expected false for unknown environment, got true")
	}
}

func TestEvaluateFlag_PercentageZero(t *testing.T) {
	flag := Flag{
		Environments: map[string]Environment{
			"production": {
				Enabled: true,
				Rollout: Rollout{Percentage: 0, Stickiness: "user_id"},
			},
		},
	}
	user := User{ID: "user1"}

	if !EvaluateFlag(flag, user, "production") {
		t.Errorf("expected true when rollout is 0%% or 100%%, got false")
	}
}

func TestEvaluateFlag_Percentage100(t *testing.T) {
	flag := Flag{
		Environments: map[string]Environment{
			"production": {
				Enabled: true,
				Rollout: Rollout{Percentage: 100, Stickiness: "user_id"},
			},
		},
	}
	user := User{ID: "user1"}

	if !EvaluateFlag(flag, user, "production") {
		t.Errorf("expected true for 100%% rollout")
	}
}

func TestEvaluateFlag_ConsistentRollout(t *testing.T) {
	user := User{ID: "user42"}
	flag := Flag{
		Environments: map[string]Environment{
			"production": {
				Enabled: true,
				Rollout: Rollout{Percentage: 50, Stickiness: "user_id"},
			},
		},
	}

	// Should always return the same result for the same user
	result1 := EvaluateFlag(flag, user, "production")
	result2 := EvaluateFlag(flag, user, "production")

	if result1 != result2 {
		t.Errorf("expected consistent result for same user, got %v and %v", result1, result2)
	}
}

func TestEvaluateFlag_WithinRollout(t *testing.T) {
	// Find a user ID that gives a bucket under 10
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

	user := User{ID: userID}
	flag := Flag{
		Environments: map[string]Environment{
			"production": {
				Enabled: true,
				Rollout: Rollout{Percentage: 10, Stickiness: "user_id"},
			},
		},
	}

	if !EvaluateFlag(flag, user, "production") {
		t.Errorf("expected true for user within rollout range")
	}
}
