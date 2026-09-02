# Lesson 2 complete — slog and graceful shutdown

The learner completed lesson 2: slog JSON logging, signal handling (SIGINT/SIGTERM), goroutine + channels for Run/shutdown, httpserver Run/Shutdown delegating to http.Server, clean Ctrl+C. Fixed the recursive Shutdown bug themselves after guidance. Stretch: config_test.go added.

**Evidence:** Code review Sep 2025; user confirmed "done lesson 2".

**Implications:** Unlocks lesson 3 (Docker). Channels/select pattern will recur in lesson 14 (async workers) — reinforce with Redis Streams context later.
