package ability

type weightRange struct {
	min weightValue
	max weightValue
}
type WeightRange = weightRange

type weightValue struct {
	grams int
}
type WeightValue = weightValue

type strengthCarryingCapacity struct {
	lightLoadMax weightValue
	mediumLoad   weightRange
	heavyLoad    weightRange
}
type StrengthCarryingCapacity = strengthCarryingCapacity

type spellcastingAbilityProfile struct {
	abilityModifier int
	maxSpellLevel   int
}
type SpellcastingAbilityProfile = spellcastingAbilityProfile

func (w weightValue) GetKilograms() float64 {
	return float64(w.grams) / 1000
}

func (w weightValue) GetPounds() float64 {
	return float64(w.grams) / gramsPerPound
}

func (r weightRange) GetMin() WeightValue {
	return r.min
}

func (r weightRange) GetMax() WeightValue {
	return r.max
}

func (c strengthCarryingCapacity) GetLightLoadMax() WeightValue {
	return c.lightLoadMax
}

func (c strengthCarryingCapacity) GetMediumLoad() WeightRange {
	return c.mediumLoad
}

func (c strengthCarryingCapacity) GetHeavyLoad() WeightRange {
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

func (a abilityScore) GetCarryingCapacity() (StrengthCarryingCapacity, bool) {
	if a.id != StrengthScore {
		return strengthCarryingCapacity{}, false
	}

	score, valid := a.value.GetValue()
	if !valid {
		return strengthCarryingCapacity{}, false
	}

	return resolveStrengthCarryingCapacity(score), true
}

func (a abilityScore) GetSpellcastingProfile() (SpellcastingAbilityProfile, bool) {
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

func carryingCapacityRange(minPounds float64, maxPounds float64) weightRange {
	return weightRange{
		min: pounds(minPounds),
		max: pounds(maxPounds),
	}
}

var strengthCarryingCapacityTable = map[int]strengthCarryingCapacity{
	1: {
		lightLoadMax: pounds(3),
		mediumLoad:   carryingCapacityRange(4, 6),
		heavyLoad:    carryingCapacityRange(7, 10),
	},
	2: {
		lightLoadMax: pounds(6),
		mediumLoad:   carryingCapacityRange(7, 13),
		heavyLoad:    carryingCapacityRange(14, 20),
	},
	3: {
		lightLoadMax: pounds(10),
		mediumLoad:   carryingCapacityRange(11, 20),
		heavyLoad:    carryingCapacityRange(21, 30),
	},
	4: {
		lightLoadMax: pounds(13),
		mediumLoad:   carryingCapacityRange(14, 26),
		heavyLoad:    carryingCapacityRange(27, 40),
	},
	5: {
		lightLoadMax: pounds(16),
		mediumLoad:   carryingCapacityRange(17, 33),
		heavyLoad:    carryingCapacityRange(34, 50),
	},
	6: {
		lightLoadMax: pounds(20),
		mediumLoad:   carryingCapacityRange(21, 40),
		heavyLoad:    carryingCapacityRange(41, 60),
	},
	7: {
		lightLoadMax: pounds(23),
		mediumLoad:   carryingCapacityRange(24, 46),
		heavyLoad:    carryingCapacityRange(47, 70),
	},
	8: {
		lightLoadMax: pounds(26),
		mediumLoad:   carryingCapacityRange(27, 53),
		heavyLoad:    carryingCapacityRange(54, 80),
	},
	9: {
		lightLoadMax: pounds(30),
		mediumLoad:   carryingCapacityRange(31, 60),
		heavyLoad:    carryingCapacityRange(61, 90),
	},
	10: {
		lightLoadMax: pounds(33),
		mediumLoad:   carryingCapacityRange(34, 66),
		heavyLoad:    carryingCapacityRange(67, 100),
	},
	11: {
		lightLoadMax: pounds(38),
		mediumLoad:   carryingCapacityRange(39, 76),
		heavyLoad:    carryingCapacityRange(77, 115),
	},
	12: {
		lightLoadMax: pounds(43),
		mediumLoad:   carryingCapacityRange(44, 86),
		heavyLoad:    carryingCapacityRange(87, 130),
	},
	13: {
		lightLoadMax: pounds(50),
		mediumLoad:   carryingCapacityRange(51, 100),
		heavyLoad:    carryingCapacityRange(101, 150),
	},
	14: {
		lightLoadMax: pounds(58),
		mediumLoad:   carryingCapacityRange(59, 116),
		heavyLoad:    carryingCapacityRange(117, 175),
	},
	15: {
		lightLoadMax: pounds(66),
		mediumLoad:   carryingCapacityRange(67, 133),
		heavyLoad:    carryingCapacityRange(134, 200),
	},
	16: {
		lightLoadMax: pounds(76),
		mediumLoad:   carryingCapacityRange(77, 153),
		heavyLoad:    carryingCapacityRange(154, 230),
	},
	17: {
		lightLoadMax: pounds(86),
		mediumLoad:   carryingCapacityRange(87, 173),
		heavyLoad:    carryingCapacityRange(174, 260),
	},
	18: {
		lightLoadMax: pounds(100),
		mediumLoad:   carryingCapacityRange(101, 200),
		heavyLoad:    carryingCapacityRange(201, 300),
	},
	19: {
		lightLoadMax: pounds(116),
		mediumLoad:   carryingCapacityRange(117, 233),
		heavyLoad:    carryingCapacityRange(234, 350),
	},
	20: {
		lightLoadMax: pounds(133),
		mediumLoad:   carryingCapacityRange(134, 266),
		heavyLoad:    carryingCapacityRange(267, 400),
	},
	21: {
		lightLoadMax: pounds(153),
		mediumLoad:   carryingCapacityRange(154, 306),
		heavyLoad:    carryingCapacityRange(307, 460),
	},
	22: {
		lightLoadMax: pounds(173),
		mediumLoad:   carryingCapacityRange(174, 346),
		heavyLoad:    carryingCapacityRange(347, 520),
	},
	23: {
		lightLoadMax: pounds(200),
		mediumLoad:   carryingCapacityRange(201, 400),
		heavyLoad:    carryingCapacityRange(401, 600),
	},
	24: {
		lightLoadMax: pounds(233),
		mediumLoad:   carryingCapacityRange(234, 466),
		heavyLoad:    carryingCapacityRange(467, 700),
	},
	25: {
		lightLoadMax: pounds(266),
		mediumLoad:   carryingCapacityRange(267, 533),
		heavyLoad:    carryingCapacityRange(534, 800),
	},
	26: {
		lightLoadMax: pounds(306),
		mediumLoad:   carryingCapacityRange(307, 613),
		heavyLoad:    carryingCapacityRange(614, 920),
	},
	27: {
		lightLoadMax: pounds(346),
		mediumLoad:   carryingCapacityRange(347, 693),
		heavyLoad:    carryingCapacityRange(694, 1040),
	},
	28: {
		lightLoadMax: pounds(400),
		mediumLoad:   carryingCapacityRange(401, 800),
		heavyLoad:    carryingCapacityRange(801, 1200),
	},
	29: {
		lightLoadMax: pounds(466),
		mediumLoad:   carryingCapacityRange(467, 933),
		heavyLoad:    carryingCapacityRange(934, 1400),
	},
}
