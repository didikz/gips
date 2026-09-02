// Package config loads application configuration from the environment.
package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

// Config holds runtime settings for the image processing service.
type Config struct {
	Addr     string // listen address, e.g. ":8080"
	LogLevel slog.Level
}

// Load reads configuration from environment variables with sensible defaults.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}

	return Config{
		Addr:     ":" + port,
		LogLevel: parseLogLevel(os.Getenv("LOG_LEVEL")),
	}
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
