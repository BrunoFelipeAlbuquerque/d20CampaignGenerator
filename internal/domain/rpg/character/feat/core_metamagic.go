package feat

const (
	EmpowerSpellFeatID  FeatID = "Empower Spell"
	EnlargeSpellFeatID  FeatID = "Enlarge Spell"
	ExtendSpellFeatID   FeatID = "Extend Spell"
	HeightenSpellFeatID FeatID = "Heighten Spell"
	MaximizeSpellFeatID FeatID = "Maximize Spell"
	QuickenSpellFeatID  FeatID = "Quicken Spell"
	SilentSpellFeatID   FeatID = "Silent Spell"
	StillSpellFeatID    FeatID = "Still Spell"
	WidenSpellFeatID    FeatID = "Widen Spell"
)

var coreMetamagicFeats = mustBuildCoreMetamagicFeats()

var coreMetamagicFeatOrder = []FeatID{
	EmpowerSpellFeatID,
	EnlargeSpellFeatID,
	ExtendSpellFeatID,
	HeightenSpellFeatID,
	MaximizeSpellFeatID,
	QuickenSpellFeatID,
	SilentSpellFeatID,
	StillSpellFeatID,
	WidenSpellFeatID,
}

func mustBuildCoreMetamagicFeats() map[FeatID]Feat {
	return map[FeatID]Feat{
		EmpowerSpellFeatID:  mustNewCoreMetamagicFeat(EmpowerSpellFeatID),
		EnlargeSpellFeatID:  mustNewCoreMetamagicFeat(EnlargeSpellFeatID),
		ExtendSpellFeatID:   mustNewCoreMetamagicFeat(ExtendSpellFeatID),
		HeightenSpellFeatID: mustNewCoreMetamagicFeat(HeightenSpellFeatID),
		MaximizeSpellFeatID: mustNewCoreMetamagicFeat(MaximizeSpellFeatID),
		QuickenSpellFeatID:  mustNewCoreMetamagicFeat(QuickenSpellFeatID),
		SilentSpellFeatID:   mustNewCoreMetamagicFeat(SilentSpellFeatID),
		StillSpellFeatID:    mustNewCoreMetamagicFeat(StillSpellFeatID),
		WidenSpellFeatID:    mustNewCoreMetamagicFeat(WidenSpellFeatID),
	}
}

func mustNewCoreMetamagicFeat(id FeatID, prerequisites ...Prerequisite) Feat {
	prerequisiteList, ok := NewPrerequisiteList(prerequisites)
	if !ok {
		panic("invalid core metamagic feat prerequisite seed")
	}

	value, ok := NewFeat(id, MetamagicFeatCategory, prerequisiteList, false, true, false)
	if !ok {
		panic("invalid core metamagic feat seed")
	}

	return value
}
