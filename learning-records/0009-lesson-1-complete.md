# Lesson 1 complete — HTTP skeleton and /version

The learner completed lesson 1: healthz endpoint, package layout, go fmt/vet, and the stretch goal — `GET /version` with a dedicated handler registered in `httpserver/server.go`.

**Evidence:** `internal/http/handler/version.go` present; route registered; `go vet ./...` passes.

**Implications:** Unlocks lesson 2 (config, slog, graceful shutdown). Can skip re-teaching mux/handler registration pattern.
