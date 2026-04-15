package ability

type casterPillar string
type CasterPillar = casterPillar

const (
	ArcaneCasterPillar CasterPillar = "Arcane"
	DivineCasterPillar CasterPillar = "Divine"
	PrimalCasterPillar CasterPillar = "Primal"
)

type casterLevel struct {
	arcane      int
	arcaneValid bool
	divine      int
	divineValid bool
	primal      int
	primalValid bool
}
type CasterLevel = casterLevel

func NewCasterLevel(arcane int, divine int, primal int) CasterLevel {
	return casterLevel{
		arcane:      arcane,
		arcaneValid: true,
		divine:      divine,
		divineValid: true,
		primal:      primal,
		primalValid: true,
	}
}

func NewImpossibleCasterLevel() CasterLevel {
	return casterLevel{}
}

func (c casterLevel) GetArcane() (int, bool) {
	return c.arcane, c.arcaneValid
}

func (c casterLevel) GetDivine() (int, bool) {
	return c.divine, c.divineValid
}

func (c casterLevel) GetPrimal() (int, bool) {
	return c.primal, c.primalValid
}

func (c casterLevel) GetCasterLevel() CasterLevel {
	return c
}

func (c *casterLevel) SetCasterLevel(cl CasterLevel) bool {
	if !isValidCasterLevel(cl) {
		return false
	}

	c.arcane = cl.arcane
	c.arcaneValid = cl.arcaneValid
	c.divine = cl.divine
	c.divineValid = cl.divineValid
	c.primal = cl.primal
	c.primalValid = cl.primalValid
	return true
}

func (c *casterLevel) SetArcaneCasterLevel(value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	c.arcane = value
	c.arcaneValid = true
	return true
}

func (c *casterLevel) SetDivineCasterLevel(value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	c.divine = value
	c.divineValid = true
	return true
}

func (c *casterLevel) SetPrimalCasterLevel(value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	c.primal = value
	c.primalValid = true
	return true
}

func (c *casterLevel) DisableArcaneCasterLevel() {
	c.arcane = 0
	c.arcaneValid = false
}

func (c *casterLevel) DisableDivineCasterLevel() {
	c.divine = 0
	c.divineValid = false
}

func (c *casterLevel) DisablePrimalCasterLevel() {
	c.primal = 0
	c.primalValid = false
}

func (c *casterLevel) AddCasterLevel(cl CasterLevel) bool {
	if !isValidCasterLevel(cl) {
		return false
	}

	if cl.arcaneValid {
		if c.arcaneValid {
			c.arcane += cl.arcane
		} else {
			c.arcane = cl.arcane
			c.arcaneValid = true
		}
	}

	if cl.divineValid {
		if c.divineValid {
			c.divine += cl.divine
		} else {
			c.divine = cl.divine
			c.divineValid = true
		}
	}

	if cl.primalValid {
		if c.primalValid {
			c.primal += cl.primal
		} else {
			c.primal = cl.primal
			c.primalValid = true
		}
	}

	return true
}

func (c *casterLevel) AddArcaneCasterLevel(value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	if c.arcaneValid {
		c.arcane += value
	} else {
		c.arcane = value
		c.arcaneValid = true
	}

	return true
}

func (c *casterLevel) AddDivineCasterLevel(value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	if c.divineValid {
		c.divine += value
	} else {
		c.divine = value
		c.divineValid = true
	}

	return true
}

func (c *casterLevel) AddPrimalCasterLevel(value int) bool {
	if !isValidCasterLevelValue(value) {
		return false
	}

	if c.primalValid {
		c.primal += value
	} else {
		c.primal = value
		c.primalValid = true
	}

	return true
}

func isValidCasterLevel(cl CasterLevel) bool {
	return (!cl.arcaneValid || isValidCasterLevelValue(cl.arcane)) &&
		(!cl.divineValid || isValidCasterLevelValue(cl.divine)) &&
		(!cl.primalValid || isValidCasterLevelValue(cl.primal))
}

func isValidCasterLevelValue(value int) bool {
	return value >= 0
}
