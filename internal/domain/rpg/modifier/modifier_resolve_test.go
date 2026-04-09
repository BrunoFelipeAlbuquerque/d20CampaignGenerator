package modifier

import "testing"

// ==============================
// MOCKS
// ==============================

type mockTarget struct{}

func (mockTarget) isTarget() {}

type mockCondition struct{}

func (mockCondition) isCondition() {}

// ==============================
// HELPERS
// ==============================

func m(t ModifierType, v int, src ModifierCircumstanceSource) Modifier {
	return Modifier{
		Type:   t,
		Value:  v,
		Source: src,
		Target: mockTarget{},
	}
}

// ==============================
// DODGE (STACK ALL)
// ==============================

func TestModifierResolve_DodgeStacks(t *testing.T) {
	mods := ModifierList{
		m(ModifierDodge, 1, "a"),
		m(ModifierDodge, 2, "b"),
		m(ModifierDodge, 3, "c"),
	}

	got := mods.ModifierResolve()
	want := 6

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// ==============================
// DEFAULT (TAKE MAX)
// ==============================

func TestModifierResolve_DefaultMax(t *testing.T) {
	mods := ModifierList{
		m("enhancement", 1, "a"),
		m("enhancement", 5, "b"),
		m("enhancement", 3, "c"),
	}

	got := mods.ModifierResolve()
	want := 5

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// ==============================
// CIRCUMSTANCE (STACK BY SOURCE)
// ==============================

func TestModifierResolve_CircumstanceDifferentSources(t *testing.T) {
	mods := ModifierList{
		m(ModifierCircumstance, 2, "flanking"),
		m(ModifierCircumstance, 1, "higher_ground"),
	}

	got := mods.ModifierResolve()
	want := 3

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_CircumstanceSameSourceTakeHighest(t *testing.T) {
	mods := ModifierList{
		m(ModifierCircumstance, 1, "flanking"),
		m(ModifierCircumstance, 3, "flanking"),
		m(ModifierCircumstance, 2, "flanking"),
	}

	got := mods.ModifierResolve()
	want := 3

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// ==============================
// MIXED TYPES
// ==============================

func TestModifierResolve_MixedTypes(t *testing.T) {
	mods := ModifierList{
		m(ModifierDodge, 1, "a"),
		m(ModifierDodge, 2, "b"),

		m("enhancement", 5, "x"),
		m("enhancement", 3, "y"),

		m(ModifierCircumstance, 2, "flanking"),
		m(ModifierCircumstance, 1, "higher_ground"),
	}

	got := mods.ModifierResolve()
	want := 3 + 5 + 3 // dodge(3) + enhancement(5) + circumstance(3)

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// ==============================
// NEGATIVE VALUES
// ==============================

func TestModifierResolve_Negatives_Default(t *testing.T) {
	mods := ModifierList{
		m("enhancement", -1, "a"),
		m("enhancement", -5, "b"),
		m("enhancement", -3, "c"),
	}

	// current behavior: picks max (less negative)
	got := mods.ModifierResolve()
	want := -1

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_Negatives_Dodge(t *testing.T) {
	mods := ModifierList{
		m(ModifierDodge, -1, "a"),
		m(ModifierDodge, -2, "b"),
	}

	got := mods.ModifierResolve()
	want := -3

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// ==============================
// EMPTY LIST
// ==============================

func TestModifierResolve_Empty(t *testing.T) {
	var mods ModifierList

	got := mods.ModifierResolve()
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

// ==============================
// SINGLE VALUE
// ==============================

func TestModifierResolve_Single(t *testing.T) {
	mods := ModifierList{
		m("enhancement", 7, "a"),
	}

	got := mods.ModifierResolve()
	if got != 7 {
		t.Errorf("expected 7, got %d", got)
	}
}

// ==============================
// CIRCUMSTANCE EDGE CASE
// ==============================

func TestModifierResolve_CircumstanceMixedSameSourceNegative(t *testing.T) {
	mods := ModifierList{
		m(ModifierCircumstance, 2, "flanking"),
		m(ModifierCircumstance, -5, "flanking"),
	}

	// current logic: picks highest (2)
	got := mods.ModifierResolve()
	want := 2

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// ==============================
// MANY MODIFIERS (STRESS)
// ==============================

func TestModifierResolve_Stress(t *testing.T) {
	var mods ModifierList

	for i := 0; i < 1000; i++ {
		mods = append(mods, m(ModifierDodge, 1, "a"))
	}

	got := mods.ModifierResolve()
	want := 1000

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}
