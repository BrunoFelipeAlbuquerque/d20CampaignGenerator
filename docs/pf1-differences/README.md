# Pathfinder 1e Differences

This project intentionally deviates from default Pathfinder 1e in a few places.

## Caster Level By Source

Caster level is tracked by spellcasting source, not by individual class.

- `Arcane`
- `Divine`
- `Primal`

This affects only caster level math and spell potency.

It does not merge or share:

- spell slots
- spells known
- prepared spells
- class spell access

Classes are expected to contribute to one of these sources on their own side of the model. The `CasterLevel` domain only stores and updates the final source totals.

## Construct Hit Points

Construct bonus hit points use this project's custom size table instead of the default PF1 values.

Current construct size bonus HP:

- `Tiny`: `+5`
- `Small`: `+10`
- `Medium`: `+20`
- `Large`: `+30`
- `Huge`: `+50`
- `Gargantuan`: `+80`
- `Colossal`: `+130`
- `Titanic`: `+210`

Hit point ledgers also use the project's fixed average hit die values:

- `d6`: `4`
- `d8`: `5`
- `d10`: `6`
- `d12`: `7`

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

Metric size values are project-authored conversions because Pathfinder 1e does not provide an official metric size table for creature dimensions.
