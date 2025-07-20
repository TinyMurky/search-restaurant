// Package config contains config.yaml related feature
package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/TinyMurky/search-restaurant/internal/pkg/customerrors"
)

// LoadConfig read config.yaml from "path" and return Config
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, customerrors.NewIOError("open config file", err)
	}

	defer file.Close()

	return loadConfigByReader(file)
}

// loadConfigByReader read config by io.Reader
func loadConfigByReader(r io.Reader) (*Config, error) {
	cfg := Config{}

	decoder := yaml.NewDecoder(r)

	if err := decoder.Decode(&cfg); err != nil {
		return nil, customerrors.NewIOError("decode config", err)
	}

	return &cfg, nil
}
