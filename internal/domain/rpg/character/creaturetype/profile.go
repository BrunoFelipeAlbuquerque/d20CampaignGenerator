package creaturetype

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

type creatureTypeProfile struct {
	hitDieType       ability.HitDieType
	babProgression   ability.BaseAttackBonusProgression
	goodSaves        []ability.SavingThrowID
	skillPointsPerHD int
	hitPointKind     ability.HitPointKind
	traitIDs         []CreatureTypeTraitID
}

type CreatureTypeProfile = creatureTypeProfile

func NewCreatureTypeProfile(
	hitDieType ability.HitDieType,
	babProgression ability.BaseAttackBonusProgression,
	goodSaves []ability.SavingThrowID,
	skillPointsPerHD int,
	hitPointKind ability.HitPointKind,
	traitIDs []CreatureTypeTraitID,
) (CreatureTypeProfile, bool) {
	if !isValidHitDieType(hitDieType) ||
		!isValidBABProgression(babProgression) ||
		!isValidHitPointKind(hitPointKind) ||
		skillPointsPerHD < 0 {
		return creatureTypeProfile{}, false
	}

	dedupedSaves, ok := dedupeGoodSaves(goodSaves)
	if !ok {
		return creatureTypeProfile{}, false
	}

	dedupedTraits, ok := dedupeTraitIDs(traitIDs)
	if !ok {
		return creatureTypeProfile{}, false
	}

	return creatureTypeProfile{
		hitDieType:       hitDieType,
		babProgression:   babProgression,
		goodSaves:        dedupedSaves,
		skillPointsPerHD: skillPointsPerHD,
		hitPointKind:     hitPointKind,
		traitIDs:         dedupedTraits,
	}, true
}

func (p creatureTypeProfile) GetHitDieType() ability.HitDieType {
	return p.hitDieType
}

func (p creatureTypeProfile) GetBABProgression() ability.BaseAttackBonusProgression {
	return p.babProgression
}

func (p creatureTypeProfile) GetGoodSaves() []ability.SavingThrowID {
	return append([]ability.SavingThrowID(nil), p.goodSaves...)
}

func (p creatureTypeProfile) GetSkillPointsPerHD() int {
	return p.skillPointsPerHD
}

func (p creatureTypeProfile) GetHitPointKind() ability.HitPointKind {
	return p.hitPointKind
}

func (p creatureTypeProfile) GetTraitIDs() []CreatureTypeTraitID {
	return append([]CreatureTypeTraitID(nil), p.traitIDs...)
}

func (p creatureTypeProfile) HasTrait(traitID CreatureTypeTraitID) bool {
	for _, current := range p.traitIDs {
		if current == traitID {
			return true
		}
	}

	return false
}

func dedupeGoodSaves(goodSaves []ability.SavingThrowID) ([]ability.SavingThrowID, bool) {
	if len(goodSaves) == 0 {
		return nil, true
	}

	seen := make(map[ability.SavingThrowID]struct{}, len(goodSaves))
	deduped := make([]ability.SavingThrowID, 0, len(goodSaves))

	for _, save := range goodSaves {
		if !isValidSavingThrowID(save) {
			return nil, false
		}

		if _, ok := seen[save]; ok {
			continue
		}

		seen[save] = struct{}{}
		deduped = append(deduped, save)
	}

	return deduped, true
}

func dedupeTraitIDs(traitIDs []CreatureTypeTraitID) ([]CreatureTypeTraitID, bool) {
	if len(traitIDs) == 0 {
		return nil, true
	}

	seen := make(map[CreatureTypeTraitID]struct{}, len(traitIDs))
	deduped := make([]CreatureTypeTraitID, 0, len(traitIDs))

	for _, traitID := range traitIDs {
		if !isValidCreatureTypeTraitID(traitID) {
			return nil, false
		}

		if _, ok := seen[traitID]; ok {
			continue
		}

		seen[traitID] = struct{}{}
		deduped = append(deduped, traitID)
	}

	return deduped, true
}

func isValidHitDieType(value ability.HitDieType) bool {
	switch value {
	case ability.D6HitDie, ability.D8HitDie, ability.D10HitDie, ability.D12HitDie:
		return true
	default:
		return false
	}
}

func isValidBABProgression(value ability.BaseAttackBonusProgression) bool {
	switch value {
	case ability.BaseAttackBonusHalf, ability.BaseAttackBonusThreeQuarters, ability.BaseAttackBonusFull:
		return true
	default:
		return false
	}
}

func isValidSavingThrowID(value ability.SavingThrowID) bool {
	switch value {
	case ability.FortitudeSave, ability.ReflexSave, ability.WillSave:
		return true
	default:
		return false
	}
}

func isValidHitPointKind(value ability.HitPointKind) bool {
	switch value {
	case ability.StandardHitPoints, ability.UndeadHitPoints, ability.ConstructHitPoints:
		return true
	default:
		return false
	}
}
