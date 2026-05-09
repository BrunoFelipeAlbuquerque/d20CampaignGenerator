---
title: "Feat Caster Level Composition"
source: "issue"
status: "accepted"
tags: ["character", "feat", "caster-level", "composition-boundary"]
created: "2026-05-09"
---

## Context

Core feats can require a minimum caster level, but class levels and spellcasting access are not the same fact as caster level.

## Decision

`character` feat prerequisite composition accepts explicit caster-level facts and evaluates `CasterLevelPrerequisite` against those facts.

Do not infer generic caster level from class levels in `CharacterFeatPrerequisiteState`.

## Reuse

When a feat prerequisite depends on a derived character fact, pass the fact explicitly through the character composition state unless the current task is specifically to build that derivation.

## Verification

`go test ./...` covers caster-level prerequisite success, low-level rejection, invalid caster-level facts, duplicate caster-level facts, and the class-level inference misuse path.
