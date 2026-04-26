# Project Rules — d20CampaignGenerator

## Rules Source

- Use only local rules documents in `internal/domain/rpg/resources/rules/`.
- Do not look up Pathfinder rules online.
- Prioritize `PFRPG_SRD.pdf` for all core rule checks.
- Use `Main35FAQv06302008.pdf` only when a rule doubt remains after checking `PFRPG_SRD.pdf`.

## Architectural Rules

### 1. Domain-first

Build domain objects first.
Do not start from API, persistence, or UI concerns.

### 2. Small domains

Prefer small validated objects over large mutable aggregates.

### 3. Composition boundary

`character` is the official composition boundary.

Meaning:

- `ability` provides primitives
- `creaturetype` provides classification and structural rule resolution
- `character` composes them

### 4. No reverse coupling

Do not make `ability` depend on `creaturetype`.
Do not make primitive domains import higher-level domains.

Allowed direction:

- `character` -> `creaturetype`
- `character` -> `ability`
- `creaturetype` -> `ability` only when structural metadata must bridge into ability-safe constructors

Disallowed direction:

- `ability` -> `creaturetype`
- primitive domain -> composition domain

## Modeling Rules

### 5. Invalid states should fail construction

Use constructors and validation.
Do not silently allow semantically broken states.

### 6. Structural rules first

Resolve:

- hit die type
- hit point kind
- save metadata
- traits
- contextual flags

Do not resolve:

- final BAB
- final saves
- full class interactions
- combat engine behavior
- ancestry/race integration
unless the current task explicitly requires it.

### 7. Contextual exceptions stay contextual

Example:

- humanoid class-rule exception must remain metadata/contextual flag
- do not hardcode final behavior too early

### 8. No fake traits

Do not encode unrelated rule metadata as traits.
Traits are for actual creature properties, not arbitrary bookkeeping.

### 9. No speculative abstractions

Do not create frameworks, generic bucket systems, or meta-layers before real pressure exists.

## Codex Execution Rules

### 10. Minimal diff

Change only what is needed.

### 11. No scope creep

Do not fix unrelated issues in the same pass.

### 12. No redesign by surprise

If the task asks for a fix, fix it.
Do not rewrite the subsystem.

### 13. Prefer one good direction

Do not return multiple design options unless explicitly requested.

## Testing Rules

### 14. Test invariants

Tests should protect:

- constructor validation
- deduplication
- resolved metadata correctness
- contextual flags
- cross-domain bridge behavior

### 15. Test misuse boundaries

Tests should prove the API is hard to misuse.

### 16. Keep tests domain-focused

Do not turn tests into integration suites for unrelated systems.

## Current Strategic Rule

### 17. Prove composition before expanding domains

The next valuable work is usually:

- integrate existing domains safely
not
- add many new traits/subtypes/systems

### 18. Stop polishing when the shape is correct

Once a domain is structurally sound, move to the next integration point.
