package creaturetype

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
)

func TestResolveCreatureRules_UsesBaseProfileWithoutSubtypeEffects(t *testing.T) {
	classification, ok := NewCreatureClassification(AnimalType, nil, nil)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(AnimalType)
	if !ok {
		t.Fatal("expected animal profile to exist")
	}

	rules, ok := ResolveCreatureRules(classification, baseProfile)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if rules.GetHitDieType() != ability.D8HitDie {
		t.Fatalf("expected hit die type %q, got %q", ability.D8HitDie, rules.GetHitDieType())
	}

	if rules.GetBABProgression() != ability.BaseAttackBonusThreeQuarters {
		t.Fatalf("expected BAB progression %q, got %q", ability.BaseAttackBonusThreeQuarters, rules.GetBABProgression())
	}

	if rules.GetHitPointKind() != ability.StandardHitPoints {
		t.Fatalf("expected hit point kind %q, got %q", ability.StandardHitPoints, rules.GetHitPointKind())
	}

	if !rules.HasTrait(LowLightVisionTrait) {
		t.Fatal("expected base profile trait to be present")
	}
}

func TestResolveCreatureRules_AppliesSubtypeTraitOverridesAndFlags(t *testing.T) {
	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{NativeSubtype, IncorporealSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(OutsiderType)
	if !ok {
		t.Fatal("expected outsider profile to exist")
	}

	effect, ok := NewCreatureSubtypeEffect(classification, baseProfile)
	if !ok {
		t.Fatal("expected subtype effect to be constructed")
	}

	rules, ok := ResolveCreatureRules(classification, baseProfile, effect)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if !rules.HasTrait(BreatheEatSleepTrait) {
		t.Fatal("expected native override to add BreatheEatSleep")
	}

	if rules.HasTrait(NoNeedToEatSleepBreatheTrait) {
		t.Fatal("expected native override to remove NoNeedToEatSleepBreathe")
	}

	if !rules.HasTrait(PrecisionDamageImmuneTrait) {
		t.Fatal("expected incorporeal subtype trait to be present")
	}

	if !rules.HasContextualFlag(IncorporealBodyRulesFlag) {
		t.Fatal("expected incorporeal structural flag to be present")
	}
}

func TestResolveCreatureRules_PreservesAugmentedFrom(t *testing.T) {
	originalType := HumanoidType

	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{AugmentedSubtype},
		&originalType,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(OutsiderType)
	if !ok {
		t.Fatal("expected outsider profile to exist")
	}

	effect, ok := NewCreatureSubtypeEffect(classification, baseProfile)
	if !ok {
		t.Fatal("expected subtype effect to be constructed")
	}

	rules, ok := ResolveCreatureRules(classification, baseProfile, effect)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	augmentedFrom, ok := rules.GetAugmentedFrom()
	if !ok || augmentedFrom != HumanoidType {
		t.Fatalf("expected augmented-from (%q, true), got (%q, %t)", HumanoidType, augmentedFrom, ok)
	}
}

func TestResolveCreatureRules_AddsHumanoidContextualFlag(t *testing.T) {
	classification, ok := NewCreatureClassification(HumanoidType, nil, nil)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(HumanoidType)
	if !ok {
		t.Fatal("expected humanoid profile to exist")
	}

	rules, ok := ResolveCreatureRules(classification, baseProfile)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if !rules.HasContextualFlag(HumanoidRacialHDUsesClassRulesFlag) {
		t.Fatal("expected humanoid contextual flag to be present")
	}
}

func TestResolveCreatureRules_RejectsMismatchedBaseProfile(t *testing.T) {
	classification, ok := NewCreatureClassification(AnimalType, nil, nil)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	mismatchedProfile, ok := GetCreatureTypeProfile(DragonType)
	if !ok {
		t.Fatal("expected dragon profile to exist")
	}

	if _, ok := ResolveCreatureRules(classification, mismatchedProfile); ok {
		t.Fatal("expected mismatched base profile to be rejected")
	}
}

func TestResolveCreatureRules_RejectsSubtypeEffectFromDifferentBaseProfile(t *testing.T) {
	classification, ok := NewCreatureClassification(AnimalType, nil, nil)
	if !ok {
		t.Fatal("expected animal classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(AnimalType)
	if !ok {
		t.Fatal("expected animal profile to exist")
	}

	otherClassification, ok := NewCreatureClassification(
		UndeadType,
		[]CreatureSubtypeID{IncorporealSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected undead classification to be constructed")
	}

	otherBaseProfile, ok := GetCreatureTypeProfile(UndeadType)
	if !ok {
		t.Fatal("expected undead profile to exist")
	}

	otherEffect, ok := NewCreatureSubtypeEffect(otherClassification, otherBaseProfile)
	if !ok {
		t.Fatal("expected undead subtype effect to be constructed")
	}

	if _, ok := ResolveCreatureRules(classification, baseProfile, otherEffect); ok {
		t.Fatal("expected mismatched subtype effect to be rejected")
	}
}

func TestResolveCreatureRules_MapsSwarmStructuralOverrideToFlag(t *testing.T) {
	classification, ok := NewCreatureClassification(
		VerminType,
		[]CreatureSubtypeID{SwarmSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(VerminType)
	if !ok {
		t.Fatal("expected vermin profile to exist")
	}

	effect, ok := NewCreatureSubtypeEffect(classification, baseProfile)
	if !ok {
		t.Fatal("expected subtype effect to be constructed")
	}

	rules, ok := ResolveCreatureRules(classification, baseProfile, effect)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if !rules.HasContextualFlag(SwarmBodyRulesFlag) {
		t.Fatal("expected swarm structural flag to be present")
	}
}
