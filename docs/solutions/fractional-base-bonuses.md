---
title: "Fractional Base Bonuses"
source: "issue"
status: "accepted"
tags: ["rules", "class", "saves", "bab"]
created: "2026-05-08"
---

## Context

The project uses core class data, but BAB and base save composition follow the Pathfinder Unchained fractional base bonuses variant.

## Decision

Keep `BaseAttackBonus` and `SavingThrow` fractional internally. For saves, apply the good-save `+2` once per save type when at least one class has a good progression.

## Reuse

Before changing class composition math, search `docs/pf1/chunks` for `fractional base bonuses` and use that local rule note instead of Core Rulebook table-stacking assumptions.

## Verification

`go test ./...` covers the current `SavingThrow.AddClassLevel` one-time good-save behavior.
