package flagly

type EvaluationService interface {
	EvaluateFlag(flag Flag, user User, environment string) (bool, error)
}
