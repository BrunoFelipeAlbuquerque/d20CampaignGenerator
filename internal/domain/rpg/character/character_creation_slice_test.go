package character

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterfeat "d20campaigngenerator/internal/domain/rpg/character/feat"
	characterrace "d20campaigngenerator/internal/domain/rpg/character/race"
)

func TestMinimumLevelOneCharacterCreationSlice_ComposesCoreRaceClassHPSpellcastingAndFeatPrerequisites(t *testing.T) {
	selectedRace := mustNewCharacterRaceForSliceTest(t, characterrace.HumanRaceID)
	race, ok := selectedRace.GetRace()
	if !ok {
		t.Fatal("expected selected race to resolve")
	}

	if race.GetSize() != ability.MediumSize {
		t.Fatalf("expected human size %q, got %q", ability.MediumSize, race.GetSize())
	}

	selectableModifier, ok := race.GetSelectableAbilityScoreModifier()
	if !ok || selectableModifier != 2 {
		t.Fatalf("expected human selectable ability modifier (2, true), got (%d, %t)", selectableModifier, ok)
	}

	selectedClass := mustNewCharacterClassForTest(t, characterclass.WizardClassID)
	class, ok := selectedClass.GetClass()
	if !ok {
		t.Fatal("expected selected class to resolve")
	}

	if class.GetHitDieType() != ability.D6HitDie {
		t.Fatalf("expected wizard hit die %q, got %q", ability.D6HitDie, class.GetHitDieType())
	}

	hp, ok := NewFirstLevelCharacterHitPoints(selectedClass, 14)
	if !ok {
		t.Fatal("expected first-level character hit points to compose")
	}

	if hp.GetTotal() != 8 || hp.GetCurrent() != 8 {
		t.Fatalf("expected first-level wizard total/current HP 8, got total %d current %d", hp.GetTotal(), hp.GetCurrent())
	}

	d6Count, ok := hp.GetHitDie().GetDieCount(ability.D6HitDie)
	if !ok || d6Count != 1 {
		t.Fatalf("expected first-level wizard hit die (1, true), got (%d, %t)", d6Count, ok)
	}

	progression, ok := NewCharacterSpellcastingProgression(selectedClass)
	if !ok {
		t.Fatal("expected wizard spellcasting progression to compose")
	}

	cantrips, ok := progression.GetSpellSlots(1, 0)
	if !ok || cantrips != 3 {
		t.Fatalf("expected wizard level 1 cantrip slots (3, true), got (%d, %t)", cantrips, ok)
	}

	firstLevelSlots, ok := progression.GetSpellSlots(1, 1)
	if !ok || firstLevelSlots != 1 {
		t.Fatalf("expected wizard level 1 spell slots (1, true), got (%d, %t)", firstLevelSlots, ok)
	}

	if _, ok := progression.GetSpellSlots(1, 2); ok {
		t.Fatal("expected wizard level 1 spell level 2 slots to be unavailable")
	}

	prerequisiteState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1)},
		nil,
		nil,
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.ArcaneStrikeFeatID, prerequisiteState); !ok {
		t.Fatal("expected arcane strike to compose from level-1 wizard spellcasting")
	}

	if _, ok := NewCharacterFeat(characterfeat.PowerAttackFeatID, prerequisiteState); ok {
		t.Fatal("expected power attack to fail without strength and base attack prerequisites")
	}
}

func TestMinimumLevelOneCharacterCreationSlice_InvalidSelectedInputsFailClosed(t *testing.T) {
	if _, ok := NewCharacterRace(characterrace.RaceID("android")); ok {
		t.Fatal("expected unknown race to fail")
	}

	if _, ok := NewCharacterClass(characterclass.ClassID("alchemist")); ok {
		t.Fatal("expected unknown class to fail")
	}

	var zeroClass CharacterClass
	if _, ok := NewFirstLevelCharacterHitPoints(zeroClass, 14); ok {
		t.Fatal("expected zero-value class hit points to fail")
	}

	selectedClass := mustNewCharacterClassForTest(t, characterclass.WizardClassID)
	if _, ok := NewFirstLevelCharacterHitPoints(selectedClass, 0); ok {
		t.Fatal("expected invalid constitution hit points to fail")
	}

	fighter := mustNewCharacterClassForTest(t, characterclass.FighterClassID)
	if _, ok := NewCharacterSpellcastingProgression(fighter); ok {
		t.Fatal("expected non-spellcasting class progression to fail")
	}

	var prerequisiteState CharacterFeatPrerequisiteState
	if _, ok := NewCharacterFeat(characterfeat.ArcaneStrikeFeatID, prerequisiteState); ok {
		t.Fatal("expected zero-value feat prerequisite state to fail")
	}
}

func mustNewCharacterRaceForSliceTest(
	t *testing.T,
	id characterrace.RaceID,
) CharacterRace {
	t.Helper()

	race, ok := NewCharacterRace(id)
	if !ok {
		t.Fatalf("expected character race %q to compose", id)
	}

	return race
}
