package ability

type baseAttackBonusProgression string
type BaseAttackBonusProgression = baseAttackBonusProgression

const (
	BaseAttackBonusHalf          BaseAttackBonusProgression = "1/2"
	BaseAttackBonusThreeQuarters BaseAttackBonusProgression = "3/4"
	BaseAttackBonusFull          BaseAttackBonusProgression = "1/1"
)

type baseAttackBonus struct {
	actualValue rationalValue
	value       int
}
type BaseAttackBonus = baseAttackBonus

func NewBaseAttackBonus(actualValue RationalValue) (BaseAttackBonus, bool) {
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

func (b baseAttackBonus) GetActualValue() RationalValue {
	return b.actualValue
}

func (b baseAttackBonus) GetValue() int {
	return b.value
}

func (b *baseAttackBonus) SetActualValue(actualValue RationalValue) bool {
	if !isNonNegativeRationalValue(actualValue) {
		return false
	}

	b.actualValue = actualValue
	b.value = actualValue.Floor()
	return true
}

func (b *baseAttackBonus) SetByClassLevel(level int, progression BaseAttackBonusProgression) bool {
	if level < 0 {
		return false
	}

	actualValue, ok := progression.toRationalValue()
	if !ok {
		return false
	}

	return b.SetActualValue(actualValue.MultiplyByInt(level))
}

func (p baseAttackBonusProgression) toRationalValue() (RationalValue, bool) {
	switch p {
	case BaseAttackBonusHalf:
		return NewRationalValue(1, 2)
	case BaseAttackBonusThreeQuarters:
		return NewRationalValue(3, 4)
	case BaseAttackBonusFull:
		return NewRationalValue(1, 1)
	default:
		return rationalValue{}, false
	}
}
