# Go Resources — Image Processing Service

## Knowledge

- [Effective Go](https://go.dev/doc/effective_go)
  Canonical style and idioms for naming, formatting, interfaces, concurrency, and error handling. Use for: every design decision in this project.
- [How to Write Go Code](https://go.dev/doc/code)
  Module layout, `go build`, testing, and documentation conventions. Use for: project structure and tooling.
- [Go Doc Comments](https://go.dev/doc/comment)
  How to write package and symbol documentation. Use for: exported APIs you want maintainable long-term.
- [net/http package documentation](https://pkg.go.dev/net/http)
  Standard-library HTTP server, handlers, middleware patterns via `Handler`/`HandlerFunc`. Use for: all HTTP layers until we deliberately add a router.
- [database/sql package documentation](https://pkg.go.dev/database/sql)
  Stdlib SQL interface — connections, transactions, prepared statements. Use for: `internal/store/` functions (lesson 4+).
- [modernc.org/sqlite driver](https://pkg.go.dev/modernc.org/sqlite)
  Pure-Go SQLite driver (no CGO). Use for: local dev and tests in this project.
- [SQLite SQL syntax](https://www.sqlite.org/lang.html)
  Primary reference for DDL and queries. Use for: migrations and store-layer SQL.
- [redis/go-redis v9](https://pkg.go.dev/github.com/redis/go-redis/v9)
  Official Redis client for Go. Use for: transform cache, Redis Streams workers, rate limits (lessons 10–12).
- [Redis caching patterns (Redis docs)](https://redis.io/docs/latest/develop/use/patterns/)
  Cache-aside, TTL, invalidation strategies. Use for: lesson 10 design.
- [Cloudflare R2 pricing](https://developers.cloudflare.com/r2/pricing/)
  Free tier: 10 GB-month storage, 1M Class A ops, 10M Class B ops, zero egress. Use for: lesson 15 object storage.
- [Redis Streams introduction](https://redis.io/docs/latest/develop/data-types/streams/)
  Consumer groups, at-least-once delivery. Use for: lesson 14 async transform workers.
- [Docker multi-stage builds](https://docs.docker.com/build/building/multi-stage/)
  Build in golang image, copy binary to minimal runtime. Use for: lesson 3 Dockerfile.
- [GitHub Actions — Go workflow](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)
  setup-go, test, vet. Use for: lesson 4 CI pipeline.
- [Working with GHCR](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
  Push container images with GITHUB_TOKEN. Use for: lesson 4 docker job.
- [Kubernetes Basics tutorial](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
  Deployments, Services, scaling, rolling updates. Use for: lesson 16 foundation.
- [kind quick start](https://kind.sigs.k8s.io/docs/user/quick-start/)
  Local K8s cluster inside Docker. Use for: lesson 16 local deploy.
- [Helm chart template guide](https://helm.sh/docs/chart_template_guide/getting_started/)
  Charts, values, install/upgrade/rollback. Use for: lesson 17.
- [k3s quick start](https://docs.k3s.io/quick-start)
  Single-node K8s on a VPS. Use for: lesson 18 VPS deploy.
- [image package documentation](https://pkg.go.dev/image)
  Core image types and decoding/encoding. Use for: transform pipeline (lessons 8+).
- [Roadmap.sh: Image Processing Service](https://roadmap.sh/projects/image-processing-service)
  Target project spec: auth, upload, transform, retrieve, list. Use for: feature checklist and endpoint contracts.
- [The Go Programming Language (book) — Donovan & Kernighan](https://www.gopl.io/)
  Deeper treatment of concurrency, reflection, and the standard library. Use for: when Effective Go raises a question it doesn't fully answer.

## Wisdom (Communities)

- [Gophers Slack](https://invite.slack.golangbridge.org/)
  High-signal channel for idiomatic Go questions. Use for: design reviews and "is this idiomatic?" checks.
- [r/golang](https://reddit.com/r/golang)
  Active community; filter for quality. Use for: ecosystem news and pattern discussions.
- [Go Time podcast (Changelog)](https://changelog.com/gotime)
  Practitioner interviews and language evolution. Use for: broader context, not day-to-day coding.

## Gaps

- Production-grade stdlib-only HTTP service walkthrough with graceful shutdown (lesson 2)
- Stdlib image transformation recipes beyond resize/crop (filters may need `golang.org/x/image` — lesson 10)

## Community preferences

- Not yet stated — user has not opted out of communities.
