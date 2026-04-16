package ability

// Package ability contains low-level Pathfinder 1e character stat domains.
//
// Project conventions:
// - invalid input is impossible to construct
// - fractional BAB and save math uses exact rational values, not floats
// - HitDie average HP uses the project's fixed averages: d6=4, d8=5, d10=6, d12=7
