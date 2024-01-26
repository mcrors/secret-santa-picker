package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/mcrors/secret-santa-picker-server/config"
	"github.com/mcrors/secret-santa-picker-server/handler"
	"github.com/mcrors/secret-santa-picker-server/repository/postgres"
	"github.com/mcrors/secret-santa-picker-server/server"
	"github.com/mcrors/secret-santa-picker-server/service"
)

func main() {
	var cfg config.Config
	if err := config.LoadConfig(&cfg); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	setupLogger(cfg)

	// Create db object

	// Create repository object
	userRepo := postgres.NewUser()
	// Create services
	userService := service.NewUser(userRepo)
	// Create handlers
	regHandler := handler.NewRegister(userService)
	// Create Server

	s, err := server.NewServer(
		cfg.Http,
		regHandler,
	)
	if err != nil {
		log.Fatalf("error creating server: %v", err)
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
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
