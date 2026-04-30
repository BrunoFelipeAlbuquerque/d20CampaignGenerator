---
title: "Feat Prerequisite Boundary"
source: "issue"
status: "accepted"
tags: ["feat", "prerequisites", "domain-boundary"]
created: "2026-04-30"
---

## Context

Core feat prerequisites are not one flat text field. They reference ability scores, BAB, character level, class level, class features, spellcasting, caster level, skill ranks, selected weapon proficiency, spell-school feat choices, same-selection feat choices, and other feats.

## Decision

Represent feat prerequisites as sealed typed value objects under `internal/domain/rpg/character/feat`.

Do not encode unsupported prerequisite shapes as opaque strings or decorated feat IDs.

## Reuse

When seeding feats, add a constructor for a missing prerequisite shape before adding seed data that needs it.

Keep feat-selection grants, such as human bonus feats, separate from prerequisite strings.

## Verification

`go test ./...` covers constructor validation, zero-value rejection, defensive-copy behavior, and public-package composition of constructor-made prerequisites.
