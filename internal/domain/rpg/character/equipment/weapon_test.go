package equipment

import "testing"

func TestNewWeapon_ConstructsValidatedWeaponChassis(t *testing.T) {
	damageProfile := mustNewWeaponDamageProfile(
		t,
		mustNewWeaponDamageDice(t, 1, 6),
		mustNewWeaponDamageDice(t, 1, 8),
	)
	criticalProfile := mustNewWeaponCriticalProfile(t, 19, 2)
	cost := mustNewEquipmentCost(t, 1500)
	weight := mustNewEquipmentWeightOunces(t, 64)

	weapon, ok := NewWeapon(
		WeaponID("longsword"),
		"Longsword",
		MartialWeaponProficiencyCategory,
		OneHandedMeleeWeaponCategory,
		damageProfile,
		criticalProfile,
		NewNoWeaponRangeIncrement(),
		cost,
		weight,
	)
	if !ok {
		t.Fatal("expected weapon chassis to be constructed")
	}

	if weapon.GetID() != WeaponID("longsword") {
		t.Fatalf("expected weapon id %q, got %q", WeaponID("longsword"), weapon.GetID())
	}

	if weapon.GetDisplayName() != "Longsword" {
		t.Fatalf("expected display name %q, got %q", "Longsword", weapon.GetDisplayName())
	}

	if weapon.GetProficiencyCategory() != MartialWeaponProficiencyCategory {
		t.Fatalf("expected proficiency category %q, got %q", MartialWeaponProficiencyCategory, weapon.GetProficiencyCategory())
	}

	if weapon.GetCategory() != OneHandedMeleeWeaponCategory {
		t.Fatalf("expected weapon category %q, got %q", OneHandedMeleeWeaponCategory, weapon.GetCategory())
	}

	if weapon.GetDamageProfile().GetSmallDamage().GetDiceCount() != 1 ||
		weapon.GetDamageProfile().GetSmallDamage().GetDieSides() != 6 {
		t.Fatal("expected small longsword damage to be 1d6")
	}

	if weapon.GetDamageProfile().GetMediumDamage().GetDiceCount() != 1 ||
		weapon.GetDamageProfile().GetMediumDamage().GetDieSides() != 8 {
		t.Fatal("expected medium longsword damage to be 1d8")
	}

	if weapon.GetDamageProfile().HasSecondaryDamage() {
		t.Fatal("expected longsword to have no secondary damage profile")
	}

	if weapon.GetCriticalProfile().GetThreatMinimum() != 19 ||
		weapon.GetCriticalProfile().GetThreatMaximum() != 20 ||
		weapon.GetCriticalProfile().GetPrimaryMultiplier() != 2 {
		t.Fatal("expected longsword critical profile to be 19-20/x2")
	}

	if weapon.GetRangeIncrement().HasRangeIncrement() {
		t.Fatal("expected longsword to have no range increment")
	}

	if weapon.GetCost().GetCopperPieces() != 1500 {
		t.Fatalf("expected cost 1500 cp, got %d cp", weapon.GetCost().GetCopperPieces())
	}

	if weapon.GetWeight().GetOunces() != 64 {
		t.Fatalf("expected weight 64 oz, got %d oz", weapon.GetWeight().GetOunces())
	}
}

func TestNewWeapon_ConstructsRangedWeaponWithRangeIncrement(t *testing.T) {
	damageProfile := mustNewWeaponDamageProfile(
		t,
		mustNewWeaponDamageDice(t, 1, 6),
		mustNewWeaponDamageDice(t, 1, 8),
	)
	criticalProfile := mustNewWeaponCriticalProfile(t, 19, 2)
	rangeIncrement := mustNewWeaponRangeIncrementFeet(t, 80)

	weapon, ok := NewWeapon(
		WeaponID("crossbow-light"),
		"Crossbow, light",
		SimpleWeaponProficiencyCategory,
		RangedWeaponCategory,
		damageProfile,
		criticalProfile,
		rangeIncrement,
		mustNewEquipmentCost(t, 3500),
		mustNewEquipmentWeightOunces(t, 64),
	)
	if !ok {
		t.Fatal("expected ranged weapon chassis to be constructed")
	}

	if !weapon.GetRangeIncrement().HasRangeIncrement() {
		t.Fatal("expected ranged weapon to have a range increment")
	}

	if weapon.GetRangeIncrement().GetFeet() != 80 {
		t.Fatalf("expected range increment 80 ft, got %d ft", weapon.GetRangeIncrement().GetFeet())
	}
}

func TestNewWeapon_ConstructsDoubleWeaponProfiles(t *testing.T) {
	damageProfile := mustNewDoubleWeaponDamageProfile(
		t,
		mustNewWeaponDamageDice(t, 1, 6),
		mustNewWeaponDamageDice(t, 1, 8),
		mustNewWeaponDamageDice(t, 1, 4),
		mustNewWeaponDamageDice(t, 1, 6),
	)
	criticalProfile := mustNewDoubleWeaponCriticalProfile(t, 20, 3, 4)

	weapon, ok := NewWeapon(
		WeaponID("hammer-gnome-hooked"),
		"Hammer, gnome hooked",
		ExoticWeaponProficiencyCategory,
		TwoHandedMeleeWeaponCategory,
		damageProfile,
		criticalProfile,
		NewNoWeaponRangeIncrement(),
		mustNewEquipmentCost(t, 2000),
		mustNewEquipmentWeightOunces(t, 96),
	)
	if !ok {
		t.Fatal("expected double weapon chassis to be constructed")
	}

	if !weapon.GetDamageProfile().HasSecondaryDamage() {
		t.Fatal("expected double weapon to have secondary damage")
	}

	if weapon.GetDamageProfile().GetSecondarySmallDamage().GetDiceCount() != 1 ||
		weapon.GetDamageProfile().GetSecondarySmallDamage().GetDieSides() != 4 {
		t.Fatal("expected secondary small damage to be 1d4")
	}

	if !weapon.GetCriticalProfile().HasSecondaryMultiplier() {
		t.Fatal("expected double weapon to have secondary critical multiplier")
	}

	if weapon.GetCriticalProfile().GetPrimaryMultiplier() != 3 ||
		weapon.GetCriticalProfile().GetSecondaryMultiplier() != 4 {
		t.Fatal("expected double weapon critical profile to be x3/x4")
	}
}

func TestNewWeapon_AllowsConstructedNoDamageNoCriticalWeapon(t *testing.T) {
	damageProfile := mustNewWeaponDamageProfile(t, NewNoWeaponDamage(), NewNoWeaponDamage())
	rangeIncrement := mustNewWeaponRangeIncrementFeet(t, 10)

	weapon, ok := NewWeapon(
		WeaponID("net"),
		"Net",
		ExoticWeaponProficiencyCategory,
		RangedWeaponCategory,
		damageProfile,
		NewNoWeaponCriticalProfile(),
		rangeIncrement,
		mustNewEquipmentCost(t, 2000),
		mustNewEquipmentWeightOunces(t, 96),
	)
	if !ok {
		t.Fatal("expected no-damage weapon chassis to be constructed")
	}

	if weapon.GetDamageProfile().HasDamage() {
		t.Fatal("expected net to have no damage")
	}

	if weapon.GetCriticalProfile().HasCritical() {
		t.Fatal("expected net to have no critical profile")
	}
}

func TestNewWeapon_ConstructsFlatDamageProfile(t *testing.T) {
	damageProfile := mustNewWeaponDamageProfile(
		t,
		mustNewWeaponFlatDamage(t, 1),
		mustNewWeaponDamageDice(t, 1, 2),
	)

	if damageProfile.GetSmallDamage().GetKind() != FlatWeaponDamageKind {
		t.Fatalf("expected flat damage kind %q, got %q", FlatWeaponDamageKind, damageProfile.GetSmallDamage().GetKind())
	}

	if damageProfile.GetSmallDamage().GetFlatPoints() != 1 {
		t.Fatalf("expected flat damage 1, got %d", damageProfile.GetSmallDamage().GetFlatPoints())
	}
}

func TestNewWeapon_RejectsInvalidInputs(t *testing.T) {
	damageProfile := mustNewWeaponDamageProfile(
		t,
		mustNewWeaponDamageDice(t, 1, 6),
		mustNewWeaponDamageDice(t, 1, 8),
	)
	criticalProfile := mustNewWeaponCriticalProfile(t, 20, 2)
	rangeIncrement := mustNewWeaponRangeIncrementFeet(t, 10)
	cost := mustNewEquipmentCost(t, 100)
	weight := mustNewEquipmentWeightOunces(t, 16)

	if _, ok := NewWeaponDamageDice(0, 6); ok {
		t.Fatal("expected zero damage dice count to be rejected")
	}

	if _, ok := NewWeaponDamageDice(1, 5); ok {
		t.Fatal("expected unsupported damage die sides to be rejected")
	}

	if _, ok := NewWeaponFlatDamage(0); ok {
		t.Fatal("expected zero flat damage to be rejected")
	}

	if _, ok := NewWeaponDamageProfile(WeaponDamage{}, mustNewWeaponDamageDice(t, 1, 6)); ok {
		t.Fatal("expected zero-value small damage to be rejected")
	}

	if _, ok := NewWeaponDamageProfile(NewNoWeaponDamage(), mustNewWeaponDamageDice(t, 1, 6)); ok {
		t.Fatal("expected mixed no-damage and damage profile to be rejected")
	}

	if _, ok := NewDoubleWeaponDamageProfile(
		NewNoWeaponDamage(),
		mustNewWeaponDamageDice(t, 1, 6),
		mustNewWeaponDamageDice(t, 1, 6),
		mustNewWeaponDamageDice(t, 1, 6),
	); ok {
		t.Fatal("expected double weapon no-damage profile to be rejected")
	}

	if _, ok := NewWeaponCriticalProfile(17, 2); ok {
		t.Fatal("expected invalid threat minimum to be rejected")
	}

	if _, ok := NewWeaponCriticalProfile(20, 5); ok {
		t.Fatal("expected invalid critical multiplier to be rejected")
	}

	if _, ok := NewDoubleWeaponCriticalProfile(20, 3, 5); ok {
		t.Fatal("expected invalid secondary critical multiplier to be rejected")
	}

	if _, ok := NewWeaponRangeIncrementFeet(0); ok {
		t.Fatal("expected zero range increment to be rejected")
	}

	for _, id := range []WeaponID{"", " longsword", "longsword ", "\tlongsword"} {
		if _, ok := NewWeapon(id, "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, criticalProfile, NewNoWeaponRangeIncrement(), cost, weight); ok {
			t.Fatalf("expected invalid weapon id %q to be rejected", id)
		}
	}

	for _, displayName := range []string{"", " Longsword", "Longsword ", "\tLongsword"} {
		if _, ok := NewWeapon(WeaponID("longsword"), displayName, MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, criticalProfile, NewNoWeaponRangeIncrement(), cost, weight); ok {
			t.Fatalf("expected invalid display name %q to be rejected", displayName)
		}
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", WeaponProficiencyCategory("Advanced"), OneHandedMeleeWeaponCategory, damageProfile, criticalProfile, NewNoWeaponRangeIncrement(), cost, weight); ok {
		t.Fatal("expected unknown weapon proficiency category to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, WeaponCategory("Vehicle"), damageProfile, criticalProfile, NewNoWeaponRangeIncrement(), cost, weight); ok {
		t.Fatal("expected unknown weapon category to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, WeaponDamageProfile{}, criticalProfile, NewNoWeaponRangeIncrement(), cost, weight); ok {
		t.Fatal("expected zero-value damage profile to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, WeaponCriticalProfile{}, NewNoWeaponRangeIncrement(), cost, weight); ok {
		t.Fatal("expected zero-value critical profile to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, criticalProfile, WeaponRangeIncrement{}, cost, weight); ok {
		t.Fatal("expected zero-value range increment to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, criticalProfile, NewNoWeaponRangeIncrement(), EquipmentCost{}, weight); ok {
		t.Fatal("expected zero-value weapon cost to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, criticalProfile, NewNoWeaponRangeIncrement(), cost, EquipmentWeight{}); ok {
		t.Fatal("expected zero-value weapon weight to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("crossbow-light"), "Crossbow, light", SimpleWeaponProficiencyCategory, RangedWeaponCategory, damageProfile, criticalProfile, NewNoWeaponRangeIncrement(), cost, weight); ok {
		t.Fatal("expected ranged weapon without range increment to be rejected")
	}

	noDamageProfile := mustNewWeaponDamageProfile(t, NewNoWeaponDamage(), NewNoWeaponDamage())
	if _, ok := NewWeapon(WeaponID("net"), "Net", ExoticWeaponProficiencyCategory, RangedWeaponCategory, noDamageProfile, criticalProfile, rangeIncrement, cost, weight); ok {
		t.Fatal("expected no-damage weapon with critical profile to be rejected")
	}

	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, NewNoWeaponCriticalProfile(), NewNoWeaponRangeIncrement(), cost, weight); ok {
		t.Fatal("expected damage weapon without critical profile to be rejected")
	}

	secondaryCriticalProfile := mustNewDoubleWeaponCriticalProfile(t, 20, 2, 3)
	if _, ok := NewWeapon(WeaponID("longsword"), "Longsword", MartialWeaponProficiencyCategory, OneHandedMeleeWeaponCategory, damageProfile, secondaryCriticalProfile, NewNoWeaponRangeIncrement(), cost, weight); ok {
		t.Fatal("expected single-damage weapon with secondary critical multiplier to be rejected")
	}
}

func mustNewWeaponDamageDice(t *testing.T, diceCount int, dieSides int) WeaponDamage {
	t.Helper()

	damage, ok := NewWeaponDamageDice(diceCount, dieSides)
	if !ok {
		t.Fatalf("expected weapon damage %dd%d to be constructed", diceCount, dieSides)
	}

	return damage
}

func mustNewWeaponFlatDamage(t *testing.T, points int) WeaponDamage {
	t.Helper()

	damage, ok := NewWeaponFlatDamage(points)
	if !ok {
		t.Fatalf("expected flat weapon damage %d to be constructed", points)
	}

	return damage
}

func mustNewWeaponDamageProfile(t *testing.T, small WeaponDamage, medium WeaponDamage) WeaponDamageProfile {
	t.Helper()

	damageProfile, ok := NewWeaponDamageProfile(small, medium)
	if !ok {
		t.Fatal("expected weapon damage profile to be constructed")
	}

	return damageProfile
}

func mustNewDoubleWeaponDamageProfile(
	t *testing.T,
	smallPrimary WeaponDamage,
	mediumPrimary WeaponDamage,
	smallSecondary WeaponDamage,
	mediumSecondary WeaponDamage,
) WeaponDamageProfile {
	t.Helper()

	damageProfile, ok := NewDoubleWeaponDamageProfile(smallPrimary, mediumPrimary, smallSecondary, mediumSecondary)
	if !ok {
		t.Fatal("expected double weapon damage profile to be constructed")
	}

	return damageProfile
}

func mustNewWeaponCriticalProfile(t *testing.T, threatMinimum int, multiplier int) WeaponCriticalProfile {
	t.Helper()

	criticalProfile, ok := NewWeaponCriticalProfile(threatMinimum, multiplier)
	if !ok {
		t.Fatalf("expected weapon critical profile %d-20/x%d to be constructed", threatMinimum, multiplier)
	}

	return criticalProfile
}

func mustNewDoubleWeaponCriticalProfile(
	t *testing.T,
	threatMinimum int,
	primaryMultiplier int,
	secondaryMultiplier int,
) WeaponCriticalProfile {
	t.Helper()

	criticalProfile, ok := NewDoubleWeaponCriticalProfile(threatMinimum, primaryMultiplier, secondaryMultiplier)
	if !ok {
		t.Fatal("expected double weapon critical profile to be constructed")
	}

	return criticalProfile
}

func mustNewWeaponRangeIncrementFeet(t *testing.T, feet int) WeaponRangeIncrement {
	t.Helper()

	rangeIncrement, ok := NewWeaponRangeIncrementFeet(feet)
	if !ok {
		t.Fatalf("expected weapon range increment %d ft to be constructed", feet)
	}

	return rangeIncrement
}

func mustNewEquipmentCost(t *testing.T, copperPieces int) EquipmentCost {
	t.Helper()

	cost, ok := NewEquipmentCost(copperPieces)
	if !ok {
		t.Fatalf("expected equipment cost %d cp to be constructed", copperPieces)
	}

	return cost
}

func mustNewEquipmentWeightOunces(t *testing.T, ounces int) EquipmentWeight {
	t.Helper()

	weight, ok := NewEquipmentWeightOunces(ounces)
	if !ok {
		t.Fatalf("expected equipment weight %d oz to be constructed", ounces)
	}

	return weight
}
