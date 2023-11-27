package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Name  string `yaml:"name" mapstructure:"name"`
	Port  string `yaml:"port" mapstructure:"http_port"`
	Url   string `yaml:"url" mapstructure:"pg_url"`
	Level string `yaml:"level" mapstructure:"log_level"`

	RequestTopic string   `yaml:"request_topic" mapstructure:"request_topic"`
	Brokers      []string `yaml:"brokers" mapstructure:"brokers"`
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
