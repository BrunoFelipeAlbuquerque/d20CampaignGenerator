package race

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

type raceID string
type RaceID = raceID

type languageID string
type LanguageID = languageID

type racialFeatureID string
type RacialFeatureID = racialFeatureID

type abilityScoreModifier struct {
	scoreID  ability.AbilityScoreID
	modifier int
}
type AbilityScoreModifier = abilityScoreModifier

type race struct {
	id                    raceID
	size                  ability.Size
	baseSpeed             int
	abilityScoreModifiers []abilityScoreModifier
	racialLanguages       []languageID
	racialFeatures        []racialFeatureID
}
type Race = race

func NewAbilityScoreModifier(scoreID ability.AbilityScoreID, modifier int) (AbilityScoreModifier, bool) {
	if scoreID.GetName() == "" || modifier == 0 {
		return abilityScoreModifier{}, false
	}

	return abilityScoreModifier{
		scoreID:  scoreID,
		modifier: modifier,
	}, true
}

func NewRace(
	id RaceID,
	size ability.Size,
	baseSpeed int,
	abilityScoreModifiers []AbilityScoreModifier,
	racialLanguages []LanguageID,
	racialFeatures []RacialFeatureID,
) (Race, bool) {
	if !isValidRaceID(id) || !isValidSize(size) || baseSpeed <= 0 {
		return race{}, false
	}

	dedupedModifiers, ok := dedupeAbilityScoreModifiers(abilityScoreModifiers)
	if !ok {
		return race{}, false
	}

	dedupedLanguages, ok := dedupeLanguageIDs(racialLanguages)
	if !ok {
		return race{}, false
	}

	dedupedFeatures, ok := dedupeRacialFeatureIDs(racialFeatures)
	if !ok {
		return race{}, false
	}

	return race{
		id:                    id,
		size:                  size,
		baseSpeed:             baseSpeed,
		abilityScoreModifiers: dedupedModifiers,
		racialLanguages:       dedupedLanguages,
		racialFeatures:        dedupedFeatures,
	}, true
}

func (m abilityScoreModifier) GetScoreID() ability.AbilityScoreID {
	return m.scoreID
}

func (m abilityScoreModifier) GetModifier() int {
	return m.modifier
}

func (r race) GetID() RaceID {
	return r.id
}

func (r race) GetSize() ability.Size {
	return r.size
}

func (r race) GetBaseSpeed() int {
	return r.baseSpeed
}

func (r race) GetAbilityScoreModifiers() []AbilityScoreModifier {
	return append([]AbilityScoreModifier(nil), r.abilityScoreModifiers...)
}

func (r race) GetRacialLanguages() []LanguageID {
	return append([]LanguageID(nil), r.racialLanguages...)
}

func (r race) GetRacialFeatures() []RacialFeatureID {
	return append([]RacialFeatureID(nil), r.racialFeatures...)
}

func (r race) HasRacialFeature(featureID RacialFeatureID) bool {
	for _, current := range r.racialFeatures {
		if current == featureID {
			return true
		}
	}

	return false
}

func isValidRaceID(value RaceID) bool {
	return value != ""
}

func isValidSize(value ability.Size) bool {
	_, ok := value.GetModifier()
	return ok
}

func isValidLanguageID(value LanguageID) bool {
	return value != ""
}

func isValidRacialFeatureID(value RacialFeatureID) bool {
	return value != ""
}

func dedupeAbilityScoreModifiers(modifiers []AbilityScoreModifier) ([]AbilityScoreModifier, bool) {
	if len(modifiers) == 0 {
		return nil, true
	}

	seen := make(map[ability.AbilityScoreID]struct{}, len(modifiers))
	deduped := make([]AbilityScoreModifier, 0, len(modifiers))

	for _, modifier := range modifiers {
		if modifier.scoreID.GetName() == "" || modifier.modifier == 0 {
			return nil, false
		}

		if _, ok := seen[modifier.scoreID]; ok {
			continue
		}

		seen[modifier.scoreID] = struct{}{}
		deduped = append(deduped, modifier)
	}

	return deduped, true
}

func dedupeLanguageIDs(languageIDs []LanguageID) ([]LanguageID, bool) {
	if len(languageIDs) == 0 {
		return nil, true
	}

	seen := make(map[LanguageID]struct{}, len(languageIDs))
	deduped := make([]LanguageID, 0, len(languageIDs))

	for _, languageID := range languageIDs {
		if !isValidLanguageID(languageID) {
			return nil, false
		}

		if _, ok := seen[languageID]; ok {
			continue
		}

		seen[languageID] = struct{}{}
		deduped = append(deduped, languageID)
	}

	return deduped, true
}

func dedupeRacialFeatureIDs(featureIDs []RacialFeatureID) ([]RacialFeatureID, bool) {
	if len(featureIDs) == 0 {
		return nil, true
	}

	seen := make(map[RacialFeatureID]struct{}, len(featureIDs))
	deduped := make([]RacialFeatureID, 0, len(featureIDs))

	for _, featureID := range featureIDs {
		if !isValidRacialFeatureID(featureID) {
			return nil, false
		}

		if _, ok := seen[featureID]; ok {
			continue
		}

		seen[featureID] = struct{}{}
		deduped = append(deduped, featureID)
	}

	return deduped, true
}
