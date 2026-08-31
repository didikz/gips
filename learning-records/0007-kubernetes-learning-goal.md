# Kubernetes added to curriculum

Learner wants Kubernetes despite acknowledging it may be overengineered for the project alone — learning K8s is an explicit goal. Fargate experience is dated; K8s replaces "out of scope" deploy guidance. Local cluster via kind (fits Docker lesson 3); manifests evolve as features land; full deploy lesson after R2 (lesson 16).

**Evidence:** User request during curriculum planning.

**Implications:** Lesson 16 covers Deployment, Service, ConfigMap, Secret, PVC (SQLite single-replica), optional Ingress. CI pushes to GHCR; K8s pulls from GHCR. Document SQLite ≠ HA — one replica for metadata, or note path to Postgres later.
