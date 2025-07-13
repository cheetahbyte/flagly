package flagly

type Target struct {
	Environments []string `json:"environments" yaml:"environments"`
	// Roles        []string `json:"roles" yaml:"roles"`
	// Users        []string `json:"users" yaml:"users"`
}
