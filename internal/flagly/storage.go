package flagly

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Environment struct {
	Enabled bool    `json:"enabled"`
	Rollout Rollout `json:"rollout"`
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

func InitStorage(configFile string) (*Storage, error) {
	cfg := &Storage{}
	dat, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(dat), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
