package feat

var coreFeatOrder = buildCoreFeatOrder()

func init() {
	feats := GetFeats()
	if !validateCoreFeatPrerequisiteReferences(feats, seededCoreFeatIDs(feats)) {
		panic("missing referenced core feat seed")
	}
}

func GetFeatByID(id FeatID) (Feat, bool) {
	if value, ok := coreGeneralFeats[id]; ok {
		return cloneFeat(value), true
	}

	if value, ok := coreCombatFeats[id]; ok {
		return cloneFeat(value), true
	}

	if value, ok := coreCriticalFeats[id]; ok {
		return cloneFeat(value), true
	}

	if value, ok := coreItemCreationFeats[id]; ok {
		return cloneFeat(value), true
	}

	if value, ok := coreMetamagicFeats[id]; ok {
		return cloneFeat(value), true
	}

	return feat{}, false
}

func GetFeats() []Feat {
	feats := make([]Feat, 0, len(coreFeatOrder))

	for _, id := range coreFeatOrder {
		value, ok := GetFeatByID(id)
		if !ok {
			panic("missing ordered core feat seed")
		}

		feats = append(feats, value)
	}

	return feats
}

func buildCoreFeatOrder() []FeatID {
	order := make(
		[]FeatID,
		0,
		len(coreGeneralFeatOrder)+
			len(coreCombatFeatOrder)+
			len(coreCriticalFeatOrder)+
			len(coreItemCreationFeatOrder)+
			len(coreMetamagicFeatOrder),
	)

	order = append(order, coreGeneralFeatOrder...)
	order = append(order, coreCombatFeatOrder...)
	order = append(order, coreCriticalFeatOrder...)
	order = append(order, coreItemCreationFeatOrder...)
	order = append(order, coreMetamagicFeatOrder...)

	return order
}

func seededCoreFeatIDs(feats []Feat) map[FeatID]struct{} {
	seededIDs := make(map[FeatID]struct{}, len(feats))

	for _, feat := range feats {
		seededIDs[feat.GetID()] = struct{}{}
	}

	return seededIDs
}

func validateCoreFeatPrerequisiteReferences(feats []Feat, seededIDs map[FeatID]struct{}) bool {
	for _, feat := range feats {
		if _, ok := seededIDs[feat.GetID()]; !ok {
			return false
		}

		for _, prerequisite := range feat.GetPrerequisites() {
			for _, referencedID := range coreFeatPrerequisiteReferences(prerequisite) {
				if _, ok := seededIDs[referencedID]; !ok {
					return false
				}
			}
		}
	}

	return true
}

func coreFeatPrerequisiteReferences(prerequisite Prerequisite) []FeatID {
	switch value := prerequisite.(type) {
	case FeatPrerequisite:
		return []FeatID{value.GetFeatID()}
	case AnyFeatPrerequisite:
		return value.GetFeatIDs()
	case SameSelectionFeatPrerequisite:
		return []FeatID{value.GetFeatID()}
	case SpellSchoolFeatPrerequisite:
		return []FeatID{value.GetFeatID()}
	default:
		return nil
	}
}

func cloneFeat(value Feat) Feat {
	return feat{
		id:               value.id,
		category:         value.category,
		prerequisites:    clonePrerequisiteList(value.prerequisites),
		fighterBonusFeat: value.fighterBonusFeat,
		metamagic:        value.metamagic,
		itemCreation:     value.itemCreation,
	}
}
