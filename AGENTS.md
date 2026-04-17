# AGENTS.md

Do not be creative. Follow backlog order and project rules.

---

## Project purpose

Model Pathfinder 1e rules using:

- small validated domains
- core-only data
- composition under `character`

Build order:

1. domain/chassis
2. core seed data
3. resolution/query logic
4. composition

Do not skip steps.

---

## Source of truth

- `BACKLOG.md` is the only task source.
- Always pick the first unchecked item.
- Do not invent tasks.
- Do not reorder priorities.

---

## Project structure

- `internal/domain/rpg/character/ability`: primitives and value objects
- `internal/domain/rpg/character/creaturetype`: structural rule resolution
- `internal/domain/rpg/character`: composition boundary

---

## Architecture rules

- Do not redesign architecture.
- Do not expand scope.
- Keep diffs small.
- Preserve invariants.
- Fail on invalid states.
- Keep composition in `character`.
- Do not make `ability` depend on higher domains.
- Keep `creaturetype` structural only.
- Thin adapters in `character` are allowed.

---

## Scope rules

- One backlog item at a time.
- Each item = one PR-sized change.
- Do not fix unrelated issues.
- Do not add extra features.

### Core-only rule

If an item says "core", use Core Rulebook only.

Do not use:

- APG
- ARG
- Ultimate Combat
- Ultimate Magic
- any supplements

---

## Task types

Each backlog item is one of:

### domain/chassis

- create minimal valid domain
- include validation
- include tests
- no large data

### core seed data

- add only requested core data
- do not add extra entries
- do not redesign domain

### resolution/query logic

- keep logic small
- use existing domains
- no new frameworks

---

## Working mode

Use:

- `/internal/ai/agents/product-owner.md` → restate task
- `/internal/ai/agents/senior-dev.md` → implement
- `/internal/ai/agents/tech-lead.md` → review

Also follow:

- `/internal/ai/skills/codex.md`
- `/internal/ai/skills/rules.md`
- `/internal/ai/skills/architecture.md`

---

## Execution pipeline

For each backlog item:

1. Read the item only.
2. Restate using Product Owner:
   - scope
   - constraints
   - acceptance criteria
3. Implement using Senior Developer.
4. Run tests.
5. Review using Tech Lead.
6. Stop.

Do not continue to next item unless explicitly told.

---

## Test command

Run:

`go test ./...`

If tests fail:

- fix only task-related failures
- do not repair unrelated systems

---

## Delivery rules

- Commit only task-related changes.
- Keep changes minimal.
- Prefer explicit code.
- No speculative abstractions.
- If blocked, implement smallest safe version and leave short note.
- Do not skip tests.

---

## Review rules

Check:

- invariants
- domain boundaries
- API misuse risk
- architectural drift
- technical debt

Do not request redesign unless rules are violated.

---

## Stop conditions

Stop when:

- backlog item is complete
- tests pass
- further work requires redesign
- further work introduces non-core content
- next step is unrelated

Do not continue beyond the item.