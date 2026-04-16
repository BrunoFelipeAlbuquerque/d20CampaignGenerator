// Package creaturetype models creature type and subtype classification rules.
//
// It resolves structural and racial-HD metadata, and can bridge that metadata
// into hit-die and hit-point domains when given the needed inputs.
// It does not compute final BAB, final saves, or wider class interactions.
// Only a limited subset of subtypes is supported intentionally.
// The package is meant to be composed by higher-level domains later.
package creaturetype
