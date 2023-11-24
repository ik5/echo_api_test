package main

import (
	"log/slog"
	"os"
)

func initLogger() *slog.Logger {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,            // Show me the source
				Level:     slog.LevelDebug, // let's see debug...
			},
		),
	)

	return logger
}
