package equipment

import "testing"

func TestCoreArmor_SeedsLightArmorAndShieldsBatchOne(t *testing.T) {
	expected := map[ArmorID]struct {
		displayName               string
		category                  ArmorCategory
		copperPieces              int
		ounces                    int
		armorClassBonus           int
		maximumDexterityBonus     int
		hasMaximumDexterityBonus  bool
		armorCheckPenalty         int
		arcaneSpellFailurePercent int
	}{
		PaddedArmorID: {
			displayName:               "Padded",
			category:                  LightArmorCategory,
			copperPieces:              500,
			ounces:                    160,
			armorClassBonus:           1,
			maximumDexterityBonus:     8,
			hasMaximumDexterityBonus:  true,
			armorCheckPenalty:         0,
			arcaneSpellFailurePercent: 5,
		},
		LeatherArmorID: {
			displayName:               "Leather",
			category:                  LightArmorCategory,
			copperPieces:              1000,
			ounces:                    240,
			armorClassBonus:           2,
			maximumDexterityBonus:     6,
			hasMaximumDexterityBonus:  true,
			armorCheckPenalty:         0,
			arcaneSpellFailurePercent: 10,
		},
		StuddedLeatherArmorID: {
			displayName:               "Studded leather",
			category:                  LightArmorCategory,
			copperPieces:              2500,
			ounces:                    320,
			armorClassBonus:           3,
			maximumDexterityBonus:     5,
			hasMaximumDexterityBonus:  true,
			armorCheckPenalty:         -1,
			arcaneSpellFailurePercent: 15,
		},
		ChainShirtArmorID: {
			displayName:               "Chain shirt",
			category:                  LightArmorCategory,
			copperPieces:              10000,
			ounces:                    400,
			armorClassBonus:           4,
			maximumDexterityBonus:     4,
			hasMaximumDexterityBonus:  true,
			armorCheckPenalty:         -2,
			arcaneSpellFailurePercent: 20,
		},
		BucklerArmorID: {
			displayName:               "Buckler",
			category:                  ShieldArmorCategory,
			copperPieces:              1500,
			ounces:                    80,
			armorClassBonus:           1,
			armorCheckPenalty:         -1,
			arcaneSpellFailurePercent: 5,
		},
		ShieldLightWoodenArmorID: {
			displayName:               "Shield, light wooden",
			category:                  ShieldArmorCategory,
			copperPieces:              300,
			ounces:                    80,
			armorClassBonus:           1,
			armorCheckPenalty:         -1,
			arcaneSpellFailurePercent: 5,
		},
		ShieldLightSteelArmorID: {
			displayName:               "Shield, light steel",
			category:                  ShieldArmorCategory,
			copperPieces:              900,
			ounces:                    96,
			armorClassBonus:           1,
			armorCheckPenalty:         -1,
			arcaneSpellFailurePercent: 5,
		},
		ShieldHeavyWoodenArmorID: {
			displayName:               "Shield, heavy wooden",
			category:                  ShieldArmorCategory,
			copperPieces:              700,
			ounces:                    160,
			armorClassBonus:           2,
			armorCheckPenalty:         -2,
			arcaneSpellFailurePercent: 15,
		},
		ShieldHeavySteelArmorID: {
			displayName:               "Shield, heavy steel",
			category:                  ShieldArmorCategory,
			copperPieces:              2000,
			ounces:                    240,
			armorClassBonus:           2,
			armorCheckPenalty:         -2,
			arcaneSpellFailurePercent: 15,
		},
		ShieldTowerArmorID: {
			displayName:               "Shield, tower",
			category:                  TowerShieldArmorCategory,
			copperPieces:              3000,
			ounces:                    720,
			armorClassBonus:           4,
			maximumDexterityBonus:     2,
			hasMaximumDexterityBonus:  true,
			armorCheckPenalty:         -10,
			arcaneSpellFailurePercent: 50,
		},
	}

	if len(coreArmor) != len(expected) {
		t.Fatalf("expected %d core armor seeds, got %d", len(expected), len(coreArmor))
	}

	for id, expectation := range expected {
		armor, ok := coreArmor[id]
		if !ok {
			t.Fatalf("expected core armor seed %q", id)
		}

		if armor.GetID() != id {
			t.Fatalf("expected armor id %q, got %q", id, armor.GetID())
		}

		if armor.GetDisplayName() != expectation.displayName {
			t.Fatalf("expected %q display name %q, got %q", id, expectation.displayName, armor.GetDisplayName())
		}

		if armor.GetCategory() != expectation.category {
			t.Fatalf("expected %q category %q, got %q", id, expectation.category, armor.GetCategory())
		}

		if armor.GetArmorClassBonus().GetPoints() != expectation.armorClassBonus {
			t.Fatalf("expected %q armor class bonus %d, got %d", id, expectation.armorClassBonus, armor.GetArmorClassBonus().GetPoints())
		}

		if armor.GetMaximumDexterityBonus().HasMaximum() != expectation.hasMaximumDexterityBonus {
			t.Fatalf("expected %q maximum Dexterity presence %t", id, expectation.hasMaximumDexterityBonus)
		}

		if expectation.hasMaximumDexterityBonus &&
			armor.GetMaximumDexterityBonus().GetPoints() != expectation.maximumDexterityBonus {
			t.Fatalf("expected %q maximum Dexterity bonus %d, got %d", id, expectation.maximumDexterityBonus, armor.GetMaximumDexterityBonus().GetPoints())
		}

		if armor.GetArmorCheckPenalty().GetPenalty() != expectation.armorCheckPenalty {
			t.Fatalf("expected %q armor check penalty %d, got %d", id, expectation.armorCheckPenalty, armor.GetArmorCheckPenalty().GetPenalty())
		}

		if armor.GetArcaneSpellFailureChance().GetPercent() != expectation.arcaneSpellFailurePercent {
			t.Fatalf("expected %q arcane spell failure chance %d, got %d", id, expectation.arcaneSpellFailurePercent, armor.GetArcaneSpellFailureChance().GetPercent())
		}

		if armor.GetSpeedImpact().HasImpact() {
			t.Fatalf("expected %q not to have speed impact metadata", id)
		}

		if armor.GetCost().GetCopperPieces() != expectation.copperPieces {
			t.Fatalf("expected %q cost %d cp, got %d cp", id, expectation.copperPieces, armor.GetCost().GetCopperPieces())
		}

		if armor.GetWeight().GetOunces() != expectation.ounces {
			t.Fatalf("expected %q weight %d oz, got %d oz", id, expectation.ounces, armor.GetWeight().GetOunces())
		}
	}
}

func TestCoreArmorOrder_ContainsOnlySeededBatchOneIDs(t *testing.T) {
	expectedOrder := []ArmorID{
		PaddedArmorID,
		LeatherArmorID,
		StuddedLeatherArmorID,
		ChainShirtArmorID,
		BucklerArmorID,
		ShieldLightWoodenArmorID,
		ShieldLightSteelArmorID,
		ShieldHeavyWoodenArmorID,
		ShieldHeavySteelArmorID,
		ShieldTowerArmorID,
	}

	if len(coreArmorOrder) != len(expectedOrder) {
		t.Fatalf("expected order length %d, got %d", len(expectedOrder), len(coreArmorOrder))
	}

	seen := make(map[ArmorID]struct{}, len(coreArmorOrder))
	for index, id := range coreArmorOrder {
		if id != expectedOrder[index] {
			t.Fatalf("expected ordered armor id at index %d to be %q, got %q", index, expectedOrder[index], id)
		}

		if _, ok := coreArmor[id]; !ok {
			t.Fatalf("expected ordered armor id %q to have a seed", id)
		}

		if _, ok := seen[id]; ok {
			t.Fatalf("expected ordered armor id %q not to be duplicated", id)
		}

		seen[id] = struct{}{}
	}
}
