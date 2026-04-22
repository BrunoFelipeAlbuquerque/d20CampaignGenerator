package creaturetype

type subtypeResolution struct {
	addedTraitIDs   []CreatureTypeTraitID
	removedTraitIDs []CreatureTypeTraitID
	contextualFlags []ResolvedCreatureRuleFlag
	augmentedFrom   *CreatureTypeID
}

type subtypeResolutionSpec struct {
	requiredBaseType              *CreatureTypeID
	addedTraitIDs                 []CreatureTypeTraitID
	removedTraitIDs               []CreatureTypeTraitID
	contextualFlags               []ResolvedCreatureRuleFlag
	preservesOriginalTypeMetadata bool
}

var subtypeResolutionSpecs = map[CreatureSubtypeID]subtypeResolutionSpec{
	AquaticSubtype: {
		addedTraitIDs: []CreatureTypeTraitID{
			BreathesWaterTrait,
			SwimWithoutChecksTrait,
			SwimAlwaysClassSkillTrait,
		},
	},
	AugmentedSubtype: {
		preservesOriginalTypeMetadata: true,
	},
	ElementalSubtype: {
		addedTraitIDs: []CreatureTypeTraitID{
			NoNeedToEatSleepBreatheTrait,
			ImmunityBleedTrait,
			ImmunityParalysisTrait,
			ImmunityPoisonTrait,
			ImmunitySleepTrait,
			ImmunityStunTrait,
			NotSubjectToCriticalHitsTrait,
			NotSubjectToFlankingTrait,
		},
		removedTraitIDs: []CreatureTypeTraitID{
			BreatheEatSleepTrait,
			BreatheNoNeedToEatSleepTrait,
		},
	},
	IncorporealSubtype: {
		addedTraitIDs: []CreatureTypeTraitID{
			NotSubjectToCriticalHitsTrait,
			PrecisionDamageImmuneTrait,
		},
		contextualFlags: []ResolvedCreatureRuleFlag{
			IncorporealBodyRulesFlag,
		},
	},
	NativeSubtype: {
		requiredBaseType: creatureTypePointer(OutsiderType),
		addedTraitIDs: []CreatureTypeTraitID{
			BreatheEatSleepTrait,
		},
		removedTraitIDs: []CreatureTypeTraitID{
			BreatheNoNeedToEatSleepTrait,
		},
	},
	SwarmSubtype: {
		contextualFlags: []ResolvedCreatureRuleFlag{
			SwarmBodyRulesFlag,
		},
	},
}

func buildSubtypeResolution(classification CreatureClassification) (subtypeResolution, bool) {
	var added []CreatureTypeTraitID
	var removed []CreatureTypeTraitID
	var contextualFlags []ResolvedCreatureRuleFlag
	var augmentedFrom *CreatureTypeID

	for _, subtype := range classification.GetSubtypes() {
		spec, ok := subtypeResolutionSpecs[subtype]
		if !ok {
			return subtypeResolution{}, false
		}

		if spec.requiredBaseType != nil && classification.GetBaseType() != *spec.requiredBaseType {
			return subtypeResolution{}, false
		}

		added = append(added, spec.addedTraitIDs...)
		removed = append(removed, spec.removedTraitIDs...)
		contextualFlags = append(contextualFlags, spec.contextualFlags...)

		if spec.preservesOriginalTypeMetadata {
			originalType, ok := classification.GetAugmentedFrom()
			if !ok {
				return subtypeResolution{}, false
			}

			originalCopy := originalType
			augmentedFrom = &originalCopy
		}
	}

	dedupedAdded, ok := dedupeResolvedTraitIDs(added)
	if !ok {
		return subtypeResolution{}, false
	}

	dedupedRemoved, ok := dedupeResolvedTraitIDs(removed)
	if !ok {
		return subtypeResolution{}, false
	}

	dedupedContextualFlags, ok := dedupeResolvedRuleFlags(contextualFlags)
	if !ok {
		return subtypeResolution{}, false
	}

	return subtypeResolution{
		addedTraitIDs:   dedupedAdded,
		removedTraitIDs: dedupedRemoved,
		contextualFlags: dedupedContextualFlags,
		augmentedFrom:   augmentedFrom,
	}, true
}

func dedupeResolvedTraitIDs(traitIDs []CreatureTypeTraitID) ([]CreatureTypeTraitID, bool) {
	if len(traitIDs) == 0 {
		return nil, true
	}

	seen := make(map[CreatureTypeTraitID]struct{}, len(traitIDs))
	deduped := make([]CreatureTypeTraitID, 0, len(traitIDs))

	for _, traitID := range traitIDs {
		if !isValidResolvedTraitID(traitID) {
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

func creatureTypePointer(value CreatureTypeID) *CreatureTypeID {
	return &value
}
