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

func TestGetWeaponByID_ReturnsSeededCoreWeapon(t *testing.T) {
	weapon, ok := GetWeaponByID(DaggerWeaponID)
	if !ok {
		t.Fatal("expected dagger to be returned from core weapon lookup")
	}

	if weapon.GetID() != DaggerWeaponID {
		t.Fatalf("expected weapon id %q, got %q", DaggerWeaponID, weapon.GetID())
	}

	if weapon.GetDisplayName() != "Dagger" {
		t.Fatalf("expected display name %q, got %q", "Dagger", weapon.GetDisplayName())
	}

	if weapon.GetProficiencyCategory() != SimpleWeaponProficiencyCategory {
		t.Fatalf("expected proficiency category %q, got %q", SimpleWeaponProficiencyCategory, weapon.GetProficiencyCategory())
	}

	if weapon.GetCategory() != LightMeleeWeaponCategory {
		t.Fatalf("expected category %q, got %q", LightMeleeWeaponCategory, weapon.GetCategory())
	}

	if weapon.GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected dagger cost 200 cp, got %d cp", weapon.GetCost().GetCopperPieces())
	}

	if weapon.GetWeight().GetOunces() != 16 {
		t.Fatalf("expected dagger weight 16 oz, got %d oz", weapon.GetWeight().GetOunces())
	}

	assertCoreWeaponDamage(t, DaggerWeaponID, weapon.GetDamageProfile().GetSmallDamage(), 1, 3, "small")
	assertCoreWeaponDamage(t, DaggerWeaponID, weapon.GetDamageProfile().GetMediumDamage(), 1, 4, "medium")

	if weapon.GetCriticalProfile().GetThreatMinimum() != 19 {
		t.Fatalf("expected dagger threat minimum 19, got %d", weapon.GetCriticalProfile().GetThreatMinimum())
	}

	if weapon.GetCriticalProfile().GetPrimaryMultiplier() != 2 {
		t.Fatalf("expected dagger critical multiplier 2, got %d", weapon.GetCriticalProfile().GetPrimaryMultiplier())
	}

	if weapon.GetRangeIncrement().GetFeet() != 10 {
		t.Fatalf("expected dagger range increment 10 ft, got %d ft", weapon.GetRangeIncrement().GetFeet())
	}
}

func TestGetWeaponByID_ReturnsDetachedCopy(t *testing.T) {
	first, ok := GetWeaponByID(DaggerWeaponID)
	if !ok {
		t.Fatal("expected dagger to be returned from core weapon lookup")
	}

	first.id = "changed"
	first.displayName = "Changed"
	first.proficiencyCategory = MartialWeaponProficiencyCategory
	first.category = RangedWeaponCategory
	first.damageProfile.mediumPrimary.dieSides = 12
	first.criticalProfile.threatMinimum = 20
	first.rangeIncrement.feet = 999
	first.cost.copperPieces = 999
	first.weight.ounces = 999

	second, ok := GetWeaponByID(DaggerWeaponID)
	if !ok {
		t.Fatal("expected dagger to be returned from core weapon lookup")
	}

	if second.GetID() != DaggerWeaponID {
		t.Fatalf("expected stored weapon id to remain %q, got %q", DaggerWeaponID, second.GetID())
	}

	if second.GetDisplayName() != "Dagger" {
		t.Fatalf("expected stored display name to remain %q, got %q", "Dagger", second.GetDisplayName())
	}

	if second.GetProficiencyCategory() != SimpleWeaponProficiencyCategory {
		t.Fatalf("expected stored proficiency category to remain %q, got %q", SimpleWeaponProficiencyCategory, second.GetProficiencyCategory())
	}

	if second.GetCategory() != LightMeleeWeaponCategory {
		t.Fatalf("expected stored category to remain %q, got %q", LightMeleeWeaponCategory, second.GetCategory())
	}

	if second.GetDamageProfile().GetMediumDamage().GetDieSides() != 4 {
		t.Fatalf("expected stored medium damage die to remain d4, got d%d", second.GetDamageProfile().GetMediumDamage().GetDieSides())
	}

	if second.GetCriticalProfile().GetThreatMinimum() != 19 {
		t.Fatalf("expected stored threat minimum to remain 19, got %d", second.GetCriticalProfile().GetThreatMinimum())
	}

	if second.GetRangeIncrement().GetFeet() != 10 {
		t.Fatalf("expected stored range increment to remain 10 ft, got %d ft", second.GetRangeIncrement().GetFeet())
	}

	if second.GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected stored cost to remain 200 cp, got %d cp", second.GetCost().GetCopperPieces())
	}

	if second.GetWeight().GetOunces() != 16 {
		t.Fatalf("expected stored weight to remain 16 oz, got %d oz", second.GetWeight().GetOunces())
	}
}

func TestGetWeaponByID_RejectsUnknownWeapon(t *testing.T) {
	if _, ok := GetWeaponByID(WeaponID("longbow")); ok {
		t.Fatal("expected unknown weapon lookup to fail")
	}
}

func TestGetWeapons_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	weapons := GetWeapons()
	if len(weapons) != len(coreSimpleWeaponOrder) {
		t.Fatalf("expected %d queried weapon entries, got %d", len(coreSimpleWeaponOrder), len(weapons))
	}

	for i, expectedID := range coreSimpleWeaponOrder {
		if weapons[i].GetID() != expectedID {
			t.Fatalf("expected weapon at index %d to be %q, got %q", i, expectedID, weapons[i].GetID())
		}
	}
}

func TestGetWeapons_ReturnsDetachedCopies(t *testing.T) {
	first := GetWeapons()
	second := GetWeapons()

	first[0].id = "changed"
	first[0].displayName = "Changed"
	first[0].proficiencyCategory = MartialWeaponProficiencyCategory
	first[0].category = RangedWeaponCategory
	first[0].damageProfile.mediumPrimary.dieSides = 12
	first[0].criticalProfile.threatMinimum = 19
	first[0].cost.copperPieces = 999
	first[0].weight.ounces = 999

	if second[0].GetID() != GauntletWeaponID {
		t.Fatalf("expected stored weapon id to remain %q, got %q", GauntletWeaponID, second[0].GetID())
	}

	if second[0].GetDisplayName() != "Gauntlet" {
		t.Fatalf("expected stored display name to remain %q, got %q", "Gauntlet", second[0].GetDisplayName())
	}

	if second[0].GetProficiencyCategory() != SimpleWeaponProficiencyCategory {
		t.Fatalf("expected stored proficiency category to remain %q, got %q", SimpleWeaponProficiencyCategory, second[0].GetProficiencyCategory())
	}

	if second[0].GetCategory() != UnarmedAttackWeaponCategory {
		t.Fatalf("expected stored category to remain %q, got %q", UnarmedAttackWeaponCategory, second[0].GetCategory())
	}

	if second[0].GetDamageProfile().GetMediumDamage().GetDieSides() != 3 {
		t.Fatalf("expected stored medium damage die to remain d3, got d%d", second[0].GetDamageProfile().GetMediumDamage().GetDieSides())
	}

	if second[0].GetCriticalProfile().GetThreatMinimum() != 20 {
		t.Fatalf("expected stored threat minimum to remain 20, got %d", second[0].GetCriticalProfile().GetThreatMinimum())
	}

	if second[0].GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected stored cost to remain 200 cp, got %d cp", second[0].GetCost().GetCopperPieces())
	}

	if second[0].GetWeight().GetOunces() != 16 {
		t.Fatalf("expected stored weight to remain 16 oz, got %d oz", second[0].GetWeight().GetOunces())
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
