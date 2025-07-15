package flagly

type Rollout struct {
	Percentage int `json:"percentage" yaml:"percentage"`
	// Start      time.Time `json:"date" yaml:"date"`
	Stickiness string `json:"stickyness" yaml:"stickyness"`
}

type Environment struct {
	Enabled bool    `json:"enabled"`
	Rollout Rollout `json:"rollout"`
}

type Flag struct {
	Key          string                 `json:"key" yaml:"key"`
	Description  string                 `json:"description" yaml:"description"`
	Environments map[string]Environment `json:"environments" yaml:"environments"`
}

type User struct {
	ID string
}
