package spell

import (
	"testing"

	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

func TestNewSpellListEntry_ConstructsValidatedEntry(t *testing.T) {
	entry, ok := NewSpellListEntry(
		SpellID("light"),
		characterclass.WizardClassID,
		0,
	)
	if !ok {
		t.Fatal("expected spell list entry to be constructed")
	}

	if entry.GetSpellID() != SpellID("light") {
		t.Fatalf("expected spell id %q, got %q", SpellID("light"), entry.GetSpellID())
	}

	if entry.GetClassID() != characterclass.WizardClassID {
		t.Fatalf("expected class id %q, got %q", characterclass.WizardClassID, entry.GetClassID())
	}

	if entry.GetSpellLevel() != 0 {
		t.Fatalf("expected spell level 0, got %d", entry.GetSpellLevel())
	}
}

func TestNewSpellListEntry_AcceptsDelayedCasterFirstLevelEntries(t *testing.T) {
	entry, ok := NewSpellListEntry(
		SpellID("bless weapon"),
		characterclass.PaladinClassID,
		1,
	)
	if !ok {
		t.Fatal("expected delayed-caster 1st-level spell list entry to be constructed")
	}

	if entry.GetSpellLevel() != 1 {
		t.Fatalf("expected spell level 1, got %d", entry.GetSpellLevel())
	}
}

func TestNewSpellListEntry_RejectsInvalidInputs(t *testing.T) {
	if _, ok := NewSpellListEntry(
		"",
		characterclass.WizardClassID,
		0,
	); ok {
		t.Fatal("expected empty spell id to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID(" light"),
		characterclass.WizardClassID,
		0,
	); ok {
		t.Fatal("expected spell id with surrounding whitespace to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID("light"),
		characterclass.ClassID("oracle"),
		0,
	); ok {
		t.Fatal("expected unknown class id to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID("light"),
		characterclass.FighterClassID,
		0,
	); ok {
		t.Fatal("expected non-spellcasting class id to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID("light"),
		characterclass.WizardClassID,
		-1,
	); ok {
		t.Fatal("expected negative spell level to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID("light"),
		characterclass.WizardClassID,
		10,
	); ok {
		t.Fatal("expected spell level above 9 to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID("limited wish"),
		characterclass.BardClassID,
		7,
	); ok {
		t.Fatal("expected bard spell level above 6 to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID("light"),
		characterclass.PaladinClassID,
		0,
	); ok {
		t.Fatal("expected delayed-caster 0-level spell list entry to be rejected")
	}

	if _, ok := NewSpellListEntry(
		SpellID("break enchantment"),
		characterclass.PaladinClassID,
		5,
	); ok {
		t.Fatal("expected paladin spell level above 4 to be rejected")
	}
}
