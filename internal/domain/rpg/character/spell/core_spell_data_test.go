package spell

import "testing"

func TestCoreSpellData_SeedsAllCoreCantripsAndOrisons(t *testing.T) {
	if len(coreSpells) != 28 {
		t.Fatalf("expected 28 core cantrip/orison spell seeds, got %d", len(coreSpells))
	}

	for _, entry := range coreSpellListEntries {
		if entry.GetSpellLevel() != 0 {
			continue
		}

		if _, ok := coreSpells[entry.GetSpellID()]; !ok {
			t.Fatalf("expected spell data seed for core 0-level spell %q", entry.GetSpellID())
		}
	}
}

func TestCoreSpellData_SeededSpellsRemainValid(t *testing.T) {
	for id, spell := range coreSpells {
		if id != spell.GetID() {
			t.Fatalf("expected map key %q to match spell id %q", id, spell.GetID())
		}

		if _, ok := NewSpell(
			spell.GetID(),
			spell.GetSchool(),
			spell.GetDescriptors(),
			spell.GetComponents(),
			spell.GetCastingTime(),
			spell.GetRange(),
			spell.GetTargetEffect(),
			spell.GetDuration(),
			spell.GetSavingThrow(),
			spell.GetSpellResistance(),
		); !ok {
			t.Fatalf("expected seeded spell %q to remain valid", id)
		}
	}
}

func TestCoreSpellData_KnownCantripAndOrisonHeaders(t *testing.T) {
	testCases := []struct {
		id              SpellID
		school          SchoolID
		descriptor      DescriptorID
		component       ComponentID
		castingTime     string
		spellRange      string
		targetEffect    string
		duration        string
		savingThrow     string
		spellResistance string
	}{
		{
			id:              SpellID("Acid Splash"),
			school:          ConjurationSchoolID,
			descriptor:      DescriptorID("Acid"),
			component:       SomaticComponentID,
			castingTime:     "1 standard action",
			spellRange:      "close (25 ft. + 5 ft./2 levels)",
			targetEffect:    "one missile of acid",
			duration:        "instantaneous",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Detect Magic"),
			school:          DivinationSchoolID,
			component:       VerbalComponentID,
			castingTime:     "1 standard action",
			spellRange:      "60 ft.",
			targetEffect:    "cone-shaped emanation",
			duration:        "concentration, up to 1 min./level (D)",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Light"),
			school:          EvocationSchoolID,
			descriptor:      DescriptorID("Light"),
			component:       DivineFocusComponentID,
			castingTime:     "1 standard action",
			spellRange:      "touch",
			targetEffect:    "object touched",
			duration:        "10 min./level",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Stabilize"),
			school:          ConjurationSchoolID,
			component:       SomaticComponentID,
			castingTime:     "1 standard action",
			spellRange:      "close (25 ft. + 5 ft./2 levels)",
			targetEffect:    "one living creature",
			duration:        "instantaneous",
			savingThrow:     "Will negates (harmless)",
			spellResistance: "yes (harmless)",
		},
	}

	for _, tc := range testCases {
		spell, ok := coreSpells[tc.id]
		if !ok {
			t.Fatalf("expected core spell seed %q", tc.id)
		}

		if spell.GetSchool() != tc.school {
			t.Fatalf("expected %q school %q, got %q", tc.id, tc.school, spell.GetSchool())
		}

		if tc.descriptor != "" && !hasDescriptor(spell, tc.descriptor) {
			t.Fatalf("expected %q descriptor %q, got %v", tc.id, tc.descriptor, spell.GetDescriptors())
		}

		if !hasComponent(spell, tc.component) {
			t.Fatalf("expected %q component %q, got %v", tc.id, tc.component, spell.GetComponents())
		}

		if spell.GetCastingTime() != tc.castingTime {
			t.Fatalf("expected %q casting time %q, got %q", tc.id, tc.castingTime, spell.GetCastingTime())
		}

		if spell.GetRange() != tc.spellRange {
			t.Fatalf("expected %q range %q, got %q", tc.id, tc.spellRange, spell.GetRange())
		}

		if spell.GetTargetEffect() != tc.targetEffect {
			t.Fatalf("expected %q target/effect %q, got %q", tc.id, tc.targetEffect, spell.GetTargetEffect())
		}

		if spell.GetDuration() != tc.duration {
			t.Fatalf("expected %q duration %q, got %q", tc.id, tc.duration, spell.GetDuration())
		}

		if spell.GetSavingThrow() != tc.savingThrow {
			t.Fatalf("expected %q saving throw %q, got %q", tc.id, tc.savingThrow, spell.GetSavingThrow())
		}

		if spell.GetSpellResistance() != tc.spellResistance {
			t.Fatalf("expected %q spell resistance %q, got %q", tc.id, tc.spellResistance, spell.GetSpellResistance())
		}
	}
}

func hasDescriptor(spell Spell, descriptor DescriptorID) bool {
	for _, value := range spell.GetDescriptors() {
		if value == descriptor {
			return true
		}
	}

	return false
}

func hasComponent(spell Spell, component ComponentID) bool {
	for _, value := range spell.GetComponents() {
		if value == component {
			return true
		}
	}

	return false
}
