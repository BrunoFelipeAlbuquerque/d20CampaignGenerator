package character

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

func TestNewCharacterClassHitPointFacts_ComposesSingleClassLedger(t *testing.T) {
	facts, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 3),
		},
		characterclass.FighterClassID,
		14,
		[]CharacterClassHitPointEntry{
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 2, 6),
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 3, 5),
		},
	)
	if !ok {
		t.Fatal("expected class HP ledger to compose")
	}

	firstLevelClassID, ok := facts.GetFirstLevelClassID()
	if !ok || firstLevelClassID != characterclass.FighterClassID {
		t.Fatalf("expected first-level class fighter, got %q and %t", firstLevelClassID, ok)
	}

	hp, ok := facts.GetHitPoints()
	if !ok {
		t.Fatal("expected composed hit points to resolve")
	}

	if hp.GetTotal() != 27 || hp.GetCurrent() != 27 {
		t.Fatalf("expected total/current HP 27, got total %d current %d", hp.GetTotal(), hp.GetCurrent())
	}

	d10Count, ok := hp.GetHitDie().GetDieCount(ability.D10HitDie)
	if !ok || d10Count != 3 {
		t.Fatalf("expected three d10 hit dice, got (%d, %t)", d10Count, ok)
	}

	assertClassHitPointSources(t, hp, 21, 6)
}

func TestNewCharacterClassHitPointFacts_ComposesMulticlassLedger(t *testing.T) {
	facts, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
		},
		characterclass.WizardClassID,
		12,
		[]CharacterClassHitPointEntry{
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 1, 7),
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 2, 8),
		},
	)
	if !ok {
		t.Fatal("expected multiclass HP ledger to compose")
	}

	hp, ok := facts.GetHitPoints()
	if !ok {
		t.Fatal("expected composed hit points to resolve")
	}

	if hp.GetTotal() != 24 || hp.GetCurrent() != 24 {
		t.Fatalf("expected total/current HP 24, got total %d current %d", hp.GetTotal(), hp.GetCurrent())
	}

	d6Count, ok := hp.GetHitDie().GetDieCount(ability.D6HitDie)
	if !ok || d6Count != 1 {
		t.Fatalf("expected one d6 hit die, got (%d, %t)", d6Count, ok)
	}

	d10Count, ok := hp.GetHitDie().GetDieCount(ability.D10HitDie)
	if !ok || d10Count != 2 {
		t.Fatalf("expected two d10 hit dice, got (%d, %t)", d10Count, ok)
	}

	assertClassHitPointSources(t, hp, 21, 3)
}

func TestNewCharacterClassHitPointFacts_ExposesDefensiveEntryCopy(t *testing.T) {
	facts, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
		},
		characterclass.FighterClassID,
		10,
		[]CharacterClassHitPointEntry{
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 2, 5),
		},
	)
	if !ok {
		t.Fatal("expected class HP ledger to compose")
	}

	entries := facts.GetEntries()
	if len(entries) != 1 {
		t.Fatalf("expected one HP entry, got %d", len(entries))
	}
	entries[0] = mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 2, 1)

	storedEntries := facts.GetEntries()
	if storedEntries[0].GetBaseHitPoints() != 5 {
		t.Fatalf("expected stored entry to remain 5, got %d", storedEntries[0].GetBaseHitPoints())
	}
}

func TestCharacterClassHitPointFacts_ZeroValueDoesNotResolve(t *testing.T) {
	var facts CharacterClassHitPointFacts

	if _, ok := facts.GetFirstLevelClassID(); ok {
		t.Fatal("expected zero-value first-level class lookup to fail")
	}

	if _, ok := facts.GetHitPoints(); ok {
		t.Fatal("expected zero-value hit points lookup to fail")
	}

	if entries := facts.GetEntries(); entries != nil {
		t.Fatalf("expected nil zero-value entries, got %d", len(entries))
	}
}

func TestNewCharacterClassHitPointEntry_RejectsOutOfRangeBaseHitPoints(t *testing.T) {
	if _, ok := NewCharacterClassHitPointEntry(characterclass.WizardClassID, 2, 0); ok {
		t.Fatal("expected zero base HP to be rejected")
	}

	if _, ok := NewCharacterClassHitPointEntry(characterclass.WizardClassID, 2, 7); ok {
		t.Fatal("expected base HP above wizard d6 maximum to be rejected")
	}

	if _, ok := NewCharacterClassHitPointEntry(characterclass.ClassID("alchemist"), 2, 4); ok {
		t.Fatal("expected non-core class HP entry constructor input to be rejected")
	}
}

func TestNewCharacterClassHitPointFacts_RejectsMissingEntries(t *testing.T) {
	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 3),
		},
		characterclass.FighterClassID,
		10,
		[]CharacterClassHitPointEntry{
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 2, 5),
		},
	); ok {
		t.Fatal("expected missing class HP entry to be rejected")
	}
}

func TestNewCharacterClassHitPointFacts_RejectsDuplicateEntries(t *testing.T) {
	entry := mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 2, 5)

	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 3),
		},
		characterclass.FighterClassID,
		10,
		[]CharacterClassHitPointEntry{
			entry,
			entry,
		},
	); ok {
		t.Fatal("expected duplicate class HP entries to be rejected")
	}
}

func TestNewCharacterClassHitPointFacts_RejectsFirstLevelClassEntry(t *testing.T) {
	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1),
		},
		characterclass.FighterClassID,
		10,
		[]CharacterClassHitPointEntry{
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 1, 5),
		},
	); ok {
		t.Fatal("expected explicit first-level class HP entry to be rejected")
	}
}

func TestNewCharacterClassHitPointFacts_RejectsEntriesOutsideSelectedClassLevels(t *testing.T) {
	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
		},
		characterclass.FighterClassID,
		10,
		[]CharacterClassHitPointEntry{
			mustNewCharacterClassHitPointEntryForTest(t, characterclass.FighterClassID, 3, 5),
		},
	); ok {
		t.Fatal("expected class HP entry beyond selected class level to be rejected")
	}
}

func TestNewCharacterClassHitPointFacts_RejectsMalformedEntries(t *testing.T) {
	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 2),
		},
		characterclass.FighterClassID,
		10,
		[]CharacterClassHitPointEntry{
			{classID: characterclass.FighterClassID, classLevel: 2, baseHitPoints: 0},
		},
	); ok {
		t.Fatal("expected malformed class HP entry to be rejected")
	}
}

func TestNewCharacterClassHitPointFacts_RejectsInvalidInputs(t *testing.T) {
	fighterLevel := mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)

	if _, ok := NewCharacterClassHitPointFacts(nil, characterclass.FighterClassID, 10, nil); ok {
		t.Fatal("expected missing class levels to be rejected")
	}

	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{fighterLevel},
		characterclass.WizardClassID,
		10,
		nil,
	); ok {
		t.Fatal("expected unselected first-level class to be rejected")
	}

	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{fighterLevel},
		characterclass.FighterClassID,
		0,
		nil,
	); ok {
		t.Fatal("expected invalid constitution score to be rejected")
	}

	if _, ok := NewCharacterClassHitPointFacts(
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 20),
			mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1),
		},
		characterclass.FighterClassID,
		10,
		nil,
	); ok {
		t.Fatal("expected total character level above core range to be rejected")
	}
}

func mustNewCharacterClassHitPointEntryForTest(
	t *testing.T,
	id characterclass.ClassID,
	classLevel int,
	baseHitPoints int,
) CharacterClassHitPointEntry {
	t.Helper()

	entry, ok := NewCharacterClassHitPointEntry(id, classLevel, baseHitPoints)
	if !ok {
		t.Fatalf("expected class HP entry %q level %d HP %d to compose", id, classLevel, baseHitPoints)
	}

	return entry
}

func assertClassHitPointSources(
	t *testing.T,
	hp ability.HitPoints,
	expectedBaseDice int,
	expectedConstitution int,
) {
	t.Helper()

	sources := hp.GetSources()
	if len(sources) != 2 {
		t.Fatalf("expected 2 HP sources, got %d", len(sources))
	}

	if sources[0].GetName() != "Base Dice" || sources[0].GetValue() != expectedBaseDice {
		t.Fatalf("expected base dice source %d, got %q = %d", expectedBaseDice, sources[0].GetName(), sources[0].GetValue())
	}

	if sources[1].GetName() != "Constitution" || sources[1].GetValue() != expectedConstitution {
		t.Fatalf("expected constitution source %d, got %q = %d", expectedConstitution, sources[1].GetName(), sources[1].GetValue())
	}
}
