package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/shanth1/golang-microservices/url-shortener/internal/config"
	"github.com/shanth1/golang-microservices/url-shortener/internal/lib/slogger"
	"github.com/shanth1/golang-microservices/url-shortener/internal/storage/sqlite"
)

const (
	envLocal   = "local"
	envDevelop = "develop"
	envProd    = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	logger := setupLogger(cfg.Env)
	logger.Info("starting url shortener", slog.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		logger.Error("init storage error", slogger.Error(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init router

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envDevelop:
		logger = slog.New(slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
	}

	return logger
}
