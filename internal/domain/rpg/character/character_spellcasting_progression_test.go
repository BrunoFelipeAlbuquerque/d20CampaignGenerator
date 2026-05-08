package character

import (
	"testing"

	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

func TestNewCharacterSpellcastingProgression_ComposesCasterClassProgression(t *testing.T) {
	selectedClass := mustNewCharacterClassForTest(t, characterclass.WizardClassID)

	progression, ok := NewCharacterSpellcastingProgression(selectedClass)
	if !ok {
		t.Fatal("expected caster class spellcasting progression to compose")
	}

	if progression.GetClassID() != characterclass.WizardClassID {
		t.Fatalf("expected progression class id %q, got %q", characterclass.WizardClassID, progression.GetClassID())
	}

	class, ok := progression.GetClass()
	if !ok {
		t.Fatal("expected progression class to resolve")
	}

	if class.GetSpellcasting().GetKind() != characterclass.ArcanePreparedSpellcastingKind {
		t.Fatalf("expected spellcasting kind %q, got %q", characterclass.ArcanePreparedSpellcastingKind, class.GetSpellcasting().GetKind())
	}

	table, ok := progression.GetProgression()
	if !ok {
		t.Fatal("expected progression table to resolve")
	}

	if table.GetClassID() != characterclass.WizardClassID {
		t.Fatalf("expected table class id %q, got %q", characterclass.WizardClassID, table.GetClassID())
	}

	spellSlots, ok := progression.GetSpellSlots(5, 3)
	if !ok || spellSlots != 1 {
		t.Fatalf("expected wizard level 5 spell level 3 slots (1, true), got (%d, %t)", spellSlots, ok)
	}
}

func TestNewCharacterSpellcastingProgression_ComposesDelayedCasterZeroSlotUnlocks(t *testing.T) {
	selectedClass := mustNewCharacterClassForTest(t, characterclass.PaladinClassID)

	progression, ok := NewCharacterSpellcastingProgression(selectedClass)
	if !ok {
		t.Fatal("expected delayed caster progression to compose")
	}

	if _, ok := progression.GetSpellSlots(4, 0); ok {
		t.Fatal("expected paladin 0-level spell lookup to fail")
	}

	spellSlots, ok := progression.GetSpellSlots(4, 1)
	if !ok || spellSlots != 0 {
		t.Fatalf("expected paladin level 4 spell level 1 zero-slot unlock (0, true), got (%d, %t)", spellSlots, ok)
	}
}

func TestNewCharacterSpellcastingProgression_RejectsNonSpellcastingClass(t *testing.T) {
	selectedClass := mustNewCharacterClassForTest(t, characterclass.FighterClassID)

	if _, ok := NewCharacterSpellcastingProgression(selectedClass); ok {
		t.Fatal("expected non-spellcasting class progression composition to fail")
	}
}

func TestNewCharacterSpellcastingProgression_RejectsZeroValueClass(t *testing.T) {
	var selectedClass CharacterClass

	if _, ok := NewCharacterSpellcastingProgression(selectedClass); ok {
		t.Fatal("expected zero-value class progression composition to fail")
	}
}

func TestCharacterSpellcastingProgression_ZeroValueDoesNotResolve(t *testing.T) {
	var progression CharacterSpellcastingProgression

	if _, ok := progression.GetClass(); ok {
		t.Fatal("expected zero-value progression class not to resolve")
	}

	if _, ok := progression.GetProgression(); ok {
		t.Fatal("expected zero-value progression table not to resolve")
	}

	if _, ok := progression.GetSpellSlots(1, 0); ok {
		t.Fatal("expected zero-value progression spell slots not to resolve")
	}
}

func mustNewCharacterClassForTest(t *testing.T, id characterclass.ClassID) CharacterClass {
	t.Helper()

	class, ok := NewCharacterClass(id)
	if !ok {
		t.Fatalf("expected character class %q to compose", id)
	}

	return class
}
