package flagly

type Rollout struct {
	Percentage int `json:"percentage" yaml:"percentage"`
	// Start      time.Time `json:"date" yaml:"date"`
	Stickiness string `json:"stickyness" yaml:"stickyness"`
}
