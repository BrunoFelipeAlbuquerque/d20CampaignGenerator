package character

import characterspell "d20campaigngenerator/internal/domain/rpg/character/spell"

type characterSelectedSpellSchool struct {
	id    characterspell.SchoolID
	valid bool
}
type CharacterSelectedSpellSchool = characterSelectedSpellSchool

func NewCharacterSelectedSpellSchool(
	id characterspell.SchoolID,
) (CharacterSelectedSpellSchool, bool) {
	if !isValidCharacterSpellSchoolID(id) {
		return characterSelectedSpellSchool{}, false
	}

	return characterSelectedSpellSchool{
		id:    id,
		valid: true,
	}, true
}

func (s characterSelectedSpellSchool) GetSchoolID() characterspell.SchoolID {
	if !s.valid {
		return ""
	}

	return s.id
}

func buildCharacterSelectedSpellSchool(
	value CharacterSelectedSpellSchool,
) (characterSelectedSpellSchool, bool) {
	if isEmptyCharacterSelectedSpellSchool(value) {
		return characterSelectedSpellSchool{}, true
	}

	if !value.valid {
		return characterSelectedSpellSchool{}, false
	}

	return NewCharacterSelectedSpellSchool(value.id)
}

func isEmptyCharacterSelectedSpellSchool(value CharacterSelectedSpellSchool) bool {
	return !value.valid && value.id == ""
}

func isValidCharacterSpellSchoolID(id characterspell.SchoolID) bool {
	switch id {
	case characterspell.AbjurationSchoolID,
		characterspell.ConjurationSchoolID,
		characterspell.DivinationSchoolID,
		characterspell.EnchantmentSchoolID,
		characterspell.EvocationSchoolID,
		characterspell.IllusionSchoolID,
		characterspell.NecromancySchoolID,
		characterspell.TransmutationSchoolID,
		characterspell.UniversalSchoolID:
		return true
	default:
		return false
	}
}
