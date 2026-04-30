package class

import "testing"

func TestCoreSpellcastingProgressionTables_SeedSevenCoreCastingClasses(t *testing.T) {
	if len(coreSpellcastingProgressionTables) != 7 {
		t.Fatalf("expected 7 core spellcasting progression tables, got %d", len(coreSpellcastingProgressionTables))
	}

	if len(coreSpellcastingProgressionClassOrder) != 7 {
		t.Fatalf("expected 7 ordered core spellcasting progression class ids, got %d", len(coreSpellcastingProgressionClassOrder))
	}

	for _, classID := range coreSpellcastingProgressionClassOrder {
		progression, ok := coreSpellcastingProgressionTables[classID]
		if !ok {
			t.Fatalf("expected core spellcasting progression for %q to be seeded", classID)
		}

		if progression.GetClassID() != classID {
			t.Fatalf("expected progression class id %q, got %q", classID, progression.GetClassID())
		}

		if progression.GetMaxClassLevel() != maxCoreClassLevel {
			t.Fatalf(
				"expected progression for %q to have %d class levels, got %d",
				classID,
				maxCoreClassLevel,
				progression.GetMaxClassLevel(),
			)
		}
	}
}

func TestCoreSpellcastingProgressionTables_KnownBreakpoints(t *testing.T) {
	testCases := []struct {
		classID    ClassID
		classLevel int
		spellLevel int
		expected   int
	}{
		{BardClassID, 1, 1, 1},
		{BardClassID, 4, 2, 1},
		{BardClassID, 16, 6, 1},
		{ClericClassID, 1, 0, 3},
		{ClericClassID, 3, 2, 1},
		{ClericClassID, 17, 9, 1},
		{DruidClassID, 20, 9, 4},
		{PaladinClassID, 4, 1, 0},
		{PaladinClassID, 6, 1, 1},
		{PaladinClassID, 7, 2, 0},
		{RangerClassID, 4, 1, 0},
		{PaladinClassID, 14, 4, 1},
		{RangerClassID, 10, 3, 0},
		{RangerClassID, 20, 4, 3},
		{SorcererClassID, 1, 1, 3},
		{SorcererClassID, 4, 2, 3},
		{SorcererClassID, 18, 9, 3},
		{WizardClassID, 1, 0, 3},
		{WizardClassID, 5, 3, 1},
		{WizardClassID, 20, 9, 4},
	}

	for _, tc := range testCases {
		progression, ok := coreSpellcastingProgressionTables[tc.classID]
		if !ok {
			t.Fatalf("expected core spellcasting progression for %q to be seeded", tc.classID)
		}

		actual, ok := progression.GetSpellSlots(tc.classLevel, tc.spellLevel)
		if !ok {
			t.Fatalf(
				"expected class %q level %d spell level %d spell slots lookup to succeed",
				tc.classID,
				tc.classLevel,
				tc.spellLevel,
			)
		}

		if actual != tc.expected {
			t.Fatalf(
				"expected class %q level %d spell level %d spell slots %d, got %d",
				tc.classID,
				tc.classLevel,
				tc.spellLevel,
				tc.expected,
				actual,
			)
		}
	}
}

func TestCoreSpellcastingProgressionTables_DelayedCastersDistinguishUnavailableLevelsFromZeroSlotUnlocks(t *testing.T) {
	testCases := []struct {
		classID    ClassID
		classLevel int
		spellLevel int
	}{
		{PaladinClassID, 4, 0},
		{PaladinClassID, 4, 2},
		{RangerClassID, 4, 0},
		{RangerClassID, 9, 3},
	}

	for _, tc := range testCases {
		progression, ok := coreSpellcastingProgressionTables[tc.classID]
		if !ok {
			t.Fatalf("expected core spellcasting progression for %q to be seeded", tc.classID)
		}

		if _, ok := progression.GetSpellSlots(tc.classLevel, tc.spellLevel); ok {
			t.Fatalf(
				"expected class %q level %d spell level %d lookup to fail before that spell level unlocks",
				tc.classID,
				tc.classLevel,
				tc.spellLevel,
			)
		}
	}

	zeroSlotUnlockCases := []struct {
		classID    ClassID
		classLevel int
		spellLevel int
	}{
		{PaladinClassID, 4, 1},
		{PaladinClassID, 7, 2},
		{RangerClassID, 10, 3},
		{RangerClassID, 13, 4},
	}

	for _, tc := range zeroSlotUnlockCases {
		progression, ok := coreSpellcastingProgressionTables[tc.classID]
		if !ok {
			t.Fatalf("expected core spellcasting progression for %q to be seeded", tc.classID)
		}

		actual, ok := progression.GetSpellSlots(tc.classLevel, tc.spellLevel)
		if !ok || actual != 0 {
			t.Fatalf(
				"expected class %q level %d spell level %d zero-slot unlock (0, true), got (%d, %t)",
				tc.classID,
				tc.classLevel,
				tc.spellLevel,
				actual,
				ok,
			)
		}
	}
}

func TestGetSpellcastingProgressionByClassID_ReturnsSeededCoreProgression(t *testing.T) {
	progression, ok := GetSpellcastingProgressionByClassID(WizardClassID)
	if !ok {
		t.Fatal("expected wizard spellcasting progression lookup to succeed")
	}

	if progression.GetClassID() != WizardClassID {
		t.Fatalf("expected wizard progression class id %q, got %q", WizardClassID, progression.GetClassID())
	}

	spellSlots, ok := progression.GetSpellSlots(5, 3)
	if !ok || spellSlots != 1 {
		t.Fatalf("expected wizard level 5 spell level 3 slots (1, true), got (%d, %t)", spellSlots, ok)
	}
}

func TestGetSpellcastingProgressionByClassID_ReturnsDetachedCopy(t *testing.T) {
	first, ok := GetSpellcastingProgressionByClassID(BardClassID)
	if !ok {
		t.Fatal("expected bard spellcasting progression lookup to succeed")
	}

	first.slotsByClassLevel[0][0] = 99

	second, ok := GetSpellcastingProgressionByClassID(BardClassID)
	if !ok {
		t.Fatal("expected bard spellcasting progression lookup to succeed")
	}

	spellSlots, ok := second.GetSpellSlots(1, 0)
	if !ok || spellSlots != 0 {
		t.Fatalf("expected stored bard level 1 spell level 0 slots (0, true), got (%d, %t)", spellSlots, ok)
	}
}

func TestGetSpellcastingProgressionByClassID_RejectsUnknownOrNonSpellcastingClass(t *testing.T) {
	if _, ok := GetSpellcastingProgressionByClassID(ClassID("oracle")); ok {
		t.Fatal("expected unknown class progression lookup to fail")
	}

	if _, ok := GetSpellcastingProgressionByClassID(FighterClassID); ok {
		t.Fatal("expected non-spellcasting class progression lookup to fail")
	}
}

func TestGetSpellcastingProgressions_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	progressions := GetSpellcastingProgressions()
	if len(progressions) != len(coreSpellcastingProgressionClassOrder) {
		t.Fatalf(
			"expected %d queried spellcasting progressions, got %d",
			len(coreSpellcastingProgressionClassOrder),
			len(progressions),
		)
	}

	for i, expectedClassID := range coreSpellcastingProgressionClassOrder {
		if progressions[i].GetClassID() != expectedClassID {
			t.Fatalf("expected progression at index %d to be %q, got %q", i, expectedClassID, progressions[i].GetClassID())
		}
	}
}

func TestGetSpellcastingProgressions_ReturnsDetachedCopies(t *testing.T) {
	first := GetSpellcastingProgressions()
	second := GetSpellcastingProgressions()

	first[0].slotsByClassLevel[0][0] = 99

	spellSlots, ok := second[0].GetSpellSlots(1, 0)
	if !ok || spellSlots != 0 {
		t.Fatalf("expected stored first progression level 1 spell level 0 slots (0, true), got (%d, %t)", spellSlots, ok)
	}
}
