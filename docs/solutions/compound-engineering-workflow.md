---
title: "Compound Engineering Workflow"
source: "user-request"
status: "accepted"
tags: ["workflow", "agents", "review"]
created: "2026-04-24"
---

## Context

The project already requires issue-first planning, Product Owner restatement, Senior Developer implementation, Tech Lead review, tests, commit, push, and PR output.

The requested compound-engineering adaptation must not create a second task source or weaken backlog order.

## Decision

Add a small compound learning step after Tech Lead approval.

Capture only reusable project-specific lessons in `docs/solutions/`, using YAML frontmatter for discovery.

## Reuse

Use this pattern when a completed task clarifies a rule boundary, invalid state, core-rules correction, or misuse path.

Skip the note when the lesson is generic, already documented, or would create new backlog work.

## Verification

Future work still starts with `ISSUES.md`, then `BACKLOG.md`.

Solution notes are context only. They do not reorder work.
