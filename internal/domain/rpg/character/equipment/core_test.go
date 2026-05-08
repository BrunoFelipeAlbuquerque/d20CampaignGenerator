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
