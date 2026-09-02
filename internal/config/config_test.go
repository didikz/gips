package config_test

import (
	"log/slog"
	"testing"

	"github.com/didikz/gips/internal/config"
)

func TestConfigLoadEmptyEnv(t *testing.T) {
	config := config.Load()
	addr := config.Addr
	if addr != ":8080" {
		t.Errorf(`config address = %q, want %q`, addr, ":8080")
	}
}

func TestConfigLoadAddrEnv(t *testing.T) {
	t.Setenv("PORT", "8000")
	config := config.Load()
	addr := config.Addr
	if addr != ":8000" {
		t.Errorf(`config address = %q, want %q`, addr, ":8000")
	}
}

func TestConfigLoadLevel(t *testing.T) {
	tests := []struct {
		env  string
		want slog.Level
	}{
		{"info", slog.LevelInfo},
		{"debug", slog.LevelDebug},
		{"WARN", slog.LevelWarn},
		{"error", slog.LevelError},
		{"none", slog.LevelInfo},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			t.Setenv("LOG_LEVEL", tt.env)
			got := config.Load().LogLevel
			if got != tt.want {
				t.Errorf("parse log level got %v, want %v", got, tt.want)
			}
		})
	}
}
