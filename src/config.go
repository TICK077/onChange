package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	WatchDir string   `yaml:"watch_dir"`
	WorkDir  string   `yaml:"work_dir"`
	DelaySec int      `yaml:"delay_sec"`
	Command  []string `yaml:"command"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
