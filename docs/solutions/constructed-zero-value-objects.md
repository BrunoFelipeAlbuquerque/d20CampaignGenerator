---
title: "Constructed Zero Value Objects"
source: "backlog"
status: "accepted"
tags: ["equipment", "value-object", "misuse-boundary"]
created: "2026-05-08"
---

## Context

Core equipment facts can legitimately have zero cost or zero weight, but an unconstructed Go zero-value struct can look identical unless the value object records construction.

## Decision

When zero is a valid rule value for an exported value object, store an unexported validity marker set only by the constructor. Parent constructors should reject unconstructed zero values while accepting constructed zero values.

## Reuse

Use this shape for future rule values where zero is meaningful, such as item cost, item weight, or similar optional numeric facts. Do not use it for IDs or fields where zero is always invalid.

## Verification

`internal/domain/rpg/character/equipment` tests cover constructed zero cost and weight being accepted, while zero-value cost and weight structs are rejected by `NewEquipment`.
