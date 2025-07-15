package storage

import (
	"os"

	"github.com/cheetahbyte/flagly/pkg/flagly"
	"github.com/goccy/go-yaml"
)

type Storage struct {
	Flags        []flagly.Flag `json:"flags" yaml:"flags"`
	Environments []string      `json:"environments" yaml:"environments"`
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
