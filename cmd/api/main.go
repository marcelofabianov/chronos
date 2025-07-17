package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/marcelofabianov/chronos/internal/app"
	"github.com/marcelofabianov/chronos/internal/platform/config"
	"github.com/marcelofabianov/chronos/internal/platform/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("application startup failed: %v", err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	lg := logger.NewSlogLogger(cfg.Logger)

	app, err := app.New(cfg, lg)
	if err != nil {
		return err
	}

	if err := app.Run(); err != nil {
		return err
	}

	return nil
}
