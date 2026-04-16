package ability

import "testing"

func TestNewSavingThrow_PreservesActualAndRoundsDownValue(t *testing.T) {
	save, ok := NewSavingThrow(ReflexSave, 2.3333333333)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if save.GetID() != ReflexSave {
		t.Fatalf("expected save id %q, got %q", ReflexSave, save.GetID())
	}

	if !almostEqual(save.GetActualValue(), 2.3333333333) {
		t.Fatalf("expected actual save 2.33, got %.10f", save.GetActualValue())
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

	if !almostEqual(save.GetActualValue(), 2.5) {
		t.Fatalf("expected actual save 2.5, got %.2f", save.GetActualValue())
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

	if !almostEqual(save.GetActualValue(), 1.6666666667) {
		t.Fatalf("expected actual save about 1.67, got %.10f", save.GetActualValue())
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

	if !almostEqual(save.GetActualValue(), 3) {
		t.Fatalf("expected actual save 3.0, got %.2f", save.GetActualValue())
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

	if !almostEqual(save.GetActualValue(), 0.6666666667) {
		t.Fatalf("expected actual save about 0.67, got %.10f", save.GetActualValue())
	}

	if save.GetValue() != 0 {
		t.Fatalf("expected rounded save 0, got %d", save.GetValue())
	}

	if save.HasGoodBaseBonusApplied() {
		t.Fatal("expected reset save progression to clear the one-time +2 tracking flag")
	}
}

func TestSavingThrowAddClassLevel_RejectsInvalidProgression(t *testing.T) {
	save, ok := NewSavingThrow(ReflexSave, 1)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if ok := save.AddClassLevel(2, SavingThrowProgression(0.4)); ok {
		t.Fatal("expected invalid save progression to be rejected")
	}

	if !almostEqual(save.GetActualValue(), 1) || save.GetValue() != 1 {
		t.Fatalf("expected save to remain unchanged, got actual %.2f and value %d", save.GetActualValue(), save.GetValue())
	}
}

func TestSavingThrowSetID_RejectsInvalidIDs(t *testing.T) {
	save, ok := NewSavingThrow(WillSave, 0.5)
	if !ok {
		t.Fatal("expected saving throw to be constructed")
	}

	if _, ok := NewSavingThrow(SavingThrowID("Luck"), 0.5); ok {
		t.Fatal("expected invalid saving throw id to be rejected at construction")
	}

	if ok := save.SetID(SavingThrowID("Luck")); ok {
		t.Fatal("expected invalid save id to be rejected")
	}

	if save.GetID() != WillSave {
		t.Fatalf("expected save id to remain %q, got %q", WillSave, save.GetID())
	}
}
