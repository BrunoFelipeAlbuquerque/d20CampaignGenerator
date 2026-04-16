package creaturetype

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

func ResolveCreatureRules(
	classification CreatureClassification,
	baseProfile CreatureTypeProfile,
	effects ...CreatureSubtypeEffect,
) (ResolvedCreatureRules, bool) {
	canonicalBaseProfile, ok := GetCreatureTypeProfile(classification.GetBaseType())
	if !ok || !profilesEqual(baseProfile, canonicalBaseProfile) {
		return resolvedCreatureRules{}, false
	}

	traitIDs := baseProfile.GetTraitIDs()
	contextualFlags := make([]ResolvedCreatureRuleFlag, 0, len(effects)+1)
	var augmentedFrom *CreatureTypeID

	for _, effect := range effects {
		if !profilesEqual(effect.GetBaseProfile(), baseProfile) {
			return resolvedCreatureRules{}, false
		}

		traitIDs = applyTraitOverrides(
			traitIDs,
			effect.GetRemovedTraitIDs(),
			effect.GetAddedTraitIDs(),
		)

		for _, structuralEffectID := range effect.GetStructuralEffectIDs() {
			flag, ok := structuralEffectToFlag(structuralEffectID)
			if !ok {
				return resolvedCreatureRules{}, false
			}

			contextualFlags = append(contextualFlags, flag)
		}

		if originalType, ok := effect.GetPreservedOriginalType(); ok {
			if augmentedFrom != nil && *augmentedFrom != originalType {
				return resolvedCreatureRules{}, false
			}

			originalCopy := originalType
			augmentedFrom = &originalCopy
		}
	}

	if classification.GetBaseType() == HumanoidType {
		contextualFlags = append(contextualFlags, HumanoidRacialHDUsesClassRulesFlag)
	}

	dedupedTraits, ok := dedupeResolvedCreatureTraitIDs(traitIDs)
	if !ok {
		return resolvedCreatureRules{}, false
	}

	dedupedFlags, ok := dedupeResolvedCreatureRuleFlags(contextualFlags)
	if !ok {
		return resolvedCreatureRules{}, false
	}

	return resolvedCreatureRules{
		hitDieType:       baseProfile.GetHitDieType(),
		babProgression:   baseProfile.GetBABProgression(),
		goodSaves:        baseProfile.GetGoodSaves(),
		skillPointsPerHD: baseProfile.GetSkillPointsPerHD(),
		hitPointKind:     baseProfile.GetHitPointKind(),
		traitIDs:         dedupedTraits,
		contextualFlags:  dedupedFlags,
		augmentedFrom:    augmentedFrom,
	}, true
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

func structuralEffectToFlag(
	effectID CreatureSubtypeStructuralEffectID,
) (ResolvedCreatureRuleFlag, bool) {
	switch effectID {
	case IncorporealBodyEffect:
		return IncorporealBodyRulesFlag, true
	case SwarmBodyEffect:
		return SwarmBodyRulesFlag, true
	default:
		return "", false
	}
}

func dedupeResolvedCreatureRuleFlags(
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

func dedupeResolvedCreatureTraitIDs(
	traitIDs []CreatureTypeTraitID,
) ([]CreatureTypeTraitID, bool) {
	if len(traitIDs) == 0 {
		return nil, true
	}

	seen := make(map[CreatureTypeTraitID]struct{}, len(traitIDs))
	deduped := make([]CreatureTypeTraitID, 0, len(traitIDs))

	for _, traitID := range traitIDs {
		if !isValidSubtypeEffectTraitID(traitID) {
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

func profilesEqual(left CreatureTypeProfile, right CreatureTypeProfile) bool {
	if left.GetHitDieType() != right.GetHitDieType() ||
		left.GetBABProgression() != right.GetBABProgression() ||
		left.GetSkillPointsPerHD() != right.GetSkillPointsPerHD() ||
		left.GetHitPointKind() != right.GetHitPointKind() {
		return false
	}

	if !savingThrowIDsEqual(left.GetGoodSaves(), right.GetGoodSaves()) {
		return false
	}

	return creatureTypeTraitIDsEqual(left.GetTraitIDs(), right.GetTraitIDs())
}

func savingThrowIDsEqual(left []ability.SavingThrowID, right []ability.SavingThrowID) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func creatureTypeTraitIDsEqual(left []CreatureTypeTraitID, right []CreatureTypeTraitID) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}
