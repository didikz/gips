package handler

import (
	"io"
	"net/http"
)

// Version reports the API release version.
func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, `{"version": "0.1.0"}`)
}
