---
title: "Carryable Item Boundary"
source: "issue"
status: "accepted"
tags: ["equipment", "inventory", "composition-boundary"]
created: "2026-05-09"
---

## Context

Adventuring gear, weapons, armor, and shields all carry display, cost, and weight facts, but they do not share one full rules model.

## Decision

Expose a carryable item projection under `equipment` for inventory and carried-weight composition.

Keep combat-facing weapon and armor facts in their own chassis. `character` inventory should carry a validated `CarryableItemRef` and resolve weight through the shared carryable lookup.

## Reuse

When core weapon or armor seeds are added, extend the carryable lookup instead of adding inventory-specific side tables or special casing carried weight in `character`.

## Verification

`go test ./...` covers carryable refs, gear lookup, weapon/armor/shield carryable projections, fail-closed unseeded weapon and armor refs, and carried-weight composition through the shared lookup.
