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
