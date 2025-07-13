package flagly

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Environment struct {
	Enabled bool
	Rollout Rollout
}

type Flag struct {
	Key          string                 `json:"key" yaml:"key"`
	Description  string                 `json:"description" yaml:"description"`
	Environments map[string]Environment `json:"environments" yaml:"environments"`
}

type Storage struct {
	Flags        []Flag   `json:"flags" yaml:"flags"`
	Environments []string `json:"environments" yaml:"environments"`
}

var Store *Storage

func InitStorage(configFile string) error {
	cfg := &Storage{}
	dat, err := os.ReadFile("./flagly.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(dat), cfg)
	if err != nil {
		return err
	}
	Store = cfg
	return nil
}
