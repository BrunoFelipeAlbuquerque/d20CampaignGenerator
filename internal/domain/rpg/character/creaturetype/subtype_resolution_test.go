package creaturetype

import "testing"

func TestBuildSubtypeResolution_AquaticAddsWaterMovementTraits(t *testing.T) {
	classification, ok := NewCreatureClassification(
		AnimalType,
		[]CreatureSubtypeID{AquaticSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	resolution, ok := buildSubtypeResolution(classification)
	if !ok {
		t.Fatal("expected subtype resolution to be constructed")
	}

	if len(resolution.addedTraitIDs) != 3 {
		t.Fatalf("expected aquatic to add 3 traits, got %d", len(resolution.addedTraitIDs))
	}

	if !hasTraitID(resolution.addedTraitIDs, BreathesWaterTrait) {
		t.Fatal("expected aquatic effect to add BreathesWater")
	}

	if !hasTraitID(resolution.addedTraitIDs, SwimWithoutChecksTrait) {
		t.Fatal("expected aquatic effect to add SwimWithoutChecks")
	}

	if !hasTraitID(resolution.addedTraitIDs, SwimAlwaysClassSkillTrait) {
		t.Fatal("expected aquatic effect to add SwimAlwaysClassSkill")
	}
}

func TestBuildSubtypeResolution_AugmentedPreservesOriginalTypeMetadata(t *testing.T) {
	originalType := HumanoidType

	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{AugmentedSubtype},
		&originalType,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	resolution, ok := buildSubtypeResolution(classification)
	if !ok {
		t.Fatal("expected subtype resolution to be constructed")
	}

	preservedType, ok := getOptionalCreatureTypeID(resolution.augmentedFrom)
	if !ok || preservedType != HumanoidType {
		t.Fatalf("expected preserved original type (%q, true), got (%q, %t)", HumanoidType, preservedType, ok)
	}
}

func TestBuildSubtypeResolution_ElementalOverridesLivingNeedsAndAddsStructuralTraits(t *testing.T) {
	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{ElementalSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	resolution, ok := buildSubtypeResolution(classification)
	if !ok {
		t.Fatal("expected subtype resolution to be constructed")
	}

	if !hasTraitID(resolution.addedTraitIDs, NoNeedToEatSleepBreatheTrait) {
		t.Fatal("expected elemental effect to add NoNeedToEatSleepBreathe")
	}

	if !hasTraitID(resolution.removedTraitIDs, BreatheEatSleepTrait) {
		t.Fatal("expected elemental effect to remove BreatheEatSleep")
	}

	if !hasTraitID(resolution.addedTraitIDs, ImmunityBleedTrait) {
		t.Fatal("expected elemental effect to add ImmunityBleed")
	}

	if !hasTraitID(resolution.addedTraitIDs, NotSubjectToFlankingTrait) {
		t.Fatal("expected elemental effect to add NotSubjectToFlanking")
	}
}

func TestBuildSubtypeResolution_IncorporealAddsStructuralBodyEffect(t *testing.T) {
	classification, ok := NewCreatureClassification(
		UndeadType,
		[]CreatureSubtypeID{IncorporealSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	resolution, ok := buildSubtypeResolution(classification)
	if !ok {
		t.Fatal("expected subtype resolution to be constructed")
	}

	if !hasResolvedRuleFlag(resolution.contextualFlags, IncorporealBodyRulesFlag) {
		t.Fatal("expected incorporeal effect to add IncorporealBodyRules flag")
	}

	if !hasTraitID(resolution.addedTraitIDs, PrecisionDamageImmuneTrait) {
		t.Fatal("expected incorporeal effect to add precision damage immunity")
	}
}

func TestBuildSubtypeResolution_NativeIsOnlyValidForOutsiders(t *testing.T) {
	outsiderClassification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{NativeSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected outsider classification to be constructed")
	}

	resolution, ok := buildSubtypeResolution(outsiderClassification)
	if !ok {
		t.Fatal("expected native outsider subtype resolution to be constructed")
	}

	if !hasTraitID(resolution.addedTraitIDs, BreatheEatSleepTrait) {
		t.Fatal("expected native effect to add BreatheEatSleep")
	}

	if !hasTraitID(resolution.removedTraitIDs, NoNeedToEatSleepBreatheTrait) {
		t.Fatal("expected native effect to remove NoNeedToEatSleepBreathe")
	}

	animalClassification, ok := NewCreatureClassification(
		AnimalType,
		[]CreatureSubtypeID{NativeSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected animal classification to be constructed")
	}

	if _, ok := buildSubtypeResolution(animalClassification); ok {
		t.Fatal("expected native subtype on non-outsider to be rejected")
	}
}

func TestBuildSubtypeResolution_SwarmUsesStructuralOverride(t *testing.T) {
	classification, ok := NewCreatureClassification(
		VerminType,
		[]CreatureSubtypeID{SwarmSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	resolution, ok := buildSubtypeResolution(classification)
	if !ok {
		t.Fatal("expected subtype resolution to be constructed")
	}

	if !hasResolvedRuleFlag(resolution.contextualFlags, SwarmBodyRulesFlag) {
		t.Fatal("expected swarm effect to add SwarmBodyRules flag")
	}
}

func TestBuildSubtypeResolution_DedupesResolvedTraitsAndFlags(t *testing.T) {
	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{ElementalSubtype, IncorporealSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	resolution, ok := buildSubtypeResolution(classification)
	if !ok {
		t.Fatal("expected subtype resolution to be constructed")
	}

	if countTraitID(resolution.addedTraitIDs, NotSubjectToCriticalHitsTrait) != 1 {
		t.Fatalf("expected NotSubjectToCriticalHits to be deduped, got %d", countTraitID(resolution.addedTraitIDs, NotSubjectToCriticalHitsTrait))
	}

	flags, ok := dedupeResolvedRuleFlags([]ResolvedCreatureRuleFlag{
		SwarmBodyRulesFlag,
		SwarmBodyRulesFlag,
	})
	if !ok {
		t.Fatal("expected resolved flags to dedupe successfully")
	}

	if len(flags) != 1 {
		t.Fatalf("expected deduped flags length 1, got %d", len(flags))
	}
}

func hasTraitID(traitIDs []CreatureTypeTraitID, want CreatureTypeTraitID) bool {
	for _, current := range traitIDs {
		if current == want {
			return true
		}
	}

	return false
}

func countTraitID(traitIDs []CreatureTypeTraitID, want CreatureTypeTraitID) int {
	count := 0
	for _, current := range traitIDs {
		if current == want {
			count++
		}
	}

	return count
}

func hasResolvedRuleFlag(flags []ResolvedCreatureRuleFlag, want ResolvedCreatureRuleFlag) bool {
	for _, current := range flags {
		if current == want {
			return true
		}
	}

	return false
}

func getOptionalCreatureTypeID(value *CreatureTypeID) (CreatureTypeID, bool) {
	if value == nil {
		return "", false
	}

	return *value, true
}
