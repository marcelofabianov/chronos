package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/dig"

	v1 "github.com/marcelofabianov/chronos/internal/handler/v1"
	"github.com/marcelofabianov/chronos/internal/platform/config"
	"github.com/marcelofabianov/chronos/internal/platform/web"
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

func (a *App) Run() error {
	router := web.NewRouter(a.config, a.logger)

	router.Get("/", DefaultHandler)

	router.Mount("/api/v1", v1.LoadRoutes())

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		a.logger.Info("server is starting", "address", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	<-stopChan

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.logger.Info("shutting down server gracefully")

	return nil
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": ".",
		"status":  "ok",
	}

	jsonData, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(jsonData)
}
