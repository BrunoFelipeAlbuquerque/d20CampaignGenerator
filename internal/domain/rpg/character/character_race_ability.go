package character

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

type characterSelectedAbilityScore struct {
	id    ability.AbilityScoreID
	valid bool
}
type CharacterSelectedAbilityScore = characterSelectedAbilityScore

func NewCharacterSelectedAbilityScore(
	id ability.AbilityScoreID,
) (CharacterSelectedAbilityScore, bool) {
	if !isCoreCharacterAbilityScoreID(id) {
		return characterSelectedAbilityScore{}, false
	}

	return characterSelectedAbilityScore{
		id:    id,
		valid: true,
	}, true
}

func NewFixedRacialCharacterAbilityScores(
	selectedRace CharacterRace,
	baseScores []CharacterAbilityScore,
) ([]CharacterAbilityScore, bool) {
	race, ok := selectedRace.GetRace()
	if !ok {
		return nil, false
	}

	if _, ok := race.GetSelectableAbilityScoreModifier(); ok {
		return nil, false
	}

	composedScores, ok := buildCompleteCharacterAbilityScoreMap(baseScores)
	if !ok {
		return nil, false
	}

	for _, modifier := range race.GetAbilityScoreModifiers() {
		scoreID := modifier.GetScoreID()
		baseScore, ok := composedScores[scoreID]
		if !ok {
			return nil, false
		}

		score, ok := NewCharacterAbilityScore(scoreID, baseScore+modifier.GetModifier())
		if !ok {
			return nil, false
		}

		composedScores[scoreID] = score.GetScore()
	}

	return characterAbilityScoresFromMap(composedScores)
}

func NewSelectableRacialCharacterAbilityScores(
	selectedRace CharacterRace,
	baseScores []CharacterAbilityScore,
	selectedAbilities []CharacterSelectedAbilityScore,
) ([]CharacterAbilityScore, bool) {
	race, ok := selectedRace.GetRace()
	if !ok {
		return nil, false
	}

	selectableModifier, ok := race.GetSelectableAbilityScoreModifier()
	if !ok || len(race.GetAbilityScoreModifiers()) != 0 {
		return nil, false
	}

	selectedAbility, ok := buildSingleCharacterSelectedAbilityScore(selectedAbilities)
	if !ok {
		return nil, false
	}

	composedScores, ok := buildCompleteCharacterAbilityScoreMap(baseScores)
	if !ok {
		return nil, false
	}

	scoreID := selectedAbility.GetAbilityScoreID()
	baseScore, ok := composedScores[scoreID]
	if !ok {
		return nil, false
	}

	score, ok := NewCharacterAbilityScore(scoreID, baseScore+selectableModifier)
	if !ok {
		return nil, false
	}

	composedScores[scoreID] = score.GetScore()

	return characterAbilityScoresFromMap(composedScores)
}

func (s characterSelectedAbilityScore) GetAbilityScoreID() ability.AbilityScoreID {
	if !s.valid {
		return ""
	}

	return s.id
}

func coreCharacterAbilityScoreIDs() []ability.AbilityScoreID {
	return []ability.AbilityScoreID{
		ability.StrengthScore,
		ability.DexterityScore,
		ability.ConstitutionScore,
		ability.IntelligenceScore,
		ability.WisdomScore,
		ability.CharismaScore,
	}
}

func characterAbilityScoresFromMap(
	values map[ability.AbilityScoreID]int,
) ([]CharacterAbilityScore, bool) {
	coreScoreIDs := coreCharacterAbilityScoreIDs()
	result := make([]CharacterAbilityScore, 0, len(coreScoreIDs))

	for _, scoreID := range coreScoreIDs {
		value, ok := values[scoreID]
		if !ok {
			return nil, false
		}

		score, ok := NewCharacterAbilityScore(scoreID, value)
		if !ok {
			return nil, false
		}

		result = append(result, score)
	}

	return result, true
}

func buildCompleteCharacterAbilityScoreMap(
	values []CharacterAbilityScore,
) (map[ability.AbilityScoreID]int, bool) {
	coreScoreIDs := coreCharacterAbilityScoreIDs()
	result, ok := buildCharacterAbilityScoreMap(values)
	if !ok || len(result) != len(coreScoreIDs) {
		return nil, false
	}

	for _, scoreID := range coreScoreIDs {
		if _, ok := result[scoreID]; !ok {
			return nil, false
		}
	}

	return result, true
}

func buildSingleCharacterSelectedAbilityScore(
	values []CharacterSelectedAbilityScore,
) (CharacterSelectedAbilityScore, bool) {
	if len(values) != 1 {
		return characterSelectedAbilityScore{}, false
	}

	value := values[0]
	if !value.valid {
		return characterSelectedAbilityScore{}, false
	}

	return NewCharacterSelectedAbilityScore(value.id)
}

func isCoreCharacterAbilityScoreID(id ability.AbilityScoreID) bool {
	for _, scoreID := range coreCharacterAbilityScoreIDs() {
		if id == scoreID {
			return true
		}
	}

	return false
}
