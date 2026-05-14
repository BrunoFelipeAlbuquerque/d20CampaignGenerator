package character

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

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

	coreScoreIDs := coreCharacterAbilityScoreIDs()
	result := make([]CharacterAbilityScore, 0, len(coreScoreIDs))
	for _, scoreID := range coreScoreIDs {
		score, ok := NewCharacterAbilityScore(scoreID, composedScores[scoreID])
		if !ok {
			return nil, false
		}

		result = append(result, score)
	}

	return result, true
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
