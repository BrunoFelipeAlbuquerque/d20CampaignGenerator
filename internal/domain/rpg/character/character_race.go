package character

import characterrace "d20campaigngenerator/internal/domain/rpg/character/race"

type characterRace struct {
	id characterrace.RaceID
}
type CharacterRace = characterRace

func NewCharacterRace(id characterrace.RaceID) (CharacterRace, bool) {
	if _, ok := characterrace.GetRaceByID(id); !ok {
		return characterRace{}, false
	}

	return characterRace{id: id}, true
}

func (r characterRace) GetRaceID() characterrace.RaceID {
	return r.id
}

func (r characterRace) GetRace() (characterrace.Race, bool) {
	return characterrace.GetRaceByID(r.id)
}
