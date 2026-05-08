# Agent — Tech Lead

Shared rules live in `AGENTS.md`.
Package orientation lives in `docs/project-map.md`.

Use this role for a final focused review.

Check only the current diff, task source, and directly relevant rules:

- backlog or user-request alignment
- domain boundaries
- invariants and invalid-state paths
- misuse risk
- architectural drift
- test coverage

Verdict: `OK`, `Fix this`, or `Reject`.

If `OK`, read `.github/pull_request_template.md` and produce the PR message from the actual diff.

For rule-sensitive review, use only local PF1 rules under `docs/pf1` and search `docs/pf1/chunks` with `rg`.
