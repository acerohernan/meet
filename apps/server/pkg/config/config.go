package config

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port   int           `yaml:"port"`
	Logger *LoggerConfig `yaml:"logger"`
	JWT    *JWTConfig    `yaml:"jwt"`
	Router *RouterConfig `yaml:"router"`
	Redis  *RedisConfig  `yaml:"redis"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
}
type RouterConfig struct {
	Region string `yaml:"region"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	UseTLS   bool   `yaml:"use_tls"`
	DB       int    `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
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
