package config_test

import (
	"os"
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
	err := os.Setenv("PORT", "8000")
	if err != nil {
		t.Error("error setting env", err)
	}

	config := config.Load()
	addr := config.Addr
	if addr != ":8000" {
		t.Errorf(`config address = %q, want %q`, addr, ":8000")
	}
}
