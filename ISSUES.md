# ISSUES

## Rules for issue ordering

- issues are grouped by criticality
- `NEED` = blocking or correctness-critical
- `SHOULD` = important soon
- `CAN` = useful later

---

## NEED

- [X] Restore backlog alignment before adding more domains or composition:
  - the first unchecked backlog item is still Race query helpers
  - repo work has already moved into composition and off-backlog surfaces
  - `BACKLOG.md` must stay trustworthy as the project source of truth

- [X] Complete the missing Race query slice:
  - add `GetRaceByID`
  - expose the backlog query surface for seeded core races
  - keep returned race data defensive-copy safe

- [X] Fix core race seed correctness:
  - gnome is missing `Keen Senses`
  - half-elf is missing `Elven Immunities`
  - half-elf is missing `Keen Senses`
  - half-orc is missing `Orc Ferocity`
  - human, half-elf, and half-orc still represent the variable `+2` racial bonus as marker text instead of rules data

- [X] Stop duplicating structural race facts as ad-hoc racial features:
  - size already exists as `size`
  - speed already exists as `baseSpeed`
  - current seeds mix `Medium`, `Small`, `Normal Speed`, and `Slow Speed` into feature lists inconsistently

- [X] Correct the `Fine` size rules entry:
  - current space is `0 ft`
  - Pathfinder size data expects `0.5 ft` space

---

## SHOULD

- [X] Quarantine project house rules from upcoming core backlog work:
  - `TitanicSize`
  - custom construct bonus HP table
  - source-bucket caster levels
  - these should not blur what is core-only work vs project-specific rules

- [ ] Replace free-form race language and feature strings with canonical IDs before expanding lookup usage:
  - current validation only rejects empty strings
  - typo and casing drift will leak into query helpers

- [ ] Bring repository documentation back in sync with repo rules:
  - `README.md` emphasizes extensibility and homebrew
  - `AGENTS.md` and `BACKLOG.md` enforce backlog-first, core-first delivery

- [ ] Make backlog status match repo reality:
  - some unchecked Race helper and test work already exists
  - repo packages like `modifier` are not represented in `BACKLOG.md`

---

## CAN

- [ ] Run a later audit pass over the already-built foundation domains after P0 Race, Skill, and Class are aligned

- [X] Collapse duplicated helper logic only if it becomes active maintenance pain:
  - example: ability modifier math is duplicated in more than one place
