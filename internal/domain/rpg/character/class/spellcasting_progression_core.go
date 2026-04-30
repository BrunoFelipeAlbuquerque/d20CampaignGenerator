package class

var coreSpellcastingProgressionTables = mustBuildCoreSpellcastingProgressionTables(coreClasses)

var coreSpellcastingProgressionClassOrder = []ClassID{
	BardClassID,
	ClericClassID,
	DruidClassID,
	PaladinClassID,
	RangerClassID,
	SorcererClassID,
	WizardClassID,
}

func GetSpellcastingProgressionByClassID(classID ClassID) (SpellcastingProgressionTable, bool) {
	value, ok := coreSpellcastingProgressionTables[classID]
	if !ok {
		return spellcastingProgressionTable{}, false
	}

	return cloneSpellcastingProgressionTable(value), true
}

func GetSpellcastingProgressions() []SpellcastingProgressionTable {
	progressions := make([]SpellcastingProgressionTable, 0, len(coreSpellcastingProgressionClassOrder))

	for _, classID := range coreSpellcastingProgressionClassOrder {
		progressions = append(
			progressions,
			cloneSpellcastingProgressionTable(coreSpellcastingProgressionTables[classID]),
		)
	}

	return progressions
}

func mustBuildCoreSpellcastingProgressionTables(
	classes map[ClassID]Class,
) map[ClassID]SpellcastingProgressionTable {
	_ = classes

	return map[ClassID]SpellcastingProgressionTable{
		BardClassID: mustNewSpellcastingProgressionTable(
			BardClassID,
			[][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 3, 1},
				{0, 4, 2},
				{0, 4, 3},
				{0, 4, 3, 1},
				{0, 4, 4, 2},
				{0, 5, 4, 3},
				{0, 5, 4, 3, 1},
				{0, 5, 4, 4, 2},
				{0, 5, 5, 4, 3},
				{0, 5, 5, 4, 3, 1},
				{0, 5, 5, 4, 4, 2},
				{0, 5, 5, 5, 4, 3},
				{0, 5, 5, 5, 4, 3, 1},
				{0, 5, 5, 5, 4, 4, 2},
				{0, 5, 5, 5, 5, 4, 3},
				{0, 5, 5, 5, 5, 5, 4},
				{0, 5, 5, 5, 5, 5, 5},
			},
		),
		// Cleric domain spell slots stay outside this unrestricted progression table.
		ClericClassID: mustNewSpellcastingProgressionTable(
			ClericClassID,
			fullPreparedCasterSpellSlotsByClassLevel(),
		),
		DruidClassID: mustNewSpellcastingProgressionTable(
			DruidClassID,
			fullPreparedCasterSpellSlotsByClassLevel(),
		),
		PaladinClassID: mustNewSpellcastingProgressionTable(
			PaladinClassID,
			delayedFourthLevelCasterSpellSlotsByClassLevel(),
		),
		RangerClassID: mustNewSpellcastingProgressionTable(
			RangerClassID,
			delayedFourthLevelCasterSpellSlotsByClassLevel(),
		),
		SorcererClassID: mustNewSpellcastingProgressionTable(
			SorcererClassID,
			[][]int{
				{0, 3},
				{0, 4},
				{0, 5},
				{0, 6, 3},
				{0, 6, 4},
				{0, 6, 5, 3},
				{0, 6, 6, 4},
				{0, 6, 6, 5, 3},
				{0, 6, 6, 6, 4},
				{0, 6, 6, 6, 5, 3},
				{0, 6, 6, 6, 6, 4},
				{0, 6, 6, 6, 6, 5, 3},
				{0, 6, 6, 6, 6, 6, 4},
				{0, 6, 6, 6, 6, 6, 5, 3},
				{0, 6, 6, 6, 6, 6, 6, 4},
				{0, 6, 6, 6, 6, 6, 6, 5, 3},
				{0, 6, 6, 6, 6, 6, 6, 6, 4},
				{0, 6, 6, 6, 6, 6, 6, 6, 5, 3},
				{0, 6, 6, 6, 6, 6, 6, 6, 6, 4},
				{0, 6, 6, 6, 6, 6, 6, 6, 6, 6},
			},
		),
		WizardClassID: mustNewSpellcastingProgressionTable(
			WizardClassID,
			fullPreparedCasterSpellSlotsByClassLevel(),
		),
	}
}

func mustNewSpellcastingProgressionTable(
	classID ClassID,
	slotsByClassLevel [][]int,
) SpellcastingProgressionTable {
	value, ok := NewSpellcastingProgressionTable(classID, slotsByClassLevel)
	if !ok {
		panic("invalid core spellcasting progression seed")
	}

	return value
}

func fullPreparedCasterSpellSlotsByClassLevel() [][]int {
	return [][]int{
		{3, 1},
		{4, 2},
		{4, 2, 1},
		{4, 3, 2},
		{4, 3, 2, 1},
		{4, 3, 3, 2},
		{4, 4, 3, 2, 1},
		{4, 4, 3, 3, 2},
		{4, 4, 4, 3, 2, 1},
		{4, 4, 4, 3, 3, 2},
		{4, 4, 4, 4, 3, 2, 1},
		{4, 4, 4, 4, 3, 3, 2},
		{4, 4, 4, 4, 4, 3, 2, 1},
		{4, 4, 4, 4, 4, 3, 3, 2},
		{4, 4, 4, 4, 4, 4, 3, 2, 1},
		{4, 4, 4, 4, 4, 4, 3, 3, 2},
		{4, 4, 4, 4, 4, 4, 4, 3, 2, 1},
		{4, 4, 4, 4, 4, 4, 4, 3, 3, 2},
		{4, 4, 4, 4, 4, 4, 4, 4, 3, 3},
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	}
}

func delayedFourthLevelCasterSpellSlotsByClassLevel() [][]int {
	return [][]int{
		nil,
		nil,
		nil,
		{0},
		{0},
		{1},
		{1, 0},
		{1, 1},
		{2, 1},
		{2, 1, 0},
		{2, 1, 1},
		{2, 2, 1},
		{3, 2, 1, 0},
		{3, 2, 1, 1},
		{3, 2, 2, 1},
		{3, 3, 2, 1},
		{4, 3, 2, 1},
		{4, 3, 2, 2},
		{4, 3, 3, 2},
		{4, 4, 3, 3},
	}
}

func cloneSpellcastingProgressionTable(value SpellcastingProgressionTable) SpellcastingProgressionTable {
	slotsByClassLevel := make([][]int, 0, len(value.slotsByClassLevel))
	for _, row := range value.slotsByClassLevel {
		slotsByClassLevel = append(slotsByClassLevel, append([]int(nil), row...))
	}

	return spellcastingProgressionTable{
		classID:           value.classID,
		slotsByClassLevel: slotsByClassLevel,
	}
}
