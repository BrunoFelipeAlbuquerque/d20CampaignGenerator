package ability

const metersPerFoot = 0.3048

type size string
type Size = size

type bodyShape string
type BodyShape = bodyShape

const (
	TallBodyShape BodyShape = "Tall"
	LongBodyShape BodyShape = "Long"
)

const (
	FineSize       Size = "Fine"
	DiminutiveSize Size = "Diminutive"
	TinySize       Size = "Tiny"
	SmallSize      Size = "Small"
	MediumSize     Size = "Medium"
	LargeSize      Size = "Large"
	HugeSize       Size = "Huge"
	GargantuanSize Size = "Gargantuan"
	ColossalSize   Size = "Colossal"
	TitanicSize    Size = "Titanic"
)

type lengthValue struct {
	feet float64
}
type LengthValue = lengthValue

type lengthRange struct {
	min           lengthValue
	max           lengthValue
	hasUpperBound bool
}
type LengthRange = lengthRange

type sizeWeightRange struct {
	min           weightValue
	max           weightValue
	hasUpperBound bool
}
type SizeWeightRange = sizeWeightRange

type sizeProfile struct {
	modifier         int
	specialModifier  int
	flyModifier      int
	stealthModifier  int
	constructBonusHP int
	spaceTall        lengthValue
	spaceLong        lengthValue
	reachTall        lengthValue
	reachLong        lengthValue
	heightRange      lengthRange
	weightRange      sizeWeightRange
}

func (l lengthValue) GetFeet() float64 {
	return l.feet
}

func (l lengthValue) GetMeters() float64 {
	return l.feet * metersPerFoot
}

func (r lengthRange) GetMin() LengthValue {
	return r.min
}

func (r lengthRange) GetMax() LengthValue {
	return r.max
}

func (r lengthRange) HasUpperBound() bool {
	return r.hasUpperBound
}

func (r sizeWeightRange) GetMin() WeightValue {
	return r.min
}

func (r sizeWeightRange) GetMax() WeightValue {
	return r.max
}

func (r sizeWeightRange) HasUpperBound() bool {
	return r.hasUpperBound
}

func (s size) GetAttackAndACModifier() (int, bool) {
	profile, ok := getSizeProfile(s)
	if !ok {
		return 0, false
	}

	return profile.modifier, true
}

func (s size) GetCMBAndCMDModifier() (int, bool) {
	profile, ok := getSizeProfile(s)
	if !ok {
		return 0, false
	}

	return profile.specialModifier, true
}

func (s size) GetModifier() (int, bool) {
	return s.GetAttackAndACModifier()
}

func (s size) GetSpecialModifier() (int, bool) {
	return s.GetCMBAndCMDModifier()
}

func (s size) GetFlyModifier() (int, bool) {
	profile, ok := getSizeProfile(s)
	if !ok {
		return 0, false
	}

	return profile.flyModifier, true
}

func (s size) GetStealthModifier() (int, bool) {
	profile, ok := getSizeProfile(s)
	if !ok {
		return 0, false
	}

	return profile.stealthModifier, true
}

func (s size) GetConstructBonusHP() (int, bool) {
	profile, ok := getSizeProfile(s)
	if !ok {
		return 0, false
	}

	return profile.constructBonusHP, true
}

func (s size) GetSpace(shape BodyShape) (LengthValue, bool) {
	profile, ok := getSizeProfile(s)
	if !ok || !isValidBodyShape(shape) {
		return lengthValue{}, false
	}

	if shape == LongBodyShape {
		return profile.spaceLong, true
	}

	return profile.spaceTall, true
}

func (s size) GetNaturalReach(shape BodyShape) (LengthValue, bool) {
	profile, ok := getSizeProfile(s)
	if !ok || !isValidBodyShape(shape) {
		return lengthValue{}, false
	}

	if shape == LongBodyShape {
		return profile.reachLong, true
	}

	return profile.reachTall, true
}

func (s size) GetTypicalHeightRange() (LengthRange, bool) {
	profile, ok := getSizeProfile(s)
	if !ok {
		return lengthRange{}, false
	}

	return profile.heightRange, true
}

func (s size) GetTypicalWeightRange() (SizeWeightRange, bool) {
	profile, ok := getSizeProfile(s)
	if !ok {
		return sizeWeightRange{}, false
	}

	return profile.weightRange, true
}

func isValidSize(value Size) bool {
	_, ok := getSizeProfile(value)
	return ok
}

func isValidBodyShape(value BodyShape) bool {
	switch value {
	case TallBodyShape, LongBodyShape:
		return true
	default:
		return false
	}
}

func getSizeProfile(value Size) (sizeProfile, bool) {
	profile, ok := sizeProfiles[value]
	return profile, ok
}

func feet(value float64) lengthValue {
	return lengthValue{feet: value}
}

func heightRange(minFeet float64, maxFeet float64, hasUpperBound bool) lengthRange {
	return lengthRange{
		min:           feet(minFeet),
		max:           feet(maxFeet),
		hasUpperBound: hasUpperBound,
	}
}

func pounds(value float64) weightValue {
	return weightValue{grams: int(value * gramsPerPound)}
}

func tons(value float64) weightValue {
	return pounds(value * 2000)
}

func weightRangePounds(minPounds float64, maxPounds float64, hasUpperBound bool) sizeWeightRange {
	return sizeWeightRange{
		min:           pounds(minPounds),
		max:           pounds(maxPounds),
		hasUpperBound: hasUpperBound,
	}
}

func weightRangeTons(minTons float64, maxTons float64, hasUpperBound bool) sizeWeightRange {
	return sizeWeightRange{
		min:           tons(minTons),
		max:           tons(maxTons),
		hasUpperBound: hasUpperBound,
	}
}

var sizeProfiles = map[Size]sizeProfile{
	FineSize: {
		modifier:         8,
		specialModifier:  -8,
		flyModifier:      8,
		stealthModifier:  16,
		constructBonusHP: 0,
		spaceTall:        feet(0.5),
		spaceLong:        feet(0.5),
		reachTall:        feet(0),
		reachLong:        feet(0),
		heightRange:      lengthRange{max: feet(0.5), hasUpperBound: true},
		weightRange:      weightRangePounds(0, 0.125, true),
	},
	DiminutiveSize: {
		modifier:         4,
		specialModifier:  -4,
		flyModifier:      6,
		stealthModifier:  12,
		constructBonusHP: 0,
		spaceTall:        feet(1),
		spaceLong:        feet(1),
		reachTall:        feet(0),
		reachLong:        feet(0),
		heightRange:      lengthRange{min: feet(0.5), max: feet(1), hasUpperBound: true},
		weightRange:      weightRangePounds(0.125, 1, true),
	},
	TinySize: {
		modifier:         2,
		specialModifier:  -2,
		flyModifier:      4,
		stealthModifier:  8,
		constructBonusHP: 5,
		spaceTall:        feet(2.5),
		spaceLong:        feet(2.5),
		reachTall:        feet(0),
		reachLong:        feet(0),
		heightRange:      heightRange(1, 2, true),
		weightRange:      weightRangePounds(1, 8, true),
	},
	SmallSize: {
		modifier:         1,
		specialModifier:  -1,
		flyModifier:      2,
		stealthModifier:  4,
		constructBonusHP: 10,
		spaceTall:        feet(5),
		spaceLong:        feet(5),
		reachTall:        feet(5),
		reachLong:        feet(5),
		heightRange:      heightRange(2, 4, true),
		weightRange:      weightRangePounds(8, 60, true),
	},
	MediumSize: {
		modifier:         0,
		specialModifier:  0,
		flyModifier:      0,
		stealthModifier:  0,
		constructBonusHP: 20,
		spaceTall:        feet(5),
		spaceLong:        feet(5),
		reachTall:        feet(5),
		reachLong:        feet(5),
		heightRange:      heightRange(4, 8, true),
		weightRange:      weightRangePounds(60, 500, true),
	},
	LargeSize: {
		modifier:         -1,
		specialModifier:  1,
		flyModifier:      -2,
		stealthModifier:  -4,
		constructBonusHP: 30,
		spaceTall:        feet(10),
		spaceLong:        feet(10),
		reachTall:        feet(10),
		reachLong:        feet(5),
		heightRange:      heightRange(8, 16, true),
		weightRange:      weightRangePounds(500, 4000, true),
	},
	HugeSize: {
		modifier:         -2,
		specialModifier:  2,
		flyModifier:      -4,
		stealthModifier:  -8,
		constructBonusHP: 50,
		spaceTall:        feet(15),
		spaceLong:        feet(15),
		reachTall:        feet(15),
		reachLong:        feet(10),
		heightRange:      heightRange(16, 32, true),
		weightRange:      weightRangeTons(2, 16, true),
	},
	GargantuanSize: {
		modifier:         -4,
		specialModifier:  4,
		flyModifier:      -6,
		stealthModifier:  -12,
		constructBonusHP: 80,
		spaceTall:        feet(20),
		spaceLong:        feet(20),
		reachTall:        feet(20),
		reachLong:        feet(15),
		heightRange:      heightRange(32, 64, true),
		weightRange:      weightRangeTons(16, 125, true),
	},
	ColossalSize: {
		modifier:         -8,
		specialModifier:  8,
		flyModifier:      -8,
		stealthModifier:  -16,
		constructBonusHP: 130,
		spaceTall:        feet(30),
		spaceLong:        feet(30),
		reachTall:        feet(30),
		reachLong:        feet(20),
		heightRange:      lengthRange{min: feet(64), max: feet(64), hasUpperBound: false},
		weightRange:      sizeWeightRange{min: tons(125), max: tons(125), hasUpperBound: false},
	},
	TitanicSize: {
		modifier:         -16,
		specialModifier:  12,
		flyModifier:      -10,
		stealthModifier:  -20,
		constructBonusHP: 210,
		spaceTall:        feet(40),
		spaceLong:        feet(40),
		reachTall:        feet(40),
		reachLong:        feet(30),
		heightRange:      lengthRange{min: feet(128), max: feet(128), hasUpperBound: false},
		weightRange:      sizeWeightRange{min: tons(1000), max: tons(1000), hasUpperBound: false},
	},
}
