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
	App  App  `mapstructure:"app"`
	Http HTTP `mapstructure:"http"`
}

func (p Config) Defaults() io.Reader {
	return bytes.NewReader(defaultConfig)
}

type App struct {
	LogLevel string `mapstructure:"log_level"`
}

type HTTP struct {
	Port int32  `mapstructure:"port"`
	Host string `mapstructure:"host"`
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
