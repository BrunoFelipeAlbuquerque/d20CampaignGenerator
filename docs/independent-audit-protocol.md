# Independent Audit Protocol

Use this protocol when the user asks for an independent audit, full system audit,
or a check before choosing the next domain or backlog item.

## Purpose

Audit the project as an independent reviewer. `AGENTS.md`, `BACKLOG.md`,
`ISSUES.md`, and `docs/project-map.md` are inputs to verify, not proof that the
system is correct.

The audit should find contradictions between the intended project shape and the
actual code, tests, seed data, and local Pathfinder 1e rules.

## Scope Rules

- Do not fix issues during audit unless the user explicitly asks for fixes.
- Do not reorder `BACKLOG.md` from audit findings.
- Only edit `ISSUES.md` when the user explicitly asks to record findings.
- Keep reads targeted. Use `rg` before opening broad file sets.
- Do not bulk-read `internal/ai/**`, `docs/solutions/**`, `docs/pf1/**`, or
  unrelated domains unless the audit question requires it.

## Audit Steps

1. State the audit question and scope.
2. Read `ISSUES.md`, the relevant `BACKLOG.md` section, and
   `docs/project-map.md`.
3. Extract the claims those files make about current priorities, package
   boundaries, completed work, and known risks.
4. Verify those claims against the code and tests with targeted searches.
5. For rule-sensitive behavior, check local rule text with
   `rg -i "<rule term>" docs/pf1/chunks`.
6. Run `go test ./...` when feasible.
7. Report findings separately from suggestions.

## What To Look For

- Source-of-truth drift between docs, backlog, issues, and code.
- Invariants that are described but not enforced by constructors or public APIs.
- Domain boundary leaks, especially logic outside `character` composition that
  should remain domain-local or composition-local.
- Fail-open paths where invalid rules data can be represented silently.
- Public API misuse paths that are easy to call and not covered by tests.
- Seed data that does not match local Core Rulebook text.
- Backlog steps that jump order, become non-core, or depend on missing chassis.
- Composition assumptions that are not represented in tests.
- Test gaps around the wrong modeling path, not only happy paths.

## Finding Severity

Use the project issue language when recording findings:

- `NEED`: Blocks correct delivery, breaks tests/builds, violates a core
  invariant, or makes the current backlog unsafe to continue.
- `SHOULD`: Important correctness, boundary, or coverage risk that should be
  handled soon but does not block the immediate backlog path.
- `CAN`: Optional cleanup, documentation improvement, or non-blocking ergonomics.

## Finding Format

Each finding should include:

- severity
- concise title
- evidence with file paths, symbols, tests, or local rule chunks
- why it matters
- smallest recommended next action

If no findings are found, say what was checked and what was not checked.

## Recording Findings

When the user asks to record audit findings, update `ISSUES.md` only. Use the
same `NEED`, `SHOULD`, and `CAN` language above.

Do not turn audit findings into backlog work unless the user explicitly asks to
reshape the backlog.

## Audit Anti-Patterns

- Declaring the system healthy only because tests pass.
- Restating `BACKLOG.md` without verifying it against code.
- Treating `AGENTS.md` or `docs/project-map.md` as correctness evidence.
- Expanding into unrelated packages because they are nearby.
- Fixing findings while in audit mode without explicit user approval.
- Recording speculative tasks that are not tied to concrete evidence.

## Output Shape

Use this order:

1. Scope
2. Findings
3. Evidence Checked
4. Recommended Next Step
5. Not Checked
