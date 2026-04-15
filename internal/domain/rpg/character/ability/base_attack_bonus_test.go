package ability

import "testing"

func TestNewBaseAttackBonus_PreservesActualAndRoundsDownValue(t *testing.T) {
	bab := NewBaseAttackBonus(0.75)

	if !almostEqual(bab.GetActualValue(), 0.75) {
		t.Fatalf("expected actual BAB 0.75, got %.2f", bab.GetActualValue())
	}

	if bab.GetValue() != 0 {
		t.Fatalf("expected rounded BAB 0, got %d", bab.GetValue())
	}
}

func TestNewBaseAttackBonusByClassLevel_UsesFractionalClassProgression(t *testing.T) {
	bab := NewBaseAttackBonusByClassLevel(2, BaseAttackBonusThreeQuarters)

	if !almostEqual(bab.GetActualValue(), 1.5) {
		t.Fatalf("expected actual BAB 1.5, got %.2f", bab.GetActualValue())
	}

	if bab.GetValue() != 1 {
		t.Fatalf("expected rounded BAB 1, got %d", bab.GetValue())
	}
}

func TestBaseAttackBonusSetByClassLevel_RejectsInvalidProgression(t *testing.T) {
	bab := NewBaseAttackBonus(1)

	if ok := bab.SetByClassLevel(3, BaseAttackBonusProgression(0.6)); ok {
		t.Fatal("expected invalid BAB progression to be rejected")
	}

	if !almostEqual(bab.GetActualValue(), 1) || bab.GetValue() != 1 {
		t.Fatalf("expected BAB to remain unchanged, got actual %.2f and value %d", bab.GetActualValue(), bab.GetValue())
	}
}

func TestBaseAttackBonusSetActualValue_RejectsNegativeValues(t *testing.T) {
	bab := NewBaseAttackBonus(2.25)

	if ok := bab.SetActualValue(-0.5); ok {
		t.Fatal("expected negative BAB to be rejected")
	}

	if !almostEqual(bab.GetActualValue(), 2.25) || bab.GetValue() != 2 {
		t.Fatalf("expected BAB to remain unchanged, got actual %.2f and value %d", bab.GetActualValue(), bab.GetValue())
	}
}
