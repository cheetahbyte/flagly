package internal

import (
	"slices"
)

func CheckEnvironment(environment string, flagConditions []Condition) bool {
	// TODO: change the way Conditions are parsed from yaml
	return slices.Contains(flagConditions[0].Environments, environment)
}
