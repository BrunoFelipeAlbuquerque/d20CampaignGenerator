# d20campaing

Pathfinder 1e character-building and rules-domain project.

This repository is aiming for a rules-aware domain model, not just a loose bag of structs. The goal is to represent character data in a way that is:

- explicit
- validated
- extensible
- friendly to homebrew
- safe to evolve as more Pathfinder systems are added

The project is still growing, but the direction is already clear: build small domain objects with strong invariants, then compose them into larger character systems later.

## What This Project Is Trying To Achieve

The long-term goal is to model Pathfinder 1e characters and creature logic in a way that stays maintainable even as the rules get ugly.

Pathfinder is full of systems that look simple on the surface but become hard once you combine:

- multiclassing
- racial rules
- monster rules
- exceptional cases
- optional rules
- homebrew

Because of that, this project prefers a domain-first approach:

- keep each rules concept in its own value object or small aggregate
- validate at construction time
- avoid magic numbers leaking all over the codebase
- avoid coupling rule containers to unrelated concepts too early
- preserve enough detail to explain where a number came from

This matters because many later systems will depend on these primitives:

- creature type and subtype
- classes
- race and ancestry
- spellcasting
- combat state
- monster advancement

If the early stat objects are sloppy, everything above them becomes brittle.

## Current Architectural Direction

The current design favors narrow, well-defined domain types under `internal/domain/rpg/character/ability`.

Each type should answer one specific rules question well.

Examples:

- `AbilityScore` answers what a score is, whether it is currently valid, and what modifier it produces
- `Alignment` answers which axes are present and how they should be named
- `BaseAttackBonus` answers both the exact fractional progression and the game-facing rounded value
- `SavingThrow` answers the same, while also handling the one-time good-save bonus correctly across multiclassing
- `CasterLevel` stores source-based caster level totals without deciding how classes map into them
- `HitPoints` is a ledger of raw HP totals and sources, not a combat-state engine
- `Size` is a rules table for size-derived modifiers and physical dimensions

The project is intentionally trying not to jump too quickly into giant “character” structs that know everything. That kind of design usually feels productive at first, then becomes painful once exceptions appear.

## Main Design Decisions

### Invalid Input Is Impossible To Construct

This is one of the most important project rules.

If a domain value is invalid, construction should fail instead of silently creating a fake default object.

In practice, that means constructors usually return:

- `(value, true)` when valid
- `(zero, false)` when invalid

Why this matters:

- invalid state is easier to catch near the source
- later code can trust constructed values more
- bugs do not hide behind harmless-looking zero values

### Domain Objects Should Be Small And Purposeful

Types should not own concerns that belong elsewhere.

For example:

- `HitPoints` owns totals, sources, temporary HP, nonlethal damage, and recalculation
- `HitPoints` does not own the entire combat-state model

That means concepts like:

- disabled
- dying
- unconscious
- dead state transitions

should live in a separate combat-state domain later.

This keeps the HP ledger useful without turning it into a god object.

### Preserve Exact Rules Math Internally

When Pathfinder uses fractional systems, this project tries to preserve the exact math internally and only round when the rules actually require it.

That is why `BaseAttackBonus` and `SavingThrow` use exact rational values instead of floats.

This avoids:

- float drift
- awkward equality checks
- hidden rounding errors in multiclass calculations

### Support Homebrew Without Forcing Core Refactors

A recurring design goal is flexibility.

If a future homebrew class, creature type, or spellcasting source is added, the project should not require rewriting core stat containers just to make room for it.

That is why some types are intentionally generic.

Example:

- `CasterLevel` stores `Arcane`, `Divine`, and `Primal` source totals
- it does not hardcode class names into itself
- classes are expected to contribute to those totals elsewhere

This makes the system more reusable and less brittle.

### Keep Exact Value And Display Value When The Rules Need Both

Some Pathfinder stats have:

- a real internal value
- a displayed or applied value

That is why:

- `BaseAttackBonus` stores exact rational progression and rounded integer BAB
- `SavingThrow` stores exact rational progression and rounded integer save

This keeps the math honest without losing the number that actually matters at the table.

### Keep A Ledger When Provenance Matters

For some systems, just knowing the final total is not enough.

`HitPoints` is the clearest example.

It keeps sources like:

- `Base Dice`
- `Constitution`
- `Charisma (Undead)`
- `Construct Size Bonus`
- temporary HP sources

That makes recalculation and debugging much easier than storing only one final integer.

## Current House Rules And Project Conventions

Some rules in this project intentionally differ from default Pathfinder 1e.

These are documented in:

[PF1 Differences](/home/brunoalbuquerquemeta/Documentos/stuff/d20campaing/docs/pf1-differences/README.md)

Current notable differences include:

- caster level tracked by source instead of by individual class
- custom construct bonus HP table
- `Titanic` as an officialized project creature size
- project-authored metric conversions for creature size ranges
- fixed average hit die values for HP calculations

## Current Domain Snapshot

### `AbilityScore`

Represents one ability score and its validity.

Important notes:

- score validity is explicit
- score value and score availability are separate concepts
- modifier calculation depends on a valid visible score

### `Alignment`

Represents the order and morality axes.

Important notes:

- construction is validated
- true neutral is rendered as `Neutral`, not `Neutral Neutral`

### `BaseAttackBonus`

Represents BAB as both exact fraction and rounded final value.

Important notes:

- class progressions are stored as rational values
- rounding only happens when deriving the displayed BAB

### `SavingThrow`

Represents an individual save with exact and rounded values.

Important notes:

- good and poor progressions are exact fractions
- the one-time `+2` good-save base bonus is tracked so multiclassing does not repeat it incorrectly

### `CasterLevel`

Represents source-based caster level totals.

Important notes:

- sources are currently `Arcane`, `Divine`, and `Primal`
- a source can be impossible for a character instead of merely zero
- this domain stores totals, not class mapping rules

### `HitDie`

Represents the hit-die composition of a character or creature.

Important notes:

- semantic validity matters, not just field presence
- zero-total hit dice are rejected
- average HP uses the project’s fixed averages

### `HitPoints`

Represents a raw HP ledger.

Important notes:

- keeps a source breakdown
- supports temporary HP
- supports nonlethal damage
- recalculates when the relevant underlying stat changes
- does not own full combat-state semantics

### `Size`

Represents size-based rules data.

Important notes:

- includes attack and AC modifier
- includes CMB and CMD modifier
- includes Fly and Stealth modifiers
- includes construct HP bonus
- includes space and reach by body shape
- includes typical height and weight ranges in imperial and metric
- includes the homebrew `Titanic` size

## Glossary

### Valid

When this project says a value is `valid`, it means the value exists and is currently usable under the rules for that domain.

Example:

- an `AbilityScoreValue` can store a number but still be invalid for use

### Invalid

An invalid value is one that should not be constructed or accepted by a validated setter.

Examples:

- negative hit dice
- unknown saving throw ids
- malformed size body shapes

### Impossible

`Impossible` is stronger than zero.

It means a stat or source does not exist for that character or cannot currently be used at all.

This is used in `CasterLevel`.

Example:

- `0, valid` means the source exists and is currently zero
- `0, invalid` means the source is impossible or unavailable

### Source

`Source` is the project term for a spellcasting origin bucket in `CasterLevel`.

Current sources:

- `Arcane`
- `Divine`
- `Primal`

This was chosen because it is more general and future-friendly than tying caster level directly to class names.

### Exact Value

The mathematically exact internal value before the game-facing rounding is applied.

Examples:

- BAB of `3/4`
- save total of `10/3`

### Display Value

The rounded or directly used table-facing number.

Examples:

- BAB `0` from exact `3/4`
- Fortitude `3` from exact `10/3`

### Ledger

A ledger is a domain structure that preserves the pieces that make up a total instead of storing only the final result.

This is especially useful when values must be recalculated after related stats change.

### Body Shape

The movement/reach profile used by `Size` when determining space and natural reach.

Current supported shapes:

- `Tall`
- `Long`

Invalid body shapes are rejected.

## Near-Term Next Steps

The next major area is likely:

- creature type and subtype

That system will likely interact with:

- hit dice
- hit points
- size
- immunities
- special traits

So it is important to keep the current domains clean before building on top of them.

## Philosophy In One Sentence

Model Pathfinder rules as small validated domains first, then compose them into bigger systems only when the boundaries are clear.
