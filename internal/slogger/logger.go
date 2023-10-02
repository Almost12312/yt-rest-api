package slogger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	var slogger *slog.Logger

	switch env {
	case "local":
		slogger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		slogger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	default:
		slogger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	slogger.Info("starting api server!")

	return slogger
}
