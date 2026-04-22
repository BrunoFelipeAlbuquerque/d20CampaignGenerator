package ability

import "testing"

func TestSizeGetModifiers_UsesPathfinderTable(t *testing.T) {
	largeModifier, ok := LargeSize.GetAttackAndACModifier()
	if !ok || largeModifier != -1 {
		t.Fatalf("expected large size modifier (-1, true), got (%d, %t)", largeModifier, ok)
	}

	largeSpecialModifier, ok := LargeSize.GetCMBAndCMDModifier()
	if !ok || largeSpecialModifier != 1 {
		t.Fatalf("expected large special size modifier (1, true), got (%d, %t)", largeSpecialModifier, ok)
	}

	hugeFlyModifier, ok := HugeSize.GetFlyModifier()
	if !ok || hugeFlyModifier != -4 {
		t.Fatalf("expected huge fly modifier (-4, true), got (%d, %t)", hugeFlyModifier, ok)
	}

	gargantuanStealthModifier, ok := GargantuanSize.GetStealthModifier()
	if !ok || gargantuanStealthModifier != -12 {
		t.Fatalf("expected gargantuan stealth modifier (-12, true), got (%d, %t)", gargantuanStealthModifier, ok)
	}
}

func TestSizeGetSpaceAndReach_RespectsTallVsLongBodies(t *testing.T) {
	fineSpace, ok := FineSize.GetSpace(TallBodyShape)
	if !ok || fineSpace.GetFeet() != 0.5 {
		t.Fatalf("expected fine space (0.5, true), got (%.1f, %t)", fineSpace.GetFeet(), ok)
	}

	largeTallReach, ok := LargeSize.GetNaturalReach(TallBodyShape)
	if !ok || largeTallReach.GetFeet() != 10 {
		t.Fatalf("expected large tall reach (10, true), got (%.1f, %t)", largeTallReach.GetFeet(), ok)
	}

	largeLongReach, ok := LargeSize.GetNaturalReach(LongBodyShape)
	if !ok || largeLongReach.GetFeet() != 5 {
		t.Fatalf("expected large long reach (5, true), got (%.1f, %t)", largeLongReach.GetFeet(), ok)
	}

	colossalSpace, ok := ColossalSize.GetSpace(TallBodyShape)
	if !ok || colossalSpace.GetFeet() != 30 {
		t.Fatalf("expected colossal space (30, true), got (%.1f, %t)", colossalSpace.GetFeet(), ok)
	}
}

func TestSizeGetSpaceAndReach_RejectInvalidBodyShapes(t *testing.T) {
	if _, ok := LargeSize.GetSpace(BodyShape("Blob")); ok {
		t.Fatal("expected invalid body shape to be rejected for space")
	}

	if _, ok := LargeSize.GetNaturalReach(BodyShape("Blob")); ok {
		t.Fatal("expected invalid body shape to be rejected for reach")
	}
}

func TestSizeGetTypicalRanges_ExposeImperialAndMetricValues(t *testing.T) {
	mediumHeight, ok := MediumSize.GetTypicalHeightRange()
	if !ok {
		t.Fatal("expected medium size height range to exist")
	}

	if mediumHeight.GetMin().GetFeet() != 4 || mediumHeight.GetMax().GetFeet() != 8 || !mediumHeight.HasUpperBound() {
		t.Fatalf(
			"expected medium height range 4-8ft, got %.1f-%.1fft (bounded=%t)",
			mediumHeight.GetMin().GetFeet(),
			mediumHeight.GetMax().GetFeet(),
			mediumHeight.HasUpperBound(),
		)
	}

	mediumWeight, ok := MediumSize.GetTypicalWeightRange()
	if !ok {
		t.Fatal("expected medium size weight range to exist")
	}

	if !almostEqual(mediumWeight.GetMin().GetPounds(), 60) || !almostEqual(mediumWeight.GetMax().GetPounds(), 500) || !mediumWeight.HasUpperBound() {
		t.Fatalf(
			"expected medium weight range 60-500lb, got %.2f-%.2flb (bounded=%t)",
			mediumWeight.GetMin().GetPounds(),
			mediumWeight.GetMax().GetPounds(),
			mediumWeight.HasUpperBound(),
		)
	}

	if !almostEqual(mediumHeight.GetMin().GetMeters(), 1.2192) {
		t.Fatalf("expected medium minimum height about 1.2192m, got %.4fm", mediumHeight.GetMin().GetMeters())
	}
}

func TestSizeGetConstructBonusHP_UsesCoreTableAndProjectTitanic(t *testing.T) {
	colossalBonus, ok := ColossalSize.GetConstructBonusHP()
	if !ok || colossalBonus != 80 {
		t.Fatalf("expected colossal construct bonus (80, true), got (%d, %t)", colossalBonus, ok)
	}

	titanicBonus, ok := TitanicSize.GetConstructBonusHP()
	if !ok || titanicBonus != 210 {
		t.Fatalf("expected titanic construct bonus (210, true), got (%d, %t)", titanicBonus, ok)
	}
}

func TestTitanicSize_UsesOfficializedHomebrewProfile(t *testing.T) {
	modifier, _ := TitanicSize.GetAttackAndACModifier()
	specialModifier, _ := TitanicSize.GetCMBAndCMDModifier()
	reach, _ := TitanicSize.GetNaturalReach(LongBodyShape)
	heightRange, _ := TitanicSize.GetTypicalHeightRange()
	weightRange, _ := TitanicSize.GetTypicalWeightRange()

	if modifier != -16 || specialModifier != 12 {
		t.Fatalf("expected titanic modifiers (-16, +12), got (%d, %+d)", modifier, specialModifier)
	}

	if reach.GetFeet() != 30 {
		t.Fatalf("expected titanic long reach 30ft, got %.1fft", reach.GetFeet())
	}

	if heightRange.HasUpperBound() {
		t.Fatal("expected titanic height range to be open-ended")
	}

	if heightRange.GetMin().GetFeet() != 128 {
		t.Fatalf("expected titanic minimum height 128ft, got %.1fft", heightRange.GetMin().GetFeet())
	}

	if !almostEqual(weightRange.GetMin().GetPounds(), 2000000) {
		t.Fatalf("expected titanic minimum weight about 2,000,000lb, got %.2flb", weightRange.GetMin().GetPounds())
	}
}
