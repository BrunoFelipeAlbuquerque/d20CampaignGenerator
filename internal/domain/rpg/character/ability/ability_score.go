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

func NewAbilityScoreValue(value int, valid bool) (AbilityScoreValue, bool) {
	if value < 0 {
		return abilityScoreValue{}, false
	}

	return abilityScoreValue{
		value: value,
		valid: valid,
	}, true
}

func NewAbilityScore(id AbilityScoreID, value AbilityScoreValue) (AbilityScore, bool) {
	if id.GetName() == "" || !isValidAbilityScoreValue(value) {
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

func (v abilityScoreValue) WithValue(value int) (AbilityScoreValue, bool) {
	if value < 0 {
		return abilityScoreValue{}, false
	}

	v.value = value
	return v, true
}

func (v abilityScoreValue) WithValid(valid bool) AbilityScoreValue {
	v.valid = valid
	return v
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

func (a *abilityScore) SetValue(value AbilityScoreValue) bool {
	if !isValidAbilityScoreValue(value) {
		return false
	}

	a.value = value
	return true
}

func (a *abilityScore) SetScoreValue(value int) bool {
	if value < 0 {
		return false
	}

	a.value.value = value
	return true
}

func (a *abilityScore) SetValueValidity(valid bool) bool {
	a.value.valid = valid
	return true
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

func isValidAbilityScoreValue(value AbilityScoreValue) bool {
	return value.value >= 0
}
