package logger

import (
	"context"
	"log/slog"
	"os"
)

var (
	Logger  *slog.Logger
)

func Init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}


func WithContext(ctx context.Context) *slog.Logger {
    return Logger.With("context", ctx)
}
