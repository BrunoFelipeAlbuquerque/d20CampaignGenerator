package character

import (
	"testing"

	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterspell "d20campaigngenerator/internal/domain/rpg/character/spell"
)

func TestNewCharacterSpellListEntry_ComposesSeededEntryWithClassSpellcasting(t *testing.T) {
	progression := mustNewCharacterSpellcastingProgressionForTest(t, characterclass.WizardClassID)
	entry := mustCoreSpellListEntryForTest(
		t,
		characterspell.SpellID("Fireball"),
		characterclass.WizardClassID,
		3,
	)

	value, ok := NewCharacterSpellListEntry(progression, entry)
	if !ok {
		t.Fatal("expected seeded spell list entry to compose with class spellcasting")
	}

	if value.GetSpellID() != characterspell.SpellID("Fireball") {
		t.Fatalf("expected spell id %q, got %q", characterspell.SpellID("Fireball"), value.GetSpellID())
	}

	if value.GetClassID() != characterclass.WizardClassID {
		t.Fatalf("expected class id %q, got %q", characterclass.WizardClassID, value.GetClassID())
	}

	if value.GetSpellLevel() != 3 {
		t.Fatalf("expected spell level 3, got %d", value.GetSpellLevel())
	}

	spell, ok := value.GetSpell()
	if !ok {
		t.Fatal("expected composed spell list entry spell to resolve")
	}

	if spell.GetID() != characterspell.SpellID("Fireball") {
		t.Fatalf("expected resolved spell %q, got %q", characterspell.SpellID("Fireball"), spell.GetID())
	}

	resolvedEntry, ok := value.GetSpellListEntry()
	if !ok {
		t.Fatal("expected composed spell list entry to resolve")
	}

	if resolvedEntry.GetSpellID() != characterspell.SpellID("Fireball") ||
		resolvedEntry.GetClassID() != characterclass.WizardClassID ||
		resolvedEntry.GetSpellLevel() != 3 {
		t.Fatalf("expected resolved spell list entry Fireball/Wizard/3, got %v", resolvedEntry)
	}
}

func TestGetCharacterSpellListEntriesBySpellLevel_ReturnsSelectedClassEntries(t *testing.T) {
	progression := mustNewCharacterSpellcastingProgressionForTest(t, characterclass.WizardClassID)

	entries, ok := GetCharacterSpellListEntriesBySpellLevel(progression, 3)
	if !ok {
		t.Fatal("expected wizard 3rd-level spell list entries to resolve")
	}

	if !hasCharacterSpellListEntry(entries, characterspell.SpellID("Fireball")) {
		t.Fatalf("expected wizard 3rd-level spell list to include %q", characterspell.SpellID("Fireball"))
	}

	if hasCharacterSpellListEntry(entries, characterspell.SpellID("Cure Serious Wounds")) {
		t.Fatalf("expected wizard 3rd-level spell list not to include %q", characterspell.SpellID("Cure Serious Wounds"))
	}
}

func TestGetCharacterSpellListEntries_ReturnsDetachedSelectedClassEntries(t *testing.T) {
	progression := mustNewCharacterSpellcastingProgressionForTest(t, characterclass.BardClassID)

	first, ok := GetCharacterSpellListEntries(progression)
	if !ok {
		t.Fatal("expected bard spell list entries to resolve")
	}

	if len(first) == 0 {
		t.Fatal("expected bard spell list entries")
	}

	first[0] = CharacterSpellListEntry{}

	second, ok := GetCharacterSpellListEntries(progression)
	if !ok {
		t.Fatal("expected bard spell list entries to resolve again")
	}

	if second[0].GetSpellID() != characterspell.SpellID("Dancing Lights") {
		t.Fatalf("expected stored bard first spell list entry to remain Dancing Lights, got %q", second[0].GetSpellID())
	}
}

func TestGetCharacterSpellListEntries_ComposesDelayedCasterSpellList(t *testing.T) {
	progression := mustNewCharacterSpellcastingProgressionForTest(t, characterclass.PaladinClassID)

	if _, ok := GetCharacterSpellListEntriesBySpellLevel(progression, 0); ok {
		t.Fatal("expected paladin 0-level spell list lookup to fail")
	}

	entries, ok := GetCharacterSpellListEntriesBySpellLevel(progression, 1)
	if !ok {
		t.Fatal("expected paladin 1st-level spell list entries to resolve")
	}

	if !hasCharacterSpellListEntry(entries, characterspell.SpellID("Cure Light Wounds")) {
		t.Fatalf("expected paladin 1st-level spell list to include %q", characterspell.SpellID("Cure Light Wounds"))
	}
}

func TestNewCharacterSpellListEntry_RejectsMismatchedClassSpellcasting(t *testing.T) {
	progression := mustNewCharacterSpellcastingProgressionForTest(t, characterclass.WizardClassID)
	entry := mustCoreSpellListEntryForTest(
		t,
		characterspell.SpellID("Cure Light Wounds"),
		characterclass.PaladinClassID,
		1,
	)

	if _, ok := NewCharacterSpellListEntry(progression, entry); ok {
		t.Fatal("expected spell list entry from a different class to be rejected")
	}
}

func TestNewCharacterSpellListEntry_RejectsShapeValidUnseededClassListEntry(t *testing.T) {
	progression := mustNewCharacterSpellcastingProgressionForTest(t, characterclass.WizardClassID)
	entry, ok := characterspell.NewSpellListEntry(
		characterspell.SpellID("Cure Light Wounds"),
		characterclass.WizardClassID,
		1,
	)
	if !ok {
		t.Fatal("expected shape-valid spell list entry to construct")
	}

	if _, ok := NewCharacterSpellListEntry(progression, entry); ok {
		t.Fatal("expected unseeded class spell list entry to be rejected")
	}
}

func TestGetCharacterSpellListEntries_RejectsZeroValueProgression(t *testing.T) {
	var progression CharacterSpellcastingProgression

	if _, ok := GetCharacterSpellListEntries(progression); ok {
		t.Fatal("expected zero-value progression spell list lookup to fail")
	}

	if _, ok := GetCharacterSpellListEntriesBySpellLevel(progression, 0); ok {
		t.Fatal("expected zero-value progression spell list level lookup to fail")
	}
}

func TestCharacterSpellListEntry_ZeroValueDoesNotResolve(t *testing.T) {
	var entry CharacterSpellListEntry

	if _, ok := entry.GetSpell(); ok {
		t.Fatal("expected zero-value spell list entry spell not to resolve")
	}

	if _, ok := entry.GetSpellListEntry(); ok {
		t.Fatal("expected zero-value spell list entry not to resolve")
	}
}

func mustNewCharacterSpellcastingProgressionForTest(
	t *testing.T,
	classID characterclass.ClassID,
) CharacterSpellcastingProgression {
	t.Helper()

	class := mustNewCharacterClassForTest(t, classID)
	progression, ok := NewCharacterSpellcastingProgression(class)
	if !ok {
		t.Fatalf("expected spellcasting progression for class %q to compose", classID)
	}

	return progression
}

func mustCoreSpellListEntryForTest(
	t *testing.T,
	spellID characterspell.SpellID,
	classID characterclass.ClassID,
	spellLevel int,
) characterspell.SpellListEntry {
	t.Helper()

	entries, ok := characterspell.GetSpellListEntriesByClassAndLevel(classID, spellLevel)
	if !ok {
		t.Fatalf("expected spell list entries for class %q spell level %d", classID, spellLevel)
	}

	for _, entry := range entries {
		if entry.GetSpellID() == spellID {
			return entry
		}
	}

	t.Fatalf("expected spell list entry (%q, %q, %d)", spellID, classID, spellLevel)
	return characterspell.SpellListEntry{}
}

func hasCharacterSpellListEntry(entries []CharacterSpellListEntry, spellID characterspell.SpellID) bool {
	for _, entry := range entries {
		if entry.GetSpellID() == spellID {
			return true
		}
	}

	return false
}
