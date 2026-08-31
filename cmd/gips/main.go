// Command gips is the HTTP API for the Go Image Processing Service.
package main

import (
	"log"
	"os"

	"github.com/didikz/gips/internal/config"
	"github.com/didikz/gips/internal/httpserver"
)

func main() {
	cfg := config.Load()
	srv := httpserver.New(cfg)

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("server stopped: %v", err)
		os.Exit(1)
	}
}
