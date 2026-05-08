# BACKLOG

## Rules for backlog items

- each item must be implementable in one PR
- each item must have tests
- no item may require redesign unless explicitly stated
- each item must be explicitly one of:
  - domain/chassis
  - core seed data
  - resolution/query logic

---

## Established foundation already present in repo

These packages already exist and are part of the current project baseline. They are not pending backlog items in this file:

- `internal/domain/rpg/character/ability`
  - validated foundational stat and chassis domains already established
- `internal/domain/rpg/character/creaturetype`
  - structural creature rule resolution already established
  - intentionally partial to current project scope
- `internal/domain/rpg/character`
  - current composition boundary already established

The tracked backlog below starts at Race / Modifier / Skill / Class because those are the remaining source-of-truth delivery items after that existing foundation.

---

## Tooling and documentation

- [x] Formalize local Pathfinder rules lookup workflow
  - Add `docs/pf1/README.md`
  - Add `pf1-extract`, `pf1-chunk`, and `pf1-search` Makefile targets
  - Update `AGENTS.md` with Pathfinder rules lookup policy
  - Ensure rule-sensitive tasks require local `rg` lookup before implementation

---

## P0 — Core foundation

### Race

- [X] Create the Race domain chassis:
  - RaceID
  - validated constructor
  - size
  - base speed
  - ability score modifiers
  - racial languages
  - racial features container (no APG traits)

- [X] Seed the 7 core races only:
  - dwarf
  - elf
  - gnome
  - half-elf
  - half-orc
  - halfling
  - human

- [X] Add race query helpers:
  - GetRaceByID
  - HasFeature
  - defensive-copy getters

- [X] Add race tests:
  - valid IDs
  - invalid construction
  - core lookup
  - feature presence

---

### Modifier

- [X] Create the Modifier domain chassis:
  - ModifierType
  - ModifierSource
  - Modifier entry with target and condition slots
  - circumstance source registry

- [X] Add modifier tests:
  - source normalization and validation
  - circumstance registry behavior
  - stacking and penalty resolution

---

### Skill

- [X] Create the Skill domain chassis:
  - SkillID
  - trained-only flag
  - armor-check-penalty flag

- [X] Model grouped skills:
  - Craft
  - Knowledge
  - Perform
  - Profession

- [X] Seed the core skill catalog only

- [X] Add skill tests:
  - valid IDs
  - grouped skill handling
  - invalid rejection

---

### Class

- [X] Create the Class domain chassis:
  - ClassID
  - hit die type
  - BAB progression
  - save progression metadata
  - skill ranks per level
  - class skills
  - weapon/armor proficiency metadata
  - spellcasting kind
  - key ability metadata

- [X] Seed the 11 core classes only:
  - barbarian
  - bard
  - cleric
  - druid
  - fighter
  - monk
  - paladin
  - ranger
  - rogue
  - sorcerer
  - wizard

- [X] Add class tests:
  - valid IDs
  - progression correctness
  - lookup correctness

---

### Spellcasting progression

- [X] Create core spellcasting progression tables:
  - class → spell slots per level

- [X] Seed progression for:
  - bard
  - cleric
  - druid
  - paladin
  - ranger
  - sorcerer
  - wizard

- [X] Add progression tests:
  - known breakpoints
  - caster vs non-caster validation

---

### Spell

- [X] Create the Spell domain chassis:
  - SpellID
  - school
  - descriptors
  - components
  - casting time
  - range
  - target/effect
  - duration
  - saving throw
  - spell resistance

- [X] Create Spell List Entry:
  - spell id
  - class id
  - spell level

- [X] Seed core spell list bindings:
  - class ↔ spell level mapping

- [X] Seed spell data (batch 1):
  - all cantrips/orisons

- [X] Seed spell data (batch 2):
  - levels 1–3

- [X] Seed spell data (batch 3):
  - levels 4–6

- [X] Seed spell data (batch 4):
  - levels 7–9

- [X] Add spell tests:
  - valid construction
  - class list lookup
  - invalid rejection

---

### Feat

- [X] Create the Feat domain chassis:
  - FeatID
  - category
  - prerequisite model
  - fighter bonus feat flag
  - metamagic flag
  - item creation flag

- [X] Seed core general feats

- [X] Seed core combat feats

- [X] Seed core critical feats

- [X] Seed core item creation feats

- [X] Seed core metamagic feats

- [X] Add feat tests:
  - prerequisite validation
  - category correctness
  - invalid rejection

---

## P1 — Core composition

### Minimum level-1 core character creation slice

Near-term goal: prove the existing core domains compose into one small, reviewable character creation path before adding more systems.

The slice should answer:

- can a core race and core class be selected and resolved through `character`?
- can HP use the existing race/class/ability foundations without a broad character aggregate?
- can caster spell slots resolve for a caster class?
- can chosen feats be accepted or rejected from existing prerequisite state?

Do not expand this slice into:

- equipment, wealth, or encumbrance
- skill-rank allocation rules
- spell preparation, spellbooks, or known-spell selection
- combat state
- a full mutable character aggregate unless the slice proves it is required
- non-core sources
- broad folder or package reorganization

- [X] Compose Race with character boundary (no redesign)

- [X] Compose Class with character boundary (no redesign)

- [X] Compose spellcasting progression with Class

- [X] Compose spell list entry with Class spellcasting

- [X] Compose feat prerequisites with:
  - ability scores
  - BAB
  - class features
  - skill ranks
  - other feats

- [ ] Add minimum level-1 character creation slice tests (resolution/query logic):
  - race + class + HP through existing character adapters
  - caster spell slots through existing class spellcasting progression
  - feat prerequisite checks through existing feat prerequisite state
  - invalid selected inputs fail closed

---

## P2 — Later (non-core)

- [ ] Archetype / Alternate Class Feature

- [ ] Prestige Classes

- [ ] Non-core races

- [ ] Non-core feats

- [ ] Non-core spells
