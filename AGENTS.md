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
- `docs/project-map.md` is the compact orientation guide for package boundaries and current repo shape.
- Do not invent tasks.
- Do not reorder priorities unless required by `ISSUES.md`.

---

## Context budget rules

Keep project context small.

For normal delivery work, read only:

1. `ISSUES.md`, unless explicitly skipped for the current turn
2. the relevant `BACKLOG.md` item or user-requested task
3. `docs/project-map.md` when package boundaries or current repo shape are unclear
4. the files directly touched by the task

Do not bulk-read `internal/ai/**`, `docs/solutions/**`, `docs/pf1/**`, or unrelated domain packages unless the current task needs them.

Use `rg` for targeted lookup instead of broad file scans.

---

## Issue-driven workflow

Before touching the backlog, process `ISSUES.md` unless the user explicitly asks to skip `ISSUES.md` for the current turn.

### Explicit skip option

- If the user explicitly asks to skip `ISSUES.md`, do not read or process it for that turn.
- Do not infer this skip from urgency, wording, or task type.
- The skip applies only to the current turn.
- The skip does not permit backlog reordering, scope expansion, or ignoring the user's requested task.

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

For delivery tasks that edit files, always create/reset the branch first with `git checkout -B`.

Planning mode and read-only audit mode do not require a branch.

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
- `internal/domain/rpg/character/race`: core race seed/query domain
- `internal/domain/rpg/character/skill`: core skill seed/query domain
- `internal/domain/rpg/character/class`: core class and spellcasting progression domain
- `internal/domain/rpg/character/spell`: spell chassis, core spell data, and spell-list bindings
- `internal/domain/rpg/character/feat`: feat chassis, prerequisites, and core feat seed/query domain
- `internal/domain/rpg/character`: composition boundary
- `internal/domain/rpg/modifier`: modifier entries, refs, and stacking resolution

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

## Pathfinder rules lookup

- Local PF1 rules are under `docs/pf1`.
- Use `rg -i "<rule term>" docs/pf1/chunks` before implementing rule-sensitive behavior.
- Use local rule text as source of truth.
- Do not rely on memory when a rule can be checked locally.
- Do not make domain code depend on PDFs, text, chunks, or tooling.

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

## Working modes

Use the lightest mode that satisfies the user's request.

### Delivery mode

Use when the user asks to make changes, continue the backlog, fix an issue, or perform the next item.

- Follow the execution pipeline.
- Create/reset a branch before editing.
- Run tests.
- Commit, push, and output the PR message.

### Planning mode

Use when the user asks to review steps, simplify direction, discuss architecture, or plan.

- Read only the minimum docs/files needed.
- Do not create a branch.
- Do not edit files.
- Do not commit or push.

### Audit mode

Use when the user asks for findings or a system audit.

- Read targeted files and report findings.
- For independent or full system audits, follow `docs/independent-audit-protocol.md`.
- Treat project docs as claims to verify against code, tests, and local PF1 rules,
  not as proof by themselves.
- Only edit `ISSUES.md` if the user explicitly asks to record findings.
- Do not reorder backlog items from audit findings.

### Internal role reminders

The files under `/internal/ai/agents/` and `/internal/ai/skills/` are short reminders only.

- `AGENTS.md` remains the workflow source of truth.
- `docs/project-map.md` remains the package orientation source.
- Do not read all internal role/skill files by default.

---

## Execution pipeline

For each item:

1. Read `ISSUES.md` first, unless the user explicitly asked to skip `ISSUES.md` for the current turn.
2. Decide task source:
   - `NEED` issue → mandatory
   - `SHOULD` issue → only if user chooses it
   - backlog → only if no blocking `NEED` exists
   - explicit `ISSUES.md` skip → requested task or backlog item only
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
   - `git push -u origin <branch-name>`
10. Read `.github/pull_request_template.md` and output the PR message.
11. Stop.

Do not continue automatically to the next item.

---

## Compound learning

After Tech Lead approval, capture reusable project learning only when the current task taught a rule, boundary, or misuse pattern that should make later work easier.

Rules:

- Do not create new product tasks from the learning note.
- Do not change backlog priority from the learning note.
- Do not document generic lessons already covered by this file.
- Keep the note tied to the completed task source.
- Store notes under `docs/solutions/` with YAML frontmatter.
- If no reusable lesson exists, skip the note.

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

## Anti-loop rules

Use these to reduce the pattern of:

- finding issues
- fixing issues
- the fix introducing more issues

For every fix, prefer closing a whole misuse path, not only today's symptom.

### Fix shape before behavior

- If the problem comes from an invalid model shape, fix the constructor or public API first.
- Do not leave a semantically wrong shape available just because callers can "remember" not to use it.
- Prefer removing the wrong convenience path over documenting the caveat.

### Make the violated rule explicit

- When fixing an issue, identify the rule that was violated:
  - invariant
  - domain boundary
  - composition boundary
  - misuse boundary
  - core-rules correctness
- Encode that rule in code or tests so the same mistake is harder to reintroduce.

### Required anti-regression coverage for issue fixes

An issue fix should usually include all applicable items below:

- constructor or API guard
- corrected seed or resolved metadata
- regression test for the observed bug
- misuse-boundary test for the wrong modeling path

Do not close an issue with only a happy-path test if the wrong path is still easy to call.

### Review each fix for issue-creation risk

Before finishing a task, check:

- what wrong modeling path is now impossible?
- what caller misuse is still possible?
- did this change add a new public surface that still permits the old mistake in another form?
- did this fix move responsibility to the correct domain instead of smearing it across callers?

If the answer shows the wrong path is still open, the fix is incomplete.

---

## Review rules

Check:

- invariants
- domain boundaries
- API misuse risk
- architectural drift
- technical debt
- whether the fix removed the wrong modeling path or only patched the symptom

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
