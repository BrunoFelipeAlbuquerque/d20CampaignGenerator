package spell

import "testing"

func TestCoreSpellData_SeedsAllCoreSpells(t *testing.T) {
	if len(coreSpells) != 623 {
		t.Fatalf("expected 623 core spell seeds, got %d", len(coreSpells))
	}

	for _, entry := range coreSpellListEntries {
		if _, ok := coreSpells[entry.GetSpellID()]; !ok {
			t.Fatalf("expected spell data seed for core spell %q at level %d", entry.GetSpellID(), entry.GetSpellLevel())
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

func TestCoreSpellData_KnownCoreHeaders(t *testing.T) {
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
		{
			id:              SpellID("Alarm"),
			school:          AbjurationSchoolID,
			component:       FocusComponentID,
			castingTime:     "1 standard action",
			spellRange:      "close (25 ft. + 5 ft./2 levels)",
			targetEffect:    "20-ft.-radius emanation centered on a point in space",
			duration:        "2 hours/level (D)",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Fireball"),
			school:          EvocationSchoolID,
			descriptor:      DescriptorID("Fire"),
			component:       MaterialComponentID,
			castingTime:     "1 standard action",
			spellRange:      "long (400 ft. + 40 ft./level)",
			targetEffect:    "20-ft.-radius spread",
			duration:        "instantaneous",
			savingThrow:     "Reflex half",
			spellResistance: "yes",
		},
		{
			id:              SpellID("Summon Monster 3"),
			school:          ConjurationSchoolID,
			component:       FocusComponentID,
			castingTime:     "1 round",
			spellRange:      "close (25 ft. + 5 ft./2 levels)",
			targetEffect:    "one summoned creature",
			duration:        "1 round/level (D)",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Break Enchantment"),
			school:          AbjurationSchoolID,
			component:       SomaticComponentID,
			castingTime:     "1 minute",
			spellRange:      "close (25 ft. + 5 ft./2 levels)",
			targetEffect:    "up to one creature per level, all within 30 ft. of each other",
			duration:        "instantaneous",
			savingThrow:     "see text",
			spellResistance: "no",
		},
		{
			id:              SpellID("Disintegrate"),
			school:          TransmutationSchoolID,
			component:       MaterialComponentID,
			castingTime:     "1 standard action",
			spellRange:      "medium (100 ft. + 10 ft./level)",
			targetEffect:    "ray",
			duration:        "instantaneous",
			savingThrow:     "Fortitude partial (object)",
			spellResistance: "yes",
		},
		{
			id:              SpellID("Heal"),
			school:          ConjurationSchoolID,
			component:       SomaticComponentID,
			castingTime:     "1 standard action",
			spellRange:      "touch",
			targetEffect:    "creature touched",
			duration:        "instantaneous",
			savingThrow:     "Will negates (harmless)",
			spellResistance: "yes (harmless)",
		},
		{
			id:              SpellID("Summon Monster 6"),
			school:          ConjurationSchoolID,
			component:       FocusComponentID,
			castingTime:     "1 round",
			spellRange:      "close (25 ft. + 5 ft./2 levels)",
			targetEffect:    "one summoned creature",
			duration:        "1 round/level (D)",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Gate"),
			school:          ConjurationSchoolID,
			component:       MaterialComponentID,
			castingTime:     "1 standard action",
			spellRange:      "medium (100 ft. + 10 ft./level)",
			targetEffect:    "see text",
			duration:        "instantaneous or concentration (up to 1 round/level); see text",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Meteor Swarm"),
			school:          EvocationSchoolID,
			descriptor:      DescriptorID("Fire"),
			component:       SomaticComponentID,
			castingTime:     "1 standard action",
			spellRange:      "long (400 ft. + 40 ft./level)",
			targetEffect:    "four 40-ft.-radius spreads, see text",
			duration:        "instantaneous",
			savingThrow:     "none or Reflex half, see text",
			spellResistance: "yes",
		},
		{
			id:              SpellID("Time Stop"),
			school:          TransmutationSchoolID,
			component:       VerbalComponentID,
			castingTime:     "1 standard action",
			spellRange:      "personal",
			targetEffect:    "you",
			duration:        "1d4+1 rounds (apparent time); see text",
			savingThrow:     "none",
			spellResistance: "no",
		},
		{
			id:              SpellID("Wish"),
			school:          UniversalSchoolID,
			component:       MaterialComponentID,
			castingTime:     "1 standard action",
			spellRange:      "see text",
			targetEffect:    "see text",
			duration:        "see text",
			savingThrow:     "none, see text",
			spellResistance: "yes",
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
