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
	modifier, ok := NewModifier(t, v, src, mockTarget{}, nil)
	if !ok {
		panic("invalid test modifier")
	}

	return modifier
}

func mustResolve(mods ModifierList) int {
	total, ok := mods.ModifierResolve()
	if !ok {
		panic("expected modifier list to resolve")
	}

	return total
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
	want := 4

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestModifierResolve_RejectsEmptyType(t *testing.T) {
	mods := ModifierList{
		{value: -1},
		{value: -5},
	}

	if _, ok := mods.ModifierResolve(); ok {
		t.Fatal("expected empty modifier type to be rejected")
	}
}

func TestModifierResolve_Negatives_Default(t *testing.T) {
	mods := ModifierList{
		m("enhancement", -1, "a"),
		m("enhancement", -5, "b"),
		m("enhancement", -3, "c"),
	}

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
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

	got := mustResolve(mods)
	want := -5

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

// ==============================
// VALIDATION
// ==============================

func TestNewModifier_RejectsUnknownType(t *testing.T) {
	if _, ok := NewModifier(ModifierType("typo"), 1, "", mockTarget{}, nil); ok {
		t.Fatal("expected unknown modifier type to be rejected")
	}
}

func TestNewModifier_RejectsInvalidCircumstanceSource(t *testing.T) {
	if _, ok := NewModifier(ModifierCircumstance, 1, "flanking", mockTarget{}, nil); ok {
		t.Fatal("expected invalid circumstance source to be rejected")
	}
}

func TestModifierResolve_RejectsUnknownModifierEntry(t *testing.T) {
	mods := ModifierList{
		{modifierType: ModifierType("typo"), value: 1},
	}

	if _, ok := mods.ModifierResolve(); ok {
		t.Fatal("expected invalid modifier entry to be rejected")
	}
}

func TestNewModifier_DefensivelyCopiesConditions(t *testing.T) {
	condition := ModifierCondition{mockCondition{}}

	modifier, ok := NewModifier(ModifierDodge, 1, "", mockTarget{}, condition)
	if !ok {
		t.Fatal("expected modifier to be constructed")
	}

	condition[0] = nil

	gotCondition := modifier.GetCondition()
	if len(gotCondition) != 1 || gotCondition[0] == nil {
		t.Fatal("expected modifier to keep an internal copy of conditions")
	}

	gotCondition[0] = nil
	if modifier.GetCondition()[0] == nil {
		t.Fatal("expected condition getter to return a defensive copy")
	}
}

func TestNewModifier_ExposesStoredFields(t *testing.T) {
	modifier, ok := NewModifier(ModifierDodge, 2, "", mockTarget{}, nil)
	if !ok {
		t.Fatal("expected modifier to be constructed")
	}

	if modifier.GetType() != ModifierDodge {
		t.Fatalf("expected type %q, got %q", ModifierDodge, modifier.GetType())
	}

	if modifier.GetValue() != 2 {
		t.Fatalf("expected value 2, got %d", modifier.GetValue())
	}

	if modifier.GetTarget() == nil {
		t.Fatal("expected target to be preserved")
	}

	if modifier.GetSource() != "" {
		t.Fatalf("expected empty source, got %q", modifier.GetSource())
	}

	if len(modifier.GetCondition()) != 0 {
		t.Fatalf("expected empty condition list, got %d entries", len(modifier.GetCondition()))
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

	got := mustResolve(mods)
	want := 1000

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}
