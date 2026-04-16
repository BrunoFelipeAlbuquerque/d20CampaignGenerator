package ability

import "testing"

func mustNewRationalValue(t *testing.T, numerator int, denominator int) RationalValue {
	t.Helper()

	value, ok := NewRationalValue(numerator, denominator)
	if !ok {
		t.Fatalf("expected rational value %d/%d to be constructed", numerator, denominator)
	}

	return value
}
