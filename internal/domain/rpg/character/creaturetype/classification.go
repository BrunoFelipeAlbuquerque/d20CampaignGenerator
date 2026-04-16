package creaturetype

type creatureClassification struct {
	baseType      CreatureTypeID
	subtypes      []CreatureSubtypeID
	augmentedFrom *CreatureTypeID
}
type CreatureClassification = creatureClassification

func NewCreatureClassification(
	baseType CreatureTypeID,
	subtypes []CreatureSubtypeID,
	augmentedFrom *CreatureTypeID,
) (CreatureClassification, bool) {
	if !isValidCreatureTypeID(baseType) {
		return creatureClassification{}, false
	}

	dedupedSubtypes, ok := dedupeSubtypes(subtypes)
	if !ok {
		return creatureClassification{}, false
	}

	if !isValidAugmentedConfiguration(baseType, dedupedSubtypes, augmentedFrom) {
		return creatureClassification{}, false
	}

	classification := creatureClassification{
		baseType: baseType,
		subtypes: dedupedSubtypes,
	}

	if augmentedFrom != nil {
		from := *augmentedFrom
		classification.augmentedFrom = &from
	}

	return classification, true
}

func (c creatureClassification) GetBaseType() CreatureTypeID {
	return c.baseType
}

func (c creatureClassification) GetSubtypes() []CreatureSubtypeID {
	return append([]CreatureSubtypeID(nil), c.subtypes...)
}

func (c creatureClassification) GetAugmentedFrom() (CreatureTypeID, bool) {
	if c.augmentedFrom == nil {
		return "", false
	}

	return *c.augmentedFrom, true
}

func (c creatureClassification) HasSubtype(subtype CreatureSubtypeID) bool {
	for _, current := range c.subtypes {
		if current == subtype {
			return true
		}
	}

	return false
}

func dedupeSubtypes(subtypes []CreatureSubtypeID) ([]CreatureSubtypeID, bool) {
	if len(subtypes) == 0 {
		return nil, true
	}

	seen := make(map[CreatureSubtypeID]struct{}, len(subtypes))
	deduped := make([]CreatureSubtypeID, 0, len(subtypes))

	for _, subtype := range subtypes {
		if !isValidCreatureSubtypeID(subtype) {
			return nil, false
		}

		if _, ok := seen[subtype]; ok {
			continue
		}

		seen[subtype] = struct{}{}
		deduped = append(deduped, subtype)
	}

	return deduped, true
}

func isValidAugmentedConfiguration(
	baseType CreatureTypeID,
	subtypes []CreatureSubtypeID,
	augmentedFrom *CreatureTypeID,
) bool {
	hasAugmented := false
	for _, subtype := range subtypes {
		if subtype == AugmentedSubtype {
			hasAugmented = true
			break
		}
	}

	if !hasAugmented {
		return augmentedFrom == nil
	}

	if augmentedFrom == nil {
		return false
	}

	return isValidCreatureTypeID(*augmentedFrom) && *augmentedFrom != baseType
}
