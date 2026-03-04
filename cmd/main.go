package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/mcrors/secret-santa-picker-server/config"
	"github.com/mcrors/secret-santa-picker-server/database"
	"github.com/mcrors/secret-santa-picker-server/repository"
	"github.com/mcrors/secret-santa-picker-server/server"
)

func main() {
	var cfg config.Config
	if err := config.LoadConfig(&cfg); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	setupLogger(cfg)

	// Create db object
	db, err := database.GetPostgresDB(cfg.DB)
	if err != nil {
		log.Fatalf("error creating postgres db: %v", err)
	}
	defer db.Close()

	// Create repository objects
	_ = repository.NewGroupRepository(db)

	// Create services

	// Create handlers

	// Create Server
	s, err := server.NewServer(
		cfg.HTTP,
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
