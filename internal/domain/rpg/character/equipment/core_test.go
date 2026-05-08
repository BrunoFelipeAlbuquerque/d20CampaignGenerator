package equipment

import "testing"

func TestCoreEquipment_SeedsAdventuringGearBatchOne(t *testing.T) {
	expected := map[EquipmentID]struct {
		displayName  string
		copperPieces int
		ounces       int
	}{
		BackpackEmptyEquipmentID:      {displayName: "Backpack (empty)", copperPieces: 200, ounces: 32},
		BedrollEquipmentID:            {displayName: "Bedroll", copperPieces: 10, ounces: 80},
		FlintAndSteelEquipmentID:      {displayName: "Flint and steel", copperPieces: 100, ounces: 0},
		PouchBeltEmptyEquipmentID:     {displayName: "Pouch, belt (empty)", copperPieces: 100, ounces: 8},
		RationsTrailPerDayEquipmentID: {displayName: "Rations, trail (per day)", copperPieces: 50, ounces: 16},
		RopeHemp50FeetEquipmentID:     {displayName: "Rope, hemp (50 ft.)", copperPieces: 100, ounces: 160},
		TorchEquipmentID:              {displayName: "Torch", copperPieces: 1, ounces: 16},
		WaterskinEquipmentID:          {displayName: "Waterskin", copperPieces: 100, ounces: 64},
	}

	if len(coreEquipment) != len(expected) {
		t.Fatalf("expected %d core equipment seeds, got %d", len(expected), len(coreEquipment))
	}

	for id, expectation := range expected {
		equipment, ok := coreEquipment[id]
		if !ok {
			t.Fatalf("expected core equipment seed %q", id)
		}

		if equipment.GetID() != id {
			t.Fatalf("expected equipment id %q, got %q", id, equipment.GetID())
		}

		if equipment.GetDisplayName() != expectation.displayName {
			t.Fatalf("expected display name %q, got %q", expectation.displayName, equipment.GetDisplayName())
		}

		if equipment.GetCategory() != AdventuringGearEquipmentCategory {
			t.Fatalf("expected adventuring gear category for %q, got %q", id, equipment.GetCategory())
		}

		if equipment.GetCost().GetCopperPieces() != expectation.copperPieces {
			t.Fatalf("expected %q cost %d cp, got %d cp", id, expectation.copperPieces, equipment.GetCost().GetCopperPieces())
		}

		if equipment.GetWeight().GetOunces() != expectation.ounces {
			t.Fatalf("expected %q weight %d oz, got %d oz", id, expectation.ounces, equipment.GetWeight().GetOunces())
		}
	}
}

func TestCoreEquipment_OrderContainsOnlySeededBatchOneIDs(t *testing.T) {
	expectedOrder := []EquipmentID{
		BackpackEmptyEquipmentID,
		BedrollEquipmentID,
		FlintAndSteelEquipmentID,
		PouchBeltEmptyEquipmentID,
		RationsTrailPerDayEquipmentID,
		RopeHemp50FeetEquipmentID,
		TorchEquipmentID,
		WaterskinEquipmentID,
	}

	if len(coreEquipmentOrder) != len(expectedOrder) {
		t.Fatalf("expected order length %d, got %d", len(expectedOrder), len(coreEquipmentOrder))
	}

	seen := make(map[EquipmentID]struct{}, len(coreEquipmentOrder))
	for index, id := range coreEquipmentOrder {
		if id != expectedOrder[index] {
			t.Fatalf("expected ordered equipment id at index %d to be %q, got %q", index, expectedOrder[index], id)
		}

		if _, ok := coreEquipment[id]; !ok {
			t.Fatalf("expected ordered equipment id %q to have a seed", id)
		}

		if _, ok := seen[id]; ok {
			t.Fatalf("expected ordered equipment id %q not to be duplicated", id)
		}

		seen[id] = struct{}{}
	}
}

func TestGetEquipmentByID_ReturnsSeededCoreEquipment(t *testing.T) {
	equipment, ok := GetEquipmentByID(BackpackEmptyEquipmentID)
	if !ok {
		t.Fatal("expected backpack to be returned from core equipment lookup")
	}

	if equipment.GetID() != BackpackEmptyEquipmentID {
		t.Fatalf("expected equipment id %q, got %q", BackpackEmptyEquipmentID, equipment.GetID())
	}

	if equipment.GetDisplayName() != "Backpack (empty)" {
		t.Fatalf("expected display name %q, got %q", "Backpack (empty)", equipment.GetDisplayName())
	}

	if equipment.GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected backpack cost 200 cp, got %d cp", equipment.GetCost().GetCopperPieces())
	}

	if equipment.GetWeight().GetOunces() != 32 {
		t.Fatalf("expected backpack weight 32 oz, got %d oz", equipment.GetWeight().GetOunces())
	}
}

func TestGetEquipmentByID_ReturnsDetachedCopy(t *testing.T) {
	first, ok := GetEquipmentByID(BackpackEmptyEquipmentID)
	if !ok {
		t.Fatal("expected backpack to be returned from core equipment lookup")
	}

	first.id = "changed"
	first.displayName = "Changed"
	first.category = "Changed"
	first.cost.copperPieces = 999
	first.weight.ounces = 999

	second, ok := GetEquipmentByID(BackpackEmptyEquipmentID)
	if !ok {
		t.Fatal("expected backpack to be returned from core equipment lookup")
	}

	if second.GetID() != BackpackEmptyEquipmentID {
		t.Fatalf("expected stored equipment id to remain %q, got %q", BackpackEmptyEquipmentID, second.GetID())
	}

	if second.GetDisplayName() != "Backpack (empty)" {
		t.Fatalf("expected stored display name to remain %q, got %q", "Backpack (empty)", second.GetDisplayName())
	}

	if second.GetCategory() != AdventuringGearEquipmentCategory {
		t.Fatalf("expected stored category to remain %q, got %q", AdventuringGearEquipmentCategory, second.GetCategory())
	}

	if second.GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected stored cost to remain 200 cp, got %d cp", second.GetCost().GetCopperPieces())
	}

	if second.GetWeight().GetOunces() != 32 {
		t.Fatalf("expected stored weight to remain 32 oz, got %d oz", second.GetWeight().GetOunces())
	}
}

func TestGetEquipmentByID_RejectsUnknownEquipment(t *testing.T) {
	if _, ok := GetEquipmentByID(EquipmentID("ten-foot-pole")); ok {
		t.Fatal("expected unknown equipment lookup to fail")
	}
}

func TestGetEquipment_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	equipment := GetEquipment()
	if len(equipment) != len(coreEquipmentOrder) {
		t.Fatalf("expected %d queried equipment entries, got %d", len(coreEquipmentOrder), len(equipment))
	}

	for i, expectedID := range coreEquipmentOrder {
		if equipment[i].GetID() != expectedID {
			t.Fatalf("expected equipment at index %d to be %q, got %q", i, expectedID, equipment[i].GetID())
		}
	}
}

func TestGetEquipment_ReturnsDetachedCopies(t *testing.T) {
	first := GetEquipment()
	second := GetEquipment()

	first[0].id = "changed"
	first[0].displayName = "Changed"
	first[0].category = "Changed"
	first[0].cost.copperPieces = 999
	first[0].weight.ounces = 999

	if second[0].GetID() != BackpackEmptyEquipmentID {
		t.Fatalf("expected stored equipment id to remain %q, got %q", BackpackEmptyEquipmentID, second[0].GetID())
	}

	if second[0].GetDisplayName() != "Backpack (empty)" {
		t.Fatalf("expected stored display name to remain %q, got %q", "Backpack (empty)", second[0].GetDisplayName())
	}

	if second[0].GetCategory() != AdventuringGearEquipmentCategory {
		t.Fatalf("expected stored category to remain %q, got %q", AdventuringGearEquipmentCategory, second[0].GetCategory())
	}

	if second[0].GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected stored cost to remain 200 cp, got %d cp", second[0].GetCost().GetCopperPieces())
	}

	if second[0].GetWeight().GetOunces() != 32 {
		t.Fatalf("expected stored weight to remain 32 oz, got %d oz", second[0].GetWeight().GetOunces())
	}
}
