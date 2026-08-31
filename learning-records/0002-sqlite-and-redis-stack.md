# SQLite + Redis chosen as data stack

The learner chose SQLite for primary persistence and Redis for caching. SQLite keeps local dev zero-config while teaching plain `database/sql` with flat store functions in `internal/store/`. Redis enters after the transform pipeline — cache-aside for transformed image bytes, matching the roadmap.sh caching tip.

**Evidence:** User request during mission setup.

**Implications:** Lesson 4 uses `database/sql` + pure-Go SQLite driver (`modernc.org/sqlite`). Lesson 10 is a dedicated Redis lesson (not a vague "caching" bullet in hardening). Image bytes stay on local disk until optional object-storage lesson; SQLite holds users and image metadata; Redis holds hot transform results.
