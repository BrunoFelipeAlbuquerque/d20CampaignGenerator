package class

import "testing"

func TestNewSpellcastingProgressionTable_ConstructsValidatedTable(t *testing.T) {
	progression, ok := NewSpellcastingProgressionTable(
		WizardClassID,
		[][]int{
			{3, 1},
			{4, 2},
			{4, 2, 1},
		},
	)
	if !ok {
		t.Fatal("expected spellcasting progression table to be constructed")
	}

	if progression.GetClassID() != WizardClassID {
		t.Fatalf("expected progression class id %q, got %q", WizardClassID, progression.GetClassID())
	}

	if progression.GetMaxClassLevel() != 3 {
		t.Fatalf("expected max class level 3, got %d", progression.GetMaxClassLevel())
	}

	if progression.GetMinimumSpellLevel() != 0 {
		t.Fatalf("expected minimum spell level 0, got %d", progression.GetMinimumSpellLevel())
	}

	levelOneSlots, ok := progression.GetSpellSlotsByClassLevel(1)
	if !ok {
		t.Fatal("expected first class level row to be available")
	}

	if len(levelOneSlots) != 2 || levelOneSlots[0] != 3 || levelOneSlots[1] != 1 {
		t.Fatalf("expected first class level row [3 1], got %v", levelOneSlots)
	}

	levelOneSlots[0] = 99

	levelOneSlotsAgain, ok := progression.GetSpellSlotsByClassLevel(1)
	if !ok {
		t.Fatal("expected first class level row to still be available")
	}

	if len(levelOneSlotsAgain) != 2 || levelOneSlotsAgain[0] != 3 {
		t.Fatalf("expected defensive class level row copy [3 1], got %v", levelOneSlotsAgain)
	}

	spellSlots, ok := progression.GetSpellSlots(3, 2)
	if !ok || spellSlots != 1 {
		t.Fatalf("expected class level 3 spell level 2 slots (1, true), got (%d, %t)", spellSlots, ok)
	}

	if _, ok := progression.GetSpellSlots(2, 3); ok {
		t.Fatal("expected unavailable spell level lookup to fail")
	}
}

func TestNewSpellcastingProgressionTable_AcceptsLeadingNoncastingLevels(t *testing.T) {
	progression, ok := NewSpellcastingProgressionTable(
		PaladinClassID,
		[][]int{
			nil,
			nil,
			nil,
			{0},
			{0},
			{1},
			{1, 0},
		},
	)
	if !ok {
		t.Fatal("expected delayed spellcasting progression table to be constructed")
	}

	levelOneSlots, ok := progression.GetSpellSlotsByClassLevel(1)
	if !ok {
		t.Fatal("expected first class level row to be available")
	}

	if len(levelOneSlots) != 0 {
		t.Fatalf("expected first paladin class level row to have no spell slots, got %v", levelOneSlots)
	}

	if progression.GetMinimumSpellLevel() != 1 {
		t.Fatalf("expected delayed-caster minimum spell level 1, got %d", progression.GetMinimumSpellLevel())
	}

	levelFourSlots, ok := progression.GetSpellSlotsByClassLevel(4)
	if !ok {
		t.Fatal("expected fourth class level row to be available")
	}

	if len(levelFourSlots) != 1 || levelFourSlots[0] != 0 {
		t.Fatalf("expected fourth paladin class level row [0], got %v", levelFourSlots)
	}

	if _, ok := progression.GetSpellSlots(1, 1); ok {
		t.Fatal("expected unavailable delayed-caster spell level lookup to fail before spellcasting starts")
	}

	if _, ok := progression.GetSpellSlots(4, 0); ok {
		t.Fatal("expected delayed-caster 0-level spell lookup to fail")
	}

	spellSlots, ok := progression.GetSpellSlots(4, 1)
	if !ok || spellSlots != 0 {
		t.Fatalf("expected class level 4 spell level 1 slots (0, true), got (%d, %t)", spellSlots, ok)
	}

	spellSlots, ok = progression.GetSpellSlots(7, 2)
	if !ok || spellSlots != 0 {
		t.Fatalf("expected class level 7 spell level 2 slots (0, true), got (%d, %t)", spellSlots, ok)
	}
}

func TestNewSpellcastingProgressionTable_RejectsInvalidInputs(t *testing.T) {
	if _, ok := NewSpellcastingProgressionTable(ClassID("oracle"), [][]int{{3, 1}}); ok {
		t.Fatal("expected unknown class id to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(FighterClassID, [][]int{{3, 1}}); ok {
		t.Fatal("expected non-spellcasting class progression to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, nil); ok {
		t.Fatal("expected missing class level rows to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, make([][]int, maxCoreClassLevel+1)); ok {
		t.Fatal("expected too many class levels to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, [][]int{{0, 0, 0}}); ok {
		t.Fatal("expected all-zero spell slot table to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, [][]int{{3, -1}}); ok {
		t.Fatal("expected negative spell slots to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, [][]int{{3, 1, 0, 1}}); ok {
		t.Fatal("expected spell level gap after positive slots to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}); ok {
		t.Fatal("expected spell slot rows above core spell level 9 to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, [][]int{{3, 1, 0, 0}}); ok {
		t.Fatal("expected impossible first-row spell level availability jump to be rejected")
	}

	if _, ok := NewSpellcastingProgressionTable(WizardClassID, [][]int{{3, 1}, nil}); ok {
		t.Fatal("expected spellcasting progression rows to reject availability disappearing after it starts")
	}

	if _, ok := NewSpellcastingProgressionTable(PaladinClassID, [][]int{nil, nil, nil, {0, 0}}); ok {
		t.Fatal("expected delayed-caster first spell row to reject unlocking beyond 1st-level spells")
	}
}

func TestNewSpellcastingProgressionTable_ValidatesSpellcastingClassBoundary(t *testing.T) {
	testCases := []struct {
		classID ClassID
		ok      bool
	}{
		{BarbarianClassID, false},
		{BardClassID, true},
		{ClericClassID, true},
		{DruidClassID, true},
		{FighterClassID, false},
		{MonkClassID, false},
		{PaladinClassID, true},
		{RangerClassID, true},
		{RogueClassID, false},
		{SorcererClassID, true},
		{WizardClassID, true},
	}

	for _, tc := range testCases {
		_, ok := NewSpellcastingProgressionTable(tc.classID, [][]int{{1}})
		if ok != tc.ok {
			t.Fatalf("expected spellcasting progression construction for class %q to be %t, got %t", tc.classID, tc.ok, ok)
		}
	}
}

func TestSpellcastingProgressionTable_GetSpellSlots_RejectsInvalidLevels(t *testing.T) {
	progression := mustSpellcastingProgressionTableForTest(
		t,
		SorcererClassID,
		[][]int{
			{5, 3},
		},
	)

	if _, ok := progression.GetSpellSlotsByClassLevel(0); ok {
		t.Fatal("expected class level 0 row lookup to fail")
	}

	if _, ok := progression.GetSpellSlotsByClassLevel(2); ok {
		t.Fatal("expected missing class level row lookup to fail")
	}

	if _, ok := progression.GetSpellSlots(0, 1); ok {
		t.Fatal("expected class level 0 spell slot lookup to fail")
	}

	if _, ok := progression.GetSpellSlots(1, -1); ok {
		t.Fatal("expected negative spell level lookup to fail")
	}

	if _, ok := progression.GetSpellSlots(1, 10); ok {
		t.Fatal("expected spell level above 9 lookup to fail")
	}
}

func TestSpellcastingProgressionTable_GetSpellSlots_DistinguishesUnavailableLevelsFromUnlockedZeroSlots(t *testing.T) {
	progression := mustSpellcastingProgressionTableForTest(
		t,
		PaladinClassID,
		[][]int{
			nil,
			nil,
			nil,
			{0},
			{0},
			{1},
			{1, 0},
		},
	)

	if _, ok := progression.GetSpellSlots(4, 0); ok {
		t.Fatal("expected delayed-caster 0-level spell lookup to fail")
	}

	if _, ok := progression.GetSpellSlots(4, 2); ok {
		t.Fatal("expected unavailable delayed-caster spell level lookup to fail")
	}

	spellSlots, ok := progression.GetSpellSlots(4, 1)
	if !ok || spellSlots != 0 {
		t.Fatalf("expected class level 4 spell level 1 slots (0, true), got (%d, %t)", spellSlots, ok)
	}

	spellSlots, ok = progression.GetSpellSlots(7, 2)
	if !ok || spellSlots != 0 {
		t.Fatalf("expected class level 7 spell level 2 slots (0, true), got (%d, %t)", spellSlots, ok)
	}
}

func mustSpellcastingProgressionTableForTest(
	t *testing.T,
	classID ClassID,
	slotsByClassLevel [][]int,
) SpellcastingProgressionTable {
	t.Helper()

	value, ok := NewSpellcastingProgressionTable(classID, slotsByClassLevel)
	if !ok {
		t.Fatal("expected spellcasting progression table to be constructed")
	}

	return value
}
