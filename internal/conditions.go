package internal

import (
	"slices"

	"github.com/cheetahbyte/flagly/internal/storage"
)

func CheckEnvironment(environment string, flagConditions []storage.Condition) bool {
	// TODO: change the way Conditions are parsed from yaml
	return slices.Contains(flagConditions[0].Environments, environment)
}
