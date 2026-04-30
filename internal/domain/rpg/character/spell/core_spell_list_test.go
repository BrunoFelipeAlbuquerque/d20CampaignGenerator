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

func TestGetSpellListEntries_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	entries := GetSpellListEntries()
	if len(entries) != len(coreSpellListEntries) {
		t.Fatalf("expected %d queried spell list entries, got %d", len(coreSpellListEntries), len(entries))
	}

	for i, expected := range coreSpellListEntries {
		if entries[i].GetSpellID() != expected.GetSpellID() ||
			entries[i].GetClassID() != expected.GetClassID() ||
			entries[i].GetSpellLevel() != expected.GetSpellLevel() {
			t.Fatalf("expected spell list entry at index %d to be %v, got %v", i, expected, entries[i])
		}
	}
}

func TestGetSpellListEntries_ReturnsDetachedSlice(t *testing.T) {
	first := GetSpellListEntries()
	second := GetSpellListEntries()

	first[0].spellID = "Changed"

	if second[0].GetSpellID() != SpellID("Dancing Lights") {
		t.Fatalf("expected stored spell list entry to remain Dancing Lights, got %q", second[0].GetSpellID())
	}
}

func TestGetSpellListEntriesByClass_ReturnsSeededCoreClassList(t *testing.T) {
	entries, ok := GetSpellListEntriesByClass(characterclass.WizardClassID)
	if !ok {
		t.Fatal("expected wizard spell list lookup to succeed")
	}

	if len(entries) != 396 {
		t.Fatalf("expected 396 wizard spell list entries, got %d", len(entries))
	}

	if !hasSpellListEntry(entries, SpellID("Fireball")) {
		t.Fatalf("expected wizard spell list to include %q", SpellID("Fireball"))
	}

	if hasSpellListEntry(entries, SpellID("Cure Serious Wounds")) {
		t.Fatalf("expected wizard spell list not to include %q", SpellID("Cure Serious Wounds"))
	}
}

func TestGetSpellListEntriesByClass_ReturnsDetachedSlice(t *testing.T) {
	first, ok := GetSpellListEntriesByClass(characterclass.BardClassID)
	if !ok {
		t.Fatal("expected bard spell list lookup to succeed")
	}

	first[0].spellID = "Changed"

	second, ok := GetSpellListEntriesByClass(characterclass.BardClassID)
	if !ok {
		t.Fatal("expected bard spell list lookup to succeed")
	}

	if second[0].GetSpellID() != SpellID("Dancing Lights") {
		t.Fatalf("expected stored bard spell list entry to remain Dancing Lights, got %q", second[0].GetSpellID())
	}
}

func TestGetSpellListEntriesByClass_RejectsInvalidClass(t *testing.T) {
	if _, ok := GetSpellListEntriesByClass(characterclass.FighterClassID); ok {
		t.Fatal("expected non-spellcasting class spell list lookup to fail")
	}

	if _, ok := GetSpellListEntriesByClass(characterclass.ClassID("oracle")); ok {
		t.Fatal("expected unknown class spell list lookup to fail")
	}
}

func TestCoreSpellListBindings_ClassListLookupByLevel(t *testing.T) {
	wizardThirdLevelEntries, ok := GetSpellListEntriesByClassAndLevel(
		characterclass.WizardClassID,
		3,
	)
	if !ok {
		t.Fatal("expected wizard 3rd-level spell list lookup to succeed")
	}

	if len(wizardThirdLevelEntries) == 0 {
		t.Fatal("expected wizard 3rd-level spell list entries")
	}

	if !hasSpellListEntry(wizardThirdLevelEntries, SpellID("Fireball")) {
		t.Fatalf("expected wizard 3rd-level spell list to include %q", SpellID("Fireball"))
	}

	if hasSpellListEntry(wizardThirdLevelEntries, SpellID("Cure Serious Wounds")) {
		t.Fatalf("expected wizard 3rd-level spell list not to include %q", SpellID("Cure Serious Wounds"))
	}

	clericThirdLevelEntries, ok := GetSpellListEntriesByClassAndLevel(
		characterclass.ClericClassID,
		3,
	)
	if !ok {
		t.Fatal("expected cleric 3rd-level spell list lookup to succeed")
	}

	if len(clericThirdLevelEntries) == 0 {
		t.Fatal("expected cleric 3rd-level spell list entries")
	}

	if !hasSpellListEntry(clericThirdLevelEntries, SpellID("Cure Serious Wounds")) {
		t.Fatalf("expected cleric 3rd-level spell list to include %q", SpellID("Cure Serious Wounds"))
	}

	if hasSpellListEntry(clericThirdLevelEntries, SpellID("Fireball")) {
		t.Fatalf("expected cleric 3rd-level spell list not to include %q", SpellID("Fireball"))
	}
}

func TestGetSpellListEntriesByClassAndLevel_ReturnsDetachedSlice(t *testing.T) {
	first, ok := GetSpellListEntriesByClassAndLevel(characterclass.BardClassID, 0)
	if !ok {
		t.Fatal("expected bard cantrip spell list lookup to succeed")
	}

	first[0].spellID = "Changed"

	second, ok := GetSpellListEntriesByClassAndLevel(characterclass.BardClassID, 0)
	if !ok {
		t.Fatal("expected bard cantrip spell list lookup to succeed")
	}

	if second[0].GetSpellID() != SpellID("Dancing Lights") {
		t.Fatalf("expected stored bard cantrip entry to remain Dancing Lights, got %q", second[0].GetSpellID())
	}
}

func TestGetSpellListEntriesByClassAndLevel_RejectsInvalidClassOrLevel(t *testing.T) {
	testCases := []struct {
		classID    characterclass.ClassID
		spellLevel int
	}{
		{characterclass.FighterClassID, 0},
		{characterclass.ClassID("oracle"), 0},
		{characterclass.WizardClassID, -1},
		{characterclass.WizardClassID, 10},
		{characterclass.BardClassID, 7},
		{characterclass.PaladinClassID, 0},
	}

	for _, tc := range testCases {
		if _, ok := GetSpellListEntriesByClassAndLevel(tc.classID, tc.spellLevel); ok {
			t.Fatalf("expected spell list lookup for class %q level %d to fail", tc.classID, tc.spellLevel)
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

func hasSpellListEntry(entries []SpellListEntry, spellID SpellID) bool {
	for _, entry := range entries {
		if entry.GetSpellID() == spellID {
			return true
		}
	}

	return false
}
