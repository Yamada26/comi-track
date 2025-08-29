package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// Init initializes the global logger
func init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
