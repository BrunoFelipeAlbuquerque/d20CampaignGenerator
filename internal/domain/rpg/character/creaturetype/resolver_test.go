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

	rules, ok := ResolveCreatureRules(classification)
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

	if len(rules.GetFixedGoodSaves()) != 2 {
		t.Fatalf("expected 2 fixed good saves, got %d", len(rules.GetFixedGoodSaves()))
	}

	if !rules.HasTrait(LowLightVisionTrait) {
		t.Fatal("expected base profile trait to be present")
	}
}

func TestResolveCreatureRules_AutomaticallyAppliesSubtypeEffects(t *testing.T) {
	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{NativeSubtype, IncorporealSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	rules, ok := ResolveCreatureRules(classification)
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

	rules, ok := ResolveCreatureRules(classification)
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

	rules, ok := ResolveCreatureRules(classification)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if !rules.HasContextualFlag(HumanoidRacialHDUsesClassRulesFlag) {
		t.Fatal("expected humanoid contextual flag to be present")
	}

	if rules.GetSelectableGoodSaveCount() != 1 {
		t.Fatalf("expected humanoid selectable good save count 1, got %d", rules.GetSelectableGoodSaveCount())
	}

	if len(rules.GetFixedGoodSaves()) != 0 {
		t.Fatalf("expected humanoid fixed good saves to be empty, got %v", rules.GetFixedGoodSaves())
	}

	if rules.HasTrait(CreatureTypeTraitID("OneGoodSaveChoice")) {
		t.Fatal("expected humanoid save-choice metadata not to be encoded as a fake trait")
	}
}

func TestResolveCreatureRules_ExposesOutsiderSaveChoiceMetadata(t *testing.T) {
	classification, ok := NewCreatureClassification(OutsiderType, nil, nil)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	rules, ok := ResolveCreatureRules(classification)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if rules.GetSelectableGoodSaveCount() != 2 {
		t.Fatalf("expected outsider selectable good save count 2, got %d", rules.GetSelectableGoodSaveCount())
	}

	if len(rules.GetFixedGoodSaves()) != 0 {
		t.Fatalf("expected outsider fixed good saves to be empty, got %v", rules.GetFixedGoodSaves())
	}

	if rules.HasTrait(CreatureTypeTraitID("TwoGoodSaveChoices")) {
		t.Fatal("expected outsider save-choice metadata not to be encoded as a fake trait")
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

	rules, ok := ResolveCreatureRules(classification)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if !rules.HasContextualFlag(SwarmBodyRulesFlag) {
		t.Fatal("expected swarm structural flag to be present")
	}
}

func TestResolveCreatureRules_RejectsNativeOnNonOutsider(t *testing.T) {
	classification, ok := NewCreatureClassification(
		AnimalType,
		[]CreatureSubtypeID{NativeSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	if _, ok := ResolveCreatureRules(classification); ok {
		t.Fatal("expected native subtype on non-outsider to be rejected")
	}
}

func TestResolveCreatureRules_DedupesResolvedTraits(t *testing.T) {
	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{ElementalSubtype, IncorporealSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	rules, ok := ResolveCreatureRules(classification)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	if countResolvedTrait(rules.GetTraitIDs(), NotSubjectToCriticalHitsTrait) != 1 {
		t.Fatalf("expected NotSubjectToCriticalHits to be deduped, got %d", countResolvedTrait(rules.GetTraitIDs(), NotSubjectToCriticalHitsTrait))
	}
}

func TestResolveCreatureRules_DedupesResolvedFlags(t *testing.T) {
	flags, ok := dedupeResolvedRuleFlags([]ResolvedCreatureRuleFlag{
		SwarmBodyRulesFlag,
		SwarmBodyRulesFlag,
	})
	if !ok {
		t.Fatal("expected resolved flag dedupe to succeed")
	}

	if len(flags) != 1 {
		t.Fatalf("expected 1 deduped flag, got %d", len(flags))
	}
}

func TestIsValidResolvedCreatureRules_RejectsInvalidResolvedMetadata(t *testing.T) {
	if isValidResolvedCreatureRules(resolvedCreatureRules{
		fixedGoodSaves:          []ability.SavingThrowID{ability.FortitudeSave, ability.ReflexSave},
		selectableGoodSaveCount: 2,
	}) {
		t.Fatal("expected impossible good save metadata to be rejected")
	}

	if isValidResolvedCreatureRules(resolvedCreatureRules{
		fixedGoodSaves: []ability.SavingThrowID{ability.SavingThrowID("Luck")},
	}) {
		t.Fatal("expected invalid fixed good saves to be rejected")
	}

	if isValidResolvedCreatureRules(resolvedCreatureRules{
		traitIDs: []CreatureTypeTraitID{CreatureTypeTraitID("FastHealing")},
	}) {
		t.Fatal("expected invalid resolved trait ids to be rejected")
	}

	if isValidResolvedCreatureRules(resolvedCreatureRules{
		contextualFlags: []ResolvedCreatureRuleFlag{ResolvedCreatureRuleFlag("Invalid")},
	}) {
		t.Fatal("expected invalid contextual flags to be rejected")
	}
}

func countResolvedTrait(traitIDs []CreatureTypeTraitID, want CreatureTypeTraitID) int {
	count := 0
	for _, current := range traitIDs {
		if current == want {
			count++
		}
	}

	return count
}
