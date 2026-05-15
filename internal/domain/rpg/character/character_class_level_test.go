package character

import (
	"testing"

	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

func TestNewCharacterLevelFacts_ComposesCoreClassLevels(t *testing.T) {
	facts, ok := NewCharacterLevelFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 5),
	})
	if !ok {
		t.Fatal("expected core class levels to compose")
	}

	if facts.GetTotalCharacterLevel() != 6 {
		t.Fatalf("expected total character level 6, got %d", facts.GetTotalCharacterLevel())
	}

	fighterLevel, ok := facts.GetClassLevel(characterclass.FighterClassID)
	if !ok || fighterLevel != 5 {
		t.Fatalf("expected fighter level (5, true), got (%d, %t)", fighterLevel, ok)
	}

	wizardLevel, ok := facts.GetClassLevel(characterclass.WizardClassID)
	if !ok || wizardLevel != 1 {
		t.Fatalf("expected wizard level (1, true), got (%d, %t)", wizardLevel, ok)
	}

	if _, ok := facts.GetClassLevel(characterclass.RogueClassID); ok {
		t.Fatal("expected unselected rogue class level lookup to fail")
	}
}

func TestNewCharacterLevelFacts_ExposesDefensiveClassLevelCopy(t *testing.T) {
	facts, ok := NewCharacterLevelFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
	})
	if !ok {
		t.Fatal("expected core class levels to compose")
	}

	classLevels := facts.GetClassLevels()
	if len(classLevels) != 1 {
		t.Fatalf("expected one class level fact, got %d", len(classLevels))
	}
	classLevels[0] = mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1)

	fighterLevel, ok := facts.GetClassLevel(characterclass.FighterClassID)
	if !ok || fighterLevel != 2 {
		t.Fatalf("expected fighter level to remain (2, true), got (%d, %t)", fighterLevel, ok)
	}

	if _, ok := facts.GetClassLevel(characterclass.WizardClassID); ok {
		t.Fatal("expected copied wizard level mutation not to affect facts")
	}
}

func TestNewCharacterLevelFacts_RejectsMissingClassLevels(t *testing.T) {
	if _, ok := NewCharacterLevelFacts(nil); ok {
		t.Fatal("expected missing class levels to be rejected")
	}

	if _, ok := NewCharacterLevelFacts([]CharacterClassLevel{}); ok {
		t.Fatal("expected empty class levels to be rejected")
	}
}

func TestNewCharacterLevelFacts_RejectsDuplicateClassLevels(t *testing.T) {
	fighterLevel := mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)

	if _, ok := NewCharacterLevelFacts([]CharacterClassLevel{fighterLevel, fighterLevel}); ok {
		t.Fatal("expected duplicate class levels to be rejected")
	}
}

func TestNewCharacterLevelFacts_RejectsMalformedClassLevels(t *testing.T) {
	if _, ok := NewCharacterClassLevel(characterclass.FighterClassID, 0); ok {
		t.Fatal("expected zero class level constructor input to be rejected")
	}

	if _, ok := NewCharacterClassLevel(characterclass.FighterClassID, 21); ok {
		t.Fatal("expected class level above the core range to be rejected")
	}

	if _, ok := NewCharacterLevelFacts([]CharacterClassLevel{{classID: characterclass.FighterClassID, level: 0}}); ok {
		t.Fatal("expected malformed zero class level fact to be rejected")
	}
}

func TestNewCharacterLevelFacts_RejectsNonCoreClassLevels(t *testing.T) {
	if _, ok := NewCharacterClassLevel(characterclass.ClassID("alchemist"), 1); ok {
		t.Fatal("expected non-core class level constructor input to be rejected")
	}

	if _, ok := NewCharacterLevelFacts([]CharacterClassLevel{{classID: characterclass.ClassID("alchemist"), level: 1}}); ok {
		t.Fatal("expected non-core class level fact to be rejected")
	}
}

func TestNewCharacterLevelFacts_RejectsTotalCharacterLevelAboveCoreRange(t *testing.T) {
	if _, ok := NewCharacterLevelFacts([]CharacterClassLevel{
		mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 20),
		mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
	}); ok {
		t.Fatal("expected total character level above the core range to be rejected")
	}
}
