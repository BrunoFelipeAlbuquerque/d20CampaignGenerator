package character

type characterSelectedFamiliarEligibility struct {
	eligible bool
	valid    bool
}
type CharacterSelectedFamiliarEligibility = characterSelectedFamiliarEligibility

func NewCharacterSelectedFamiliarEligibility() CharacterSelectedFamiliarEligibility {
	return characterSelectedFamiliarEligibility{
		eligible: true,
		valid:    true,
	}
}

func (e characterSelectedFamiliarEligibility) IsEligible() bool {
	return e.valid && e.eligible
}

func buildCharacterSelectedFamiliarEligibility(
	value CharacterSelectedFamiliarEligibility,
) (characterSelectedFamiliarEligibility, bool) {
	if isEmptyCharacterSelectedFamiliarEligibility(value) {
		return characterSelectedFamiliarEligibility{}, true
	}

	if !value.valid || !value.eligible {
		return characterSelectedFamiliarEligibility{}, false
	}

	return NewCharacterSelectedFamiliarEligibility(), true
}

func isEmptyCharacterSelectedFamiliarEligibility(value CharacterSelectedFamiliarEligibility) bool {
	return !value.valid && !value.eligible
}
