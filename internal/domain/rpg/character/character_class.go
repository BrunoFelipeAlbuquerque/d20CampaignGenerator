package character

import characterclass "d20campaigngenerator/internal/domain/rpg/character/class"

type characterClass struct {
	id characterclass.ClassID
}
type CharacterClass = characterClass

func NewCharacterClass(id characterclass.ClassID) (CharacterClass, bool) {
	if _, ok := characterclass.GetClassByID(id); !ok {
		return characterClass{}, false
	}

	return characterClass{id: id}, true
}

func (c characterClass) GetClassID() characterclass.ClassID {
	return c.id
}

func (c characterClass) GetClass() (characterclass.Class, bool) {
	return characterclass.GetClassByID(c.id)
}
