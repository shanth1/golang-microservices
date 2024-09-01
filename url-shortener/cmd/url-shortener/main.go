package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/shanth1/golang-microservices/url-shortener/internal/config"
)

const (
	envLocal   = "local"
	envDevelop = "develop"
	envProd    = "prod"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init jlogger
	logger := setupLogger(cfg.Env)
	logger.Info("starting url shortener", slog.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	// TODO: init storage

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
