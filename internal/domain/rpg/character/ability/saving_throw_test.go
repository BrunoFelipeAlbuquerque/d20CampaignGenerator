package ability

import "testing"

func TestNewSavingThrow_PreservesActualAndRoundsDownValue(t *testing.T) {
	save, ok := NewSavingThrow(ReflexSave, mustNewRationalValue(t, 7, 3))
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if save.GetID() != ReflexSave {
		t.Fatalf("expected save id %q, got %q", ReflexSave, save.GetID())
	}

	if save.GetActualValue().GetNumerator() != 7 || save.GetActualValue().GetDenominator() != 3 {
		t.Fatalf(
			"expected actual save 7/3, got %d/%d",
			save.GetActualValue().GetNumerator(),
			save.GetActualValue().GetDenominator(),
		)
	}

	if save.GetValue() != 2 {
		t.Fatalf("expected rounded save 2, got %d", save.GetValue())
	}

	if save.HasGoodBaseBonusApplied() {
		t.Fatal("expected direct actual-value constructor to not track a good-save base bonus")
	}
}

func TestNewSavingThrowByClassLevel_GoodProgressionIncludesOneTimeBonus(t *testing.T) {
	save, ok := NewSavingThrowByClassLevel(FortitudeSave, 1, SavingThrowGood)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if save.GetActualValue().GetNumerator() != 5 || save.GetActualValue().GetDenominator() != 2 {
		t.Fatalf(
			"expected actual save 5/2, got %d/%d",
			save.GetActualValue().GetNumerator(),
			save.GetActualValue().GetDenominator(),
		)
	}

	if save.GetValue() != 2 {
		t.Fatalf("expected rounded save 2, got %d", save.GetValue())
	}

	if !save.HasGoodBaseBonusApplied() {
		t.Fatal("expected good save progression to mark the one-time +2 as applied")
	}
}

func TestNewSavingThrowByClassLevel_PoorProgressionRoundsDownFraction(t *testing.T) {
	save, ok := NewSavingThrowByClassLevel(WillSave, 5, SavingThrowPoor)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if save.GetActualValue().GetNumerator() != 5 || save.GetActualValue().GetDenominator() != 3 {
		t.Fatalf(
			"expected actual save 5/3, got %d/%d",
			save.GetActualValue().GetNumerator(),
			save.GetActualValue().GetDenominator(),
		)
	}

	if save.GetValue() != 1 {
		t.Fatalf("expected rounded save 1, got %d", save.GetValue())
	}
}

func TestSavingThrowAddClassLevel_DoesNotRepeatGoodSaveBaseBonusAcrossMulticlassing(t *testing.T) {
	save, ok := NewSavingThrowByClassLevel(FortitudeSave, 1, SavingThrowGood)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if ok := save.AddClassLevel(1, SavingThrowGood); !ok {
		t.Fatal("expected second good save class level to be accepted")
	}

	if save.GetActualValue().GetNumerator() != 3 || save.GetActualValue().GetDenominator() != 1 {
		t.Fatalf(
			"expected actual save 3/1, got %d/%d",
			save.GetActualValue().GetNumerator(),
			save.GetActualValue().GetDenominator(),
		)
	}

	if save.GetValue() != 3 {
		t.Fatalf("expected rounded save 3, got %d", save.GetValue())
	}
}

func TestSavingThrowSetByClassLevel_ReplacesExistingState(t *testing.T) {
	save, ok := NewSavingThrowByClassLevel(FortitudeSave, 1, SavingThrowGood)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if ok := save.SetByClassLevel(2, SavingThrowPoor); !ok {
		t.Fatal("expected save to be reset from a new class progression")
	}

	if save.GetActualValue().GetNumerator() != 2 || save.GetActualValue().GetDenominator() != 3 {
		t.Fatalf(
			"expected actual save 2/3, got %d/%d",
			save.GetActualValue().GetNumerator(),
			save.GetActualValue().GetDenominator(),
		)
	}

	if save.GetValue() != 0 {
		t.Fatalf("expected rounded save 0, got %d", save.GetValue())
	}

	if save.HasGoodBaseBonusApplied() {
		t.Fatal("expected reset save progression to clear the one-time +2 tracking flag")
	}
}

func TestSavingThrowAddClassLevel_RejectsInvalidProgression(t *testing.T) {
	save, ok := NewSavingThrow(ReflexSave, mustNewRationalValue(t, 1, 1))
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if ok := save.AddClassLevel(2, SavingThrowProgression("0.4")); ok {
		t.Fatal("expected invalid save progression to be rejected")
	}

	if save.GetActualValue().GetNumerator() != 1 || save.GetActualValue().GetDenominator() != 1 || save.GetValue() != 1 {
		t.Fatalf(
			"expected save to remain 1/1 and value 1, got %d/%d and value %d",
			save.GetActualValue().GetNumerator(),
			save.GetActualValue().GetDenominator(),
			save.GetValue(),
		)
	}
}

func TestSavingThrowSetID_RejectsInvalidIDs(t *testing.T) {
	save, ok := NewSavingThrow(WillSave, mustNewRationalValue(t, 1, 2))
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if _, ok := NewSavingThrow(SavingThrowID("Luck"), mustNewRationalValue(t, 1, 2)); ok {
		t.Fatal("expected invalid saving throw id to be rejected at construction")
	}

	if ok := save.SetID(SavingThrowID("Luck")); ok {
		t.Fatal("expected invalid save id to be rejected")
	}

	if save.GetID() != WillSave {
		t.Fatalf("expected save id to remain %q, got %q", WillSave, save.GetID())
	}
}

func TestNewSavingThrow_RejectsNegativeActualValue(t *testing.T) {
	invalidActualValue, ok := NewRationalValue(-1, 2)
	if !ok {
		t.Fatal("expected negative rational value to still be constructible for validation tests")
	}

	if _, ok := NewSavingThrow(FortitudeSave, invalidActualValue); ok {
		t.Fatal("expected invalid actual value to be rejected")
	}
}
