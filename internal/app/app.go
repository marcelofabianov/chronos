package app

import (
	"log/slog"

	"go.uber.org/dig"

	"github.com/marcelofabianov/chronos/internal/platform/config"
)

type App struct {
	container *dig.Container
	config    *config.AppConfig
	logger    *slog.Logger
}

func New(cfg *config.AppConfig, logger *slog.Logger) (*App, error) {
	container := dig.New()

	app := &App{
		container: container,
		config:    cfg,
		logger:    logger,
	}

	app.logger.Info("application container built successfully")

	return app, nil
}
