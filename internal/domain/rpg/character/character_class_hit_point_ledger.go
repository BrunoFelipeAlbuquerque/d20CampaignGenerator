package character

import (
	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
)

type characterClassHitPointEntry struct {
	classID       characterclass.ClassID
	classLevel    int
	baseHitPoints int
}
type CharacterClassHitPointEntry = characterClassHitPointEntry

type characterClassHitPointFacts struct {
	valid             bool
	firstLevelClassID characterclass.ClassID
	entries           []CharacterClassHitPointEntry
	hitPoints         ability.HitPoints
}
type CharacterClassHitPointFacts = characterClassHitPointFacts

func NewCharacterClassHitPointEntry(
	id characterclass.ClassID,
	classLevel int,
	baseHitPoints int,
) (CharacterClassHitPointEntry, bool) {
	class, ok := characterclass.GetClassByID(id)
	if !ok || classLevel <= 0 || classLevel > maxCoreCharacterLevel {
		return characterClassHitPointEntry{}, false
	}

	maximumHitPoints, ok := maximumClassHitDieValue(class.GetHitDieType())
	if !ok || baseHitPoints <= 0 || baseHitPoints > maximumHitPoints {
		return characterClassHitPointEntry{}, false
	}

	return characterClassHitPointEntry{
		classID:       id,
		classLevel:    classLevel,
		baseHitPoints: baseHitPoints,
	}, true
}

func NewCharacterClassHitPointFacts(
	classLevels []CharacterClassLevel,
	firstLevelClassID characterclass.ClassID,
	constitutionScore int,
	entries []CharacterClassHitPointEntry,
) (CharacterClassHitPointFacts, bool) {
	levelFacts, ok := NewCharacterLevelFacts(classLevels)
	if !ok {
		return characterClassHitPointFacts{}, false
	}

	firstClass, ok := characterclass.GetClassByID(firstLevelClassID)
	if !ok {
		return characterClassHitPointFacts{}, false
	}

	if firstLevel, ok := levelFacts.GetClassLevel(firstLevelClassID); !ok || firstLevel <= 0 {
		return characterClassHitPointFacts{}, false
	}

	hitDie, ok := newCharacterClassHitDie(levelFacts.GetClassLevels())
	if !ok {
		return characterClassHitPointFacts{}, false
	}

	firstLevelBaseHitPoints, ok := maximumClassHitDieValue(firstClass.GetHitDieType())
	if !ok {
		return characterClassHitPointFacts{}, false
	}

	laterBaseHitPoints, normalizedEntries, ok := resolveLaterClassHitPoints(
		levelFacts,
		firstLevelClassID,
		entries,
	)
	if !ok {
		return characterClassHitPointFacts{}, false
	}

	hitPoints, ok := ability.NewExplicitStandardHitPoints(
		hitDie,
		constitutionScore,
		firstLevelBaseHitPoints+laterBaseHitPoints,
	)
	if !ok {
		return characterClassHitPointFacts{}, false
	}

	return characterClassHitPointFacts{
		valid:             true,
		firstLevelClassID: firstLevelClassID,
		entries:           normalizedEntries,
		hitPoints:         hitPoints,
	}, true
}

func (e characterClassHitPointEntry) GetClassID() characterclass.ClassID {
	return e.classID
}

func (e characterClassHitPointEntry) GetClassLevel() int {
	return e.classLevel
}

func (e characterClassHitPointEntry) GetBaseHitPoints() int {
	return e.baseHitPoints
}

func (f characterClassHitPointFacts) GetFirstLevelClassID() (characterclass.ClassID, bool) {
	return f.firstLevelClassID, f.valid
}

func (f characterClassHitPointFacts) GetEntries() []CharacterClassHitPointEntry {
	if !f.valid {
		return nil
	}

	return append([]CharacterClassHitPointEntry(nil), f.entries...)
}

func (f characterClassHitPointFacts) GetHitPoints() (ability.HitPoints, bool) {
	return f.hitPoints, f.valid
}

func newCharacterClassHitDie(classLevels []CharacterClassLevel) (ability.HitDie, bool) {
	d6Count := 0
	d8Count := 0
	d10Count := 0
	d12Count := 0

	for _, classLevel := range classLevels {
		class, ok := characterclass.GetClassByID(classLevel.GetClassID())
		if !ok {
			return ability.HitDie{}, false
		}

		switch class.GetHitDieType() {
		case ability.D6HitDie:
			d6Count += classLevel.GetLevel()
		case ability.D8HitDie:
			d8Count += classLevel.GetLevel()
		case ability.D10HitDie:
			d10Count += classLevel.GetLevel()
		case ability.D12HitDie:
			d12Count += classLevel.GetLevel()
		default:
			return ability.HitDie{}, false
		}
	}

	return ability.NewHitDie(d6Count, d8Count, d10Count, d12Count)
}

func resolveLaterClassHitPoints(
	levelFacts CharacterLevelFacts,
	firstLevelClassID characterclass.ClassID,
	entries []CharacterClassHitPointEntry,
) (int, []CharacterClassHitPointEntry, bool) {
	expectedEntries := expectedClassHitPointEntrySet(levelFacts, firstLevelClassID)
	if len(entries) != len(expectedEntries) {
		return 0, nil, false
	}

	seenEntries := make(map[characterClassHitPointEntryKey]struct{}, len(entries))
	normalizedEntries := make([]CharacterClassHitPointEntry, 0, len(entries))
	totalBaseHitPoints := 0

	for _, entry := range entries {
		normalizedEntry, ok := NewCharacterClassHitPointEntry(
			entry.classID,
			entry.classLevel,
			entry.baseHitPoints,
		)
		if !ok {
			return 0, nil, false
		}

		key := characterClassHitPointEntryKey{
			classID:    normalizedEntry.classID,
			classLevel: normalizedEntry.classLevel,
		}
		if _, ok := expectedEntries[key]; !ok {
			return 0, nil, false
		}

		if _, ok := seenEntries[key]; ok {
			return 0, nil, false
		}

		seenEntries[key] = struct{}{}
		normalizedEntries = append(normalizedEntries, normalizedEntry)
		totalBaseHitPoints += normalizedEntry.baseHitPoints
	}

	return totalBaseHitPoints, normalizedEntries, true
}

type characterClassHitPointEntryKey struct {
	classID    characterclass.ClassID
	classLevel int
}

func expectedClassHitPointEntrySet(
	levelFacts CharacterLevelFacts,
	firstLevelClassID characterclass.ClassID,
) map[characterClassHitPointEntryKey]struct{} {
	result := make(map[characterClassHitPointEntryKey]struct{})

	for _, classLevel := range levelFacts.GetClassLevels() {
		for level := 1; level <= classLevel.GetLevel(); level++ {
			if classLevel.GetClassID() == firstLevelClassID && level == 1 {
				continue
			}

			result[characterClassHitPointEntryKey{
				classID:    classLevel.GetClassID(),
				classLevel: level,
			}] = struct{}{}
		}
	}

	return result
}

func maximumClassHitDieValue(hitDieType ability.HitDieType) (int, bool) {
	switch hitDieType {
	case ability.D6HitDie:
		return 6, true
	case ability.D8HitDie:
		return 8, true
	case ability.D10HitDie:
		return 10, true
	case ability.D12HitDie:
		return 12, true
	default:
		return 0, false
	}
}
