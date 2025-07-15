package storage

import (
	"os"

	"github.com/cheetahbyte/flagly/pkg/flagly"
	"github.com/goccy/go-yaml"
)

func InitStorage(configFile string) (*flagly.Storage, error) {
	cfg := &flagly.Storage{}
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
