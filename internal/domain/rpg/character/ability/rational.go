package ability

type rationalValue struct {
	numerator   int
	denominator int
}
type RationalValue = rationalValue

func NewRationalValue(numerator int, denominator int) (RationalValue, bool) {
	if denominator <= 0 {
		return rationalValue{}, false
	}

	return reduceRationalValue(numerator, denominator), true
}

func (r rationalValue) GetNumerator() int {
	return r.numerator
}

func (r rationalValue) GetDenominator() int {
	return r.denominator
}

func (r rationalValue) GetFloat64() float64 {
	return float64(r.numerator) / float64(r.denominator)
}

func (r rationalValue) Add(other RationalValue) RationalValue {
	return reduceRationalValue(
		(r.numerator*other.denominator)+(other.numerator*r.denominator),
		r.denominator*other.denominator,
	)
}

func (r rationalValue) MultiplyByInt(value int) RationalValue {
	return reduceRationalValue(r.numerator*value, r.denominator)
}

func (r rationalValue) Floor() int {
	return r.numerator / r.denominator
}

func isValidRationalValue(value RationalValue) bool {
	return value.denominator > 0
}

func isNonNegativeRationalValue(value RationalValue) bool {
	return isValidRationalValue(value) && value.numerator >= 0
}

func zeroRationalValue() RationalValue {
	return rationalValue{denominator: 1}
}

func reduceRationalValue(numerator int, denominator int) RationalValue {
	if numerator == 0 {
		return rationalValue{denominator: 1}
	}

	gcdValue := greatestCommonDivisor(absInt(numerator), absInt(denominator))
	return rationalValue{
		numerator:   numerator / gcdValue,
		denominator: denominator / gcdValue,
	}
}

func greatestCommonDivisor(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	if a == 0 {
		return 1
	}

	return a
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}

	return value
}
