package character

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

type characterLoadCategory string
type CharacterLoadCategory = characterLoadCategory

const (
	LightLoadCategory       CharacterLoadCategory = "Light"
	MediumLoadCategory      CharacterLoadCategory = "Medium"
	HeavyLoadCategory       CharacterLoadCategory = "Heavy"
	OverMaximumLoadCategory CharacterLoadCategory = "Over Maximum"
)

type characterCarriedWeight struct {
	totalOunces int
	load        characterLoadCategory
}
type CharacterCarriedWeight = characterCarriedWeight

func NewCharacterCarriedWeight(
	strength ability.AbilityScore,
	equipment []CharacterEquipment,
) (CharacterCarriedWeight, bool) {
	capacity, ok := strength.GetCarryingCapacity()
	if !ok {
		return characterCarriedWeight{}, false
	}

	totalOunces, ok := totalCarriedEquipmentOunces(equipment)
	if !ok {
		return characterCarriedWeight{}, false
	}

	load, ok := resolveCharacterLoadCategory(totalOunces, capacity)
	if !ok {
		return characterCarriedWeight{}, false
	}

	return characterCarriedWeight{
		totalOunces: totalOunces,
		load:        load,
	}, true
}

func (w characterCarriedWeight) GetTotalOunces() int {
	return w.totalOunces
}

func (w characterCarriedWeight) GetTotalPounds() float64 {
	return float64(w.totalOunces) / 16
}

func (w characterCarriedWeight) GetLoadCategory() CharacterLoadCategory {
	return w.load
}

func totalCarriedEquipmentOunces(equipment []CharacterEquipment) (int, bool) {
	total := 0

	for _, selectedEquipment := range equipment {
		if selectedEquipment.quantity <= 0 {
			return 0, false
		}

		resolvedEquipment, ok := selectedEquipment.GetEquipment()
		if !ok {
			return 0, false
		}

		weightOunces := resolvedEquipment.GetWeight().GetOunces()
		if weightOunces < 0 || wouldOverflowWeightTotal(total, weightOunces, selectedEquipment.quantity) {
			return 0, false
		}

		total += weightOunces * selectedEquipment.quantity
	}

	return total, true
}

func resolveCharacterLoadCategory(
	totalOunces int,
	capacity ability.StrengthCarryingCapacity,
) (CharacterLoadCategory, bool) {
	if totalOunces < 0 {
		return "", false
	}

	totalPounds := float64(totalOunces) / 16
	if poundsWithinCapacity(totalPounds, capacity.GetLightLoadMax().GetPounds()) {
		return LightLoadCategory, true
	}

	mediumLoad := capacity.GetMediumLoad()
	if poundsWithinCapacity(totalPounds, mediumLoad.GetMax().GetPounds()) {
		return MediumLoadCategory, true
	}

	heavyLoad := capacity.GetHeavyLoad()
	if poundsWithinCapacity(totalPounds, heavyLoad.GetMax().GetPounds()) {
		return HeavyLoadCategory, true
	}

	return OverMaximumLoadCategory, true
}

func poundsWithinCapacity(totalPounds float64, capacityPounds float64) bool {
	const epsilon = 0.01

	return totalPounds <= capacityPounds+epsilon
}

func wouldOverflowWeightTotal(total int, weightOunces int, quantity int) bool {
	const maxInt = int(^uint(0) >> 1)

	if weightOunces == 0 {
		return false
	}

	return quantity > (maxInt-total)/weightOunces
}
