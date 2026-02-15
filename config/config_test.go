package config

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	var config Config
	err := LoadConfig(&config)
	if err != nil {
		t.Errorf("Error loading config: %s", err)
	}

	if config.App.LogLevel != "DEBUG" {
		t.Errorf("Log level is not DEBUG")
	}

	if config.HTTP.Port != 8080 {
		t.Errorf("Port is not 8080")
	}

	if config.HTTP.Host != "localhost" {
		t.Errorf("Host is not localhost")
	}

	if config.DB.Host != "localhost" {
		t.Errorf("Db host is not localhost")
	}

	if config.DB.Port != 5432 {
		t.Errorf("Db port is not 5432")
	}

	if config.DB.Username != "secret_santa_user" {
		t.Errorf("Db user is not postgres")
	}

	if config.DB.Password != "secret_santa_password" {
		t.Errorf("Db password is not secret_santa_password")
	}

	if config.DB.SSLMode != "disable" {
		t.Errorf("SSL mode is not disable")
	}
}

func TestOverrideDefaultConfig(t *testing.T) {
	// Create a temp file
	configFilePattern := "config_test_*.yaml"
	fileName, cleanup, err := createConfigfile(configFilePattern)
	if err != nil {
		t.Errorf("Error creating temp file: %s", err)
	}
	defer cleanup()

	os.Setenv("SECRET_SANTA_CONFIG_FILE", fileName)

	var config Config
	err = LoadConfig(&config)
	if err != nil {
		t.Errorf("Error loading config: %s", err)
	}

	fmt.Println(config.App.LogLevel)
	if config.App.LogLevel != "INFO" {
		t.Errorf("Expected log level to be INFO, got %s", config.App.LogLevel)
	}
}

func createConfigfile(filename string) (string, func(), error) {
	// Create a temp file
	f, err := os.CreateTemp("", filename)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	// Write to the file
	_, err = f.WriteString("app:\n  log_level: INFO")
	if err != nil {
		return "", nil, err
	}

	name := f.Name()
	// return a function to clean up the file
	return name, func() { os.Remove(name) }, nil
}
