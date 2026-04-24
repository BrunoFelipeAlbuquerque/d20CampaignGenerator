package spell

import "testing"

func TestNewSpell_ConstructsValidatedSpellChassis(t *testing.T) {
	spell, ok := NewSpell(
		SpellID("light"),
		EvocationSchoolID,
		[]DescriptorID{"Light"},
		[]ComponentID{VerbalComponentID, MaterialComponentID},
		"1 standard action",
		"touch",
		"object touched",
		"10 min./level",
		"none",
		"no",
	)
	if !ok {
		t.Fatal("expected spell chassis to be constructed")
	}

	if spell.GetID() != SpellID("light") {
		t.Fatalf("expected spell id %q, got %q", SpellID("light"), spell.GetID())
	}

	if spell.GetSchool() != EvocationSchoolID {
		t.Fatalf("expected spell school %q, got %q", EvocationSchoolID, spell.GetSchool())
	}

	descriptors := spell.GetDescriptors()
	if len(descriptors) != 1 || descriptors[0] != DescriptorID("Light") {
		t.Fatalf("expected light descriptor, got %v", descriptors)
	}

	components := spell.GetComponents()
	if len(components) != 2 ||
		components[0] != VerbalComponentID ||
		components[1] != MaterialComponentID {
		t.Fatalf("expected verbal and material components, got %v", components)
	}

	if spell.GetCastingTime() != "1 standard action" {
		t.Fatalf("expected casting time %q, got %q", "1 standard action", spell.GetCastingTime())
	}

	if spell.GetRange() != "touch" {
		t.Fatalf("expected range %q, got %q", "touch", spell.GetRange())
	}

	if spell.GetTargetEffect() != "object touched" {
		t.Fatalf("expected target/effect %q, got %q", "object touched", spell.GetTargetEffect())
	}

	if spell.GetDuration() != "10 min./level" {
		t.Fatalf("expected duration %q, got %q", "10 min./level", spell.GetDuration())
	}

	if spell.GetSavingThrow() != "none" {
		t.Fatalf("expected saving throw %q, got %q", "none", spell.GetSavingThrow())
	}

	if spell.GetSpellResistance() != "no" {
		t.Fatalf("expected spell resistance %q, got %q", "no", spell.GetSpellResistance())
	}
}

func TestNewSpell_AllowsSpellsWithoutDescriptors(t *testing.T) {
	spell, ok := NewSpell(
		SpellID("mage armor"),
		ConjurationSchoolID,
		nil,
		[]ComponentID{VerbalComponentID, SomaticComponentID, FocusComponentID},
		"1 standard action",
		"touch",
		"creature touched",
		"1 hour/level",
		"Will negates (harmless)",
		"no",
	)
	if !ok {
		t.Fatal("expected spell without descriptors to be constructed")
	}

	if len(spell.GetDescriptors()) != 0 {
		t.Fatalf("expected no descriptors, got %v", spell.GetDescriptors())
	}
}

func TestNewSpell_DedupesDescriptorsAndComponents(t *testing.T) {
	spell, ok := NewSpell(
		SpellID("light"),
		EvocationSchoolID,
		[]DescriptorID{"Light", "Light"},
		[]ComponentID{VerbalComponentID, VerbalComponentID, MaterialComponentID},
		"1 standard action",
		"touch",
		"object touched",
		"10 min./level",
		"none",
		"no",
	)
	if !ok {
		t.Fatal("expected spell with duplicated metadata to be constructed")
	}

	if len(spell.GetDescriptors()) != 1 {
		t.Fatalf("expected deduped descriptors length 1, got %d", len(spell.GetDescriptors()))
	}

	if len(spell.GetComponents()) != 2 {
		t.Fatalf("expected deduped components length 2, got %d", len(spell.GetComponents()))
	}
}

func TestNewSpell_RejectsInvalidInputs(t *testing.T) {
	if _, ok := NewSpell(
		"",
		DivinationSchoolID,
		nil,
		[]ComponentID{VerbalComponentID},
		"1 standard action",
		"60 ft.",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected empty spell id to be rejected")
	}

	if _, ok := NewSpell(
		SpellID(" detect magic"),
		DivinationSchoolID,
		nil,
		[]ComponentID{VerbalComponentID},
		"1 standard action",
		"60 ft.",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected spell id with surrounding whitespace to be rejected")
	}

	if _, ok := NewSpell(
		SpellID("detect magic"),
		SchoolID("Chronomancy"),
		nil,
		[]ComponentID{VerbalComponentID},
		"1 standard action",
		"60 ft.",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected unknown school to be rejected")
	}

	if _, ok := NewSpell(
		SpellID("detect magic"),
		DivinationSchoolID,
		[]DescriptorID{""},
		[]ComponentID{VerbalComponentID},
		"1 standard action",
		"60 ft.",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected empty descriptor to be rejected")
	}

	if _, ok := NewSpell(
		SpellID("detect magic"),
		DivinationSchoolID,
		nil,
		nil,
		"1 standard action",
		"60 ft.",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected missing components to be rejected")
	}

	if _, ok := NewSpell(
		SpellID("detect magic"),
		DivinationSchoolID,
		nil,
		[]ComponentID{ComponentID("XP")},
		"1 standard action",
		"60 ft.",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected unknown component to be rejected")
	}

	if _, ok := NewSpell(
		SpellID("detect magic"),
		DivinationSchoolID,
		nil,
		[]ComponentID{VerbalComponentID},
		"",
		"60 ft.",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected empty casting time to be rejected")
	}

	if _, ok := NewSpell(
		SpellID("detect magic"),
		DivinationSchoolID,
		nil,
		[]ComponentID{VerbalComponentID},
		"1 standard action",
		"60 ft. ",
		"cone-shaped emanation",
		"concentration, up to 1 min./level",
		"none",
		"no",
	); ok {
		t.Fatal("expected range with surrounding whitespace to be rejected")
	}
}

func TestSpell_GettersReturnDefensiveCopies(t *testing.T) {
	spell, ok := NewSpell(
		SpellID("light"),
		EvocationSchoolID,
		[]DescriptorID{"Light"},
		[]ComponentID{VerbalComponentID, MaterialComponentID},
		"1 standard action",
		"touch",
		"object touched",
		"10 min./level",
		"none",
		"no",
	)
	if !ok {
		t.Fatal("expected spell chassis to be constructed")
	}

	descriptors := spell.GetDescriptors()
	components := spell.GetComponents()

	descriptors[0] = DescriptorID("Darkness")
	components[0] = SomaticComponentID

	if spell.GetDescriptors()[0] != DescriptorID("Light") {
		t.Fatalf("expected defensive descriptor copy to preserve %q", DescriptorID("Light"))
	}

	if spell.GetComponents()[0] != VerbalComponentID {
		t.Fatalf("expected defensive component copy to preserve %q", VerbalComponentID)
	}
}
