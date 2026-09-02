// Package httpserver wires routes and serves HTTP for gips.
package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/didikz/gips/internal/config"
	"github.com/didikz/gips/internal/http/handler"
)

// Server wraps the HTTP server and its dependencies.
type Server struct {
	cfg    config.Config
	server *http.Server
	log    *slog.Logger
}

// New constructs a Server with routes registered.
func New(cfg config.Config, log *slog.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handler.Health)
	mux.HandleFunc("GET /version", handler.Version)

	return &Server{
		cfg: cfg,
		server: &http.Server{
			Addr:    cfg.Addr,
			Handler: mux,
		},
		log: log,
	}
}

func (s *Server) Run() error {
	s.log.Info("server listening", "addr", s.cfg.Addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("server shutting down")
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
