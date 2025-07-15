package flagly

type Storage struct {
	Flags        []Flag   `json:"flags" yaml:"flags"`
	Environments []string `json:"environments" yaml:"environments"`
}
