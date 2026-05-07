package feat

const (
	BleedingCriticalFeatID   FeatID = "Bleeding Critical"
	BlindingCriticalFeatID   FeatID = "Blinding Critical"
	DeafeningCriticalFeatID  FeatID = "Deafening Critical"
	ExhaustingCriticalFeatID FeatID = "Exhausting Critical"
	SickeningCriticalFeatID  FeatID = "Sickening Critical"
	StaggeringCriticalFeatID FeatID = "Staggering Critical"
	StunningCriticalFeatID   FeatID = "Stunning Critical"
	TiringCriticalFeatID     FeatID = "Tiring Critical"
)

var coreCriticalFeats = mustBuildCoreCriticalFeats()

var coreCriticalFeatOrder = []FeatID{
	BleedingCriticalFeatID,
	BlindingCriticalFeatID,
	DeafeningCriticalFeatID,
	ExhaustingCriticalFeatID,
	SickeningCriticalFeatID,
	StaggeringCriticalFeatID,
	StunningCriticalFeatID,
	TiringCriticalFeatID,
}

func mustBuildCoreCriticalFeats() map[FeatID]Feat {
	return map[FeatID]Feat{
		BleedingCriticalFeatID:   mustNewCoreCriticalFeat(BleedingCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustBaseAttackBonusPrerequisite(11)),
		BlindingCriticalFeatID:   mustNewCoreCriticalFeat(BlindingCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustBaseAttackBonusPrerequisite(15)),
		DeafeningCriticalFeatID:  mustNewCoreCriticalFeat(DeafeningCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustBaseAttackBonusPrerequisite(13)),
		ExhaustingCriticalFeatID: mustNewCoreCriticalFeat(ExhaustingCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustFeatPrerequisite(TiringCriticalFeatID), mustBaseAttackBonusPrerequisite(15)),
		SickeningCriticalFeatID:  mustNewCoreCriticalFeat(SickeningCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustBaseAttackBonusPrerequisite(11)),
		StaggeringCriticalFeatID: mustNewCoreCriticalFeat(StaggeringCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustBaseAttackBonusPrerequisite(13)),
		StunningCriticalFeatID:   mustNewCoreCriticalFeat(StunningCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustFeatPrerequisite(StaggeringCriticalFeatID), mustBaseAttackBonusPrerequisite(17)),
		TiringCriticalFeatID:     mustNewCoreCriticalFeat(TiringCriticalFeatID, mustFeatPrerequisite(CriticalFocusFeatID), mustBaseAttackBonusPrerequisite(13)),
	}
}

func mustNewCoreCriticalFeat(id FeatID, prerequisites ...Prerequisite) Feat {
	prerequisiteList, ok := NewPrerequisiteList(prerequisites)
	if !ok {
		panic("invalid core critical feat prerequisite seed")
	}

	value, ok := NewFeat(id, CriticalFeatCategory, prerequisiteList, true, false, false)
	if !ok {
		panic("invalid core critical feat seed")
	}

	return value
}
