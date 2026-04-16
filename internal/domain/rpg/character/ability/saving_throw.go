package ability

const goodSavingThrowBaseBonus = 2

type savingThrowID string
type SavingThrowID = savingThrowID

const (
	FortitudeSave SavingThrowID = "Fortitude"
	ReflexSave    SavingThrowID = "Reflex"
	WillSave      SavingThrowID = "Will"
)

type savingThrowProgression float64
type SavingThrowProgression = savingThrowProgression

const (
	SavingThrowPoor SavingThrowProgression = 1.0 / 3.0
	SavingThrowGood SavingThrowProgression = 0.5
)

type savingThrow struct {
	id                   savingThrowID
	actualValue          float64
	value                int
	goodBaseBonusApplied bool
}
type SavingThrow = savingThrow

func NewSavingThrow(id SavingThrowID, actualValue float64) (SavingThrow, bool) {
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

func (s savingThrow) GetActualValue() float64 {
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

func (s *savingThrow) SetActualValue(actualValue float64) bool {
	if actualValue < 0 {
		return false
	}

	s.actualValue = actualValue
	s.value = roundDown(actualValue)
	s.goodBaseBonusApplied = false
	return true
}

func (s *savingThrow) SetByClassLevel(level int, progression SavingThrowProgression) bool {
	s.actualValue = 0
	s.value = 0
	s.goodBaseBonusApplied = false

	return s.AddClassLevel(level, progression)
}

func (s *savingThrow) AddClassLevel(level int, progression SavingThrowProgression) bool {
	if level < 0 || !isValidSavingThrowProgression(progression) {
		return false
	}

	actualValue := float64(level) * float64(progression)
	if progression == SavingThrowGood && !s.goodBaseBonusApplied {
		actualValue += goodSavingThrowBaseBonus
		s.goodBaseBonusApplied = true
	}

	s.actualValue += actualValue
	s.value = roundDown(s.actualValue)
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

func isValidSavingThrowProgression(progression SavingThrowProgression) bool {
	switch progression {
	case SavingThrowPoor, SavingThrowGood:
		return true
	default:
		return false
	}
}
