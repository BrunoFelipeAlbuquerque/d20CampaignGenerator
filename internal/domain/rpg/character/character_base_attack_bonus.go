package character

import (
	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

type characterBaseAttackBonusFacts struct {
	baseAttackBonus ability.BaseAttackBonus
}
type CharacterBaseAttackBonusFacts = characterBaseAttackBonusFacts

func NewCharacterBaseAttackBonusFacts(
	classLevels []CharacterClassLevel,
) (CharacterBaseAttackBonusFacts, bool) {
	levelFacts, ok := NewCharacterLevelFacts(classLevels)
	if !ok {
		return characterBaseAttackBonusFacts{}, false
	}

	actualValue, ok := ability.NewRationalValue(0, 1)
	if !ok {
		return characterBaseAttackBonusFacts{}, false
	}

	for _, classLevel := range levelFacts.GetClassLevels() {
		class, ok := characterclass.GetClassByID(classLevel.GetClassID())
		if !ok {
			return characterBaseAttackBonusFacts{}, false
		}

		classBaseAttackBonus, ok := ability.NewBaseAttackBonusByClassLevel(
			classLevel.GetLevel(),
			class.GetBaseAttackBonusProgression(),
		)
		if !ok {
			return characterBaseAttackBonusFacts{}, false
		}

		actualValue = actualValue.Add(classBaseAttackBonus.GetActualValue())
	}

	baseAttackBonus, ok := ability.NewBaseAttackBonus(actualValue)
	if !ok {
		return characterBaseAttackBonusFacts{}, false
	}

	return characterBaseAttackBonusFacts{
		baseAttackBonus: baseAttackBonus,
	}, true
}

func (f characterBaseAttackBonusFacts) GetBaseAttackBonus() ability.BaseAttackBonus {
	return f.baseAttackBonus
}

func (f characterBaseAttackBonusFacts) GetActualValue() ability.RationalValue {
	return f.baseAttackBonus.GetActualValue()
}

func (f characterBaseAttackBonusFacts) GetValue() int {
	return f.baseAttackBonus.GetValue()
}
