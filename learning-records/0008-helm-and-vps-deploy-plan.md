# Helm lesson 17; VPS for cloud deploy later

Helm added as lesson 17 — templatize lesson 16 manifests into a chart with values for local (kind) vs VPS (prod). Learner has no cloud infra now; will provision empty VPS later. Plan: kind + Docker for lessons 3–17 locally; VPS becomes optional lesson 18 or self-guided deploy target once cluster exists.

**Evidence:** User request; no VPS available yet.

**Implications:** Lesson 17 produces `deploy/helm/gips/`. Document VPS path: single-node k3s recommended over raw kubeadm for solo learning. kind remains primary for lessons 16–17. VPS deploy doc as reference/stretch, not blocker to start lesson 1.
