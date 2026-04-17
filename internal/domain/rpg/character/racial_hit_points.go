package character

import (
	"d20campaigngenerator/internal/domain/rpg/character/ability"
	"d20campaigngenerator/internal/domain/rpg/character/creaturetype"
)

func NewRacialHitPoints(
	rules creaturetype.ResolvedCreatureRules,
	racialHitDieCount int,
	constitutionScore int,
	charismaScore int,
	size ability.Size,
) (ability.HitPoints, bool) {
	return rules.NewRacialHitPoints(racialHitDieCount, constitutionScore, charismaScore, size)
}
