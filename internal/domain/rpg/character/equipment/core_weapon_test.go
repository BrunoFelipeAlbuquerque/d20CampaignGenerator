package equipment

import "testing"

func TestCoreSimpleWeapons_SeedsBatchOne(t *testing.T) {
	expected := map[WeaponID]struct {
		displayName      string
		category         WeaponCategory
		copperPieces     int
		ounces           int
		smallDiceCount   int
		smallDieSides    int
		mediumDiceCount  int
		mediumDieSides   int
		threatMinimum    int
		multiplier       int
		rangeIncrement   int
		secondaryProfile bool
	}{
		GauntletWeaponID: {
			displayName:     "Gauntlet",
			category:        UnarmedAttackWeaponCategory,
			copperPieces:    200,
			ounces:          16,
			smallDiceCount:  1,
			smallDieSides:   2,
			mediumDiceCount: 1,
			mediumDieSides:  3,
			threatMinimum:   20,
			multiplier:      2,
		},
		UnarmedStrikeWeaponID: {
			displayName:     "Unarmed strike",
			category:        UnarmedAttackWeaponCategory,
			smallDiceCount:  1,
			smallDieSides:   2,
			mediumDiceCount: 1,
			mediumDieSides:  3,
			threatMinimum:   20,
			multiplier:      2,
		},
		DaggerWeaponID: {
			displayName:     "Dagger",
			category:        LightMeleeWeaponCategory,
			copperPieces:    200,
			ounces:          16,
			smallDiceCount:  1,
			smallDieSides:   3,
			mediumDiceCount: 1,
			mediumDieSides:  4,
			threatMinimum:   19,
			multiplier:      2,
			rangeIncrement:  10,
		},
		MaceLightWeaponID: {
			displayName:     "Mace, light",
			category:        LightMeleeWeaponCategory,
			copperPieces:    500,
			ounces:          64,
			smallDiceCount:  1,
			smallDieSides:   4,
			mediumDiceCount: 1,
			mediumDieSides:  6,
			threatMinimum:   20,
			multiplier:      2,
		},
		SickleWeaponID: {
			displayName:     "Sickle",
			category:        LightMeleeWeaponCategory,
			copperPieces:    600,
			ounces:          32,
			smallDiceCount:  1,
			smallDieSides:   4,
			mediumDiceCount: 1,
			mediumDieSides:  6,
			threatMinimum:   20,
			multiplier:      2,
		},
		ClubWeaponID: {
			displayName:     "Club",
			category:        OneHandedMeleeWeaponCategory,
			ounces:          48,
			smallDiceCount:  1,
			smallDieSides:   4,
			mediumDiceCount: 1,
			mediumDieSides:  6,
			threatMinimum:   20,
			multiplier:      2,
			rangeIncrement:  10,
		},
		MaceHeavyWeaponID: {
			displayName:     "Mace, heavy",
			category:        OneHandedMeleeWeaponCategory,
			copperPieces:    1200,
			ounces:          128,
			smallDiceCount:  1,
			smallDieSides:   6,
			mediumDiceCount: 1,
			mediumDieSides:  8,
			threatMinimum:   20,
			multiplier:      2,
		},
		MorningstarWeaponID: {
			displayName:     "Morningstar",
			category:        OneHandedMeleeWeaponCategory,
			copperPieces:    800,
			ounces:          96,
			smallDiceCount:  1,
			smallDieSides:   6,
			mediumDiceCount: 1,
			mediumDieSides:  8,
			threatMinimum:   20,
			multiplier:      2,
		},
		ShortspearWeaponID: {
			displayName:     "Shortspear",
			category:        OneHandedMeleeWeaponCategory,
			copperPieces:    100,
			ounces:          48,
			smallDiceCount:  1,
			smallDieSides:   4,
			mediumDiceCount: 1,
			mediumDieSides:  6,
			threatMinimum:   20,
			multiplier:      2,
			rangeIncrement:  20,
		},
		LongspearWeaponID: {
			displayName:     "Longspear",
			category:        TwoHandedMeleeWeaponCategory,
			copperPieces:    500,
			ounces:          144,
			smallDiceCount:  1,
			smallDieSides:   6,
			mediumDiceCount: 1,
			mediumDieSides:  8,
			threatMinimum:   20,
			multiplier:      3,
		},
		QuarterstaffWeaponID: {
			displayName:      "Quarterstaff",
			category:         TwoHandedMeleeWeaponCategory,
			ounces:           64,
			smallDiceCount:   1,
			smallDieSides:    4,
			mediumDiceCount:  1,
			mediumDieSides:   6,
			threatMinimum:    20,
			multiplier:       2,
			secondaryProfile: true,
		},
		SpearWeaponID: {
			displayName:     "Spear",
			category:        TwoHandedMeleeWeaponCategory,
			copperPieces:    200,
			ounces:          96,
			smallDiceCount:  1,
			smallDieSides:   6,
			mediumDiceCount: 1,
			mediumDieSides:  8,
			threatMinimum:   20,
			multiplier:      3,
			rangeIncrement:  20,
		},
		CrossbowHeavyWeaponID: {
			displayName:     "Crossbow, heavy",
			category:        RangedWeaponCategory,
			copperPieces:    5000,
			ounces:          128,
			smallDiceCount:  1,
			smallDieSides:   8,
			mediumDiceCount: 1,
			mediumDieSides:  10,
			threatMinimum:   19,
			multiplier:      2,
			rangeIncrement:  120,
		},
		CrossbowLightWeaponID: {
			displayName:     "Crossbow, light",
			category:        RangedWeaponCategory,
			copperPieces:    3500,
			ounces:          64,
			smallDiceCount:  1,
			smallDieSides:   6,
			mediumDiceCount: 1,
			mediumDieSides:  8,
			threatMinimum:   19,
			multiplier:      2,
			rangeIncrement:  80,
		},
		DartWeaponID: {
			displayName:     "Dart",
			category:        RangedWeaponCategory,
			copperPieces:    50,
			ounces:          8,
			smallDiceCount:  1,
			smallDieSides:   3,
			mediumDiceCount: 1,
			mediumDieSides:  4,
			threatMinimum:   20,
			multiplier:      2,
			rangeIncrement:  20,
		},
		JavelinWeaponID: {
			displayName:     "Javelin",
			category:        RangedWeaponCategory,
			copperPieces:    100,
			ounces:          32,
			smallDiceCount:  1,
			smallDieSides:   4,
			mediumDiceCount: 1,
			mediumDieSides:  6,
			threatMinimum:   20,
			multiplier:      2,
			rangeIncrement:  30,
		},
		SlingWeaponID: {
			displayName:     "Sling",
			category:        RangedWeaponCategory,
			smallDiceCount:  1,
			smallDieSides:   3,
			mediumDiceCount: 1,
			mediumDieSides:  4,
			threatMinimum:   20,
			multiplier:      2,
			rangeIncrement:  50,
		},
	}

	if len(coreSimpleWeapons) != len(expected) {
		t.Fatalf("expected %d core simple weapon seeds, got %d", len(expected), len(coreSimpleWeapons))
	}

	for id, expectation := range expected {
		weapon, ok := coreSimpleWeapons[id]
		if !ok {
			t.Fatalf("expected core simple weapon seed %q", id)
		}

		if weapon.GetID() != id {
			t.Fatalf("expected weapon id %q, got %q", id, weapon.GetID())
		}

		if weapon.GetDisplayName() != expectation.displayName {
			t.Fatalf("expected %q display name %q, got %q", id, expectation.displayName, weapon.GetDisplayName())
		}

		if weapon.GetProficiencyCategory() != SimpleWeaponProficiencyCategory {
			t.Fatalf("expected %q proficiency category %q, got %q", id, SimpleWeaponProficiencyCategory, weapon.GetProficiencyCategory())
		}

		if weapon.GetCategory() != expectation.category {
			t.Fatalf("expected %q category %q, got %q", id, expectation.category, weapon.GetCategory())
		}

		if weapon.GetCost().GetCopperPieces() != expectation.copperPieces {
			t.Fatalf("expected %q cost %d cp, got %d cp", id, expectation.copperPieces, weapon.GetCost().GetCopperPieces())
		}

		if weapon.GetWeight().GetOunces() != expectation.ounces {
			t.Fatalf("expected %q weight %d oz, got %d oz", id, expectation.ounces, weapon.GetWeight().GetOunces())
		}

		assertCoreWeaponDamage(
			t,
			id,
			weapon.GetDamageProfile().GetSmallDamage(),
			expectation.smallDiceCount,
			expectation.smallDieSides,
			"small",
		)
		assertCoreWeaponDamage(
			t,
			id,
			weapon.GetDamageProfile().GetMediumDamage(),
			expectation.mediumDiceCount,
			expectation.mediumDieSides,
			"medium",
		)

		if weapon.GetDamageProfile().HasSecondaryDamage() != expectation.secondaryProfile {
			t.Fatalf("expected %q secondary damage profile %t", id, expectation.secondaryProfile)
		}

		if expectation.secondaryProfile {
			assertCoreWeaponDamage(
				t,
				id,
				weapon.GetDamageProfile().GetSecondarySmallDamage(),
				expectation.smallDiceCount,
				expectation.smallDieSides,
				"secondary small",
			)
			assertCoreWeaponDamage(
				t,
				id,
				weapon.GetDamageProfile().GetSecondaryMediumDamage(),
				expectation.mediumDiceCount,
				expectation.mediumDieSides,
				"secondary medium",
			)
		}

		if !weapon.GetCriticalProfile().HasCritical() {
			t.Fatalf("expected %q to have a critical profile", id)
		}

		if weapon.GetCriticalProfile().GetThreatMinimum() != expectation.threatMinimum {
			t.Fatalf("expected %q threat minimum %d, got %d", id, expectation.threatMinimum, weapon.GetCriticalProfile().GetThreatMinimum())
		}

		if weapon.GetCriticalProfile().GetPrimaryMultiplier() != expectation.multiplier {
			t.Fatalf("expected %q critical multiplier %d, got %d", id, expectation.multiplier, weapon.GetCriticalProfile().GetPrimaryMultiplier())
		}

		if expectation.rangeIncrement == 0 {
			if weapon.GetRangeIncrement().HasRangeIncrement() {
				t.Fatalf("expected %q not to have a range increment", id)
			}

			continue
		}

		if !weapon.GetRangeIncrement().HasRangeIncrement() {
			t.Fatalf("expected %q to have a range increment", id)
		}

		if weapon.GetRangeIncrement().GetFeet() != expectation.rangeIncrement {
			t.Fatalf("expected %q range increment %d ft, got %d ft", id, expectation.rangeIncrement, weapon.GetRangeIncrement().GetFeet())
		}
	}
}

func TestCoreSimpleWeaponOrder_ContainsOnlySeededBatchOneIDs(t *testing.T) {
	expectedOrder := []WeaponID{
		GauntletWeaponID,
		UnarmedStrikeWeaponID,
		DaggerWeaponID,
		MaceLightWeaponID,
		SickleWeaponID,
		ClubWeaponID,
		MaceHeavyWeaponID,
		MorningstarWeaponID,
		ShortspearWeaponID,
		LongspearWeaponID,
		QuarterstaffWeaponID,
		SpearWeaponID,
		CrossbowHeavyWeaponID,
		CrossbowLightWeaponID,
		DartWeaponID,
		JavelinWeaponID,
		SlingWeaponID,
	}

	if len(coreSimpleWeaponOrder) != len(expectedOrder) {
		t.Fatalf("expected order length %d, got %d", len(expectedOrder), len(coreSimpleWeaponOrder))
	}

	seen := make(map[WeaponID]struct{}, len(coreSimpleWeaponOrder))
	for index, id := range coreSimpleWeaponOrder {
		if id != expectedOrder[index] {
			t.Fatalf("expected ordered weapon id at index %d to be %q, got %q", index, expectedOrder[index], id)
		}

		if _, ok := coreSimpleWeapons[id]; !ok {
			t.Fatalf("expected ordered weapon id %q to have a seed", id)
		}

		if _, ok := seen[id]; ok {
			t.Fatalf("expected ordered weapon id %q not to be duplicated", id)
		}

		seen[id] = struct{}{}
	}
}

func assertCoreWeaponDamage(
	t *testing.T,
	id WeaponID,
	damage WeaponDamage,
	diceCount int,
	dieSides int,
	label string,
) {
	t.Helper()

	if damage.GetKind() != DiceWeaponDamageKind {
		t.Fatalf("expected %q %s damage kind %q, got %q", id, label, DiceWeaponDamageKind, damage.GetKind())
	}

	if damage.GetDiceCount() != diceCount || damage.GetDieSides() != dieSides {
		t.Fatalf("expected %q %s damage %dd%d, got %dd%d", id, label, diceCount, dieSides, damage.GetDiceCount(), damage.GetDieSides())
	}
}
