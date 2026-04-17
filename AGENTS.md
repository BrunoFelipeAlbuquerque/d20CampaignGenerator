# AGENTS.md

Do not be creative. Follow backlog, issues, and project rules.

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

- `BACKLOG.md` is the primary task source.
- `ISSUES.md` defines interruptions and corrections.
- Do not invent tasks.
- Do not reorder priorities unless required by `ISSUES.md`.

---

## Issue-driven workflow

Before touching the backlog, ALWAYS process `ISSUES.md`.

### 1. NEED issues (mandatory)

- Check if any `NEED` issue exists.
- If yes:
  - Treat it as a blocking item.
  - Execute it immediately.
  - Do not proceed to backlog.
- `NEED` overrides backlog order.

### 2. SHOULD issues (conditional)

- If there is no `NEED` issue and there are `SHOULD` issues:
  - Ask the user whether to:
    - continue with the backlog
    - or tackle the `SHOULD` issue
- Do not decide autonomously.

### 3. CAN issues (explicit only)

- Ignore `CAN` issues by default.
- Only tackle `CAN` issues if explicitly requested by the user.

---

## Backlog rules

- Only execute backlog work if no blocking `NEED` issue exists.
- Always pick the first unchecked backlog item.
- Do not skip items.
- Do not reorder items.

---

## Branching rules

When starting any task, always create/reset the branch first with `git checkout -B`.

### If the task comes from an issue

~~~bash
git checkout -B issue/{issue-name}
~~~

### If the task comes from the backlog

~~~bash
git checkout -B feat/{backlog-name}
~~~

Rules:

- `{issue-name}` and `{backlog-name}` must be short, descriptive, and kebab-case.
- Always run branch creation before making changes.
- Do not reuse unrelated branch names.

---

## Commit and push rules

After the task is done, always run:

~~~bash
git add .
git commit -m "{commit message according what was done}"
git push -u origin <branch-name>
~~~

Rules:

- Commit message must describe exactly what was done.
- Do not use generic commit messages.

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

- One item at a time.
- Each item = one PR-sized change.
- Do not fix unrelated problems.
- Do not add extra features.

### Core-only rule

If an item says `core`, use Core Rulebook only.

Do not use:

- APG
- ARG
- Ultimate Combat
- Ultimate Magic
- any supplements

---

## Task types

Each item must be one of:

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

For each item:

1. Read `ISSUES.md` first.
2. Decide task source:
   - `NEED` issue → mandatory
   - `SHOULD` issue → only if user chooses it
   - backlog → only if no blocking `NEED` exists
3. Create/reset the branch:
   - issue → `git checkout -B issue/{issue-name}`
   - backlog → `git checkout -B feat/{backlog-name}`
4. Restate using Product Owner:
   - scope
   - constraints
   - acceptance criteria
5. Implement using Senior Developer.
6. Run tests.
7. Review using Tech Lead.
8. If approved:
   - if backlog item, update `BACKLOG.md` by marking only that item as done
   - if issue, update `ISSUES.md` accordingly
9. Run:
   - `git add .`
   - `git commit -m "{commit message according what was done}"`
10. If running locally:
   - `git push -u origin <branch-name>`
11. . If running in Codex Cloud:
   - Create PR through connected GitHub
12. Read `.github/pull_request_template.md` and output the PR message.
13. Stop.

Do not continue automatically to the next item.

---

## Test command

Run:

~~~bash
go test ./...
~~~

If tests fail:

- fix only task-related failures
- do not repair unrelated systems

---

## Delivery rules

- Commit only task-related changes.
- Keep changes minimal.
- Prefer explicit code.
- No speculative abstractions.
- If blocked, implement the smallest safe version and leave a short note.
- Do not skip tests.

---

## Review rules

Check:

- invariants
- domain boundaries
- API misuse risk
- architectural drift
- technical debt

Do not request redesign unless project rules are violated.

---

## Stop conditions

Stop when:

- the item is complete
- tests pass
- further work requires redesign
- further work introduces non-core content
- the next step is unrelated

Do not continue beyond the current item.