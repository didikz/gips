# Teaching notes

## Learner profile

- Senior software engineer / tech lead / solo product engineer — 10+ years
- **Go in production:** Flip For Business — monolith → Go microservices, 200k+ req/day, p95 latency work, SNAP compliance
- **Infra:** AWS Fargate (dated), GCP/AWS solo stack, load testing, 99.956% uptime systems
- **Wants to learn:** Kubernetes (kind locally), Go concurrency, stdlib image pipeline, Effective Go style
- Target project: roadmap.sh Image Processing Service

## Pace

- Treat lessons 1–3 as quick review unless gaps show up in exercises
- Do not teach "what is JWT" or "why cache" — teach Go-specific mechanics
- Testing and observability woven per lesson, not one late "hardening" dump

## Preferences

- Project-based learning over isolated syntax drills
- Start simple; add complexity only when the project demands it
- **SQLite** for database — zero-config local dev, plain SQL via `database/sql`
- **Redis** for caching — especially transformed images (roadmap.sh tip)
- **Flat Go layout** — no ORM, no repository pattern, no framework-style layering; simple packages and functions
- **No cloud infra yet** — learn deploy on **kind** locally; empty **VPS + k3s** later for lesson 18
- **Docker + GitHub Actions + K8s + Helm** — full deploy path on laptop first

## Stack decisions

| Layer | Choice | Why |
|-------|--------|-----|
| Primary DB | SQLite via `database/sql` | Single file, no server to run for early lessons |
| SQLite driver | `modernc.org/sqlite` | Pure Go, no CGO — `go test` works everywhere |
| Data access | `internal/store/` functions | Flat: `store.CreateUser(db, …)` — not interfaces or ORM |
| Cache | Redis via `redis/go-redis/v9` | Direct client in store/cache helpers when needed |
| Image files | Local disk → **Cloudflare R2** (lesson 15) | R2 free tier covers learning; S3-compatible API |
| Async queue | **Redis Streams** | Same Redis as cache; XADD consumer workers |
| Container | **Docker** multi-stage + distroless | Lesson 3, after graceful shutdown |
| CI/CD | **GitHub Actions** → GHCR | Lesson 4: test on PR, push image on main |
| Orchestration | **Kubernetes (kind)** | Lesson 16: plain manifests |
| Packaging | **Helm 3** | Lesson 17: chart + values |
| Production host | **VPS + k3s** | Lesson 18 when VPS ready; 2 GB+ RAM |

## Teaching style

- Do **not** use Laravel/ORM/repository as comparison points in lessons
- Explain Go on its own terms: explicit, flat, stdlib-first
- **Do not edit learner's code** — propose fixes and explain; let them implement (explicit preference, Sep 2025)


