package feat

var coreFeatOrder = buildCoreFeatOrder()

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
