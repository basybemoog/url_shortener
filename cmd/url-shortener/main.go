package main

import (
	"fmt"
	"log/slog"
	"os"
	config2 "urlshortner/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config2.MustLoad()
	fmt.Println(config)
	log := setupLogger(config.Env)

	log.Info("starting url-shortener", slog.String("env", config.Env))
	log.Debug("debug messages enabled")

	//TODO: init router: chi, render

	//TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
