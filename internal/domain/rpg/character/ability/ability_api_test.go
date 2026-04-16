package ability_test

import (
	"testing"

	"d20campaing/internal/domain/rpg/character/ability"
)

func TestExportedAbilityTypesAreUsableOutsidePackage(t *testing.T) {
	value, ok := ability.NewAbilityScoreValue(18, true)
	if !ok {
		t.Fatal("expected ability score value to be constructed")
	}
	score, ok := ability.NewAbilityScore(ability.StrengthScore, value)
	if !ok {
		t.Fatal("expected ability score to be constructed")
	}
	bab, ok := ability.NewBaseAttackBonusByClassLevel(2, ability.BaseAttackBonusThreeQuarters)
	if !ok {
		t.Fatal("expected BAB to be constructed")
	}
	save, ok := ability.NewSavingThrowByClassLevel(ability.FortitudeSave, 1, ability.SavingThrowGood)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}
	casterLevel, ok := ability.NewCasterLevel(5, 0, 0)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}
	hd, ok := ability.NewHitDie(0, 1, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := ability.NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}
	creatureSize := ability.HugeSize

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
	arcaneCasterLevel, ok := exportedCasterLevel.GetSourceLevel(ability.ArcaneCasterSource)
	if !ok {
		t.Fatal("expected exported arcane caster level to be available")
	}

	if arcaneCasterLevel != 5 {
		t.Fatalf("expected exported arcane caster level 5, got %d", arcaneCasterLevel)
	}

	var exportedHD ability.HitDie = hd
	if exportedHD.GetAverageBaseHP() != 5 {
		t.Fatalf("expected exported average base HP 5, got %d", exportedHD.GetAverageBaseHP())
	}

	var exportedHP ability.HitPoints = hp
	if exportedHP.GetTotal() != 6 {
		t.Fatalf("expected exported HP total 6, got %d", exportedHP.GetTotal())
	}

	var exportedSize ability.Size = creatureSize
	sizeModifier, ok := exportedSize.GetModifier()
	if !ok {
		t.Fatal("expected exported size modifier to be available")
	}

	if sizeModifier != -2 {
		t.Fatalf("expected exported size modifier -2, got %d", sizeModifier)
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
