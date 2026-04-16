package creaturetype

import "testing"

func TestNewCreatureClassification_UsesBaseTypeSubtypesAndAugmentedFrom(t *testing.T) {
	originalType := HumanoidType

	classification, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{NativeSubtype, AugmentedSubtype},
		&originalType,
	)
	if !ok {
		t.Fatal("expected creature classification to be constructed")
	}

	if classification.GetBaseType() != OutsiderType {
		t.Fatalf("expected base type %q, got %q", OutsiderType, classification.GetBaseType())
	}

	if !classification.HasSubtype(NativeSubtype) {
		t.Fatal("expected Native subtype to be present")
	}

	if !classification.HasSubtype(AugmentedSubtype) {
		t.Fatal("expected Augmented subtype to be present")
	}

	augmentedFrom, ok := classification.GetAugmentedFrom()
	if !ok || augmentedFrom != HumanoidType {
		t.Fatalf("expected augmented-from (%q, true), got (%q, %t)", HumanoidType, augmentedFrom, ok)
	}
}

func TestNewCreatureClassification_RejectsInvalidBaseType(t *testing.T) {
	if _, ok := NewCreatureClassification(
		CreatureTypeID("Vehicle"),
		[]CreatureSubtypeID{AquaticSubtype},
		nil,
	); ok {
		t.Fatal("expected invalid base type to be rejected")
	}
}

func TestNewCreatureClassification_RejectsInvalidSubtype(t *testing.T) {
	if _, ok := NewCreatureClassification(
		AnimalType,
		[]CreatureSubtypeID{CreatureSubtypeID("Vehicle")},
		nil,
	); ok {
		t.Fatal("expected invalid subtype to be rejected")
	}
}

func TestNewCreatureClassification_DedupesSubtypes(t *testing.T) {
	classification, ok := NewCreatureClassification(
		AnimalType,
		[]CreatureSubtypeID{
			AquaticSubtype,
			AquaticSubtype,
			SwarmSubtype,
			SwarmSubtype,
		},
		nil,
	)
	if !ok {
		t.Fatal("expected creature classification to be constructed")
	}

	subtypes := classification.GetSubtypes()
	if len(subtypes) != 2 {
		t.Fatalf("expected 2 deduped subtypes, got %d", len(subtypes))
	}

	if subtypes[0] != AquaticSubtype || subtypes[1] != SwarmSubtype {
		t.Fatalf("expected deduped subtype order [Aquatic Swarm], got %v", subtypes)
	}
}

func TestNewCreatureClassification_RejectsAugmentedFromWithoutAugmentedSubtype(t *testing.T) {
	originalType := HumanoidType

	if _, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{NativeSubtype},
		&originalType,
	); ok {
		t.Fatal("expected augmented-from without Augmented subtype to be rejected")
	}
}

func TestNewCreatureClassification_RejectsAugmentedSubtypeWithoutAugmentedFrom(t *testing.T) {
	if _, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{AugmentedSubtype},
		nil,
	); ok {
		t.Fatal("expected Augmented subtype without augmented-from to be rejected")
	}
}

func TestNewCreatureClassification_RejectsInvalidAugmentedFromType(t *testing.T) {
	invalidType := CreatureTypeID("Vehicle")

	if _, ok := NewCreatureClassification(
		OutsiderType,
		[]CreatureSubtypeID{AugmentedSubtype},
		&invalidType,
	); ok {
		t.Fatal("expected invalid augmented-from type to be rejected")
	}
}

func TestNewCreatureClassification_RejectsAugmentedFromMatchingBaseType(t *testing.T) {
	baseType := DragonType

	if _, ok := NewCreatureClassification(
		DragonType,
		[]CreatureSubtypeID{AugmentedSubtype},
		&baseType,
	); ok {
		t.Fatal("expected augmented-from matching base type to be rejected")
	}
}
