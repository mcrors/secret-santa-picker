package config

import (
	"io/ioutil"
	"log"
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

	if config.Http.Port != 8080 {
		t.Errorf("Port is not 8080")
	}

	if config.Http.Host != "localhost" {
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
	cleanup := createConfigfile()
	defer cleanup()

	var config Config
	err := LoadConfig(&config)
	if err != nil {
		t.Errorf("Error loading config: %s", err)
	}
}

func createConfigfile() func() {
	// Create a temp file
	f, err := ioutil.TempFile("", "config")
	if err != nil {
		log.Fatal(err)
	}

	// return a function to clean up the file
	return func() {
		os.Remove(f.Name())
	}
}
