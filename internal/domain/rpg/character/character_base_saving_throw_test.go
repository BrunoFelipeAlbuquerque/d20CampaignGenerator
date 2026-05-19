package character

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

func TestNewCharacterBaseSavingThrowFacts_ComposesFractionalMulticlassSaves(t *testing.T) {
	facts, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.ClericClassID, 5),
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
	})
	if !ok {
		t.Fatal("expected class levels to compose into base saving throw facts")
	}

	assertCharacterBaseSavingThrowFacts(t, facts, ability.FortitudeSave, 11, 2, 5)
	assertCharacterBaseSavingThrowFacts(t, facts, ability.ReflexSave, 7, 3, 2)
	assertCharacterBaseSavingThrowFacts(t, facts, ability.WillSave, 31, 6, 5)
}

func TestNewCharacterBaseSavingThrowFacts_ComposesDistinctGoodSaveTypes(t *testing.T) {
	facts, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
		mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
	})
	if !ok {
		t.Fatal("expected class levels to compose into base saving throw facts")
	}

	assertCharacterBaseSavingThrowFacts(t, facts, ability.FortitudeSave, 10, 3, 3)
	assertCharacterBaseSavingThrowFacts(t, facts, ability.ReflexSave, 1, 1, 1)
	assertCharacterBaseSavingThrowFacts(t, facts, ability.WillSave, 19, 6, 3)
}

func TestCharacterBaseSavingThrowFacts_ExposesSpecificSaveGetters(t *testing.T) {
	facts, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.RogueClassID, 3),
	})
	if !ok {
		t.Fatal("expected class levels to compose into base saving throw facts")
	}

	fortitude, ok := facts.GetFortitude()
	if !ok || fortitude.GetID() != ability.FortitudeSave {
		t.Fatalf("expected Fortitude save lookup to resolve, got %q and %t", fortitude.GetID(), ok)
	}

	reflex, ok := facts.GetReflex()
	if !ok || reflex.GetID() != ability.ReflexSave {
		t.Fatalf("expected Reflex save lookup to resolve, got %q and %t", reflex.GetID(), ok)
	}

	will, ok := facts.GetWill()
	if !ok || will.GetID() != ability.WillSave {
		t.Fatalf("expected Will save lookup to resolve, got %q and %t", will.GetID(), ok)
	}
}

func TestCharacterBaseSavingThrowFacts_RejectsInvalidLookupAndZeroValue(t *testing.T) {
	facts, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1),
	})
	if !ok {
		t.Fatal("expected class levels to compose into base saving throw facts")
	}

	if _, ok := facts.GetSavingThrow(ability.SavingThrowID("Luck")); ok {
		t.Fatal("expected invalid save lookup to fail")
	}

	var zeroFacts CharacterBaseSavingThrowFacts
	if _, ok := zeroFacts.GetFortitude(); ok {
		t.Fatal("expected zero-value facts not to resolve")
	}
}

func TestNewCharacterBaseSavingThrowFacts_RejectsMissingClassLevels(t *testing.T) {
	if _, ok := NewCharacterBaseSavingThrowFacts(nil); ok {
		t.Fatal("expected missing class levels to be rejected")
	}

	if _, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{}); ok {
		t.Fatal("expected empty class levels to be rejected")
	}
}

func TestNewCharacterBaseSavingThrowFacts_RejectsDuplicateClassLevels(t *testing.T) {
	fighterLevel := mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)

	if _, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{fighterLevel, fighterLevel}); ok {
		t.Fatal("expected duplicate class levels to be rejected")
	}
}

func TestNewCharacterBaseSavingThrowFacts_RejectsMalformedClassLevels(t *testing.T) {
	if _, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{
		{classID: characterclass.FighterClassID, level: 0},
	}); ok {
		t.Fatal("expected malformed zero class level fact to be rejected")
	}

	if _, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{
		{classID: characterclass.ClassID("alchemist"), level: 1},
	}); ok {
		t.Fatal("expected non-core class level fact to be rejected")
	}
}

func TestNewCharacterBaseSavingThrowFacts_RejectsTotalCharacterLevelAboveCoreRange(t *testing.T) {
	if _, ok := NewCharacterBaseSavingThrowFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 20),
		mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
	}); ok {
		t.Fatal("expected total character level above the core range to be rejected")
	}
}

func assertCharacterBaseSavingThrowFacts(
	t *testing.T,
	facts CharacterBaseSavingThrowFacts,
	id ability.SavingThrowID,
	expectedNumerator int,
	expectedDenominator int,
	expectedValue int,
) {
	t.Helper()

	save, ok := facts.GetSavingThrow(id)
	if !ok {
		t.Fatalf("expected save %q to resolve", id)
	}

	if save.GetActualValue().GetNumerator() != expectedNumerator ||
		save.GetActualValue().GetDenominator() != expectedDenominator ||
		save.GetValue() != expectedValue {
		t.Fatalf(
			"expected %q save %d/%d and value %d, got %d/%d and value %d",
			id,
			expectedNumerator,
			expectedDenominator,
			expectedValue,
			save.GetActualValue().GetNumerator(),
			save.GetActualValue().GetDenominator(),
			save.GetValue(),
		)
	}
}
