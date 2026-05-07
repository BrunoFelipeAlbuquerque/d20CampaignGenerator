package feat

const (
	BrewPotionFeatID             FeatID = "Brew Potion"
	CraftMagicArmsAndArmorFeatID FeatID = "Craft Magic Arms and Armor"
	CraftRodFeatID               FeatID = "Craft Rod"
	CraftStaffFeatID             FeatID = "Craft Staff"
	CraftWandFeatID              FeatID = "Craft Wand"
	CraftWondrousItemFeatID      FeatID = "Craft Wondrous Item"
	ForgeRingFeatID              FeatID = "Forge Ring"
	ScribeScrollFeatID           FeatID = "Scribe Scroll"
)

var coreItemCreationFeats = mustBuildCoreItemCreationFeats()

var coreItemCreationFeatOrder = []FeatID{
	BrewPotionFeatID,
	CraftMagicArmsAndArmorFeatID,
	CraftRodFeatID,
	CraftStaffFeatID,
	CraftWandFeatID,
	CraftWondrousItemFeatID,
	ForgeRingFeatID,
	ScribeScrollFeatID,
}

func mustBuildCoreItemCreationFeats() map[FeatID]Feat {
	return map[FeatID]Feat{
		BrewPotionFeatID:             mustNewCoreItemCreationFeat(BrewPotionFeatID, mustCasterLevelPrerequisite(3)),
		CraftMagicArmsAndArmorFeatID: mustNewCoreItemCreationFeat(CraftMagicArmsAndArmorFeatID, mustCasterLevelPrerequisite(5)),
		CraftRodFeatID:               mustNewCoreItemCreationFeat(CraftRodFeatID, mustCasterLevelPrerequisite(9)),
		CraftStaffFeatID:             mustNewCoreItemCreationFeat(CraftStaffFeatID, mustCasterLevelPrerequisite(11)),
		CraftWandFeatID:              mustNewCoreItemCreationFeat(CraftWandFeatID, mustCasterLevelPrerequisite(5)),
		CraftWondrousItemFeatID:      mustNewCoreItemCreationFeat(CraftWondrousItemFeatID, mustCasterLevelPrerequisite(3)),
		ForgeRingFeatID:              mustNewCoreItemCreationFeat(ForgeRingFeatID, mustCasterLevelPrerequisite(7)),
		ScribeScrollFeatID:           mustNewCoreItemCreationFeat(ScribeScrollFeatID, mustCasterLevelPrerequisite(1)),
	}
}

func mustNewCoreItemCreationFeat(id FeatID, prerequisites ...Prerequisite) Feat {
	prerequisiteList, ok := NewPrerequisiteList(prerequisites)
	if !ok {
		panic("invalid core item creation feat prerequisite seed")
	}

	value, ok := NewFeat(id, ItemCreationFeatCategory, prerequisiteList, false, false, true)
	if !ok {
		panic("invalid core item creation feat seed")
	}

	return value
}
