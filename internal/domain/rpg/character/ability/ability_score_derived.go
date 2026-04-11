package ability

type weightRange struct {
	min weightValue
	max weightValue
}

type weightValue struct {
	grams int
}

type strengthCarryingCapacity struct {
	lightLoadMax weightValue
	mediumLoad   weightRange
	heavyLoad    weightRange
}

type spellcastingAbilityProfile struct {
	abilityModifier int
	maxSpellLevel   int
}

func (w weightValue) GetKilograms() float64 {
	return float64(w.grams) / 1000
}

func (w weightValue) GetPounds() float64 {
	return float64(w.grams) / gramsPerPound
}

func (r weightRange) GetMin() weightValue {
	return r.min
}

func (r weightRange) GetMax() weightValue {
	return r.max
}

func (c strengthCarryingCapacity) GetLightLoadMax() weightValue {
	return c.lightLoadMax
}

func (c strengthCarryingCapacity) GetMediumLoad() weightRange {
	return c.mediumLoad
}

func (c strengthCarryingCapacity) GetHeavyLoad() weightRange {
	return c.heavyLoad
}

func (p spellcastingAbilityProfile) GetMaxSpellLevel() int {
	return p.maxSpellLevel
}

func (p spellcastingAbilityProfile) GetBonusSpells(spellLevel int) int {
	if spellLevel < 1 {
		return 0
	}

	if p.abilityModifier < spellLevel {
		return 0
	}

	return 1 + ((p.abilityModifier - spellLevel) / 4)
}

func (a abilityScore) GetCarryingCapacity() (strengthCarryingCapacity, bool) {
	if a.id != StrengthScore {
		return strengthCarryingCapacity{}, false
	}

	score, valid := a.value.GetValue()
	if !valid {
		return strengthCarryingCapacity{}, false
	}

	return resolveStrengthCarryingCapacity(score), true
}

func (a abilityScore) GetSpellcastingProfile() (spellcastingAbilityProfile, bool) {
	score, valid := a.value.GetValue()
	if !valid {
		return spellcastingAbilityProfile{}, false
	}

	maxSpellLevel := 0
	if score > 10 {
		maxSpellLevel = score - 10
	}

	return spellcastingAbilityProfile{
		abilityModifier: calculateAbilityModifier(score),
		maxSpellLevel:   maxSpellLevel,
	}, true
}

func resolveStrengthCarryingCapacity(score int) strengthCarryingCapacity {
	if score <= 0 {
		return strengthCarryingCapacity{}
	}

	multiplier := 1
	for score > 29 {
		score -= 10
		multiplier *= 4
	}

	base := strengthCarryingCapacityTable[score]
	return base.multiply(multiplier)
}

func (c strengthCarryingCapacity) multiply(multiplier int) strengthCarryingCapacity {
	return strengthCarryingCapacity{
		lightLoadMax: c.lightLoadMax.multiply(multiplier),
		mediumLoad:   c.mediumLoad.multiply(multiplier),
		heavyLoad:    c.heavyLoad.multiply(multiplier),
	}
}

func (r weightRange) multiply(multiplier int) weightRange {
	return weightRange{
		min: r.min.multiply(multiplier),
		max: r.max.multiply(multiplier),
	}
}

func (w weightValue) multiply(multiplier int) weightValue {
	return weightValue{
		grams: w.grams * multiplier,
	}
}

const gramsPerPound = 453.59237

var strengthCarryingCapacityTable = map[int]strengthCarryingCapacity{
	1: {
		lightLoadMax: weightValue{grams: 1500},
		mediumLoad:   weightRange{min: weightValue{grams: 2000}, max: weightValue{grams: 3000}},
		heavyLoad:    weightRange{min: weightValue{grams: 3500}, max: weightValue{grams: 5000}},
	},
	2: {
		lightLoadMax: weightValue{grams: 3000},
		mediumLoad:   weightRange{min: weightValue{grams: 3500}, max: weightValue{grams: 6500}},
		heavyLoad:    weightRange{min: weightValue{grams: 7000}, max: weightValue{grams: 10000}},
	},
	3: {
		lightLoadMax: weightValue{grams: 5000},
		mediumLoad:   weightRange{min: weightValue{grams: 5500}, max: weightValue{grams: 10000}},
		heavyLoad:    weightRange{min: weightValue{grams: 10500}, max: weightValue{grams: 15000}},
	},
	4: {
		lightLoadMax: weightValue{grams: 6500},
		mediumLoad:   weightRange{min: weightValue{grams: 7000}, max: weightValue{grams: 13000}},
		heavyLoad:    weightRange{min: weightValue{grams: 13500}, max: weightValue{grams: 20000}},
	},
	5: {
		lightLoadMax: weightValue{grams: 8000},
		mediumLoad:   weightRange{min: weightValue{grams: 8500}, max: weightValue{grams: 16000}},
		heavyLoad:    weightRange{min: weightValue{grams: 17000}, max: weightValue{grams: 25000}},
	},
	6: {
		lightLoadMax: weightValue{grams: 10000},
		mediumLoad:   weightRange{min: weightValue{grams: 10500}, max: weightValue{grams: 20000}},
		heavyLoad:    weightRange{min: weightValue{grams: 20500}, max: weightValue{grams: 30000}},
	},
	7: {
		lightLoadMax: weightValue{grams: 11500},
		mediumLoad:   weightRange{min: weightValue{grams: 12000}, max: weightValue{grams: 23000}},
		heavyLoad:    weightRange{min: weightValue{grams: 23500}, max: weightValue{grams: 35000}},
	},
	8: {
		lightLoadMax: weightValue{grams: 13000},
		mediumLoad:   weightRange{min: weightValue{grams: 13500}, max: weightValue{grams: 26500}},
		heavyLoad:    weightRange{min: weightValue{grams: 27000}, max: weightValue{grams: 40000}},
	},
	9: {
		lightLoadMax: weightValue{grams: 15000},
		mediumLoad:   weightRange{min: weightValue{grams: 15500}, max: weightValue{grams: 30000}},
		heavyLoad:    weightRange{min: weightValue{grams: 30500}, max: weightValue{grams: 45000}},
	},
	10: {
		lightLoadMax: weightValue{grams: 16500},
		mediumLoad:   weightRange{min: weightValue{grams: 17000}, max: weightValue{grams: 33000}},
		heavyLoad:    weightRange{min: weightValue{grams: 33500}, max: weightValue{grams: 50000}},
	},
	11: {
		lightLoadMax: weightValue{grams: 19000},
		mediumLoad:   weightRange{min: weightValue{grams: 19500}, max: weightValue{grams: 38000}},
		heavyLoad:    weightRange{min: weightValue{grams: 38500}, max: weightValue{grams: 57500}},
	},
	12: {
		lightLoadMax: weightValue{grams: 21500},
		mediumLoad:   weightRange{min: weightValue{grams: 22000}, max: weightValue{grams: 43000}},
		heavyLoad:    weightRange{min: weightValue{grams: 43500}, max: weightValue{grams: 65000}},
	},
	13: {
		lightLoadMax: weightValue{grams: 25000},
		mediumLoad:   weightRange{min: weightValue{grams: 25500}, max: weightValue{grams: 50000}},
		heavyLoad:    weightRange{min: weightValue{grams: 50500}, max: weightValue{grams: 75000}},
	},
	14: {
		lightLoadMax: weightValue{grams: 29000},
		mediumLoad:   weightRange{min: weightValue{grams: 29500}, max: weightValue{grams: 58000}},
		heavyLoad:    weightRange{min: weightValue{grams: 58500}, max: weightValue{grams: 87500}},
	},
	15: {
		lightLoadMax: weightValue{grams: 33000},
		mediumLoad:   weightRange{min: weightValue{grams: 33500}, max: weightValue{grams: 66500}},
		heavyLoad:    weightRange{min: weightValue{grams: 67000}, max: weightValue{grams: 100000}},
	},
	16: {
		lightLoadMax: weightValue{grams: 38000},
		mediumLoad:   weightRange{min: weightValue{grams: 38500}, max: weightValue{grams: 76500}},
		heavyLoad:    weightRange{min: weightValue{grams: 77000}, max: weightValue{grams: 115000}},
	},
	17: {
		lightLoadMax: weightValue{grams: 43000},
		mediumLoad:   weightRange{min: weightValue{grams: 43500}, max: weightValue{grams: 86500}},
		heavyLoad:    weightRange{min: weightValue{grams: 87000}, max: weightValue{grams: 130000}},
	},
	18: {
		lightLoadMax: weightValue{grams: 50000},
		mediumLoad:   weightRange{min: weightValue{grams: 50500}, max: weightValue{grams: 100000}},
		heavyLoad:    weightRange{min: weightValue{grams: 100500}, max: weightValue{grams: 150000}},
	},
	19: {
		lightLoadMax: weightValue{grams: 58000},
		mediumLoad:   weightRange{min: weightValue{grams: 58500}, max: weightValue{grams: 116500}},
		heavyLoad:    weightRange{min: weightValue{grams: 117000}, max: weightValue{grams: 175000}},
	},
	20: {
		lightLoadMax: weightValue{grams: 66500},
		mediumLoad:   weightRange{min: weightValue{grams: 67000}, max: weightValue{grams: 133000}},
		heavyLoad:    weightRange{min: weightValue{grams: 133500}, max: weightValue{grams: 200000}},
	},
	21: {
		lightLoadMax: weightValue{grams: 76500},
		mediumLoad:   weightRange{min: weightValue{grams: 77000}, max: weightValue{grams: 153000}},
		heavyLoad:    weightRange{min: weightValue{grams: 153500}, max: weightValue{grams: 230000}},
	},
	22: {
		lightLoadMax: weightValue{grams: 86500},
		mediumLoad:   weightRange{min: weightValue{grams: 87000}, max: weightValue{grams: 173000}},
		heavyLoad:    weightRange{min: weightValue{grams: 173500}, max: weightValue{grams: 260000}},
	},
	23: {
		lightLoadMax: weightValue{grams: 100000},
		mediumLoad:   weightRange{min: weightValue{grams: 100500}, max: weightValue{grams: 200000}},
		heavyLoad:    weightRange{min: weightValue{grams: 200500}, max: weightValue{grams: 300000}},
	},
	24: {
		lightLoadMax: weightValue{grams: 116500},
		mediumLoad:   weightRange{min: weightValue{grams: 117000}, max: weightValue{grams: 233000}},
		heavyLoad:    weightRange{min: weightValue{grams: 233500}, max: weightValue{grams: 350000}},
	},
	25: {
		lightLoadMax: weightValue{grams: 133000},
		mediumLoad:   weightRange{min: weightValue{grams: 133500}, max: weightValue{grams: 266500}},
		heavyLoad:    weightRange{min: weightValue{grams: 267000}, max: weightValue{grams: 400000}},
	},
	26: {
		lightLoadMax: weightValue{grams: 153000},
		mediumLoad:   weightRange{min: weightValue{grams: 153500}, max: weightValue{grams: 306500}},
		heavyLoad:    weightRange{min: weightValue{grams: 307000}, max: weightValue{grams: 460000}},
	},
	27: {
		lightLoadMax: weightValue{grams: 173000},
		mediumLoad:   weightRange{min: weightValue{grams: 173500}, max: weightValue{grams: 346500}},
		heavyLoad:    weightRange{min: weightValue{grams: 347000}, max: weightValue{grams: 520000}},
	},
	28: {
		lightLoadMax: weightValue{grams: 200000},
		mediumLoad:   weightRange{min: weightValue{grams: 205000}, max: weightValue{grams: 400000}},
		heavyLoad:    weightRange{min: weightValue{grams: 400500}, max: weightValue{grams: 600000}},
	},
	29: {
		lightLoadMax: weightValue{grams: 233000},
		mediumLoad:   weightRange{min: weightValue{grams: 238500}, max: weightValue{grams: 466500}},
		heavyLoad:    weightRange{min: weightValue{grams: 467000}, max: weightValue{grams: 700000}},
	},
}
