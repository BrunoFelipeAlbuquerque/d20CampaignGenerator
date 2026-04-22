# Architecture Skill — d20CampaignGenerator

## Core Principles

* Domain-first design
* Small, validated value objects
* Composition over large aggregates
* Fail fast on invalid input
* Structural rules first, full system later

---

## Domain Structure

### ability

* Primitive domain objects
* Examples: HitPoints, HitDie, BAB, Saves
* No cross-domain logic

### creaturetype

* Classification + structural rules
* Resolves:

  * hitDieType
  * hitPointKind
  * traits
  * contextual flags
* Does NOT compute final character stats

### character

* Official composition boundary
* Combines multiple domains
* No duplication of logic
* Thin adapters allowed

---

## Resolver Pattern

* Input: classification
* Output: resolved structural rules
* No caller-driven behavior
* No partial application

Correct:

```go
ResolveCreatureRules(classification)
```

Avoid:

```go
ResolveCreatureRules(classification, profile, effects...)
```

---

## Integration Rules

* Domains do NOT depend on each other directly
* Use resolved data, not raw classification
* Composition happens in `character`

Example:

```go
character -> creaturetype -> ability
```

NOT:

```go
ability -> creaturetype
```

---

## HitPoints Integration

* creaturetype exposes:

  * hitPointKind
  * hitDieType

* ability handles:

  * HP math
  * constitution/charisma/size logic

* character bridges:

```go
rules.NewRacialHitPoints(...)
```

---

## Important Rule: Do Not Over-Resolve

Example: Humanoid

* Has contextual flag:

  * HumanoidRacialHDUsesClassRulesFlag

* Must NOT:

  * resolve full class behavior here
  * silently expose the wrong convenience path when the higher layer must decide

Correct:

* expose metadata
* reject misleading convenience APIs that normalize the wrong model
* let higher layer decide

---

## Anti-Loop Guard

When fixing a modeling issue:

* prefer making the wrong shape unrepresentable
* fix constructors and public APIs before patching callers
* add one regression test and one misuse-boundary test when applicable
* do not leave semantically wrong convenience methods available just because callers can avoid them manually

---

## Trait System (Current State)

Traits are currently flat:

```go
[]CreatureTypeTraitID
```

This is acceptable for now.

Future pressure:

* traits will split into categories:

  * biological
  * combat
  * sensory
  * systemic

Do NOT refactor yet.

---

## Subtype System

* Implemented as internal resolution
* Not exposed as public API surface
* Minimal supported set only

Avoid:

* full PF subtype catalog
* complex stacking systems

---

## Coding Philosophy

* Small changes
* No redesign unless required
* No abstraction before pressure
* No generic frameworks
* No “future-proofing” guesses

---

## Codex Modes

### Caveman Mode (default)

* direct
* minimal
* no explanation
* surgical changes

### Thinking Mode

* architecture review
* identify risks
* no verbosity

---

## Workflow

1. Design here (ChatGPT)
2. Execute in Codex (small scoped tasks)
3. Validate with tests
4. Move forward

---

## Current Status

* CreatureType domain: stable
* Resolver pattern: correct
* HP integration: correct
* Character composition: established

Next step:

* Expand usage, not design
