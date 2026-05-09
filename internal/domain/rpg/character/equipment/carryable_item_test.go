package equipment

import "testing"

func TestNewCarryableItem_ComposesEquipmentWeaponArmorAndShieldWeightFacts(t *testing.T) {
	backpack, ok := GetEquipmentByID(BackpackEmptyEquipmentID)
	if !ok {
		t.Fatal("expected backpack equipment seed to resolve")
	}

	equipmentItem, ok := NewCarryableItemFromEquipment(backpack)
	if !ok {
		t.Fatal("expected equipment carryable item to compose")
	}

	if equipmentItem.GetRef().GetKind() != EquipmentCarryableItemKind ||
		equipmentItem.GetRef().GetID() != string(BackpackEmptyEquipmentID) {
		t.Fatalf("expected equipment carryable ref %q/%q", EquipmentCarryableItemKind, BackpackEmptyEquipmentID)
	}

	if equipmentItem.GetDisplayName() != "Backpack (empty)" ||
		equipmentItem.GetCost().GetCopperPieces() != 200 ||
		equipmentItem.GetWeight().GetOunces() != 32 {
		t.Fatal("expected equipment carryable item to expose display, cost, and weight facts")
	}

	weaponItem, ok := NewCarryableItemFromWeapon(mustCarryableItemWeapon(t))
	if !ok {
		t.Fatal("expected weapon carryable item to compose")
	}

	if weaponItem.GetRef().GetKind() != WeaponCarryableItemKind ||
		weaponItem.GetRef().GetID() != "longsword" ||
		weaponItem.GetWeight().GetOunces() != 64 {
		t.Fatal("expected weapon carryable item to expose weapon ref and weight")
	}

	armorItem, ok := NewCarryableItemFromArmor(mustCarryableItemArmor(t))
	if !ok {
		t.Fatal("expected armor carryable item to compose")
	}

	if armorItem.GetRef().GetKind() != ArmorCarryableItemKind ||
		armorItem.GetRef().GetID() != "chain-shirt" ||
		armorItem.GetWeight().GetOunces() != 400 {
		t.Fatal("expected armor carryable item to expose armor ref and weight")
	}

	shieldItem, ok := NewCarryableItemFromArmor(mustCarryableItemShield(t))
	if !ok {
		t.Fatal("expected shield carryable item to compose through armor boundary")
	}

	if shieldItem.GetRef().GetKind() != ArmorCarryableItemKind ||
		shieldItem.GetRef().GetID() != "shield-heavy-steel" ||
		shieldItem.GetWeight().GetOunces() != 240 {
		t.Fatal("expected shield carryable item to expose armor ref and weight")
	}
}

func TestGetCarryableItemByRef_ReturnsSeededCoreEquipment(t *testing.T) {
	ref := mustNewEquipmentCarryableItemRefForTest(t, BackpackEmptyEquipmentID)

	item, ok := GetCarryableItemByRef(ref)
	if !ok {
		t.Fatal("expected backpack to resolve through carryable lookup")
	}

	if item.GetRef().GetKind() != EquipmentCarryableItemKind {
		t.Fatalf("expected carryable kind %q, got %q", EquipmentCarryableItemKind, item.GetRef().GetKind())
	}

	if item.GetRef().GetID() != string(BackpackEmptyEquipmentID) {
		t.Fatalf("expected carryable id %q, got %q", BackpackEmptyEquipmentID, item.GetRef().GetID())
	}

	if item.GetWeight().GetOunces() != 32 {
		t.Fatalf("expected backpack carryable weight 32 oz, got %d oz", item.GetWeight().GetOunces())
	}
}

func TestGetCarryableItemByRef_FailsClosedForUnseededWeaponAndArmorRefs(t *testing.T) {
	weaponRef, ok := NewWeaponCarryableItemRef(WeaponID("longsword"))
	if !ok {
		t.Fatal("expected valid weapon carryable ref to compose")
	}

	if _, ok := GetCarryableItemByRef(weaponRef); ok {
		t.Fatal("expected unseeded weapon lookup to fail closed")
	}

	armorRef, ok := NewArmorCarryableItemRef(ArmorID("chain-shirt"))
	if !ok {
		t.Fatal("expected valid armor carryable ref to compose")
	}

	if _, ok := GetCarryableItemByRef(armorRef); ok {
		t.Fatal("expected unseeded armor lookup to fail closed")
	}
}

func TestCarryableItem_RejectsInvalidInputs(t *testing.T) {
	if _, ok := NewEquipmentCarryableItemRef(EquipmentID(" backpack-empty")); ok {
		t.Fatal("expected malformed equipment carryable ref to be rejected")
	}

	if _, ok := NewWeaponCarryableItemRef(WeaponID(" longsword")); ok {
		t.Fatal("expected malformed weapon carryable ref to be rejected")
	}

	if _, ok := NewArmorCarryableItemRef(ArmorID(" chain-shirt")); ok {
		t.Fatal("expected malformed armor carryable ref to be rejected")
	}

	if _, ok := GetCarryableItemByRef(CarryableItemRef{}); ok {
		t.Fatal("expected zero-value carryable ref lookup to fail")
	}

	if _, ok := NewCarryableItemFromEquipment(Equipment{}); ok {
		t.Fatal("expected zero-value equipment carryable item to be rejected")
	}

	if _, ok := NewCarryableItemFromWeapon(Weapon{}); ok {
		t.Fatal("expected zero-value weapon carryable item to be rejected")
	}

	if _, ok := NewCarryableItemFromArmor(Armor{}); ok {
		t.Fatal("expected zero-value armor carryable item to be rejected")
	}
}

func TestGetCarryableItems_ReturnsSeededEquipmentCarryablesInCoreOrder(t *testing.T) {
	items := GetCarryableItems()
	if len(items) != len(coreEquipmentOrder) {
		t.Fatalf("expected %d carryable items, got %d", len(coreEquipmentOrder), len(items))
	}

	for i, expectedID := range coreEquipmentOrder {
		if items[i].GetRef().GetKind() != EquipmentCarryableItemKind {
			t.Fatalf("expected carryable item %d kind %q, got %q", i, EquipmentCarryableItemKind, items[i].GetRef().GetKind())
		}

		if items[i].GetRef().GetID() != string(expectedID) {
			t.Fatalf("expected carryable item %d id %q, got %q", i, expectedID, items[i].GetRef().GetID())
		}
	}
}

func mustCarryableItemWeapon(t *testing.T) Weapon {
	t.Helper()

	damageProfile := mustNewWeaponDamageProfile(
		t,
		mustNewWeaponDamageDice(t, 1, 6),
		mustNewWeaponDamageDice(t, 1, 8),
	)

	weapon, ok := NewWeapon(
		WeaponID("longsword"),
		"Longsword",
		MartialWeaponProficiencyCategory,
		OneHandedMeleeWeaponCategory,
		damageProfile,
		mustNewWeaponCriticalProfile(t, 19, 2),
		NewNoWeaponRangeIncrement(),
		mustNewEquipmentCost(t, 1500),
		mustNewEquipmentWeightOunces(t, 64),
	)
	if !ok {
		t.Fatal("expected weapon to compose")
	}

	return weapon
}

func mustCarryableItemArmor(t *testing.T) Armor {
	t.Helper()

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
		t.Fatal("expected armor to compose")
	}

	return armor
}

func mustCarryableItemShield(t *testing.T) Armor {
	t.Helper()

	shield, ok := NewArmor(
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
		t.Fatal("expected shield to compose")
	}

	return shield
}

func mustNewEquipmentCarryableItemRefForTest(t *testing.T, id EquipmentID) CarryableItemRef {
	t.Helper()

	ref, ok := NewEquipmentCarryableItemRef(id)
	if !ok {
		t.Fatalf("expected equipment carryable ref %q to compose", id)
	}

	return ref
}
