package equipment

import (
	"math"
	"strings"
)

type equipmentID string
type EquipmentID = equipmentID

type equipmentCategory string
type EquipmentCategory = equipmentCategory

const (
	AdventuringGearEquipmentCategory EquipmentCategory = "Adventuring Gear"
)

const (
	kilogramsPerOunce = 0.028349523125
	kilogramsPerPound = kilogramsPerOunce * 16
	metersPerFoot     = 0.3048
)

type equipmentCost struct {
	copperPieces int
	valid        bool
}
type EquipmentCost = equipmentCost

type equipmentWeight struct {
	ounces    int
	kilograms float64
	valid     bool
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

	return newEquipmentWeight(ounces, ouncesToKilograms(ounces))
}

func NewEquipmentWeightKilograms(kilograms float64) (EquipmentWeight, bool) {
	if !isValidMetricValue(kilograms, true) {
		return equipmentWeight{}, false
	}

	return newEquipmentWeight(kilogramsToOunces(kilograms), kilograms)
}

func newEquipmentWeight(ounces int, kilograms float64) (EquipmentWeight, bool) {
	if ounces < 0 ||
		!isValidMetricValue(kilograms, true) ||
		kilogramsToOunces(kilograms) != ounces {
		return equipmentWeight{}, false
	}

	return equipmentWeight{
		ounces:    ounces,
		kilograms: kilograms,
		valid:     true,
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
	return w.kilograms / kilogramsPerPound
}

func (w equipmentWeight) GetKilograms() float64 {
	return w.kilograms
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
	return weight.valid &&
		weight.ounces >= 0 &&
		isValidMetricValue(weight.kilograms, true) &&
		kilogramsToOunces(weight.kilograms) == weight.ounces
}

func isValidMetricValue(value float64, allowZero bool) bool {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return false
	}

	if allowZero {
		return value >= 0
	}

	return value > 0
}

func ouncesToKilograms(ounces int) float64 {
	return float64(ounces) * kilogramsPerOunce
}

func kilogramsToOunces(kilograms float64) int {
	return int(math.Round(kilograms / kilogramsPerOunce))
}

func feetToMeters(feet int) float64 {
	return float64(feet) * metersPerFoot
}

func metersToFeet(meters float64) int {
	return int(math.Round(meters / metersPerFoot))
}
