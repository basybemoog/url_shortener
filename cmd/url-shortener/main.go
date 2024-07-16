package main

import (
	"fmt"
	"log/slog"
	"os"
	config2 "urlshortner/internal/config"
	"urlshortner/internal/lib/logger/sl"
	"urlshortner/internal/storage/sqlite"
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

	storage, err := sqlite.New(config.StoragePath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		return
	}

	_ = storage

	storage.SaveURL("https://github.com/", "github")
	//TODO: init storage: sqlite3

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
