package character

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterfeat "d20campaigngenerator/internal/domain/rpg/character/feat"
)

func TestNewCharacterBaseAttackBonusFacts_ComposesFractionalClassProgressions(t *testing.T) {
	facts, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
		mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
	})
	if !ok {
		t.Fatal("expected class levels to compose into BAB facts")
	}

	assertCharacterBaseAttackBonusFacts(t, facts, 5, 2, 2)

	baseAttackBonus := facts.GetBaseAttackBonus()
	if baseAttackBonus.GetActualValue().GetNumerator() != 5 ||
		baseAttackBonus.GetActualValue().GetDenominator() != 2 ||
		baseAttackBonus.GetValue() != 2 {
		t.Fatalf(
			"expected exposed BAB 5/2 and value 2, got %d/%d and value %d",
			baseAttackBonus.GetActualValue().GetNumerator(),
			baseAttackBonus.GetActualValue().GetDenominator(),
			baseAttackBonus.GetValue(),
		)
	}
}

func TestNewCharacterBaseAttackBonusFacts_ComposesThreeQuarterProgression(t *testing.T) {
	facts, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.RogueClassID, 3),
		mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
	})
	if !ok {
		t.Fatal("expected class levels to compose into BAB facts")
	}

	assertCharacterBaseAttackBonusFacts(t, facts, 11, 4, 2)
}

func TestCharacterBaseAttackBonusFacts_ProvidesRoundedValueForFeatPrerequisites(t *testing.T) {
	fighterLevel := mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)
	facts, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{fighterLevel})
	if !ok {
		t.Fatal("expected fighter level to compose into BAB facts")
	}

	state, ok := NewCharacterFeatPrerequisiteState(
		[]CharacterAbilityScore{
			mustNewCharacterAbilityScoreForTest(t, ability.StrengthScore, 13),
		},
		facts.GetValue(),
		nil,
		[]CharacterClassLevel{fighterLevel},
		nil,
		nil,
		nil,
	)
	if !ok {
		t.Fatal("expected feat prerequisite state to compose from rounded BAB facts")
	}

	powerAttack, ok := characterfeat.GetFeatByID(characterfeat.PowerAttackFeatID)
	if !ok {
		t.Fatal("expected Power Attack to be seeded")
	}

	if !state.SatisfiesFeat(powerAttack) {
		t.Fatal("expected rounded BAB facts to satisfy Power Attack prerequisite")
	}
}

func TestNewCharacterBaseAttackBonusFacts_RejectsMissingClassLevels(t *testing.T) {
	if _, ok := NewCharacterBaseAttackBonusFacts(nil); ok {
		t.Fatal("expected missing class levels to be rejected")
	}

	if _, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{}); ok {
		t.Fatal("expected empty class levels to be rejected")
	}
}

func TestNewCharacterBaseAttackBonusFacts_RejectsDuplicateClassLevels(t *testing.T) {
	fighterLevel := mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)

	if _, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{fighterLevel, fighterLevel}); ok {
		t.Fatal("expected duplicate class levels to be rejected")
	}
}

func TestNewCharacterBaseAttackBonusFacts_RejectsMalformedClassLevels(t *testing.T) {
	if _, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{
		{classID: characterclass.FighterClassID, level: 0},
	}); ok {
		t.Fatal("expected malformed zero class level fact to be rejected")
	}

	if _, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{
		{classID: characterclass.ClassID("alchemist"), level: 1},
	}); ok {
		t.Fatal("expected non-core class level fact to be rejected")
	}
}

func TestNewCharacterBaseAttackBonusFacts_RejectsTotalCharacterLevelAboveCoreRange(t *testing.T) {
	if _, ok := NewCharacterBaseAttackBonusFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 20),
		mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
	}); ok {
		t.Fatal("expected total character level above the core range to be rejected")
	}
}

func assertCharacterBaseAttackBonusFacts(
	t *testing.T,
	facts CharacterBaseAttackBonusFacts,
	expectedNumerator int,
	expectedDenominator int,
	expectedValue int,
) {
	t.Helper()

	if facts.GetActualValue().GetNumerator() != expectedNumerator ||
		facts.GetActualValue().GetDenominator() != expectedDenominator ||
		facts.GetValue() != expectedValue {
		t.Fatalf(
			"expected BAB %d/%d and value %d, got %d/%d and value %d",
			expectedNumerator,
			expectedDenominator,
			expectedValue,
			facts.GetActualValue().GetNumerator(),
			facts.GetActualValue().GetDenominator(),
			facts.GetValue(),
		)
	}
}
