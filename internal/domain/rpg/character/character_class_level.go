package character

import characterclass "d20campaigngenerator/internal/domain/rpg/character/class"

const maxCoreCharacterLevel = 20

type characterLevelFacts struct {
	totalLevel  int
	classLevels []CharacterClassLevel
}
type CharacterLevelFacts = characterLevelFacts

func NewCharacterLevelFacts(
	classLevels []CharacterClassLevel,
) (CharacterLevelFacts, bool) {
	classLevelMap, ok := buildCharacterClassLevelMap(classLevels)
	if !ok || len(classLevelMap) == 0 {
		return characterLevelFacts{}, false
	}

	totalLevel := 0
	for _, level := range classLevelMap {
		if level > maxCoreCharacterLevel || totalLevel > maxCoreCharacterLevel-level {
			return characterLevelFacts{}, false
		}

		totalLevel += level
	}

	normalizedClassLevels, ok := characterClassLevelsFromMap(classLevelMap)
	if !ok {
		return characterLevelFacts{}, false
	}

	return characterLevelFacts{
		totalLevel:  totalLevel,
		classLevels: normalizedClassLevels,
	}, true
}

func (f characterLevelFacts) GetTotalCharacterLevel() int {
	return f.totalLevel
}

func (f characterLevelFacts) GetClassLevel(
	id characterclass.ClassID,
) (int, bool) {
	if _, ok := characterclass.GetClassByID(id); !ok {
		return 0, false
	}

	for _, classLevel := range f.classLevels {
		if classLevel.GetClassID() == id {
			return classLevel.GetLevel(), true
		}
	}

	return 0, false
}

func (f characterLevelFacts) GetClassLevels() []CharacterClassLevel {
	return append([]CharacterClassLevel(nil), f.classLevels...)
}

func characterClassLevelsFromMap(
	values map[characterclass.ClassID]int,
) ([]CharacterClassLevel, bool) {
	result := make([]CharacterClassLevel, 0, len(values))

	for _, class := range characterclass.GetClasses() {
		level, ok := values[class.GetID()]
		if !ok {
			continue
		}

		classLevel, ok := NewCharacterClassLevel(class.GetID(), level)
		if !ok {
			return nil, false
		}

		result = append(result, classLevel)
	}

	return result, len(result) == len(values)
}
