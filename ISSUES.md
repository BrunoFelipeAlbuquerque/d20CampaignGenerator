# ISSUES

## Rules for issue ordering

- issues are grouped by criticality
- `NEED` = blocking or correctness-critical
- `SHOULD` = important soon
- `CAN` = useful later

---

## NEED

- [X] Split race language metadata before character language composition hardens around the wrong shape:
  - `Race` currently stores a single `racialLanguages []LanguageID` list
  - the seeds only model automatic starting languages and cannot represent core bonus-language choices
  - core races like dwarf, elf, gnome, half-elf, half-orc, halfling, and human therefore lose rule-important language selection metadata
  - future character language handling would otherwise need side tables or overloading of the current field

- [ ] Stop treating grouped skills as family markers only before later skill modeling builds on the wrong unit:
  - `Skill` only records `grouped bool` for `Craft`, `Knowledge`, `Perform`, and `Profession`
  - there is no way to represent a concrete grouped skill entry such as `Knowledge (arcana)` or `Perform (sing)`
  - current tests explicitly reject specialized grouped entries as invalid IDs
  - later class-skill and character skill-rank work will otherwise have to either rank whole families or bolt on a second skill representation

- [ ] Expose usable modifier target and condition value objects before consumers route around the domain:
  - `Target` and `Condition` are sealed by unexported marker methods
  - the package exports no concrete target or condition types and no constructors for them
  - outside `modifier` code can currently only pass `nil` into those slots
  - this will push future composition work toward ad-hoc side metadata instead of the declared modifier chassis

- [ ] Remove the direct humanoid racial HD convenience path before Class work normalizes the wrong humanoid model:
  - `ResolvedCreatureRules.NewRacialHitDie` still constructs humanoid `d8` racial hit dice
  - `resolver_test.go` still locks in that humanoid construction succeeds while the class-rule flag is present
  - `character.NewRacialHitPoints` guards one entry point, but the lower-level `creaturetype` API still advertises the wrong convenience path
  - the next Class backlog work can easily model humanoids as having racial HD if this surface stays open

- [X] Restore core Strength carrying-capacity math before encumbrance depends on it:
  - `AbilityScore.GetCarryingCapacity` is currently backed by a custom metric table
  - the returned pound values are wrong for core PF1 load bands
  - example: `Strength 18` resolves to about `110.23 lb` light load instead of core `100 lb`
  - current tests lock in the non-core numbers

- [X] Stop shipping project HP house rules as default HP math:
  - `HitDie.GetAverageBaseHP` uses fixed averages (`d6=4`, `d8=5`, `d10=6`, `d12=7`)
  - construct HP uses the custom Large+ / `Titanic` size bonus table
  - `creaturetype` and `character` helpers already depend on those values
  - current HP totals are therefore not trustworthy as core PF1 outputs

- [X] Correct outsider breathing semantics in `creaturetype`:
  - `OutsiderType` currently gets `NoNeedToEatSleepBreathe`
  - standard PF1 outsiders still breathe
  - `NativeSubtype` currently compensates by adding `BreatheEatSleep`, which means the base outsider metadata is wrong

- [X] Validate modifier entries before they become shared composition inputs:
  - `Modifier` is a free-form exported struct with no validated constructor
  - unknown or mistyped `ModifierType` values silently create new stacking buckets
  - empty type has special resolution behavior instead of failing fast
  - this breaks the repo rule that invalid states should fail construction

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

- [X] Expose a read-only skill query surface for the seeded core catalog:
  - `coreSkills` exists but is package-private
  - outside consumers cannot ask for a skill by ID or enumerate the seeded catalog
  - later class-skill and character work will have to reach around the seed instead of using it

- [X] Tighten Race chassis invariants around selectable ability-score bonuses:
  - `NewRace` accepts any non-negative `selectableAbilityScoreModifier`
  - it currently allows semantically impossible combinations, like fixed modifiers plus a selectable bonus on the same race
  - the current core seeds are fine, but the chassis still permits invalid states

- [X] Make the current repo status readable from the source-of-truth docs:
  - `BACKLOG.md` starts at Race / Modifier / Skill / Class
  - the repo also contains established `ability`, `creaturetype`, and `character` packages that are not reflected there
  - a reviewer currently has to reconcile code, README, and git history to understand what is foundational versus backlog-delivered

- [X] Clarify the humanoid misuse boundary before class and character composition expands:
  - `ResolvedCreatureRules.NewRacialHitDie` and `NewRacialHitPoints` happily build humanoid racial HD / HP
  - the contextual flag exists, but the convenience API still makes the wrong path look normal
  - this is easy to use incorrectly once Class arrives

- [X] Make `creaturetype` scope easier to read from the package surface:
  - exported type IDs suggest broad PF1 coverage
  - only a limited subtype set and a partial trait model are actually implemented
  - most of that boundary is discoverable only by reading tests and `profile_table.go`

- [X] Quarantine project house rules from upcoming core backlog work:
  - `TitanicSize`
  - custom construct bonus HP table
  - source-bucket caster levels
  - these should not blur what is core-only work vs project-specific rules

- [X] Replace free-form race language and feature strings with canonical IDs before expanding lookup usage:
  - current validation only rejects empty strings
  - typo and casing drift will leak into query helpers

- [X] Bring repository documentation back in sync with repo rules:
  - `README.md` emphasizes extensibility and homebrew
  - `AGENTS.md` and `BACKLOG.md` enforce backlog-first, core-first delivery

- [X] Make backlog status match repo reality:
  - some unchecked Race helper and test work already exists
  - repo packages like `modifier` are not represented in `BACKLOG.md`

---

## CAN

- [X] Add a compact status matrix for specialist review:
  - what exists
  - what is core-correct
  - what is intentionally partial
  - what is project-specific

- [X] Normalize the remaining vague PF1 seed labels where they are still less explicit than the rulebook:
  - example: elf `Magic` is less precise than `Elven Magic`

- [X] Add read-only discovery helpers for seeded catalogs if lookup pressure grows:
  - race has `GetRaceByID` but no catalog iterator
  - skill has no public discovery surface yet

- [X] Run a later audit pass over the already-built foundation domains after P0 Race, Skill, and Class are aligned

- [X] Collapse duplicated helper logic only if it becomes active maintenance pain:
  - example: ability modifier math is duplicated in more than one place
