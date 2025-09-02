package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

var Logger *slog.Logger

func init() {
	// カレントディレクトリの1階層上をプロジェクトルートと仮定する場合
	// （実行ファイルが cmd/ 下にある場合など）
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// プロジェクト直下に app.log
	logPath := filepath.Join(cwd, "app.log")

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	Logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		AddSource: true,
	}))
}
