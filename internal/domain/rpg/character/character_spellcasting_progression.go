package character

import characterclass "d20campaigngenerator/internal/domain/rpg/character/class"

type characterSpellcastingProgression struct {
	classID characterclass.ClassID
}
type CharacterSpellcastingProgression = characterSpellcastingProgression

func NewCharacterSpellcastingProgression(
	class CharacterClass,
) (CharacterSpellcastingProgression, bool) {
	resolvedClass, ok := class.GetClass()
	if !ok || !resolvedClass.GetSpellcasting().HasSpellcasting() {
		return characterSpellcastingProgression{}, false
	}

	if _, ok := characterclass.GetSpellcastingProgressionByClassID(class.GetClassID()); !ok {
		return characterSpellcastingProgression{}, false
	}

	return characterSpellcastingProgression{
		classID: class.GetClassID(),
	}, true
}

func (p characterSpellcastingProgression) GetClassID() characterclass.ClassID {
	return p.classID
}

func (p characterSpellcastingProgression) GetClass() (characterclass.Class, bool) {
	return characterclass.GetClassByID(p.classID)
}

func (p characterSpellcastingProgression) GetProgression() (characterclass.SpellcastingProgressionTable, bool) {
	return characterclass.GetSpellcastingProgressionByClassID(p.classID)
}

func (p characterSpellcastingProgression) GetSpellSlots(classLevel int, spellLevel int) (int, bool) {
	progression, ok := p.GetProgression()
	if !ok {
		return 0, false
	}

	return progression.GetSpellSlots(classLevel, spellLevel)
}
