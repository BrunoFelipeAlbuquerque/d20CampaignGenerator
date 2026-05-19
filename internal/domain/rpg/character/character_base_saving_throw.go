package character

import (
	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

type characterBaseSavingThrowFacts struct {
	valid     bool
	fortitude ability.SavingThrow
	reflex    ability.SavingThrow
	will      ability.SavingThrow
}
type CharacterBaseSavingThrowFacts = characterBaseSavingThrowFacts

func NewCharacterBaseSavingThrowFacts(
	classLevels []CharacterClassLevel,
) (CharacterBaseSavingThrowFacts, bool) {
	levelFacts, ok := NewCharacterLevelFacts(classLevels)
	if !ok {
		return characterBaseSavingThrowFacts{}, false
	}

	fortitude, ok := newZeroBaseSavingThrow(ability.FortitudeSave)
	if !ok {
		return characterBaseSavingThrowFacts{}, false
	}

	reflex, ok := newZeroBaseSavingThrow(ability.ReflexSave)
	if !ok {
		return characterBaseSavingThrowFacts{}, false
	}

	will, ok := newZeroBaseSavingThrow(ability.WillSave)
	if !ok {
		return characterBaseSavingThrowFacts{}, false
	}

	for _, classLevel := range levelFacts.GetClassLevels() {
		class, ok := characterclass.GetClassByID(classLevel.GetClassID())
		if !ok {
			return characterBaseSavingThrowFacts{}, false
		}

		saveProgressions := class.GetSaveProgressions()
		if !fortitude.AddClassLevel(classLevel.GetLevel(), saveProgressions.GetFortitude()) ||
			!reflex.AddClassLevel(classLevel.GetLevel(), saveProgressions.GetReflex()) ||
			!will.AddClassLevel(classLevel.GetLevel(), saveProgressions.GetWill()) {
			return characterBaseSavingThrowFacts{}, false
		}
	}

	return characterBaseSavingThrowFacts{
		valid:     true,
		fortitude: fortitude,
		reflex:    reflex,
		will:      will,
	}, true
}

func (f characterBaseSavingThrowFacts) GetFortitude() (ability.SavingThrow, bool) {
	return f.GetSavingThrow(ability.FortitudeSave)
}

func (f characterBaseSavingThrowFacts) GetReflex() (ability.SavingThrow, bool) {
	return f.GetSavingThrow(ability.ReflexSave)
}

func (f characterBaseSavingThrowFacts) GetWill() (ability.SavingThrow, bool) {
	return f.GetSavingThrow(ability.WillSave)
}

func (f characterBaseSavingThrowFacts) GetSavingThrow(
	id ability.SavingThrowID,
) (ability.SavingThrow, bool) {
	if !f.valid {
		var empty ability.SavingThrow
		return empty, false
	}

	switch id {
	case ability.FortitudeSave:
		return f.fortitude, f.fortitude.GetID() == id
	case ability.ReflexSave:
		return f.reflex, f.reflex.GetID() == id
	case ability.WillSave:
		return f.will, f.will.GetID() == id
	default:
		var empty ability.SavingThrow
		return empty, false
	}
}

func newZeroBaseSavingThrow(id ability.SavingThrowID) (ability.SavingThrow, bool) {
	zeroValue, ok := ability.NewRationalValue(0, 1)
	if !ok {
		var empty ability.SavingThrow
		return empty, false
	}

	return ability.NewSavingThrow(id, zeroValue)
}
