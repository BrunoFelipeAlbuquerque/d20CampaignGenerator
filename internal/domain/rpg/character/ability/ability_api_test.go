package ability_test

import (
	"testing"

	"d20campaing/internal/domain/rpg/character/ability"
)

func TestExportedAbilityTypesAreUsableOutsidePackage(t *testing.T) {
	value := ability.NewAbilityScoreValue(18, true)
	score := ability.NewAbilityScore(ability.StrengthScore, value)
	bab := ability.NewBaseAttackBonusByClassLevel(2, ability.BaseAttackBonusThreeQuarters)
	save := ability.NewSavingThrowByClassLevel(ability.FortitudeSave, 1, ability.SavingThrowGood)
	casterLevel := ability.NewCasterLevel(5, 0, 0)

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

	var exportedBAB ability.BaseAttackBonus = bab
	if exportedBAB.GetValue() != 1 {
		t.Fatalf("expected exported BAB 1, got %d", exportedBAB.GetValue())
	}

	var exportedSave ability.SavingThrow = save
	if exportedSave.GetValue() != 2 {
		t.Fatalf("expected exported saving throw 2, got %d", exportedSave.GetValue())
	}

	var exportedCasterLevel ability.CasterLevel = casterLevel
	arcaneCasterLevel, ok := exportedCasterLevel.GetArcane()
	if !ok {
		t.Fatal("expected exported arcane caster level to be available")
	}

	if arcaneCasterLevel != 5 {
		t.Fatalf("expected exported arcane caster level 5, got %d", arcaneCasterLevel)
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
