package ability

import "testing"

func TestNewHitDie_CalculatesTotalsAndAverageBaseHP(t *testing.T) {
	hd, ok := NewHitDie(1, 2, 1, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}

	if hd.GetTotal() != 4 {
		t.Fatalf("expected total HD 4, got %d", hd.GetTotal())
	}

	if hd.GetAverageBaseHP() != 20 {
		t.Fatalf("expected average base HP 20, got %d", hd.GetAverageBaseHP())
	}
}

func TestNewStandardHitPoints_UsesConstitutionLedgerAndThreshold(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 14)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if hp.GetTotal() != 14 || hp.GetCurrent() != 14 {
		t.Fatalf("expected total/current 14, got total %d current %d", hp.GetTotal(), hp.GetCurrent())
	}

	if hp.GetDeathThreshold() != -14 {
		t.Fatalf("expected death threshold -14, got %d", hp.GetDeathThreshold())
	}

	sources := hp.GetSources()
	if len(sources) != 2 {
		t.Fatalf("expected 2 HP sources, got %d", len(sources))
	}

	if sources[0].GetName() != "Base Dice" || sources[0].GetValue() != 10 {
		t.Fatalf("expected base dice source 10, got %q = %d", sources[0].GetName(), sources[0].GetValue())
	}

	if sources[1].GetName() != "Constitution" || sources[1].GetValue() != 4 {
		t.Fatalf("expected constitution source 4, got %q = %d", sources[1].GetName(), sources[1].GetValue())
	}
}

func TestNewUndeadHitPoints_UsesCharismaAndIgnoresNonLethal(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewUndeadHitPoints(hd, 16)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if hp.GetTotal() != 16 || hp.GetDeathThreshold() != 0 {
		t.Fatalf("expected undead total 16 and threshold 0, got total %d and threshold %d", hp.GetTotal(), hp.GetDeathThreshold())
	}

	if !hp.IsNonLethalImmune() {
		t.Fatal("expected undead to ignore nonlethal damage")
	}
}

func TestNewConstructHitPoints_UsesSizeBonus(t *testing.T) {
	hd, ok := NewHitDie(0, 0, 2, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewConstructHitPoints(hd, MediumSize)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if hp.GetTotal() != 32 {
		t.Fatalf("expected construct total 32, got %d", hp.GetTotal())
	}

	sources := hp.GetSources()
	if len(sources) != 2 {
		t.Fatalf("expected 2 HP sources, got %d", len(sources))
	}

	if sources[1].GetName() != "Construct Size Bonus" || sources[1].GetValue() != 20 {
		t.Fatalf("expected construct size bonus 20, got %q = %d", sources[1].GetName(), sources[1].GetValue())
	}
}

func TestNewConstructHitPoints_UsesBuffedLargeSizesAndTitanic(t *testing.T) {
	hd, ok := NewHitDie(0, 0, 2, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	huge, ok := NewConstructHitPoints(hd, HugeSize)
	if !ok {
		t.Fatal("expected huge construct hit points to be constructed")
	}
	titanic, ok := NewConstructHitPoints(hd, TitanicSize)
	if !ok {
		t.Fatal("expected titanic construct hit points to be constructed")
	}

	if huge.GetTotal() != 62 {
		t.Fatalf("expected huge construct total 62, got %d", huge.GetTotal())
	}

	if titanic.GetTotal() != 222 {
		t.Fatalf("expected titanic construct total 222, got %d", titanic.GetTotal())
	}

	hugeSources := huge.GetSources()
	titanicSources := titanic.GetSources()

	if hugeSources[1].GetValue() != 50 {
		t.Fatalf("expected huge construct size bonus 50, got %d", hugeSources[1].GetValue())
	}

	if titanicSources[1].GetValue() != 210 {
		t.Fatalf("expected titanic construct size bonus 210, got %d", titanicSources[1].GetValue())
	}
}

func TestStandardHitPoints_AppliesMinimumOnePerHitDieFloor(t *testing.T) {
	hd, ok := NewHitDie(1, 0, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 1)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if hp.GetTotal() != 1 {
		t.Fatalf("expected minimum total HP 1, got %d", hp.GetTotal())
	}

	sources := hp.GetSources()
	if len(sources) != 3 {
		t.Fatalf("expected floor adjustment source, got %d sources", len(sources))
	}

	if sources[2].GetName() != "Minimum 1 HP per Hit Die" || sources[2].GetValue() != 2 {
		t.Fatalf("expected minimum floor adjustment 2, got %q = %d", sources[2].GetName(), sources[2].GetValue())
	}
}

func TestTakeDamage_ConsumesTemporaryHitPointsFirst(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.SetTemporaryHPSource("False Life (Temp)", 5); !ok {
		t.Fatal("expected temporary HP source to be added")
	}

	if ok := hp.TakeDamage(3, false); !ok {
		t.Fatal("expected lethal damage to be accepted")
	}

	if hp.GetTemporary() != 2 || hp.GetCurrent() != hp.GetTotal() {
		t.Fatalf("expected temporary 2 and no lethal damage taken, got temporary %d and current %d", hp.GetTemporary(), hp.GetCurrent())
	}

	if ok := hp.TakeDamage(4, false); !ok {
		t.Fatal("expected second lethal damage instance to be accepted")
	}

	if hp.GetTemporary() != 0 {
		t.Fatalf("expected temporary HP to be exhausted, got %d", hp.GetTemporary())
	}

	if hp.GetCurrent() != hp.GetTotal()-2 {
		t.Fatalf("expected current HP reduced by 2, got %d", hp.GetCurrent())
	}

	if len(hp.GetTemporarySources()) != 0 {
		t.Fatal("expected temporary HP ledger to remove depleted source")
	}
}

func TestSetTemporaryHPSource_ReplacesSameSourceInsteadOfStacking(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.SetTemporaryHPSource("False Life", 10); !ok {
		t.Fatal("expected first temporary HP source to be added")
	}

	if ok := hp.TakeDamage(3, false); !ok {
		t.Fatal("expected damage to consume temporary HP")
	}

	if ok := hp.SetTemporaryHPSource("False Life", 5); !ok {
		t.Fatal("expected same temporary HP source to be replaced")
	}

	if hp.GetTemporary() != 5 {
		t.Fatalf("expected temporary HP to reset to replacement source total 5, got %d", hp.GetTemporary())
	}

	if len(hp.GetTemporarySources()) != 1 {
		t.Fatalf("expected one temporary HP source after replacement, got %d", len(hp.GetTemporarySources()))
	}
}

func TestSetTemporaryHPSource_DifferentSourcesDoNotStackIntoOnePool(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.SetTemporaryHPSource("False Life", 10); !ok {
		t.Fatal("expected false life temporary HP source to be added")
	}

	if ok := hp.SetTemporaryHPSource("Aid", 5); !ok {
		t.Fatal("expected aid temporary HP source to be added")
	}

	if hp.GetTemporary() != 10 {
		t.Fatalf("expected only the highest temporary HP pool to apply, got %d", hp.GetTemporary())
	}

	if ok := hp.TakeDamage(6, false); !ok {
		t.Fatal("expected damage to consume only the active temporary HP pool")
	}

	if hp.GetTemporary() != 4 {
		t.Fatalf("expected active temporary HP pool to drop to 4, got %d", hp.GetTemporary())
	}

	if hp.GetCurrent() != hp.GetTotal() {
		t.Fatalf("expected no lethal damage while active temporary HP remained, got current %d", hp.GetCurrent())
	}

	if len(hp.GetTemporarySources()) != 2 {
		t.Fatalf("expected both temporary HP sources to remain tracked, got %d", len(hp.GetTemporarySources()))
	}
}

func TestRemoveTemporaryHPSource_PromotesNextHighestPool(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.SetTemporaryHPSource("False Life", 10); !ok {
		t.Fatal("expected false life temporary HP source to be added")
	}

	if ok := hp.SetTemporaryHPSource("Aid", 5); !ok {
		t.Fatal("expected aid temporary HP source to be added")
	}

	if ok := hp.RemoveTemporaryHPSource("False Life"); !ok {
		t.Fatal("expected active temporary HP source to be removable")
	}

	if hp.GetTemporary() != 5 {
		t.Fatalf("expected next highest temporary HP pool to become active, got %d", hp.GetTemporary())
	}

	if len(hp.GetTemporarySources()) != 1 {
		t.Fatalf("expected one temporary HP source to remain, got %d", len(hp.GetTemporarySources()))
	}

	if hp.GetTemporarySources()[0].GetName() != "Aid" {
		t.Fatalf("expected Aid to remain as the active temporary HP source, got %q", hp.GetTemporarySources()[0].GetName())
	}
}

func TestTakeDamage_NonLethalStacksAsDebtWithoutBecomingLethal(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 10)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.TakeDamage(hp.GetTotal()+2, true); !ok {
		t.Fatal("expected nonlethal damage to be accepted")
	}

	if hp.GetNonLethal() != hp.GetTotal()+2 {
		t.Fatalf("expected nonlethal damage to keep stacking as debt, got %d", hp.GetNonLethal())
	}

	if hp.GetCurrent() != hp.GetTotal() {
		t.Fatalf("expected nonlethal damage to not directly reduce current HP, got %d", hp.GetCurrent())
	}
}

func TestTakeDamage_NonLethalIgnoredForConstructsAndUndead(t *testing.T) {
	undeadHD, ok := NewHitDie(0, 1, 0, 0)
	if !ok {
		t.Fatal("expected undead hit die to be constructed")
	}
	constructHD, ok := NewHitDie(0, 0, 1, 0)
	if !ok {
		t.Fatal("expected construct hit die to be constructed")
	}
	undead, ok := NewUndeadHitPoints(undeadHD, 12)
	if !ok {
		t.Fatal("expected undead hit points to be constructed")
	}
	construct, ok := NewConstructHitPoints(constructHD, SmallSize)
	if !ok {
		t.Fatal("expected construct hit points to be constructed")
	}

	if ok := undead.TakeDamage(3, true); !ok {
		t.Fatal("expected undead nonlethal damage to be processed")
	}

	if ok := construct.TakeDamage(3, true); !ok {
		t.Fatal("expected construct nonlethal damage to be processed")
	}

	if undead.GetNonLethal() != 0 || undead.GetCurrent() != undead.GetTotal() {
		t.Fatalf("expected undead to ignore nonlethal damage, got nonlethal %d and current %d", undead.GetNonLethal(), undead.GetCurrent())
	}

	if construct.GetNonLethal() != 0 || construct.GetCurrent() != construct.GetTotal() {
		t.Fatalf("expected construct to ignore nonlethal damage, got nonlethal %d and current %d", construct.GetNonLethal(), construct.GetCurrent())
	}
}

func TestHeal_RestoresLethalAndReducesNonLethal(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.TakeDamage(4, false); !ok {
		t.Fatal("expected lethal damage to be accepted")
	}

	if ok := hp.TakeDamage(3, true); !ok {
		t.Fatal("expected nonlethal damage to be accepted")
	}

	if ok := hp.Heal(2); !ok {
		t.Fatal("expected healing to be accepted")
	}

	if hp.GetCurrent() != hp.GetTotal()-2 {
		t.Fatalf("expected current HP to recover by 2, got %d", hp.GetCurrent())
	}

	if hp.GetNonLethal() != 1 {
		t.Fatalf("expected nonlethal damage to recover by 2 as well, got %d", hp.GetNonLethal())
	}
}

func TestUpdateConstitutionScore_RecalculatesCurrentTotalAndThreshold(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.TakeDamage(4, false); !ok {
		t.Fatal("expected lethal damage to be accepted")
	}

	if ok := hp.UpdateConstitutionScore(16); !ok {
		t.Fatal("expected constitution update to be accepted")
	}

	if hp.GetTotal() != 16 {
		t.Fatalf("expected total HP 16 after constitution update, got %d", hp.GetTotal())
	}

	if hp.GetCurrent() != 12 {
		t.Fatalf("expected current HP to rise with total HP, got %d", hp.GetCurrent())
	}

	if hp.GetDeathThreshold() != -16 {
		t.Fatalf("expected death threshold -16, got %d", hp.GetDeathThreshold())
	}
}

func TestUpdateConstitutionScore_RejectsZeroForStandardHitPoints(t *testing.T) {
	hd, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 12)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	beforeTotal := hp.GetTotal()
	beforeCurrent := hp.GetCurrent()
	beforeThreshold := hp.GetDeathThreshold()

	if ok := hp.UpdateConstitutionScore(0); ok {
		t.Fatal("expected constitution 0 update to be rejected for standard hit points")
	}

	if hp.GetTotal() != beforeTotal {
		t.Fatalf("expected total HP to remain %d, got %d", beforeTotal, hp.GetTotal())
	}

	if hp.GetCurrent() != beforeCurrent {
		t.Fatalf("expected current HP to remain %d, got %d", beforeCurrent, hp.GetCurrent())
	}

	if hp.GetDeathThreshold() != beforeThreshold {
		t.Fatalf("expected death threshold to remain %d, got %d", beforeThreshold, hp.GetDeathThreshold())
	}
}

func TestUpdateCharismaAndSize_RecalculateSpecificLedgerEntries(t *testing.T) {
	undeadHD, ok := NewHitDie(0, 2, 0, 0)
	if !ok {
		t.Fatal("expected undead hit die to be constructed")
	}
	constructHD, ok := NewHitDie(0, 0, 2, 0)
	if !ok {
		t.Fatal("expected construct hit die to be constructed")
	}
	undead, ok := NewUndeadHitPoints(undeadHD, 12)
	if !ok {
		t.Fatal("expected undead hit points to be constructed")
	}
	construct, ok := NewConstructHitPoints(constructHD, SmallSize)
	if !ok {
		t.Fatal("expected construct hit points to be constructed")
	}

	if ok := undead.UpdateCharismaScore(16); !ok {
		t.Fatal("expected charisma update to be accepted")
	}

	if undead.GetTotal() != 16 {
		t.Fatalf("expected undead total HP 16 after charisma update, got %d", undead.GetTotal())
	}

	if ok := construct.UpdateSize(HugeSize); !ok {
		t.Fatal("expected size update to be accepted")
	}

	if construct.GetTotal() != 62 {
		t.Fatalf("expected construct total HP 62 after size update, got %d", construct.GetTotal())
	}

	constructSources := construct.GetSources()
	if len(constructSources) != 2 {
		t.Fatalf("expected 2 construct HP sources, got %d", len(constructSources))
	}

	if constructSources[1].GetName() != "Construct Size Bonus" || constructSources[1].GetValue() != 50 {
		t.Fatalf(
			"expected construct size source to update to 50, got %q = %d",
			constructSources[1].GetName(),
			constructSources[1].GetValue(),
		)
	}

	if ok := construct.UpdateSize(TitanicSize); !ok {
		t.Fatal("expected titanic size update to be accepted")
	}

	if construct.GetTotal() != 222 {
		t.Fatalf("expected construct total HP 222 after titanic update, got %d", construct.GetTotal())
	}

	constructSources = construct.GetSources()
	if constructSources[1].GetValue() != 210 {
		t.Fatalf("expected titanic construct size source 210, got %d", constructSources[1].GetValue())
	}
}

func TestTakeDamage_CanDropBelowDeathThresholdWithoutOwningCombatState(t *testing.T) {
	hd, ok := NewHitDie(0, 1, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}
	hp, ok := NewStandardHitPoints(hd, 10)
	if !ok {
		t.Fatal("expected hit points to be constructed")
	}

	if ok := hp.TakeDamage(hp.GetTotal()+10, false); !ok {
		t.Fatal("expected lethal damage to be accepted")
	}

	if hp.GetCurrent() > hp.GetDeathThreshold() {
		t.Fatalf("expected current HP to be at or below death threshold, got current %d and threshold %d", hp.GetCurrent(), hp.GetDeathThreshold())
	}
}

func TestHitPointConstructors_RejectInvalidInputs(t *testing.T) {
	if _, ok := NewHitDie(-1, 0, 0, 0); ok {
		t.Fatal("expected invalid hit die to be rejected")
	}

	if _, ok := NewHitDie(0, 0, 0, 0); ok {
		t.Fatal("expected zero-total hit die to be rejected")
	}

	hd, ok := NewHitDie(0, 1, 0, 0)
	if !ok {
		t.Fatal("expected hit die to be constructed")
	}

	if _, ok := NewStandardHitPoints(hd, -1); ok {
		t.Fatal("expected invalid standard HP input to be rejected")
	}

	if _, ok := NewStandardHitPoints(hd, 0); ok {
		t.Fatal("expected constitution 0 standard HP input to be rejected")
	}

	if _, ok := NewUndeadHitPoints(hd, -1); ok {
		t.Fatal("expected invalid undead HP input to be rejected")
	}

	if _, ok := NewConstructHitPoints(hd, Size("Invalid")); ok {
		t.Fatal("expected invalid construct size to be rejected")
	}

	invalidHD := HitDie{}
	if _, ok := NewStandardHitPoints(invalidHD, 10); ok {
		t.Fatal("expected semantically invalid hit die payload to be rejected")
	}
}
