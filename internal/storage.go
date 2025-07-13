package internal

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Condition struct {
	Environments []string `json:"environments" yaml:"environments"`
}

type Flag struct {
	Key         string      `json:"key" yaml:"key"`
	Description string      `json:"description" yaml:"description"`
	Enabled     bool        `json:"enabled" yaml:"enabled"`
	Conditions  []Condition `json:"conditions" yaml:"conditions"`
}

type Storage struct {
	Flags        []Flag   `json:"flags" yaml:"flags"`
	Environments []string `json:"environments" yaml:"environments"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
