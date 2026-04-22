# Pathfinder 1e Differences

This project intentionally deviates from default Pathfinder 1e in a few places.

## Caster Level By Source

Caster level is tracked by spellcasting source, not by individual class.

- `Arcane`
- `Divine`
- `Primal`

`Primal` is project-specific. Current core-only backlog work treats `Arcane` and `Divine` as the core-aligned sources.

This affects only caster level math and spell potency.

It does not merge or share:

- spell slots
- spells known
- prepared spells
- class spell access

Classes are expected to contribute to one of these sources on their own side of the model. The `CasterLevel` domain only stores and updates the final source totals.

## Titanic Size

`Titanic` is a homebrew creature size officialized by this project.

It extends the standard PF1 size ladder beyond `Colossal` and includes custom values for:

- attack and AC modifier
- CMB and CMD modifier
- Fly modifier
- Stealth modifier
- construct bonus hit points
- space
- natural reach
- typical height
- typical weight

Core construct bonus hit points use the default PF1 table for core sizes. `Titanic` keeps a project-specific construct bonus HP extension because the size itself is project-specific.

Metric size values are project-authored conversions because Pathfinder 1e does not provide an official metric size table for creature dimensions.
