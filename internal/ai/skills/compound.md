# Compound Skill — Project Learning

Use after review, before commit, only for lessons from the completed task.

## Purpose

Make later work easier by preserving reusable project-specific decisions.

Capture:

- rule boundaries that were clarified
- invalid shapes that were closed
- core-rules corrections
- testing patterns that prevent repeat mistakes
- architecture constraints that future tasks should follow

Rules source:

- Use only local rules documents in `internal/domain/rpg/resources/rules/`.
- Do not look up Pathfinder rules online.
- Prioritize `PFRPG_SRD.pdf` for all core rule checks.

Do not capture:

- generic Go advice
- generic AI workflow advice
- new feature ideas
- wishlist items
- tasks that belong in `BACKLOG.md` or `ISSUES.md`

## Output Location

Write solution notes in:

```text
docs/solutions/
```

Use one short kebab-case filename:

```text
docs/solutions/{topic}.md
```

## Required Frontmatter

```yaml
---
title: ""
source: "issue|backlog|user-request"
status: "accepted"
tags: []
created: "YYYY-MM-DD"
---
```

## Note Shape

Use these sections:

```markdown
## Context

## Decision

## Reuse

## Verification
```

Keep each section short. Link to concrete files when useful.

## Safety Rules

- `ISSUES.md` and `BACKLOG.md` remain the task sources.
- A solution note is not permission to skip backlog order.
- A solution note must not add non-core Pathfinder content.
- A solution note must not justify redesign.
- If the learning is already encoded by tests or project rules, prefer no note.
