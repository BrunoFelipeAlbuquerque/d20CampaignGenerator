package creaturetype

import "testing"

func TestNewCreatureSubtypeEffect_AquaticAddsWaterMovementTraits(t *testing.T) {
	classification, ok := NewCreatureClassification(
		AnimalType,
		[]CreatureSubtypeID{AquaticSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(AnimalType)
	if !ok {
		t.Fatal("expected animal profile to exist")
	}

	effect, ok := NewCreatureSubtypeEffect(classification, baseProfile)
	if !ok {
		t.Fatal("expected subtype effect to be constructed")
	}

	if !effect.HasAddedTrait(BreathesWaterTrait) {
		t.Fatal("expected aquatic effect to add BreathesWater")
	}

	if !effect.HasAddedTrait(SwimWithoutChecksTrait) {
		t.Fatal("expected aquatic effect to add SwimWithoutChecks")
	}

	if !effect.HasAddedTrait(SwimAlwaysClassSkillTrait) {
		t.Fatal("expected aquatic effect to add SwimAlwaysClassSkill")
	}
}

func TestNewCreatureSubtypeEffect_AugmentedPreservesOriginalTypeMetadata(t *testing.T) {
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

	preservedType, ok := effect.GetPreservedOriginalType()
	if !ok || preservedType != HumanoidType {
		t.Fatalf("expected preserved original type (%q, true), got (%q, %t)", HumanoidType, preservedType, ok)
	}
}

func TestNewCreatureSubtypeEffect_ElementalOverridesLivingNeedsAndAddsStructuralTraits(t *testing.T) {
	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{ElementalSubtype},
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

	if !effect.HasAddedTrait(NoNeedToEatSleepBreatheTrait) {
		t.Fatal("expected elemental effect to add NoNeedToEatSleepBreathe")
	}

	if !effect.HasRemovedTrait(BreatheEatSleepTrait) {
		t.Fatal("expected elemental effect to remove BreatheEatSleep")
	}

	if !effect.HasAddedTrait(ImmunityBleedTrait) {
		t.Fatal("expected elemental effect to add ImmunityBleed")
	}

	if !effect.HasAddedTrait(NotSubjectToFlankingTrait) {
		t.Fatal("expected elemental effect to add NotSubjectToFlanking")
	}
}

func TestNewCreatureSubtypeEffect_IncorporealAddsStructuralBodyEffect(t *testing.T) {
	classification, ok := NewCreatureClassification(
		UndeadType,
		[]CreatureSubtypeID{IncorporealSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(UndeadType)
	if !ok {
		t.Fatal("expected undead profile to exist")
	}

	effect, ok := NewCreatureSubtypeEffect(classification, baseProfile)
	if !ok {
		t.Fatal("expected subtype effect to be constructed")
	}

	if !effect.HasStructuralEffect(IncorporealBodyEffect) {
		t.Fatal("expected incorporeal effect to add IncorporealBody structural effect")
	}

	if !effect.HasAddedTrait(PrecisionDamageImmuneTrait) {
		t.Fatal("expected incorporeal effect to add precision damage immunity")
	}
}

func TestNewCreatureSubtypeEffect_NativeIsOnlyValidForOutsiders(t *testing.T) {
	outsiderClassification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{NativeSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected outsider classification to be constructed")
	}

	outsiderProfile, ok := GetCreatureTypeProfile(OutsiderType)
	if !ok {
		t.Fatal("expected outsider profile to exist")
	}

	effect, ok := NewCreatureSubtypeEffect(outsiderClassification, outsiderProfile)
	if !ok {
		t.Fatal("expected native outsider subtype effect to be constructed")
	}

	if !effect.HasAddedTrait(BreatheEatSleepTrait) {
		t.Fatal("expected native effect to add BreatheEatSleep")
	}

	if !effect.HasRemovedTrait(NoNeedToEatSleepBreatheTrait) {
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

	animalProfile, ok := GetCreatureTypeProfile(AnimalType)
	if !ok {
		t.Fatal("expected animal profile to exist")
	}

	if _, ok := NewCreatureSubtypeEffect(animalClassification, animalProfile); ok {
		t.Fatal("expected native subtype on non-outsider to be rejected")
	}
}

func TestNewCreatureSubtypeEffect_SwarmUsesStructuralOverride(t *testing.T) {
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

	if !effect.HasStructuralEffect(SwarmBodyEffect) {
		t.Fatal("expected swarm effect to add SwarmBody structural effect")
	}
}

func TestNewCreatureSubtypeEffect_KeepsBaseProfileReference(t *testing.T) {
	classification, ok := NewCreatureClassification(
		AnimalType,
		[]CreatureSubtypeID{AquaticSubtype},
		nil,
	)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	baseProfile, ok := GetCreatureTypeProfile(AnimalType)
	if !ok {
		t.Fatal("expected animal profile to exist")
	}

	effect, ok := NewCreatureSubtypeEffect(classification, baseProfile)
	if !ok {
		t.Fatal("expected subtype effect to be constructed")
	}

	if effect.GetBaseProfile().GetHitDieType() != baseProfile.GetHitDieType() {
		t.Fatalf(
			"expected base profile hit die type %q, got %q",
			baseProfile.GetHitDieType(),
			effect.GetBaseProfile().GetHitDieType(),
		)
	}
}
