package equipment

const (
	BackpackEmptyEquipmentID      EquipmentID = "backpack-empty"
	BedrollEquipmentID            EquipmentID = "bedroll"
	FlintAndSteelEquipmentID      EquipmentID = "flint-and-steel"
	PouchBeltEmptyEquipmentID     EquipmentID = "pouch-belt-empty"
	RationsTrailPerDayEquipmentID EquipmentID = "rations-trail-per-day"
	RopeHemp50FeetEquipmentID     EquipmentID = "rope-hemp-50-ft"
	TorchEquipmentID              EquipmentID = "torch"
	WaterskinEquipmentID          EquipmentID = "waterskin"
)

var coreEquipment = mustBuildCoreEquipment()

var coreEquipmentOrder = []EquipmentID{
	BackpackEmptyEquipmentID,
	BedrollEquipmentID,
	FlintAndSteelEquipmentID,
	PouchBeltEmptyEquipmentID,
	RationsTrailPerDayEquipmentID,
	RopeHemp50FeetEquipmentID,
	TorchEquipmentID,
	WaterskinEquipmentID,
}

func GetEquipmentByID(id EquipmentID) (Equipment, bool) {
	value, ok := coreEquipment[id]
	if !ok {
		return equipment{}, false
	}

	return cloneEquipment(value), true
}

func GetEquipment() []Equipment {
	equipment := make([]Equipment, 0, len(coreEquipmentOrder))

	for _, id := range coreEquipmentOrder {
		equipment = append(equipment, cloneEquipment(coreEquipment[id]))
	}

	return equipment
}

func mustBuildCoreEquipment() map[EquipmentID]Equipment {
	return map[EquipmentID]Equipment{
		BackpackEmptyEquipmentID:      mustNewCoreAdventuringGear(BackpackEmptyEquipmentID, "Backpack (empty)", 200, 32),
		BedrollEquipmentID:            mustNewCoreAdventuringGear(BedrollEquipmentID, "Bedroll", 10, 80),
		FlintAndSteelEquipmentID:      mustNewCoreAdventuringGear(FlintAndSteelEquipmentID, "Flint and steel", 100, 0),
		PouchBeltEmptyEquipmentID:     mustNewCoreAdventuringGear(PouchBeltEmptyEquipmentID, "Pouch, belt (empty)", 100, 8),
		RationsTrailPerDayEquipmentID: mustNewCoreAdventuringGear(RationsTrailPerDayEquipmentID, "Rations, trail (per day)", 50, 16),
		RopeHemp50FeetEquipmentID:     mustNewCoreAdventuringGear(RopeHemp50FeetEquipmentID, "Rope, hemp (50 ft.)", 100, 160),
		TorchEquipmentID:              mustNewCoreAdventuringGear(TorchEquipmentID, "Torch", 1, 16),
		WaterskinEquipmentID:          mustNewCoreAdventuringGear(WaterskinEquipmentID, "Waterskin", 100, 64),
	}
}

func mustNewCoreAdventuringGear(id EquipmentID, displayName string, copperPieces int, ounces int) Equipment {
	cost, ok := NewEquipmentCost(copperPieces)
	if !ok {
		panic("invalid core equipment cost seed")
	}

	weight, ok := NewEquipmentWeightOunces(ounces)
	if !ok {
		panic("invalid core equipment weight seed")
	}

	equipment, ok := NewEquipment(id, displayName, AdventuringGearEquipmentCategory, cost, weight)
	if !ok {
		panic("invalid core adventuring gear seed")
	}

	return equipment
}

func cloneEquipment(value Equipment) Equipment {
	return equipment{
		id:          value.id,
		displayName: value.displayName,
		category:    value.category,
		cost:        value.cost,
		weight:      value.weight,
	}
}
