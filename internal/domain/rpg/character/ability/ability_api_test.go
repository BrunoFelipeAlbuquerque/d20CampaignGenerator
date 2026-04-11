package ability_test

import (
	"testing"

	"d20campaing/internal/domain/rpg/character/ability"
)

func TestExportedAbilityTypesAreUsableOutsidePackage(t *testing.T) {
	value := ability.NewAbilityScoreValue(18, true)
	score := ability.NewAbilityScore(ability.StrengthScore, value)

	var exportedID ability.AbilityScoreID = score.GetID()
	if exportedID != ability.StrengthScore {
		t.Fatalf("expected exported id %q, got %q", ability.StrengthScore, exportedID)
	}

	var exportedValue ability.AbilityScoreValue = score.GetValue()
	rawValue, valid := exportedValue.GetValue()
	if rawValue != 18 || !valid {
		t.Fatalf("expected exported value (18, true), got (%d, %t)", rawValue, valid)
	}

	var exportedScore ability.AbilityScore = score
	capacity, ok := exportedScore.GetCarryingCapacity()
	if !ok {
		t.Fatal("expected carrying capacity from exported score")
	}

	var exportedCapacity ability.StrengthCarryingCapacity = capacity
	if exportedCapacity.GetLightLoadMax().GetKilograms() != 50 {
		t.Fatalf("expected light load 50kg, got %.1fkg", exportedCapacity.GetLightLoadMax().GetKilograms())
	}

	profile, ok := exportedScore.GetSpellcastingProfile()
	if !ok {
		t.Fatal("expected spellcasting profile from exported score")
	}

	var exportedProfile ability.SpellcastingAbilityProfile = profile
	if exportedProfile.GetBonusSpells(1) != 1 {
		t.Fatalf("expected 1 bonus first-level spell, got %d", exportedProfile.GetBonusSpells(1))
	}
}
