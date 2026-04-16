package ability

const goodSavingThrowBaseBonus = 2

type savingThrowID string
type SavingThrowID = savingThrowID

const (
	FortitudeSave SavingThrowID = "Fortitude"
	ReflexSave    SavingThrowID = "Reflex"
	WillSave      SavingThrowID = "Will"
)

type savingThrowProgression string
type SavingThrowProgression = savingThrowProgression

const (
	SavingThrowPoor SavingThrowProgression = "1/3"
	SavingThrowGood SavingThrowProgression = "1/2"
)

type savingThrow struct {
	id                   savingThrowID
	actualValue          rationalValue
	value                int
	goodBaseBonusApplied bool
}
type SavingThrow = savingThrow

func NewSavingThrow(id SavingThrowID, actualValue RationalValue) (SavingThrow, bool) {
	save := savingThrow{}
	if !save.SetID(id) || !save.SetActualValue(actualValue) {
		return savingThrow{}, false
	}

	return save, true
}

func NewSavingThrowByClassLevel(
	id SavingThrowID,
	level int,
	progression SavingThrowProgression,
) (SavingThrow, bool) {
	save := savingThrow{}
	if !save.SetID(id) || !save.SetByClassLevel(level, progression) {
		return savingThrow{}, false
	}

	return save, true
}

func (s savingThrow) GetID() SavingThrowID {
	return s.id
}

func (s savingThrow) GetActualValue() RationalValue {
	return s.actualValue
}

func (s savingThrow) GetValue() int {
	return s.value
}

func (s savingThrow) HasGoodBaseBonusApplied() bool {
	return s.goodBaseBonusApplied
}

func (s *savingThrow) SetID(id SavingThrowID) bool {
	if !isValidSavingThrowID(id) {
		return false
	}

	s.id = id
	return true
}

func (s *savingThrow) SetActualValue(actualValue RationalValue) bool {
	if !isNonNegativeRationalValue(actualValue) {
		return false
	}

	s.actualValue = actualValue
	s.value = actualValue.Floor()
	s.goodBaseBonusApplied = false
	return true
}

func (s *savingThrow) SetByClassLevel(level int, progression SavingThrowProgression) bool {
	if level < 0 {
		return false
	}

	if _, ok := progression.toRationalValue(); !ok {
		return false
	}

	s.actualValue = zeroRationalValue()
	s.value = 0
	s.goodBaseBonusApplied = false

	return s.AddClassLevel(level, progression)
}

func (s *savingThrow) AddClassLevel(level int, progression SavingThrowProgression) bool {
	if level < 0 {
		return false
	}

	actualIncrement, ok := progression.toRationalValue()
	if !ok {
		return false
	}

	actualIncrement = actualIncrement.MultiplyByInt(level)
	if progression == SavingThrowGood && !s.goodBaseBonusApplied {
		goodBaseBonus, _ := NewRationalValue(goodSavingThrowBaseBonus, 1)
		actualIncrement = actualIncrement.Add(goodBaseBonus)
		s.goodBaseBonusApplied = true
	}

	s.actualValue = s.actualValue.Add(actualIncrement)
	s.value = s.actualValue.Floor()
	return true
}

func isValidSavingThrowID(id SavingThrowID) bool {
	switch id {
	case FortitudeSave, ReflexSave, WillSave:
		return true
	default:
		return false
	}
}

func (p savingThrowProgression) toRationalValue() (RationalValue, bool) {
	switch p {
	case SavingThrowPoor:
		return NewRationalValue(1, 3)
	case SavingThrowGood:
		return NewRationalValue(1, 2)
	default:
		return rationalValue{}, false
	}
}
