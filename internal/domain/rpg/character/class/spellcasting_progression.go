package class

const (
	maxCoreSpellLevel = 9
	maxCoreClassLevel = 20
)

type spellcastingProgressionTable struct {
	classID           classID
	slotsByClassLevel [][]int
}
type SpellcastingProgressionTable = spellcastingProgressionTable

func NewSpellcastingProgressionTable(
	classID ClassID,
	slotsByClassLevel [][]int,
) (SpellcastingProgressionTable, bool) {
	class, ok := GetClassByID(classID)
	if !ok || !class.GetSpellcasting().HasSpellcasting() {
		return spellcastingProgressionTable{}, false
	}

	if len(slotsByClassLevel) == 0 || len(slotsByClassLevel) > maxCoreClassLevel {
		return spellcastingProgressionTable{}, false
	}

	minimumSpellLevel := getMinimumSpellLevelForClass(classID)
	normalizedRows := make([][]int, 0, len(slotsByClassLevel))
	hasAnyPositiveSpellSlots := false
	spellLevelsStarted := false
	highestAvailableSpellLevel := -1

	for _, row := range slotsByClassLevel {
		normalizedRow, rowHasPositiveSpellSlots, ok := normalizeSpellSlotsByClassLevel(row)
		if !ok {
			return spellcastingProgressionTable{}, false
		}

		if rowHasPositiveSpellSlots {
			hasAnyPositiveSpellSlots = true
		}

		if len(normalizedRow) == 0 {
			if spellLevelsStarted {
				return spellcastingProgressionTable{}, false
			}

			normalizedRows = append(normalizedRows, normalizedRow)
			continue
		}

		rowHighestAvailableSpellLevel := minimumSpellLevel + len(normalizedRow) - 1
		if rowHighestAvailableSpellLevel > maxCoreSpellLevel {
			return spellcastingProgressionTable{}, false
		}

		if !spellLevelsStarted {
			if rowHighestAvailableSpellLevel > 1 {
				return spellcastingProgressionTable{}, false
			}

			spellLevelsStarted = true
			highestAvailableSpellLevel = rowHighestAvailableSpellLevel
			normalizedRows = append(normalizedRows, normalizedRow)
			continue
		}

		if rowHighestAvailableSpellLevel < highestAvailableSpellLevel ||
			rowHighestAvailableSpellLevel > highestAvailableSpellLevel+1 {
			return spellcastingProgressionTable{}, false
		}

		highestAvailableSpellLevel = rowHighestAvailableSpellLevel
		normalizedRows = append(normalizedRows, normalizedRow)
	}

	if !hasAnyPositiveSpellSlots {
		return spellcastingProgressionTable{}, false
	}

	return spellcastingProgressionTable{
		classID:           classID,
		slotsByClassLevel: normalizedRows,
	}, true
}

func (t spellcastingProgressionTable) GetClassID() ClassID {
	return t.classID
}

func (t spellcastingProgressionTable) GetMaxClassLevel() int {
	return len(t.slotsByClassLevel)
}

func (t spellcastingProgressionTable) GetMinimumSpellLevel() int {
	return getMinimumSpellLevelForClass(t.classID)
}

func (t spellcastingProgressionTable) GetSpellSlotsByClassLevel(classLevel int) ([]int, bool) {
	row, ok := t.getSpellSlotsByClassLevel(classLevel)
	if !ok {
		return nil, false
	}

	return append([]int(nil), row...), true
}

func (t spellcastingProgressionTable) GetSpellSlots(classLevel int, spellLevel int) (int, bool) {
	minimumSpellLevel := getMinimumSpellLevelForClass(t.classID)
	if spellLevel < minimumSpellLevel || spellLevel > maxCoreSpellLevel {
		return 0, false
	}

	row, ok := t.getSpellSlotsByClassLevel(classLevel)
	if !ok {
		return 0, false
	}

	rowIndex := spellLevel - minimumSpellLevel
	if rowIndex >= len(row) {
		return 0, false
	}

	return row[rowIndex], true
}

func (t spellcastingProgressionTable) getSpellSlotsByClassLevel(classLevel int) ([]int, bool) {
	if classLevel < 1 || classLevel > len(t.slotsByClassLevel) {
		return nil, false
	}

	return t.slotsByClassLevel[classLevel-1], true
}

func normalizeSpellSlotsByClassLevel(slots []int) ([]int, bool, bool) {
	if len(slots) > maxCoreSpellLevel+1 {
		return nil, false, false
	}

	normalized := append([]int(nil), slots...)
	seenPositiveSlot := false
	seenGapAfterPositiveSlot := false
	hasAnyPositiveSpellSlots := false

	for _, value := range normalized {
		if value < 0 {
			return nil, false, false
		}

		if value > 0 {
			if seenGapAfterPositiveSlot {
				return nil, false, false
			}

			seenPositiveSlot = true
			hasAnyPositiveSpellSlots = true
			continue
		}

		if seenPositiveSlot {
			seenGapAfterPositiveSlot = true
		}
	}

	return normalized, hasAnyPositiveSpellSlots, true
}

func getMinimumSpellLevelForClass(classID ClassID) int {
	switch classID {
	case PaladinClassID, RangerClassID:
		return 1
	default:
		return 0
	}
}
