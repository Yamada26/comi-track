package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// Init initializes the global logger
func init() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	Logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		AddSource: true,
	}))
}
