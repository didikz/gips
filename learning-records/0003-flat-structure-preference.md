# Flat structure over layered abstractions

The learner explicitly does not want Laravel-style mental models — no ORM, no repository layers, no interface-heavy indirection. GIPS should stay flat: handlers call store functions, store functions take `*sql.DB` or `*redis.Client`, SQL lives in plain Go files. Add abstraction only when duplication forces it, not upfront.

**Evidence:** User feedback during stack planning.

**Implications:** Lesson 4 introduces `internal/store/` with functions, not interfaces. No "repository pattern" language in lessons or references. PHP background stays in learning records for pacing only — not as a contrast framework in teaching materials.
