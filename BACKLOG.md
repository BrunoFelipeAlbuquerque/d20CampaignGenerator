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

- [ ] Seed the 7 core races only:
  - dwarf
  - elf
  - gnome
  - half-elf
  - half-orc
  - halfling
  - human

- [ ] Add race query helpers:
  - GetRaceByID
  - HasFeature
  - defensive-copy getters

- [ ] Add race tests:
  - valid IDs
  - invalid construction
  - core lookup
  - feature presence

---

### Skill

- [ ] Create the Skill domain chassis:
  - SkillID
  - trained-only flag
  - armor-check-penalty flag

- [ ] Model grouped skills:
  - Craft
  - Knowledge
  - Perform
  - Profession

- [ ] Seed the core skill catalog only

- [ ] Add skill tests:
  - valid IDs
  - grouped skill handling
  - invalid rejection

---

### Class

- [ ] Create the Class domain chassis:
  - ClassID
  - hit die type
  - BAB progression
  - save progression metadata
  - skill ranks per level
  - class skills
  - weapon/armor proficiency metadata
  - spellcasting kind
  - key ability metadata

- [ ] Seed the 11 core classes only:
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

- [ ] Add class tests:
  - valid IDs
  - progression correctness
  - lookup correctness

---

### Spellcasting progression

- [ ] Create core spellcasting progression tables:
  - class → spell slots per level

- [ ] Seed progression for:
  - bard
  - cleric
  - druid
  - paladin
  - ranger
  - sorcerer
  - wizard

- [ ] Add progression tests:
  - known breakpoints
  - caster vs non-caster validation

---

### Spell

- [ ] Create the Spell domain chassis:
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

- [ ] Create Spell List Entry:
  - spell id
  - class id
  - spell level

- [ ] Seed core spell list bindings:
  - class ↔ spell level mapping

- [ ] Seed spell data (batch 1):
  - all cantrips/orisons

- [ ] Seed spell data (batch 2):
  - levels 1–3

- [ ] Seed spell data (batch 3):
  - levels 4–6

- [ ] Seed spell data (batch 4):
  - levels 7–9

- [ ] Add spell tests:
  - valid construction
  - class list lookup
  - invalid rejection

---

### Feat

- [ ] Create the Feat domain chassis:
  - FeatID
  - category
  - prerequisite model
  - fighter bonus feat flag
  - metamagic flag
  - item creation flag

- [ ] Seed core general feats

- [ ] Seed core combat feats

- [ ] Seed core critical feats

- [ ] Seed core item creation feats

- [ ] Seed core metamagic feats

- [ ] Add feat tests:
  - prerequisite validation
  - category correctness
  - invalid rejection

---

## P1 — Core composition

- [ ] Compose Race with character boundary (no redesign)

- [ ] Compose Class with character boundary (no redesign)

- [ ] Compose spellcasting progression with Class

- [ ] Compose spell list entry with Class spellcasting

- [ ] Compose feat prerequisites with:
  - ability scores
  - BAB
  - class features
  - skill ranks
  - other feats

- [ ] Add character creation slice tests:
  - race + class + HP
  - caster spell slots
  - feat prerequisite checks

---

## P2 — Later (non-core)

- [ ] Archetype / Alternate Class Feature

- [ ] Prestige Classes

- [ ] Non-core races

- [ ] Non-core feats

- [ ] Non-core spells