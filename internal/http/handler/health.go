// Package handler provides HTTP handlers for the image processing service.
package handler

import (
	"io"
	"net/http"
)

// Health reports service liveness for load balancers and orchestrators.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, `{"status":"ok"}`)
}
