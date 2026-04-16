package ability

type baseAttackBonusProgression float64
type BaseAttackBonusProgression = baseAttackBonusProgression

const (
	BaseAttackBonusHalf          BaseAttackBonusProgression = 0.5
	BaseAttackBonusThreeQuarters BaseAttackBonusProgression = 0.75
	BaseAttackBonusFull          BaseAttackBonusProgression = 1
)

type baseAttackBonus struct {
	actualValue float64
	value       int
}
type BaseAttackBonus = baseAttackBonus

func NewBaseAttackBonus(actualValue float64) (BaseAttackBonus, bool) {
	bab := baseAttackBonus{}
	if !bab.SetActualValue(actualValue) {
		return baseAttackBonus{}, false
	}

	return bab, true
}

func NewBaseAttackBonusByClassLevel(level int, progression BaseAttackBonusProgression) (BaseAttackBonus, bool) {
	bab := baseAttackBonus{}
	if !bab.SetByClassLevel(level, progression) {
		return baseAttackBonus{}, false
	}

	return bab, true
}

func (b baseAttackBonus) GetActualValue() float64 {
	return b.actualValue
}

func (b baseAttackBonus) GetValue() int {
	return b.value
}

func (b *baseAttackBonus) SetActualValue(actualValue float64) bool {
	if actualValue < 0 {
		return false
	}

	b.actualValue = actualValue
	b.value = roundDown(actualValue)
	return true
}

func (b *baseAttackBonus) SetByClassLevel(level int, progression BaseAttackBonusProgression) bool {
	if level < 0 || !isValidBaseAttackBonusProgression(progression) {
		return false
	}

	return b.SetActualValue(float64(level) * float64(progression))
}

func isValidBaseAttackBonusProgression(progression BaseAttackBonusProgression) bool {
	switch progression {
	case BaseAttackBonusHalf, BaseAttackBonusThreeQuarters, BaseAttackBonusFull:
		return true
	default:
		return false
	}
}

func roundDown(value float64) int {
	return int(value)
}
