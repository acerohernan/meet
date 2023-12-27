package config

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port   int           `yaml:"port"`
	Logger *LoggerConfig `yaml:"logger"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

func NewConfig(configString string) (*Config, error) {
	var conf Config

	decoder := yaml.NewDecoder(strings.NewReader(configString))

	decoder.KnownFields(true)

	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
