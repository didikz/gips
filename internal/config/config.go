// Package config loads application configuration from the environment.
package config

import (
	"os"
	"strconv"
)

// Config holds runtime settings for the image processing service.
type Config struct {
	Addr string // listen address, e.g. ":8080"
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
	return Config{Addr: ":" + port}
}
