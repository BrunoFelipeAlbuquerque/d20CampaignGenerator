package ability

type abilityScoreID string
type AbilityScoreID = abilityScoreID

type abilityScoreValue struct {
	value int
	valid bool
}
type AbilityScoreValue = abilityScoreValue

type abilityScore struct {
	id    abilityScoreID
	name  string
	value abilityScoreValue
}
type AbilityScore = abilityScore

const (
	StrengthScore     AbilityScoreID = "STR"
	DexterityScore    AbilityScoreID = "DEX"
	ConstitutionScore AbilityScoreID = "CON"
	IntelligenceScore AbilityScoreID = "INT"
	WisdomScore       AbilityScoreID = "WIS"
	CharismaScore     AbilityScoreID = "CHA"
)

func NewAbilityScoreValue(value int, valid bool) AbilityScoreValue {
	return abilityScoreValue{
		value: value,
		valid: valid,
	}
}

func NewAbilityScore(id AbilityScoreID, value AbilityScoreValue) (AbilityScore, bool) {
	if id.GetName() == "" {
		return abilityScore{}, false
	}

	return abilityScore{
		id:    id,
		name:  id.GetName(),
		value: value,
	}, true
}

func (id abilityScoreID) GetName() string {
	switch id {
	case StrengthScore:
		return "Strength"
	case DexterityScore:
		return "Dexterity"
	case ConstitutionScore:
		return "Constitution"
	case IntelligenceScore:
		return "Intelligence"
	case WisdomScore:
		return "Wisdom"
	case CharismaScore:
		return "Charisma"
	default:
		return ""
	}
}

func (v abilityScoreValue) GetValue() (int, bool) {
	return v.value, v.valid
}

func (v abilityScoreValue) IsValid() bool {
	return v.valid
}

func (v *abilityScoreValue) SetValue(value int) {
	v.value = value
}

func (v *abilityScoreValue) SetValid(valid bool) {
	v.valid = valid
}

func (a abilityScore) GetID() AbilityScoreID {
	return a.id
}

func (a abilityScore) GetName() string {
	return a.name
}

func (a abilityScore) GetValue() AbilityScoreValue {
	return a.value
}

func (a *abilityScore) SetValue(value AbilityScoreValue) {
	a.value = value
}

func (a *abilityScore) SetScoreValue(value int) {
	a.value.value = value
}

func (a *abilityScore) SetValueValidity(valid bool) {
	a.value.valid = valid
}

func (a abilityScore) GetModifier() (int, bool) {
	score, valid := a.value.GetValue()
	if !valid {
		return 0, false
	}

	return calculateAbilityModifier(score), true
}

func calculateAbilityModifier(score int) int {
	delta := score - 10
	if delta >= 0 || delta%2 == 0 {
		return delta / 2
	}

	return (delta / 2) - 1
}
