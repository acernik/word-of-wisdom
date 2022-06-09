package config

import (
	"os"

	"github.com/jinzhu/configor"
)

// Config holds the configuration.
type Config struct {
	App *App
}

// New returns the configuration.
func New() (cnf *Config, e error) {
	configFile := "config.yml"

	if len(os.Getenv("TEST_CONFIG")) > 0 {
		configFile = os.Getenv("TEST_CONFIG")
	}

	var cfg Config

	if err := configor.Load(&cfg, configFile); err != nil {
		return nil, err
	}

	return &cfg, nil
}
