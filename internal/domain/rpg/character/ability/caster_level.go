package ability

type casterSource string
type CasterSource = casterSource

const (
	ArcaneCasterSource CasterSource = "Arcane"
	DivineCasterSource CasterSource = "Divine"
	PrimalCasterSource CasterSource = "Primal"
)

type nullableInt struct {
	value int
	valid bool
}

type casterLevel struct {
	arcane nullableInt
	divine nullableInt
	primal nullableInt
}
type CasterLevel = casterLevel

func NewCasterLevel(arcane int, divine int, primal int) (CasterLevel, bool) {
	level := casterLevel{}
	if !level.SetSourceLevel(ArcaneCasterSource, arcane) ||
		!level.SetSourceLevel(DivineCasterSource, divine) ||
		!level.SetSourceLevel(PrimalCasterSource, primal) {
		return casterLevel{}, false
	}

	return level, true
}

func NewImpossibleCasterLevel() CasterLevel {
	return casterLevel{}
}

func (c casterLevel) GetSourceLevel(source CasterSource) (int, bool) {
	value, ok := c.getSourceValue(source)
	if !ok {
		return 0, false
	}

	return value.Get()
}

func (c *casterLevel) SetCasterLevel(level CasterLevel) bool {
	if !isValidCasterLevel(level) {
		return false
	}

	c.arcane = level.arcane
	c.divine = level.divine
	c.primal = level.primal
	return true
}

func (c *casterLevel) AddCasterLevel(level CasterLevel) bool {
	if !isValidCasterLevel(level) {
		return false
	}

	if level.arcane.valid && !c.AddSourceLevel(ArcaneCasterSource, level.arcane.value) {
		return false
	}

	if level.divine.valid && !c.AddSourceLevel(DivineCasterSource, level.divine.value) {
		return false
	}

	if level.primal.valid && !c.AddSourceLevel(PrimalCasterSource, level.primal.value) {
		return false
	}

	return true
}

func (c *casterLevel) SetSourceLevel(source CasterSource, value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	sourceValue, ok := c.getSourceValuePointer(source)
	if !ok {
		return false
	}

	sourceValue.Set(value)
	return true
}

func (c *casterLevel) AddSourceLevel(source CasterSource, value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	sourceValue, ok := c.getSourceValuePointer(source)
	if !ok {
		return false
	}

	sourceValue.Add(value)
	return true
}

func (c *casterLevel) DisableSourceLevel(source CasterSource) bool {
	sourceValue, ok := c.getSourceValuePointer(source)
	if !ok {
		return false
	}

	sourceValue.Disable()
	return true
}

func (c casterLevel) getSourceValue(source CasterSource) (nullableInt, bool) {
	switch source {
	case ArcaneCasterSource:
		return c.arcane, true
	case DivineCasterSource:
		return c.divine, true
	case PrimalCasterSource:
		return c.primal, true
	default:
		return nullableInt{}, false
	}
}

func (c *casterLevel) getSourceValuePointer(source CasterSource) (*nullableInt, bool) {
	switch source {
	case ArcaneCasterSource:
		return &c.arcane, true
	case DivineCasterSource:
		return &c.divine, true
	case PrimalCasterSource:
		return &c.primal, true
	default:
		return nil, false
	}
}

func (n nullableInt) Get() (int, bool) {
	return n.value, n.valid
}

func (n *nullableInt) Set(value int) {
	n.value = value
	n.valid = true
}

func (n *nullableInt) Add(value int) {
	if n.valid {
		n.value += value
		return
	}

	n.Set(value)
}

func (n *nullableInt) Disable() {
	n.value = 0
	n.valid = false
}

func isValidCasterLevel(level CasterLevel) bool {
	return (!level.arcane.valid || isValidCasterLevelValue(level.arcane.value)) &&
		(!level.divine.valid || isValidCasterLevelValue(level.divine.value)) &&
		(!level.primal.valid || isValidCasterLevelValue(level.primal.value))
}

func isValidCasterLevelValue(value int) bool {
	return value >= 0
}
