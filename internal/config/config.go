package config

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/spf13/viper"
)

//go:embed default_config.yaml
var defaultConfig []byte

// Root configuration for the main application.
type Config struct {
	App AppConfig `mapstructure:"app"`
}

func (p Config) Defaults() io.Reader {
	return bytes.NewReader(defaultConfig)
}

type AppConfig struct {
	LogLevel string `mapstructure:"log_level"`
}

func LoadConfig(config *Config) error {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(config.Defaults()); err != nil {
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}
