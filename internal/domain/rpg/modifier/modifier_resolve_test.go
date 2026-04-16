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

func m(t ModifierType, v int, src ModifierSource) Modifier {
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
		m(ModifierCircumstance, 2, SourceFlanking),
		m(ModifierCircumstance, 1, SourceHigherGround),
	}

	got := mods.ModifierResolve()
	want := 3

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_CircumstanceSameSourceTakeHighest(t *testing.T) {
	mods := ModifierList{
		m(ModifierCircumstance, 1, SourceFlanking),
		m(ModifierCircumstance, 3, SourceFlanking),
		m(ModifierCircumstance, 2, SourceFlanking),
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

		m(ModifierCircumstance, 2, SourceFlanking),
		m(ModifierCircumstance, 1, SourceHigherGround),
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

func TestModifierResolve_UntypedPenaltiesStack(t *testing.T) {
	mods := ModifierList{
		m(ModifierUntyped, -1, "condition.fatigued"),
		m(ModifierUntyped, -5, "condition.exhausted"),
		m(ModifierUntyped, -3, "environment.hampered"),
	}

	got := mods.ModifierResolve()
	want := -9

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_UntypedValuesStackArithmetically(t *testing.T) {
	mods := ModifierList{
		m(ModifierUntyped, 2, "custom.blessing"),
		m(ModifierUntyped, -1, "condition.shaken"),
		m(ModifierUntyped, 3, "custom.momentum"),
	}

	got := mods.ModifierResolve()
	want := 4

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_EmptyTypeDoesNotBehaveAsUntyped(t *testing.T) {
	mods := ModifierList{
		m("", -1, "condition.fatigued"),
		m("", -5, "condition.exhausted"),
	}

	got := mods.ModifierResolve()
	want := -1

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_Negatives_Default(t *testing.T) {
	mods := ModifierList{
		m("enhancement", -1, "a"),
		m("enhancement", -5, "b"),
		m("enhancement", -3, "c"),
	}

	got := mods.ModifierResolve()
	want := -9

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_DefaultMixedBonusesAndPenalties(t *testing.T) {
	mods := ModifierList{
		m("enhancement", 5, "a"),
		m("enhancement", 3, "b"),
		m("enhancement", -2, "c"),
		m("enhancement", -4, "d"),
	}

	got := mods.ModifierResolve()
	want := -1 // highest enhancement bonus (5) plus stacking penalties (-6)

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
		m(ModifierCircumstance, 2, SourceFlanking),
		m(ModifierCircumstance, -5, SourceFlanking),
	}

	got := mods.ModifierResolve()
	want := -3

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_CircumstancePenaltiesDifferentSourcesStack(t *testing.T) {
	mods := ModifierList{
		m(ModifierCircumstance, -2, SourceFlanking),
		m(ModifierCircumstance, -1, SourceHigherGround),
	}

	got := mods.ModifierResolve()
	want := -3

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_CircumstanceSameSourceKeepsWorstPenalty(t *testing.T) {
	mods := ModifierList{
		m(ModifierCircumstance, -2, SourceFlanking),
		m(ModifierCircumstance, -5, SourceFlanking),
		m(ModifierCircumstance, -1, SourceFlanking),
	}

	got := mods.ModifierResolve()
	want := -5

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
