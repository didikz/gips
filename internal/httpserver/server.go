// Package httpserver wires routes and serves HTTP for gips.
package httpserver

import (
	"fmt"
	"net/http"

	"github.com/didikz/gips/internal/config"
	"github.com/didikz/gips/internal/http/handler"
)

// Server wraps the HTTP server and its dependencies.
type Server struct {
	cfg    config.Config
	server *http.Server
}

// New constructs a Server with routes registered.
func New(cfg config.Config) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handler.Health)

	return &Server{
		cfg: cfg,
		server: &http.Server{
			Addr:    cfg.Addr,
			Handler: mux,
		},
	}
}

// ListenAndServe starts the HTTP server and blocks until it stops.
func (s *Server) ListenAndServe() error {
	fmt.Printf("gips listening on %s\n", s.cfg.Addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}
