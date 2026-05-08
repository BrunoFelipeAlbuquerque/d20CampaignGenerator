package character

import (
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterspell "d20campaigngenerator/internal/domain/rpg/character/spell"
)

type characterSpellListEntry struct {
	spellID    characterspell.SpellID
	classID    characterclass.ClassID
	spellLevel int
}
type CharacterSpellListEntry = characterSpellListEntry

func NewCharacterSpellListEntry(
	progression CharacterSpellcastingProgression,
	entry characterspell.SpellListEntry,
) (CharacterSpellListEntry, bool) {
	if !spellcastingProgressionSupportsSpellListEntry(progression, entry) {
		return characterSpellListEntry{}, false
	}

	return characterSpellListEntry{
		spellID:    entry.GetSpellID(),
		classID:    entry.GetClassID(),
		spellLevel: entry.GetSpellLevel(),
	}, true
}

func GetCharacterSpellListEntries(
	progression CharacterSpellcastingProgression,
) ([]CharacterSpellListEntry, bool) {
	if _, ok := progression.GetProgression(); !ok {
		return nil, false
	}

	entries, ok := characterspell.GetSpellListEntriesByClass(progression.GetClassID())
	if !ok {
		return nil, false
	}

	return composeCharacterSpellListEntries(progression, entries)
}

func GetCharacterSpellListEntriesBySpellLevel(
	progression CharacterSpellcastingProgression,
	spellLevel int,
) ([]CharacterSpellListEntry, bool) {
	if _, ok := progression.GetProgression(); !ok {
		return nil, false
	}

	entries, ok := characterspell.GetSpellListEntriesByClassAndLevel(
		progression.GetClassID(),
		spellLevel,
	)
	if !ok {
		return nil, false
	}

	return composeCharacterSpellListEntries(progression, entries)
}

func (e characterSpellListEntry) GetSpellID() characterspell.SpellID {
	return e.spellID
}

func (e characterSpellListEntry) GetClassID() characterclass.ClassID {
	return e.classID
}

func (e characterSpellListEntry) GetSpellLevel() int {
	return e.spellLevel
}

func (e characterSpellListEntry) GetSpell() (characterspell.Spell, bool) {
	return characterspell.GetSpellByID(e.spellID)
}

func (e characterSpellListEntry) GetSpellListEntry() (characterspell.SpellListEntry, bool) {
	entry, ok := seededSpellListEntry(e.spellID, e.classID, e.spellLevel)
	if !ok {
		return characterspell.SpellListEntry{}, false
	}

	return entry, true
}

func composeCharacterSpellListEntries(
	progression CharacterSpellcastingProgression,
	entries []characterspell.SpellListEntry,
) ([]CharacterSpellListEntry, bool) {
	composed := make([]CharacterSpellListEntry, 0, len(entries))

	for _, entry := range entries {
		value, ok := NewCharacterSpellListEntry(progression, entry)
		if !ok {
			return nil, false
		}

		composed = append(composed, value)
	}

	return composed, true
}

func spellcastingProgressionSupportsSpellListEntry(
	progression CharacterSpellcastingProgression,
	entry characterspell.SpellListEntry,
) bool {
	table, ok := progression.GetProgression()
	if !ok || entry.GetClassID() != progression.GetClassID() {
		return false
	}

	if _, ok := table.GetSpellSlots(table.GetMaxClassLevel(), entry.GetSpellLevel()); !ok {
		return false
	}

	if _, ok := seededSpellListEntry(entry.GetSpellID(), entry.GetClassID(), entry.GetSpellLevel()); !ok {
		return false
	}

	return true
}

func seededSpellListEntry(
	spellID characterspell.SpellID,
	classID characterclass.ClassID,
	spellLevel int,
) (characterspell.SpellListEntry, bool) {
	entries, ok := characterspell.GetSpellListEntriesByClassAndLevel(classID, spellLevel)
	if !ok {
		return characterspell.SpellListEntry{}, false
	}

	for _, entry := range entries {
		if entry.GetSpellID() == spellID {
			return entry, true
		}
	}

	return characterspell.SpellListEntry{}, false
}
