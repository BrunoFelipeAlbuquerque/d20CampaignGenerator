# Compound Skill — Project Learning

Shared rules live in `AGENTS.md`.

Use after review, before commit, only when the completed task produced reusable project-specific learning.

Do not write notes for generic workflow advice, normal task summaries, wishlist items, or rules already covered by `AGENTS.md` or `docs/project-map.md`.

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
