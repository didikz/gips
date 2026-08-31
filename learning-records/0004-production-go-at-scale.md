# Production Go experience at scale (Flip, infra lead)

Experience profile from didiktrisusanto.dev: not a Go novice. At Flip, migrated five back-office modules from monolith to Go microservices, optimized disbursement API p95 from >5s to <1s at 200k+ daily requests, and raised SLA compliance. As tech lead, ran load tests and scaled store infra to AWS Fargate (20k+ RPS). Current solo role ships full product loop on GCP/AWS at 99.956% uptime with high message volume.

**Evidence:** https://didiktrisusanto.dev/experience/

**Implications:** Compress HTTP/config/middleware/auth basics (review, not teach). Front-load Go-specific gaps: `context`, error wrapping, table-driven tests, concurrency for transforms. Weight lessons 7–10 (image domain) and observability/load testing heavier. Promote async queue and object storage from optional to recommended given infra background.
