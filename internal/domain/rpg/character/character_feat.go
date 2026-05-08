package character

import (
	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterfeat "d20campaigngenerator/internal/domain/rpg/character/feat"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

type characterAbilityScore struct {
	id    ability.AbilityScoreID
	score int
}
type CharacterAbilityScore = characterAbilityScore

type characterClassLevel struct {
	classID characterclass.ClassID
	level   int
}
type CharacterClassLevel = characterClassLevel

type characterSkillRanks struct {
	skillID skill.SkillID
	ranks   int
}
type CharacterSkillRanks = characterSkillRanks

type characterFeatPrerequisiteState struct {
	valid           bool
	abilityScores   map[ability.AbilityScoreID]int
	baseAttackBonus int
	classLevels     map[characterclass.ClassID]int
	classFeatures   map[characterclass.ClassFeatureID]struct{}
	skillRanks      map[skill.SkillID]int
	feats           map[characterfeat.FeatID]struct{}
}
type CharacterFeatPrerequisiteState = characterFeatPrerequisiteState

type characterFeat struct {
	id characterfeat.FeatID
}
type CharacterFeat = characterFeat

func NewCharacterAbilityScore(id ability.AbilityScoreID, score int) (CharacterAbilityScore, bool) {
	value, ok := ability.NewAbilityScoreValue(score, true)
	if !ok {
		return characterAbilityScore{}, false
	}

	if _, ok := ability.NewAbilityScore(id, value); !ok {
		return characterAbilityScore{}, false
	}

	return characterAbilityScore{
		id:    id,
		score: score,
	}, true
}

func NewCharacterClassLevel(
	id characterclass.ClassID,
	level int,
) (CharacterClassLevel, bool) {
	if _, ok := characterclass.GetClassByID(id); !ok || level <= 0 {
		return characterClassLevel{}, false
	}

	return characterClassLevel{
		classID: id,
		level:   level,
	}, true
}

func NewCharacterSkillRanks(id skill.SkillID, ranks int) (CharacterSkillRanks, bool) {
	if ranks <= 0 || !isValidCharacterSkillID(id) {
		return characterSkillRanks{}, false
	}

	return characterSkillRanks{
		skillID: id,
		ranks:   ranks,
	}, true
}

func NewCharacterFeatPrerequisiteState(
	abilityScores []CharacterAbilityScore,
	baseAttackBonus int,
	classLevels []CharacterClassLevel,
	classFeatures []characterclass.ClassFeatureID,
	skillRanks []CharacterSkillRanks,
	feats []characterfeat.FeatID,
) (CharacterFeatPrerequisiteState, bool) {
	if baseAttackBonus < 0 {
		return characterFeatPrerequisiteState{}, false
	}

	abilityScoreMap, ok := buildCharacterAbilityScoreMap(abilityScores)
	if !ok {
		return characterFeatPrerequisiteState{}, false
	}

	classLevelMap, ok := buildCharacterClassLevelMap(classLevels)
	if !ok {
		return characterFeatPrerequisiteState{}, false
	}

	classFeatureSet, ok := buildCharacterClassFeatureSet(classFeatures)
	if !ok {
		return characterFeatPrerequisiteState{}, false
	}

	skillRankMap, ok := buildCharacterSkillRankMap(skillRanks)
	if !ok {
		return characterFeatPrerequisiteState{}, false
	}

	featSet, ok := buildCharacterFeatSet(feats)
	if !ok {
		return characterFeatPrerequisiteState{}, false
	}

	return characterFeatPrerequisiteState{
		valid:           true,
		abilityScores:   abilityScoreMap,
		baseAttackBonus: baseAttackBonus,
		classLevels:     classLevelMap,
		classFeatures:   classFeatureSet,
		skillRanks:      skillRankMap,
		feats:           featSet,
	}, true
}

func NewCharacterFeat(
	id characterfeat.FeatID,
	prerequisites CharacterFeatPrerequisiteState,
) (CharacterFeat, bool) {
	feat, ok := characterfeat.GetFeatByID(id)
	if !ok || !prerequisites.SatisfiesFeat(feat) {
		return characterFeat{}, false
	}

	return characterFeat{id: id}, true
}

func (a characterAbilityScore) GetAbilityScoreID() ability.AbilityScoreID {
	return a.id
}

func (a characterAbilityScore) GetScore() int {
	return a.score
}

func (l characterClassLevel) GetClassID() characterclass.ClassID {
	return l.classID
}

func (l characterClassLevel) GetLevel() int {
	return l.level
}

func (r characterSkillRanks) GetSkillID() skill.SkillID {
	return r.skillID
}

func (r characterSkillRanks) GetRanks() int {
	return r.ranks
}

func (s characterFeatPrerequisiteState) SatisfiesFeat(feat characterfeat.Feat) bool {
	if !s.valid {
		return false
	}

	if _, ok := characterfeat.GetFeatByID(feat.GetID()); !ok {
		return false
	}

	for _, prerequisite := range feat.GetPrerequisites() {
		if !s.SatisfiesPrerequisite(prerequisite) {
			return false
		}
	}

	return true
}

func (s characterFeatPrerequisiteState) SatisfiesPrerequisite(
	prerequisite characterfeat.Prerequisite,
) bool {
	if !s.valid || prerequisite == nil {
		return false
	}

	switch value := prerequisite.(type) {
	case characterfeat.AbilityScorePrerequisite:
		score, ok := s.abilityScores[value.GetAbilityScoreID()]
		return ok && score >= value.GetMinimumScore()
	case characterfeat.BaseAttackBonusPrerequisite:
		return s.baseAttackBonus >= value.GetMinimumBonus()
	case characterfeat.SkillRanksPrerequisite:
		return s.skillRanks[value.GetSkillID()] >= value.GetMinimumRanks()
	case characterfeat.AnySkillRanksPrerequisite:
		for _, skillID := range value.GetSkillIDs() {
			if s.skillRanks[skillID] >= value.GetMinimumRanks() {
				return true
			}
		}
		return false
	case characterfeat.SpellcastingPrerequisite:
		return s.satisfiesSpellcastingAccess(value.GetAccess())
	case characterfeat.CharacterLevelPrerequisite:
		return s.totalCharacterLevel() >= value.GetMinimumLevel()
	case characterfeat.ClassLevelPrerequisite:
		return s.classLevels[value.GetClassID()] >= value.GetMinimumLevel()
	case characterfeat.ClassFeaturePrerequisite:
		_, ok := s.classFeatures[value.GetFeatureID()]
		return ok
	case characterfeat.FeatPrerequisite:
		_, ok := s.feats[value.GetFeatID()]
		return ok
	case characterfeat.AnyFeatPrerequisite:
		for _, featID := range value.GetFeatIDs() {
			if _, ok := s.feats[featID]; ok {
				return true
			}
		}
		return false
	case characterfeat.FeatCategoryCountPrerequisite:
		return s.featCategoryCount(value.GetCategory()) >= value.GetMinimumCount()
	default:
		return false
	}
}

func (f characterFeat) GetFeatID() characterfeat.FeatID {
	return f.id
}

func (f characterFeat) GetFeat() (characterfeat.Feat, bool) {
	return characterfeat.GetFeatByID(f.id)
}

func (s characterFeatPrerequisiteState) totalCharacterLevel() int {
	total := 0
	for _, level := range s.classLevels {
		total += level
	}

	return total
}

func (s characterFeatPrerequisiteState) satisfiesSpellcastingAccess(
	access characterfeat.SpellcastingAccess,
) bool {
	for classID := range s.classLevels {
		class, ok := characterclass.GetClassByID(classID)
		if !ok {
			return false
		}

		spellcasting := class.GetSpellcasting()
		if !spellcasting.HasSpellcasting() {
			continue
		}

		switch access {
		case characterfeat.AnySpellcastingAccess:
			return true
		case characterfeat.ArcaneSpellcastingAccess:
			if isArcaneSpellcastingKind(spellcasting.GetKind()) {
				return true
			}
		case characterfeat.DivineSpellcastingAccess:
			if isDivineSpellcastingKind(spellcasting.GetKind()) {
				return true
			}
		default:
			return false
		}
	}

	return false
}

func (s characterFeatPrerequisiteState) featCategoryCount(
	category characterfeat.FeatCategory,
) int {
	count := 0
	for featID := range s.feats {
		feat, ok := characterfeat.GetFeatByID(featID)
		if !ok {
			return 0
		}

		if feat.GetCategory() == category {
			count++
		}
	}

	return count
}

func buildCharacterAbilityScoreMap(
	values []CharacterAbilityScore,
) (map[ability.AbilityScoreID]int, bool) {
	result := make(map[ability.AbilityScoreID]int, len(values))

	for _, value := range values {
		if _, ok := NewCharacterAbilityScore(value.id, value.score); !ok {
			return nil, false
		}

		if _, ok := result[value.id]; ok {
			return nil, false
		}

		result[value.id] = value.score
	}

	return result, true
}

func buildCharacterClassLevelMap(
	values []CharacterClassLevel,
) (map[characterclass.ClassID]int, bool) {
	result := make(map[characterclass.ClassID]int, len(values))

	for _, value := range values {
		if _, ok := NewCharacterClassLevel(value.classID, value.level); !ok {
			return nil, false
		}

		if _, ok := result[value.classID]; ok {
			return nil, false
		}

		result[value.classID] = value.level
	}

	return result, true
}

func buildCharacterClassFeatureSet(
	values []characterclass.ClassFeatureID,
) (map[characterclass.ClassFeatureID]struct{}, bool) {
	result := make(map[characterclass.ClassFeatureID]struct{}, len(values))

	for _, value := range values {
		if value.GetName() == "" {
			return nil, false
		}

		if _, ok := result[value]; ok {
			return nil, false
		}

		result[value] = struct{}{}
	}

	return result, true
}

func buildCharacterSkillRankMap(values []CharacterSkillRanks) (map[skill.SkillID]int, bool) {
	result := make(map[skill.SkillID]int, len(values))

	for _, value := range values {
		if _, ok := NewCharacterSkillRanks(value.skillID, value.ranks); !ok {
			return nil, false
		}

		if _, ok := result[value.skillID]; ok {
			return nil, false
		}

		result[value.skillID] = value.ranks
	}

	return result, true
}

func buildCharacterFeatSet(
	values []characterfeat.FeatID,
) (map[characterfeat.FeatID]struct{}, bool) {
	result := make(map[characterfeat.FeatID]struct{}, len(values))

	for _, value := range values {
		if _, ok := characterfeat.GetFeatByID(value); !ok {
			return nil, false
		}

		if _, ok := result[value]; ok {
			return nil, false
		}

		result[value] = struct{}{}
	}

	return result, true
}

func isValidCharacterSkillID(id skill.SkillID) bool {
	if _, ok := skill.GetSkillByID(id); ok {
		return true
	}

	_, ok := skill.NewSkill(id, false, false, true)
	return ok
}

func isArcaneSpellcastingKind(kind characterclass.SpellcastingKind) bool {
	switch kind {
	case characterclass.ArcanePreparedSpellcastingKind,
		characterclass.ArcaneSpontaneousSpellcastingKind:
		return true
	default:
		return false
	}
}

func isDivineSpellcastingKind(kind characterclass.SpellcastingKind) bool {
	return kind == characterclass.DivinePreparedSpellcastingKind
}
