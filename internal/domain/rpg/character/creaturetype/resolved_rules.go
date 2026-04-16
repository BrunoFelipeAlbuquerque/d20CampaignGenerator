package creaturetype

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

type resolvedCreatureRuleFlag string
type ResolvedCreatureRuleFlag = resolvedCreatureRuleFlag

const (
	HumanoidRacialHDUsesClassRulesFlag ResolvedCreatureRuleFlag = "HumanoidRacialHDUsesClassRules"
	IncorporealBodyRulesFlag           ResolvedCreatureRuleFlag = "IncorporealBodyRules"
	SwarmBodyRulesFlag                 ResolvedCreatureRuleFlag = "SwarmBodyRules"
)

type resolvedCreatureRules struct {
	hitDieType              ability.HitDieType
	babProgression          ability.BaseAttackBonusProgression
	fixedGoodSaves          []ability.SavingThrowID
	selectableGoodSaveCount int
	skillPointsPerHD        int
	hitPointKind            ability.HitPointKind
	traitIDs                []CreatureTypeTraitID
	contextualFlags         []ResolvedCreatureRuleFlag
	augmentedFrom           *CreatureTypeID
}

type ResolvedCreatureRules = resolvedCreatureRules

func (r resolvedCreatureRules) GetHitDieType() ability.HitDieType {
	return r.hitDieType
}

func (r resolvedCreatureRules) GetBABProgression() ability.BaseAttackBonusProgression {
	return r.babProgression
}

func (r resolvedCreatureRules) GetFixedGoodSaves() []ability.SavingThrowID {
	return append([]ability.SavingThrowID(nil), r.fixedGoodSaves...)
}

func (r resolvedCreatureRules) GetSelectableGoodSaveCount() int {
	return r.selectableGoodSaveCount
}

func (r resolvedCreatureRules) GetSkillPointsPerHD() int {
	return r.skillPointsPerHD
}

func (r resolvedCreatureRules) GetHitPointKind() ability.HitPointKind {
	return r.hitPointKind
}

func (r resolvedCreatureRules) NewRacialHitDie(hitDieCount int) (ability.HitDie, bool) {
	return ability.NewUniformHitDie(r.hitDieType, hitDieCount)
}

func (r resolvedCreatureRules) NewHitPoints(
	hd ability.HitDie,
	constitutionScore int,
	charismaScore int,
	size ability.Size,
) (ability.HitPoints, bool) {
	switch r.hitPointKind {
	case ability.StandardHitPoints:
		return ability.NewStandardHitPoints(hd, constitutionScore)
	case ability.UndeadHitPoints:
		return ability.NewUndeadHitPoints(hd, charismaScore)
	case ability.ConstructHitPoints:
		return ability.NewConstructHitPoints(hd, size)
	default:
		return ability.HitPoints{}, false
	}
}

func (r resolvedCreatureRules) NewRacialHitPoints(
	hitDieCount int,
	constitutionScore int,
	charismaScore int,
	size ability.Size,
) (ability.HitPoints, bool) {
	hd, ok := r.NewRacialHitDie(hitDieCount)
	if !ok {
		return ability.HitPoints{}, false
	}

	return r.NewHitPoints(hd, constitutionScore, charismaScore, size)
}

func (r resolvedCreatureRules) GetTraitIDs() []CreatureTypeTraitID {
	return append([]CreatureTypeTraitID(nil), r.traitIDs...)
}

func (r resolvedCreatureRules) GetContextualFlags() []ResolvedCreatureRuleFlag {
	return append([]ResolvedCreatureRuleFlag(nil), r.contextualFlags...)
}

func (r resolvedCreatureRules) HasTrait(traitID CreatureTypeTraitID) bool {
	for _, current := range r.traitIDs {
		if current == traitID {
			return true
		}
	}

	return false
}

func (r resolvedCreatureRules) HasContextualFlag(flag ResolvedCreatureRuleFlag) bool {
	for _, current := range r.contextualFlags {
		if current == flag {
			return true
		}
	}

	return false
}

func (r resolvedCreatureRules) GetAugmentedFrom() (CreatureTypeID, bool) {
	if r.augmentedFrom == nil {
		return "", false
	}

	return *r.augmentedFrom, true
}
