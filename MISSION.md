# Mission: Production-Grade Go via Image Processing Service

## Why

Deepen Go expertise beyond occasional API work so you can design, build, and ship backend services at a senior level — with idiomatic Go, flat structure, and no translated framework patterns.

## Success looks like

- Ship a production-grade image processing service (auth, upload, transform, retrieve, list) modeled on the [roadmap.sh Image Processing Service](https://roadmap.sh/projects/image-processing-service) project brief
- Containerize with multi-stage Docker; CI/CD via GitHub Actions (test on PR, image to GHCR on main)
- Deploy to Kubernetes (local **kind**) with Helm; **VPS + k3s** when infra is ready (lesson 18)
- Write idiomatic Go guided by [Effective Go](https://go.dev/doc/effective_go): clear package boundaries, explicit error handling, interfaces at the right seams
- Start with the standard library; introduce third-party dependencies only when the stdlib genuinely cannot cover the requirement (document each addition)
- Be able to explain architectural choices (handler layout, context usage, concurrency model, testing strategy, cache invalidation) in plain Go terms

## Constraints

- Project-based, structured curriculum — one lesson unlocks the next
- Production-grade from day one: structured logging, graceful shutdown, validation, tests — introduced incrementally, not bolted on at the end
- Standard library first; defer frameworks (Gin, Echo, etc.) until core idioms are solid
- **Flat structure** — handlers → store functions → SQL/Redis. No ORM, no repository layer, no interface indirection until duplication demands it
- **SQLite** for primary persistence; **Redis** for cache + async queue; **Cloudflare R2** for object storage (lesson 15)
- **Docker** (3) + **GitHub Actions** (4) + **Kubernetes/kind** (16) + **Helm** (17) + **VPS/k3s** (18, when ready)
- Background: 10+ years PHP, prior Go HTTP API experience (small projects, several years ago)
- Existing module: `github.com/didikz/gips`, Go 1.27

## Out of scope

- Frontend or mobile clients
- Service mesh, GitOps (ArgoCD/Flux), multi-cluster
- Managed cloud K8s (EKS/GKE) — VPS + k3s is enough for this learning path
- Replacing PHP career knowledge — this workspace is Go-only
- Learning Go syntax from zero (Tour of Go is assumed background reading, not repeated here)
