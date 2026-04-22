// Package creaturetype resolves a limited structural subset of Pathfinder 1e
// creature rules for the current project scope.
//
// Current surface covers:
//   - base creature type structural metadata
//   - a limited supported subtype set: Aquatic, Augmented, Elemental,
//     Incorporeal, Native, and Swarm
//   - partial resolved traits and contextual flags needed by the existing
//     foundation and current backlog pressure
//
// This package is intentionally partial. It is not the full PF1 subtype catalog,
// not the full creature-trait rules engine, and not the final character-stat or
// combat-resolution layer. Composition remains in character.
package creaturetype
