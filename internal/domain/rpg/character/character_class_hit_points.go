package character

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

func NewFirstLevelCharacterHitPoints(
	class CharacterClass,
	constitutionScore int,
) (ability.HitPoints, bool) {
	resolvedClass, ok := class.GetClass()
	if !ok {
		return ability.HitPoints{}, false
	}

	hd, ok := ability.NewUniformHitDie(resolvedClass.GetHitDieType(), 1)
	if !ok {
		return ability.HitPoints{}, false
	}

	return ability.NewMaximumStandardHitPoints(hd, constitutionScore)
}
