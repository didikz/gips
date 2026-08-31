# R2 for object storage, Redis Streams for async queue

Object storage: Cloudflare R2 (S3-compatible, free tier sufficient for learning and portfolio demo). Async transform jobs: Redis Streams via existing Redis instance — no separate queue product.

**Evidence:** User confirmed Redis-first for queue; asked about R2 free plan.

**Implications:** Lesson 12 uses Redis Streams (XADD/XREADGROUP, consumer workers). Lesson 13 swaps local disk for R2 using AWS SDK v2 S3 client pointed at R2 endpoint. Local disk remains dev default until lesson 13.
