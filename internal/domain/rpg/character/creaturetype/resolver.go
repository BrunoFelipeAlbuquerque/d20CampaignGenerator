package creaturetype

func ResolveCreatureRules(classification CreatureClassification) (ResolvedCreatureRules, bool) {
	baseProfile, ok := GetCreatureTypeProfile(classification.GetBaseType())
	if !ok {
		return resolvedCreatureRules{}, false
	}

	subtypeResolution, ok := buildSubtypeResolution(classification)
	if !ok {
		return resolvedCreatureRules{}, false
	}

	traitIDs := baseProfile.GetTraitIDs()
	traitIDs = applyTraitOverrides(
		traitIDs,
		subtypeResolution.removedTraitIDs,
		subtypeResolution.addedTraitIDs,
	)

	contextualFlags := append([]ResolvedCreatureRuleFlag(nil), subtypeResolution.contextualFlags...)

	if classification.GetBaseType() == HumanoidType {
		contextualFlags = append(contextualFlags, HumanoidRacialHDUsesClassRulesFlag)
	}

	dedupedTraits, ok := dedupeResolvedTraitIDs(traitIDs)
	if !ok {
		return resolvedCreatureRules{}, false
	}

	dedupedFlags, ok := dedupeResolvedRuleFlags(contextualFlags)
	if !ok {
		return resolvedCreatureRules{}, false
	}

	resolved := resolvedCreatureRules{
		hitDieType:              baseProfile.GetHitDieType(),
		babProgression:          baseProfile.GetBABProgression(),
		fixedGoodSaves:          baseProfile.GetFixedGoodSaves(),
		selectableGoodSaveCount: baseProfile.GetSelectableGoodSaveCount(),
		skillPointsPerHD:        baseProfile.GetSkillPointsPerHD(),
		hitPointKind:            baseProfile.GetHitPointKind(),
		traitIDs:                dedupedTraits,
		contextualFlags:         dedupedFlags,
		augmentedFrom:           subtypeResolution.augmentedFrom,
	}
	if !isValidResolvedCreatureRules(resolved) {
		return resolvedCreatureRules{}, false
	}

	return resolved, true
}

func applyTraitOverrides(
	baseTraitIDs []CreatureTypeTraitID,
	removedTraitIDs []CreatureTypeTraitID,
	addedTraitIDs []CreatureTypeTraitID,
) []CreatureTypeTraitID {
	if len(baseTraitIDs) == 0 && len(addedTraitIDs) == 0 {
		return nil
	}

	removed := make(map[CreatureTypeTraitID]struct{}, len(removedTraitIDs))
	for _, traitID := range removedTraitIDs {
		removed[traitID] = struct{}{}
	}

	resolved := make([]CreatureTypeTraitID, 0, len(baseTraitIDs)+len(addedTraitIDs))
	for _, traitID := range baseTraitIDs {
		if _, ok := removed[traitID]; ok {
			continue
		}

		resolved = append(resolved, traitID)
	}

	resolved = append(resolved, addedTraitIDs...)
	return resolved
}

func dedupeResolvedRuleFlags(
	flags []ResolvedCreatureRuleFlag,
) ([]ResolvedCreatureRuleFlag, bool) {
	if len(flags) == 0 {
		return nil, true
	}

	seen := make(map[ResolvedCreatureRuleFlag]struct{}, len(flags))
	deduped := make([]ResolvedCreatureRuleFlag, 0, len(flags))

	for _, flag := range flags {
		if !isValidResolvedCreatureRuleFlag(flag) {
			return nil, false
		}

		if _, ok := seen[flag]; ok {
			continue
		}

		seen[flag] = struct{}{}
		deduped = append(deduped, flag)
	}

	return deduped, true
}

func isValidResolvedCreatureRuleFlag(flag ResolvedCreatureRuleFlag) bool {
	switch flag {
	case HumanoidRacialHDUsesClassRulesFlag,
		IncorporealBodyRulesFlag,
		SwarmBodyRulesFlag:
		return true
	default:
		return false
	}
}

func isValidResolvedCreatureRules(r resolvedCreatureRules) bool {
	if r.selectableGoodSaveCount < 0 || r.selectableGoodSaveCount > 3 {
		return false
	}

	if len(r.fixedGoodSaves)+r.selectableGoodSaveCount > 3 {
		return false
	}

	for _, save := range r.fixedGoodSaves {
		if !isValidSavingThrowID(save) {
			return false
		}
	}

	for _, traitID := range r.traitIDs {
		if !isValidResolvedTraitID(traitID) {
			return false
		}
	}

	for _, flag := range r.contextualFlags {
		if !isValidResolvedCreatureRuleFlag(flag) {
			return false
		}
	}

	return true
}
