package equipment

import "testing"

func TestNewEquipment_ConstructsValidatedEquipmentChassis(t *testing.T) {
	cost, ok := NewEquipmentCost(200)
	if !ok {
		t.Fatal("expected equipment cost to be constructed")
	}

	weight, ok := NewEquipmentWeightOunces(32)
	if !ok {
		t.Fatal("expected equipment weight to be constructed")
	}

	equipment, ok := NewEquipment(
		EquipmentID("backpack-empty"),
		"Backpack (empty)",
		AdventuringGearEquipmentCategory,
		cost,
		weight,
	)
	if !ok {
		t.Fatal("expected equipment chassis to be constructed")
	}

	if equipment.GetID() != EquipmentID("backpack-empty") {
		t.Fatalf("expected equipment id %q, got %q", EquipmentID("backpack-empty"), equipment.GetID())
	}

	if equipment.GetDisplayName() != "Backpack (empty)" {
		t.Fatalf("expected display name %q, got %q", "Backpack (empty)", equipment.GetDisplayName())
	}

	if equipment.GetCategory() != AdventuringGearEquipmentCategory {
		t.Fatalf("expected category %q, got %q", AdventuringGearEquipmentCategory, equipment.GetCategory())
	}

	if equipment.GetCategory().GetName() != "Adventuring Gear" {
		t.Fatalf("expected category name %q, got %q", "Adventuring Gear", equipment.GetCategory().GetName())
	}

	if equipment.GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected cost 200 cp, got %d cp", equipment.GetCost().GetCopperPieces())
	}

	if equipment.GetWeight().GetOunces() != 32 {
		t.Fatalf("expected weight 32 oz, got %d oz", equipment.GetWeight().GetOunces())
	}

	if equipment.GetWeight().GetPounds() != 2 {
		t.Fatalf("expected weight 2 lb, got %.2f lb", equipment.GetWeight().GetPounds())
	}
}

func TestNewEquipment_AllowsConstructedZeroCostAndWeight(t *testing.T) {
	cost, ok := NewEquipmentCost(0)
	if !ok {
		t.Fatal("expected zero equipment cost to be constructed")
	}

	weight, ok := NewEquipmentWeightOunces(0)
	if !ok {
		t.Fatal("expected zero equipment weight to be constructed")
	}

	if _, ok := NewEquipment(
		EquipmentID("zero-cost-zero-weight"),
		"Zero cost zero weight",
		AdventuringGearEquipmentCategory,
		cost,
		weight,
	); !ok {
		t.Fatal("expected equipment with constructed zero cost and weight to be valid")
	}
}

func TestNewEquipment_RejectsInvalidInputs(t *testing.T) {
	cost, ok := NewEquipmentCost(1)
	if !ok {
		t.Fatal("expected equipment cost to be constructed")
	}

	weight, ok := NewEquipmentWeightOunces(1)
	if !ok {
		t.Fatal("expected equipment weight to be constructed")
	}

	if _, ok := NewEquipmentCost(-1); ok {
		t.Fatal("expected negative equipment cost to be rejected")
	}

	if _, ok := NewEquipmentWeightOunces(-1); ok {
		t.Fatal("expected negative equipment weight to be rejected")
	}

	for _, id := range []EquipmentID{"", " backpack-empty", "backpack-empty ", "\tbackpack-empty"} {
		if _, ok := NewEquipment(id, "Backpack (empty)", AdventuringGearEquipmentCategory, cost, weight); ok {
			t.Fatalf("expected invalid equipment id %q to be rejected", id)
		}
	}

	for _, displayName := range []string{"", " Backpack (empty)", "Backpack (empty) ", "\tBackpack (empty)"} {
		if _, ok := NewEquipment(EquipmentID("backpack-empty"), displayName, AdventuringGearEquipmentCategory, cost, weight); ok {
			t.Fatalf("expected invalid display name %q to be rejected", displayName)
		}
	}

	if _, ok := NewEquipment(EquipmentID("backpack-empty"), "Backpack (empty)", EquipmentCategory("Magic Item"), cost, weight); ok {
		t.Fatal("expected unknown equipment category to be rejected")
	}

	if _, ok := NewEquipment(EquipmentID("backpack-empty"), "Backpack (empty)", AdventuringGearEquipmentCategory, EquipmentCost{}, weight); ok {
		t.Fatal("expected zero-value equipment cost to be rejected")
	}

	if _, ok := NewEquipment(EquipmentID("backpack-empty"), "Backpack (empty)", AdventuringGearEquipmentCategory, cost, EquipmentWeight{}); ok {
		t.Fatal("expected zero-value equipment weight to be rejected")
	}
}
