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

	normalizedRows := make([][]int, 0, len(slotsByClassLevel))
	hasAnySpellSlots := false

	for _, row := range slotsByClassLevel {
		normalizedRow, ok := normalizeSpellSlotsByClassLevel(row)
		if !ok {
			return spellcastingProgressionTable{}, false
		}

		if len(normalizedRow) > 0 {
			hasAnySpellSlots = true
		}

		normalizedRows = append(normalizedRows, normalizedRow)
	}

	if !hasAnySpellSlots {
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

func (t spellcastingProgressionTable) GetSpellSlotsByClassLevel(classLevel int) ([]int, bool) {
	row, ok := t.getSpellSlotsByClassLevel(classLevel)
	if !ok {
		return nil, false
	}

	return append([]int(nil), row...), true
}

func (t spellcastingProgressionTable) GetSpellSlots(classLevel int, spellLevel int) (int, bool) {
	if spellLevel < 0 || spellLevel > maxCoreSpellLevel {
		return 0, false
	}

	row, ok := t.getSpellSlotsByClassLevel(classLevel)
	if !ok {
		return 0, false
	}

	if spellLevel >= len(row) {
		return 0, true
	}

	return row[spellLevel], true
}

func (t spellcastingProgressionTable) getSpellSlotsByClassLevel(classLevel int) ([]int, bool) {
	if classLevel < 1 || classLevel > len(t.slotsByClassLevel) {
		return nil, false
	}

	return t.slotsByClassLevel[classLevel-1], true
}

func normalizeSpellSlotsByClassLevel(slots []int) ([]int, bool) {
	if len(slots) > maxCoreSpellLevel+1 {
		return nil, false
	}

	normalized := append([]int(nil), slots...)
	lastNonZeroIndex := -1
	seenPositiveSlot := false
	seenGapAfterPositiveSlot := false

	for i, value := range normalized {
		if value < 0 {
			return nil, false
		}

		if value > 0 {
			if seenGapAfterPositiveSlot {
				return nil, false
			}

			seenPositiveSlot = true
			lastNonZeroIndex = i
			continue
		}

		if seenPositiveSlot {
			seenGapAfterPositiveSlot = true
		}
	}

	if lastNonZeroIndex == -1 {
		return nil, true
	}

	return normalized[:lastNonZeroIndex+1], true
}
