package equipment

import "strings"

type equipmentID string
type EquipmentID = equipmentID

type equipmentCategory string
type EquipmentCategory = equipmentCategory

const (
	AdventuringGearEquipmentCategory EquipmentCategory = "Adventuring Gear"
)

type equipmentCost struct {
	copperPieces int
	valid        bool
}
type EquipmentCost = equipmentCost

type equipmentWeight struct {
	ounces int
	valid  bool
}
type EquipmentWeight = equipmentWeight

type equipment struct {
	id          equipmentID
	displayName string
	category    equipmentCategory
	cost        equipmentCost
	weight      equipmentWeight
}
type Equipment = equipment

func NewEquipmentCost(copperPieces int) (EquipmentCost, bool) {
	if copperPieces < 0 {
		return equipmentCost{}, false
	}

	return equipmentCost{
		copperPieces: copperPieces,
		valid:        true,
	}, true
}

func NewEquipmentWeightOunces(ounces int) (EquipmentWeight, bool) {
	if ounces < 0 {
		return equipmentWeight{}, false
	}

	return equipmentWeight{
		ounces: ounces,
		valid:  true,
	}, true
}

func NewEquipment(
	id EquipmentID,
	displayName string,
	category EquipmentCategory,
	cost EquipmentCost,
	weight EquipmentWeight,
) (Equipment, bool) {
	if !isValidEquipmentID(id) ||
		!isValidDisplayName(displayName) ||
		!isValidEquipmentCategory(category) ||
		!isValidEquipmentCost(cost) ||
		!isValidEquipmentWeight(weight) {
		return equipment{}, false
	}

	return equipment{
		id:          id,
		displayName: displayName,
		category:    category,
		cost:        cost,
		weight:      weight,
	}, true
}

func (c equipmentCategory) GetName() string {
	if !isValidEquipmentCategory(EquipmentCategory(c)) {
		return ""
	}

	return string(c)
}

func (c equipmentCost) GetCopperPieces() int {
	return c.copperPieces
}

func (w equipmentWeight) GetOunces() int {
	return w.ounces
}

func (w equipmentWeight) GetPounds() float64 {
	return float64(w.ounces) / 16
}

func (e equipment) GetID() EquipmentID {
	return e.id
}

func (e equipment) GetDisplayName() string {
	return e.displayName
}

func (e equipment) GetCategory() EquipmentCategory {
	return e.category
}

func (e equipment) GetCost() EquipmentCost {
	return e.cost
}

func (e equipment) GetWeight() EquipmentWeight {
	return e.weight
}

func isValidEquipmentID(id EquipmentID) bool {
	value := string(id)
	return value != "" && strings.TrimSpace(value) == value
}

func isValidDisplayName(value string) bool {
	return value != "" && strings.TrimSpace(value) == value
}

func isValidEquipmentCategory(category EquipmentCategory) bool {
	switch category {
	case AdventuringGearEquipmentCategory:
		return true
	default:
		return false
	}
}

func isValidEquipmentCost(cost EquipmentCost) bool {
	return cost.valid && cost.copperPieces >= 0
}

func isValidEquipmentWeight(weight EquipmentWeight) bool {
	return weight.valid && weight.ounces >= 0
}
