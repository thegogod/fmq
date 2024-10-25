package logger

import (
	"log"
	"log/slog"
	"os"
	"strings"
)

func New(name string) *slog.Logger {
	prefix := os.Getenv("LOG_PREFIX")

	if prefix != "" {
		prefix = prefix + "/"
	}

	lvl := slog.LevelVar{}
	lvl.Set(Debug.SLog())

	if level := Level(strings.ToLower(os.Getenv("LOG_LEVEL"))); level.Valid() {
		lvl.Set(level.SLog())
	}

	return slog.New(NewColorTextHandler(&slog.HandlerOptions{
		Level:     &lvl,
		AddSource: lvl.Level() == slog.LevelDebug,
	})).With("name", prefix+name)
}

func NewLog(name string, level slog.Level) *log.Logger {
	prefix := os.Getenv("LOG_PREFIX")

	if prefix != "" {
		prefix = prefix + "/"
	}

	lvl := slog.LevelVar{}
	lvl.Set(Debug.SLog())

	if level := Level(strings.ToLower(os.Getenv("LOG_LEVEL"))); level.Valid() {
		lvl.Set(level.SLog())
	}

	return slog.NewLogLogger(NewColorTextHandler(&slog.HandlerOptions{
		Level:     &lvl,
		AddSource: lvl.Level() == slog.LevelDebug,
	}).WithAttrs([]slog.Attr{slog.String("name", prefix+name)}), level)
}
