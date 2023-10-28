package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `yaml:"port" mapstructure:"http_port"`
	Url   string `yaml:"url" mapstructure:"pg_url"`
	Level string `yaml:"level" mapstructure:"log_level"`
}

func NewConfig() (cfg *Config, err error) {
	viper.SetConfigFile(".env")

	cfg = &Config{}

	err = viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	err = viper.Unmarshal(&cfg, viper.DecodeHook(mapstructure.StringToSliceHookFunc(",")))

	if err != nil {
		return nil, fmt.Errorf("error marshalling config: %w", err)
	}

	return cfg, nil
}
