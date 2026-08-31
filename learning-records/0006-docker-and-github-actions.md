# Docker + GitHub Actions CI/CD added to curriculum

Containerization and CI/CD are explicit curriculum topics — not deferred to the end. Docker follows graceful shutdown (lesson 3) so the image runs a production-ready binary. GitHub Actions (lesson 4) runs fmt/vet/test/build on every PR and pushes image to GHCR on main. Kubernetes remains out of scope.

**Evidence:** User request before starting lesson 1.

**Implications:** Lessons 3–4 inserted; feature lessons shift +2. `docker-compose.yml` will wire app + Redis for local dev. Learner has CI/CD experience (GitLab) — focus on Go-specific Dockerfile patterns and GH Actions, not pipeline theory.
