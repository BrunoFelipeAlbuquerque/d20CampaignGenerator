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

func TestGetArmorByID_ReturnsSeededCoreArmor(t *testing.T) {
	armor, ok := GetArmorByID(ChainShirtArmorID)
	if !ok {
		t.Fatal("expected chain shirt to be returned from core armor lookup")
	}

	if armor.GetID() != ChainShirtArmorID {
		t.Fatalf("expected armor id %q, got %q", ChainShirtArmorID, armor.GetID())
	}

	if armor.GetDisplayName() != "Chain shirt" {
		t.Fatalf("expected display name %q, got %q", "Chain shirt", armor.GetDisplayName())
	}

	if armor.GetCategory() != LightArmorCategory {
		t.Fatalf("expected category %q, got %q", LightArmorCategory, armor.GetCategory())
	}

	if armor.GetArmorClassBonus().GetPoints() != 4 {
		t.Fatalf("expected armor class bonus 4, got %d", armor.GetArmorClassBonus().GetPoints())
	}

	if !armor.GetMaximumDexterityBonus().HasMaximum() ||
		armor.GetMaximumDexterityBonus().GetPoints() != 4 {
		t.Fatal("expected maximum Dexterity bonus 4")
	}

	if armor.GetArmorCheckPenalty().GetPenalty() != -2 {
		t.Fatalf("expected armor check penalty -2, got %d", armor.GetArmorCheckPenalty().GetPenalty())
	}

	if armor.GetArcaneSpellFailureChance().GetPercent() != 20 {
		t.Fatalf("expected arcane spell failure chance 20, got %d", armor.GetArcaneSpellFailureChance().GetPercent())
	}

	if armor.GetSpeedImpact().HasImpact() {
		t.Fatal("expected chain shirt to have no speed impact metadata")
	}

	if armor.GetCost().GetCopperPieces() != 10000 {
		t.Fatalf("expected chain shirt cost 10000 cp, got %d cp", armor.GetCost().GetCopperPieces())
	}

	if armor.GetWeight().GetOunces() != 400 {
		t.Fatalf("expected chain shirt weight 400 oz, got %d oz", armor.GetWeight().GetOunces())
	}
}

func TestGetArmorByID_ReturnsDetachedCopy(t *testing.T) {
	first, ok := GetArmorByID(ChainShirtArmorID)
	if !ok {
		t.Fatal("expected chain shirt to be returned from core armor lookup")
	}

	first.id = "changed"
	first.displayName = "Changed"
	first.category = HeavyArmorCategory
	first.armorClassBonus.points = 99
	first.maximumDexterityBonus.points = 99
	first.armorCheckPenalty.penalty = 0
	first.arcaneSpellFailureChance.percent = 99
	first.speedImpact.hasImpact = true
	first.cost.copperPieces = 999
	first.weight.ounces = 999

	second, ok := GetArmorByID(ChainShirtArmorID)
	if !ok {
		t.Fatal("expected chain shirt to be returned from core armor lookup")
	}

	if second.GetID() != ChainShirtArmorID {
		t.Fatalf("expected stored armor id to remain %q, got %q", ChainShirtArmorID, second.GetID())
	}

	if second.GetDisplayName() != "Chain shirt" {
		t.Fatalf("expected stored display name to remain %q, got %q", "Chain shirt", second.GetDisplayName())
	}

	if second.GetCategory() != LightArmorCategory {
		t.Fatalf("expected stored category to remain %q, got %q", LightArmorCategory, second.GetCategory())
	}

	if second.GetArmorClassBonus().GetPoints() != 4 {
		t.Fatalf("expected stored armor class bonus to remain 4, got %d", second.GetArmorClassBonus().GetPoints())
	}

	if second.GetMaximumDexterityBonus().GetPoints() != 4 {
		t.Fatalf("expected stored maximum Dexterity bonus to remain 4, got %d", second.GetMaximumDexterityBonus().GetPoints())
	}

	if second.GetArmorCheckPenalty().GetPenalty() != -2 {
		t.Fatalf("expected stored armor check penalty to remain -2, got %d", second.GetArmorCheckPenalty().GetPenalty())
	}

	if second.GetArcaneSpellFailureChance().GetPercent() != 20 {
		t.Fatalf("expected stored arcane spell failure chance to remain 20, got %d", second.GetArcaneSpellFailureChance().GetPercent())
	}

	if second.GetSpeedImpact().HasImpact() {
		t.Fatal("expected stored speed impact metadata to remain absent")
	}

	if second.GetCost().GetCopperPieces() != 10000 {
		t.Fatalf("expected stored cost to remain 10000 cp, got %d cp", second.GetCost().GetCopperPieces())
	}

	if second.GetWeight().GetOunces() != 400 {
		t.Fatalf("expected stored weight to remain 400 oz, got %d oz", second.GetWeight().GetOunces())
	}
}

func TestGetArmorByID_RejectsUnknownArmor(t *testing.T) {
	if _, ok := GetArmorByID(ArmorID("breastplate")); ok {
		t.Fatal("expected unknown armor lookup to fail")
	}
}

func TestGetArmor_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	armor := GetArmor()
	if len(armor) != len(coreArmorOrder) {
		t.Fatalf("expected %d queried armor entries, got %d", len(coreArmorOrder), len(armor))
	}

	for i, expectedID := range coreArmorOrder {
		if armor[i].GetID() != expectedID {
			t.Fatalf("expected armor at index %d to be %q, got %q", i, expectedID, armor[i].GetID())
		}
	}
}

func TestGetArmor_ReturnsDetachedCopies(t *testing.T) {
	first := GetArmor()
	second := GetArmor()

	first[0].id = "changed"
	first[0].displayName = "Changed"
	first[0].category = HeavyArmorCategory
	first[0].armorClassBonus.points = 99
	first[0].maximumDexterityBonus.points = 99
	first[0].armorCheckPenalty.penalty = -9
	first[0].arcaneSpellFailureChance.percent = 99
	first[0].speedImpact.hasImpact = true
	first[0].cost.copperPieces = 999
	first[0].weight.ounces = 999

	if second[0].GetID() != PaddedArmorID {
		t.Fatalf("expected stored armor id to remain %q, got %q", PaddedArmorID, second[0].GetID())
	}

	if second[0].GetDisplayName() != "Padded" {
		t.Fatalf("expected stored display name to remain %q, got %q", "Padded", second[0].GetDisplayName())
	}

	if second[0].GetCategory() != LightArmorCategory {
		t.Fatalf("expected stored category to remain %q, got %q", LightArmorCategory, second[0].GetCategory())
	}

	if second[0].GetArmorClassBonus().GetPoints() != 1 {
		t.Fatalf("expected stored armor class bonus to remain 1, got %d", second[0].GetArmorClassBonus().GetPoints())
	}

	if second[0].GetMaximumDexterityBonus().GetPoints() != 8 {
		t.Fatalf("expected stored maximum Dexterity bonus to remain 8, got %d", second[0].GetMaximumDexterityBonus().GetPoints())
	}

	if second[0].GetArmorCheckPenalty().GetPenalty() != 0 {
		t.Fatalf("expected stored armor check penalty to remain 0, got %d", second[0].GetArmorCheckPenalty().GetPenalty())
	}

	if second[0].GetArcaneSpellFailureChance().GetPercent() != 5 {
		t.Fatalf("expected stored arcane spell failure chance to remain 5, got %d", second[0].GetArcaneSpellFailureChance().GetPercent())
	}

	if second[0].GetSpeedImpact().HasImpact() {
		t.Fatal("expected stored speed impact metadata to remain absent")
	}

	if second[0].GetCost().GetCopperPieces() != 500 {
		t.Fatalf("expected stored cost to remain 500 cp, got %d cp", second[0].GetCost().GetCopperPieces())
	}

	if second[0].GetWeight().GetOunces() != 160 {
		t.Fatalf("expected stored weight to remain 160 oz, got %d oz", second[0].GetWeight().GetOunces())
	}
}
