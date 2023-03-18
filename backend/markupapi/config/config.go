package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HTTP struct {
		Address         string        `yaml:"address"`
		GracefulTimeout time.Duration `yaml:"gracefull_timeout"`
	} `yaml:"http"`
}

func New(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file content: %w", err)
	}

	c := Config{}
	if err = yaml.UnmarshalStrict(data, &c); err != nil {
		return Config{}, fmt.Errorf("failed to parse file content: %w", err)
	}

	return c, nil
}
