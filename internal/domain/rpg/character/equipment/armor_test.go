package equipment

import "testing"

func TestNewArmor_ConstructsValidatedLightArmorChassis(t *testing.T) {
	armor, ok := NewArmor(
		ArmorID("chain-shirt"),
		"Chain shirt",
		LightArmorCategory,
		mustNewArmorClassBonus(t, 4),
		mustNewArmorMaximumDexterityBonus(t, 4),
		mustNewArmorCheckPenalty(t, -2),
		mustNewArmorArcaneSpellFailureChance(t, 20),
		NewNoArmorSpeedImpact(),
		mustNewEquipmentCost(t, 10000),
		mustNewEquipmentWeightOunces(t, 400),
	)
	if !ok {
		t.Fatal("expected light armor chassis to be constructed")
	}

	if armor.GetID() != ArmorID("chain-shirt") {
		t.Fatalf("expected armor id %q, got %q", ArmorID("chain-shirt"), armor.GetID())
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
		t.Fatal("expected light armor to have no speed impact metadata")
	}

	if armor.GetCost().GetCopperPieces() != 10000 {
		t.Fatalf("expected cost 10000 cp, got %d cp", armor.GetCost().GetCopperPieces())
	}

	if armor.GetWeight().GetOunces() != 400 {
		t.Fatalf("expected weight 400 oz, got %d oz", armor.GetWeight().GetOunces())
	}
}

func TestNewArmor_ConstructsMediumArmorWithSpeedImpact(t *testing.T) {
	speedImpact := mustNewArmorSpeedImpact(t, 20, 15, false)

	armor, ok := NewArmor(
		ArmorID("breastplate"),
		"Breastplate",
		MediumArmorCategory,
		mustNewArmorClassBonus(t, 6),
		mustNewArmorMaximumDexterityBonus(t, 3),
		mustNewArmorCheckPenalty(t, -4),
		mustNewArmorArcaneSpellFailureChance(t, 25),
		speedImpact,
		mustNewEquipmentCost(t, 20000),
		mustNewEquipmentWeightOunces(t, 480),
	)
	if !ok {
		t.Fatal("expected medium armor chassis to be constructed")
	}

	if !armor.GetSpeedImpact().HasImpact() {
		t.Fatal("expected medium armor to have speed impact metadata")
	}

	if armor.GetSpeedImpact().GetSpeedFor30FeetBase() != 20 ||
		armor.GetSpeedImpact().GetSpeedFor20FeetBase() != 15 {
		t.Fatal("expected medium armor speed impact to be 20 ft / 15 ft")
	}

	if armor.GetSpeedImpact().LimitsRunning() {
		t.Fatal("expected medium armor to avoid heavy-armor running limit metadata")
	}
}

func TestNewArmor_ConstructsHeavyArmorWithRunningLimit(t *testing.T) {
	speedImpact := mustNewArmorSpeedImpact(t, 20, 15, true)

	armor, ok := NewArmor(
		ArmorID("full-plate"),
		"Full plate",
		HeavyArmorCategory,
		mustNewArmorClassBonus(t, 9),
		mustNewArmorMaximumDexterityBonus(t, 1),
		mustNewArmorCheckPenalty(t, -6),
		mustNewArmorArcaneSpellFailureChance(t, 35),
		speedImpact,
		mustNewEquipmentCost(t, 150000),
		mustNewEquipmentWeightOunces(t, 800),
	)
	if !ok {
		t.Fatal("expected heavy armor chassis to be constructed")
	}

	if !armor.GetSpeedImpact().LimitsRunning() {
		t.Fatal("expected heavy armor to carry running limit metadata")
	}
}

func TestNewArmor_ConstructsShieldWithoutMaxDexOrSpeedImpact(t *testing.T) {
	armor, ok := NewArmor(
		ArmorID("shield-heavy-steel"),
		"Shield, heavy steel",
		ShieldArmorCategory,
		mustNewArmorClassBonus(t, 2),
		NewNoArmorMaximumDexterityBonus(),
		mustNewArmorCheckPenalty(t, -2),
		mustNewArmorArcaneSpellFailureChance(t, 15),
		NewNoArmorSpeedImpact(),
		mustNewEquipmentCost(t, 2000),
		mustNewEquipmentWeightOunces(t, 240),
	)
	if !ok {
		t.Fatal("expected shield chassis to be constructed")
	}

	if armor.GetMaximumDexterityBonus().HasMaximum() {
		t.Fatal("expected shield to have no maximum Dexterity bonus")
	}

	if armor.GetSpeedImpact().HasImpact() {
		t.Fatal("expected shield to have no speed impact metadata")
	}
}

func TestNewArmor_ConstructsTowerShieldWithMaxDexAndNoSpeedImpact(t *testing.T) {
	armor, ok := NewArmor(
		ArmorID("shield-tower"),
		"Shield, tower",
		TowerShieldArmorCategory,
		mustNewArmorClassBonus(t, 4),
		mustNewArmorMaximumDexterityBonus(t, 2),
		mustNewArmorCheckPenalty(t, -10),
		mustNewArmorArcaneSpellFailureChance(t, 50),
		NewNoArmorSpeedImpact(),
		mustNewEquipmentCost(t, 3000),
		mustNewEquipmentWeightOunces(t, 720),
	)
	if !ok {
		t.Fatal("expected tower shield chassis to be constructed")
	}

	if !armor.GetMaximumDexterityBonus().HasMaximum() ||
		armor.GetMaximumDexterityBonus().GetPoints() != 2 {
		t.Fatal("expected tower shield maximum Dexterity bonus 2")
	}

	if armor.GetSpeedImpact().HasImpact() {
		t.Fatal("expected tower shield to have no speed impact metadata")
	}
}

func TestNewArmor_AllowsZeroMaximumDexterityBonusAndZeroPenalties(t *testing.T) {
	if _, ok := NewArmor(
		ArmorID("half-plate"),
		"Half-plate",
		HeavyArmorCategory,
		mustNewArmorClassBonus(t, 8),
		mustNewArmorMaximumDexterityBonus(t, 0),
		mustNewArmorCheckPenalty(t, -7),
		mustNewArmorArcaneSpellFailureChance(t, 40),
		mustNewArmorSpeedImpact(t, 20, 15, true),
		mustNewEquipmentCost(t, 60000),
		mustNewEquipmentWeightOunces(t, 800),
	); !ok {
		t.Fatal("expected armor with maximum Dexterity bonus 0 to be valid")
	}

	if _, ok := NewArmor(
		ArmorID("padded"),
		"Padded",
		LightArmorCategory,
		mustNewArmorClassBonus(t, 1),
		mustNewArmorMaximumDexterityBonus(t, 8),
		mustNewArmorCheckPenalty(t, 0),
		mustNewArmorArcaneSpellFailureChance(t, 5),
		NewNoArmorSpeedImpact(),
		mustNewEquipmentCost(t, 500),
		mustNewEquipmentWeightOunces(t, 160),
	); !ok {
		t.Fatal("expected armor with check penalty 0 to be valid")
	}
}

func TestNewArmor_RejectsInvalidInputs(t *testing.T) {
	armorClassBonus := mustNewArmorClassBonus(t, 4)
	maximumDexterityBonus := mustNewArmorMaximumDexterityBonus(t, 4)
	armorCheckPenalty := mustNewArmorCheckPenalty(t, -2)
	arcaneSpellFailureChance := mustNewArmorArcaneSpellFailureChance(t, 20)
	speedImpact := mustNewArmorSpeedImpact(t, 20, 15, false)
	cost := mustNewEquipmentCost(t, 100)
	weight := mustNewEquipmentWeightOunces(t, 16)

	if _, ok := NewArmorClassBonus(0); ok {
		t.Fatal("expected zero armor class bonus to be rejected")
	}

	if _, ok := NewArmorMaximumDexterityBonus(-1); ok {
		t.Fatal("expected negative maximum Dexterity bonus to be rejected")
	}

	if _, ok := NewArmorCheckPenalty(1); ok {
		t.Fatal("expected positive armor check penalty to be rejected")
	}

	if _, ok := NewArmorArcaneSpellFailureChance(-1); ok {
		t.Fatal("expected negative arcane spell failure chance to be rejected")
	}

	if _, ok := NewArmorArcaneSpellFailureChance(101); ok {
		t.Fatal("expected arcane spell failure chance above 100 to be rejected")
	}

	if _, ok := NewArmorSpeedImpact(0, 15, false); ok {
		t.Fatal("expected zero 30-foot-base speed impact to be rejected")
	}

	if _, ok := NewArmorSpeedImpact(30, 15, false); ok {
		t.Fatal("expected no-op 30-foot-base speed impact to be rejected")
	}

	if _, ok := NewArmorSpeedImpact(20, 20, false); ok {
		t.Fatal("expected no-op 20-foot-base speed impact to be rejected")
	}

	for _, id := range []ArmorID{"", " chain-shirt", "chain-shirt ", "\tchain-shirt"} {
		if _, ok := NewArmor(id, "Chain shirt", LightArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
			t.Fatalf("expected invalid armor id %q to be rejected", id)
		}
	}

	for _, displayName := range []string{"", " Chain shirt", "Chain shirt ", "\tChain shirt"} {
		if _, ok := NewArmor(ArmorID("chain-shirt"), displayName, LightArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
			t.Fatalf("expected invalid armor display name %q to be rejected", displayName)
		}
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", ArmorCategory("Powered Armor"), armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected unknown armor category to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, ArmorClassBonus{}, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected zero-value armor class bonus to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, ArmorMaximumDexterityBonus{}, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected zero-value maximum Dexterity bonus to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, maximumDexterityBonus, ArmorCheckPenalty{}, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected zero-value armor check penalty to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, ArmorArcaneSpellFailureChance{}, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected zero-value arcane spell failure chance to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, ArmorSpeedImpact{}, cost, weight); ok {
		t.Fatal("expected zero-value armor speed impact to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), EquipmentCost{}, weight); ok {
		t.Fatal("expected zero-value armor cost to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, EquipmentWeight{}); ok {
		t.Fatal("expected zero-value armor weight to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, NewNoArmorMaximumDexterityBonus(), armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected light armor without maximum Dexterity bonus to be rejected")
	}

	if _, ok := NewArmor(ArmorID("chain-shirt"), "Chain shirt", LightArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, speedImpact, cost, weight); ok {
		t.Fatal("expected light armor with speed impact metadata to be rejected")
	}

	if _, ok := NewArmor(ArmorID("breastplate"), "Breastplate", MediumArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected medium armor without speed impact metadata to be rejected")
	}

	if _, ok := NewArmor(ArmorID("breastplate"), "Breastplate", MediumArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, mustNewArmorSpeedImpact(t, 20, 15, true), cost, weight); ok {
		t.Fatal("expected medium armor with heavy-armor running limit metadata to be rejected")
	}

	if _, ok := NewArmor(ArmorID("full-plate"), "Full plate", HeavyArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, speedImpact, cost, weight); ok {
		t.Fatal("expected heavy armor without running limit metadata to be rejected")
	}

	if _, ok := NewArmor(ArmorID("shield-heavy-steel"), "Shield, heavy steel", ShieldArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected shield with maximum Dexterity bonus to be rejected")
	}

	if _, ok := NewArmor(ArmorID("shield-heavy-steel"), "Shield, heavy steel", ShieldArmorCategory, armorClassBonus, NewNoArmorMaximumDexterityBonus(), armorCheckPenalty, arcaneSpellFailureChance, speedImpact, cost, weight); ok {
		t.Fatal("expected shield with speed impact metadata to be rejected")
	}

	if _, ok := NewArmor(ArmorID("shield-tower"), "Shield, tower", TowerShieldArmorCategory, armorClassBonus, NewNoArmorMaximumDexterityBonus(), armorCheckPenalty, arcaneSpellFailureChance, NewNoArmorSpeedImpact(), cost, weight); ok {
		t.Fatal("expected tower shield without maximum Dexterity bonus to be rejected")
	}

	if _, ok := NewArmor(ArmorID("shield-tower"), "Shield, tower", TowerShieldArmorCategory, armorClassBonus, maximumDexterityBonus, armorCheckPenalty, arcaneSpellFailureChance, speedImpact, cost, weight); ok {
		t.Fatal("expected tower shield with speed impact metadata to be rejected")
	}
}

func mustNewArmorClassBonus(t *testing.T, points int) ArmorClassBonus {
	t.Helper()

	bonus, ok := NewArmorClassBonus(points)
	if !ok {
		t.Fatalf("expected armor class bonus %d to be constructed", points)
	}

	return bonus
}

func mustNewArmorMaximumDexterityBonus(t *testing.T, points int) ArmorMaximumDexterityBonus {
	t.Helper()

	bonus, ok := NewArmorMaximumDexterityBonus(points)
	if !ok {
		t.Fatalf("expected maximum Dexterity bonus %d to be constructed", points)
	}

	return bonus
}

func mustNewArmorCheckPenalty(t *testing.T, penalty int) ArmorCheckPenalty {
	t.Helper()

	checkPenalty, ok := NewArmorCheckPenalty(penalty)
	if !ok {
		t.Fatalf("expected armor check penalty %d to be constructed", penalty)
	}

	return checkPenalty
}

func mustNewArmorArcaneSpellFailureChance(t *testing.T, percent int) ArmorArcaneSpellFailureChance {
	t.Helper()

	chance, ok := NewArmorArcaneSpellFailureChance(percent)
	if !ok {
		t.Fatalf("expected arcane spell failure chance %d to be constructed", percent)
	}

	return chance
}

func mustNewArmorSpeedImpact(
	t *testing.T,
	speedFor30FeetBase int,
	speedFor20FeetBase int,
	limitsRunning bool,
) ArmorSpeedImpact {
	t.Helper()

	speedImpact, ok := NewArmorSpeedImpact(speedFor30FeetBase, speedFor20FeetBase, limitsRunning)
	if !ok {
		t.Fatalf("expected armor speed impact %d/%d to be constructed", speedFor30FeetBase, speedFor20FeetBase)
	}

	return speedImpact
}
