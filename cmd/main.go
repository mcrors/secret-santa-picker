// main is the entry point into the application
package main

import (
	_ "embed"
	"log"
	"log/slog"
	"os"

	"github.com/mcrors/secret-santa-picker-server/internal/config"
	domain "github.com/mcrors/secret-santa-picker-server/internal/domain"
)

func main() {
	var cfg config.Config
	if err := config.LoadConfig(&cfg); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	setupLogger(cfg)

	// Create db object
	// Create repository object
	// Create services
	// Create handlers
	// Create Server

	u := domain.User{
		FirstName: "Rory",
		LastName:  "Houlihan",
		Email:     "rory@houli.eu",
	}

	slog.Debug("this is a debug message")
	slog.Info("current user", slog.Any("user", u))
}

func setupLogger(cfg config.Config) {
	var logLevel slog.Level
	switch cfg.App.LogLevel {
	case "INFO":
		logLevel = slog.LevelInfo
	case "DEBUG":
		logLevel = slog.LevelDebug
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}
