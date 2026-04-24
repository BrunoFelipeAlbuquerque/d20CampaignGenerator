package spell

import (
	"testing"

	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

func TestCoreSpellListBindings_SeedCoreCastingClassMappings(t *testing.T) {
	if len(coreSpellListEntries) != 1455 {
		t.Fatalf("expected 1455 core spell list entries, got %d", len(coreSpellListEntries))
	}

	expectedClassCounts := map[characterclass.ClassID]int{
		characterclass.BardClassID:     164,
		characterclass.ClericClassID:   236,
		characterclass.DruidClassID:    169,
		characterclass.PaladinClassID:  45,
		characterclass.RangerClassID:   51,
		characterclass.SorcererClassID: 394,
		characterclass.WizardClassID:   396,
	}

	actualClassCounts := make(map[characterclass.ClassID]int, len(expectedClassCounts))
	seen := make(map[spellListEntryKey]struct{}, len(coreSpellListEntries))

	for _, entry := range coreSpellListEntries {
		if _, ok := NewSpellListEntry(entry.GetSpellID(), entry.GetClassID(), entry.GetSpellLevel()); !ok {
			t.Fatalf("expected seeded spell list entry to remain valid: %v", entry)
		}

		key := spellListEntryKey{
			spellID:    entry.GetSpellID(),
			classID:    entry.GetClassID(),
			spellLevel: entry.GetSpellLevel(),
		}
		if _, ok := seen[key]; ok {
			t.Fatalf("expected no duplicate core spell list entry seed for %v", key)
		}
		seen[key] = struct{}{}

		actualClassCounts[entry.GetClassID()]++
	}

	for classID, expected := range expectedClassCounts {
		if actualClassCounts[classID] != expected {
			t.Fatalf(
				"expected class %q to have %d core spell list entries, got %d",
				classID,
				expected,
				actualClassCounts[classID],
			)
		}
	}
}

func TestCoreSpellListBindings_KnownCoreBreakpoints(t *testing.T) {
	testCases := []struct {
		spellID    SpellID
		classID    characterclass.ClassID
		spellLevel int
	}{
		{SpellID("Light"), characterclass.BardClassID, 0},
		{SpellID("Light"), characterclass.ClericClassID, 0},
		{SpellID("Light"), characterclass.DruidClassID, 0},
		{SpellID("Light"), characterclass.SorcererClassID, 0},
		{SpellID("Light"), characterclass.WizardClassID, 0},
		{SpellID("Cure Light Wounds"), characterclass.PaladinClassID, 1},
		{SpellID("Cure Light Wounds"), characterclass.RangerClassID, 2},
		{SpellID("Summon Nature's Ally 1"), characterclass.DruidClassID, 1},
		{SpellID("Summon Nature's Ally 1"), characterclass.RangerClassID, 1},
		{SpellID("Wish"), characterclass.SorcererClassID, 9},
		{SpellID("Wish"), characterclass.WizardClassID, 9},
	}

	for _, tc := range testCases {
		if !hasCoreSpellListEntry(tc.spellID, tc.classID, tc.spellLevel) {
			t.Fatalf("expected core spell list entry (%q, %q, %d)", tc.spellID, tc.classID, tc.spellLevel)
		}
	}
}

func TestCoreSpellListBindings_DoNotSeedInvalidCoreClassLevels(t *testing.T) {
	testCases := []struct {
		spellID    SpellID
		classID    characterclass.ClassID
		spellLevel int
	}{
		{SpellID("Light"), characterclass.PaladinClassID, 0},
		{SpellID("Light"), characterclass.RangerClassID, 0},
		{SpellID("Wish"), characterclass.BardClassID, 9},
		{SpellID("Cure Light Wounds"), characterclass.WizardClassID, 1},
	}

	for _, tc := range testCases {
		if hasCoreSpellListEntry(tc.spellID, tc.classID, tc.spellLevel) {
			t.Fatalf("expected no core spell list entry (%q, %q, %d)", tc.spellID, tc.classID, tc.spellLevel)
		}
	}
}

func hasCoreSpellListEntry(spellID SpellID, classID characterclass.ClassID, spellLevel int) bool {
	for _, entry := range coreSpellListEntries {
		if entry.GetSpellID() == spellID &&
			entry.GetClassID() == classID &&
			entry.GetSpellLevel() == spellLevel {
			return true
		}
	}

	return false
}
