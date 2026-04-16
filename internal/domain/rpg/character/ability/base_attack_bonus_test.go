package ability

import "testing"

func TestNewBaseAttackBonus_PreservesActualAndRoundsDownValue(t *testing.T) {
	bab, ok := NewBaseAttackBonus(mustNewRationalValue(t, 3, 4))
	if !ok {
		t.Fatal("expected BAB to be constructed")
	}

	if bab.GetActualValue().GetNumerator() != 3 || bab.GetActualValue().GetDenominator() != 4 {
		t.Fatalf(
			"expected actual BAB 3/4, got %d/%d",
			bab.GetActualValue().GetNumerator(),
			bab.GetActualValue().GetDenominator(),
		)
	}

	if bab.GetValue() != 0 {
		t.Fatalf("expected rounded BAB 0, got %d", bab.GetValue())
	}
}

func TestNewBaseAttackBonusByClassLevel_UsesFractionalClassProgression(t *testing.T) {
	bab, ok := NewBaseAttackBonusByClassLevel(2, BaseAttackBonusThreeQuarters)
	if !ok {
		t.Fatal("expected BAB to be constructed")
	}

	if bab.GetActualValue().GetNumerator() != 3 || bab.GetActualValue().GetDenominator() != 2 {
		t.Fatalf(
			"expected actual BAB 3/2, got %d/%d",
			bab.GetActualValue().GetNumerator(),
			bab.GetActualValue().GetDenominator(),
		)
	}

	if bab.GetValue() != 1 {
		t.Fatalf("expected rounded BAB 1, got %d", bab.GetValue())
	}
}

func TestBaseAttackBonusSetByClassLevel_RejectsInvalidProgression(t *testing.T) {
	bab, ok := NewBaseAttackBonus(mustNewRationalValue(t, 1, 1))
	if !ok {
		t.Fatal("expected BAB to be constructed")
	}

	if ok := bab.SetByClassLevel(3, BaseAttackBonusProgression("0.6")); ok {
		t.Fatal("expected invalid BAB progression to be rejected")
	}

	if bab.GetActualValue().GetNumerator() != 1 || bab.GetActualValue().GetDenominator() != 1 || bab.GetValue() != 1 {
		t.Fatalf(
			"expected BAB to remain 1/1 and value 1, got %d/%d and value %d",
			bab.GetActualValue().GetNumerator(),
			bab.GetActualValue().GetDenominator(),
			bab.GetValue(),
		)
	}
}

func TestBaseAttackBonusSetActualValue_RejectsNegativeValues(t *testing.T) {
	bab, ok := NewBaseAttackBonus(mustNewRationalValue(t, 9, 4))
	if !ok {
		t.Fatal("expected BAB to be constructed")
	}

	invalidActualValue, ok := NewRationalValue(-1, 2)
	if !ok {
		t.Fatal("expected negative rational value to still be constructible for validation tests")
	}

	if _, ok := NewBaseAttackBonus(invalidActualValue); ok {
		t.Fatal("expected invalid BAB constructor input to be rejected")
	}

	if ok := bab.SetActualValue(invalidActualValue); ok {
		t.Fatal("expected negative BAB to be rejected")
	}

	if bab.GetActualValue().GetNumerator() != 9 || bab.GetActualValue().GetDenominator() != 4 || bab.GetValue() != 2 {
		t.Fatalf(
			"expected BAB to remain 9/4 and value 2, got %d/%d and value %d",
			bab.GetActualValue().GetNumerator(),
			bab.GetActualValue().GetDenominator(),
			bab.GetValue(),
		)
	}
}
