package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

//go:embed default_config.yaml
var defaultConfig []byte

const (
	EnvPrefix = "SECRET_SANTA"
)

// Root configuration for the main application.
type Config struct {
	App  App      `mapstructure:"app"`
	Http HTTP     `mapstructure:"http"`
	DB   Database `mapstructure:"database"`
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

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int32  `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

func LoadConfig(config *Config) error {
	if err := loadDefaultConfig(config, config.Defaults()); err != nil {
		return fmt.Errorf("failed to load default config: %w", err)
	}

	viper.SetEnvPrefix(EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// This requires an environment variable called APP_CONFIG_PATH. This will point to
	// another config file which will override the default config values
	configPath := viper.GetString("CONFIG_PATH")
	fmt.Println("Config Path:", configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	// If the environment variable is not set, then no other file will be loaded
	if configPath != "" {
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("unable to read config, %w", err)
		}
		if err := viper.Unmarshal(config); err != nil {
			return fmt.Errorf("unable to decode config into struct, %w", err)
		}
	}

	return nil
}

func loadDefaultConfig(config interface{}, defaults io.Reader) error {
	cfgMap := make(map[string]interface{})
	if err := mapstructure.Decode(config, &cfgMap); err != nil {
		return err
	}

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(defaults); err != nil {
		return err
	}
	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}
