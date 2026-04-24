package spell

import characterclass "d20campaigngenerator/internal/domain/rpg/character/class"

const maxCoreSpellLevel = 9

type spellListEntry struct {
	spellID    spellID
	classID    characterclass.ClassID
	spellLevel int
}
type SpellListEntry = spellListEntry

func NewSpellListEntry(
	spellID SpellID,
	classID characterclass.ClassID,
	spellLevel int,
) (SpellListEntry, bool) {
	if !isValidSpellID(spellID) ||
		!isValidSpellListClassID(classID) ||
		!isValidSpellListLevel(classID, spellLevel) {
		return spellListEntry{}, false
	}

	return spellListEntry{
		spellID:    spellID,
		classID:    classID,
		spellLevel: spellLevel,
	}, true
}

func (e spellListEntry) GetSpellID() SpellID {
	return e.spellID
}

func (e spellListEntry) GetClassID() characterclass.ClassID {
	return e.classID
}

func (e spellListEntry) GetSpellLevel() int {
	return e.spellLevel
}

func isValidSpellListClassID(id characterclass.ClassID) bool {
	class, ok := characterclass.GetClassByID(id)
	return ok && class.GetSpellcasting().HasSpellcasting()
}

func isValidSpellListLevel(classID characterclass.ClassID, spellLevel int) bool {
	if spellLevel < getMinimumSpellListLevel(classID) ||
		spellLevel > getMaximumSpellListLevel(classID) {
		return false
	}

	return true
}

func getMinimumSpellListLevel(classID characterclass.ClassID) int {
	switch classID {
	case characterclass.PaladinClassID, characterclass.RangerClassID:
		return 1
	default:
		return 0
	}
}

func getMaximumSpellListLevel(classID characterclass.ClassID) int {
	switch classID {
	case characterclass.BardClassID:
		return 6
	case characterclass.PaladinClassID, characterclass.RangerClassID:
		return 4
	default:
		return maxCoreSpellLevel
	}
}
