package ability

import "testing"

func TestNewCasterLevel_PreservesThreePresentPillars(t *testing.T) {
	level := NewCasterLevel(4, 2, 7)

	arcane, arcaneValid := level.GetArcane()
	divine, divineValid := level.GetDivine()
	primal, primalValid := level.GetPrimal()

	if arcane != 4 || !arcaneValid || divine != 2 || !divineValid || primal != 7 || !primalValid {
		t.Fatalf(
			"expected present pillars (4, 2, 7), got arcane (%d, %t), divine (%d, %t), primal (%d, %t)",
			arcane, arcaneValid,
			divine, divineValid,
			primal, primalValid,
		)
	}
}

func TestNewImpossibleCasterLevel_MarksAllPillarsAsUnavailable(t *testing.T) {
	level := NewImpossibleCasterLevel()

	if _, ok := level.GetArcane(); ok {
		t.Fatal("expected arcane pillar to be unavailable")
	}

	if _, ok := level.GetDivine(); ok {
		t.Fatal("expected divine pillar to be unavailable")
	}

	if _, ok := level.GetPrimal(); ok {
		t.Fatal("expected primal pillar to be unavailable")
	}
}

func TestCasterLevelGetCasterLevel_ReturnsWholeValueObject(t *testing.T) {
	level := NewCasterLevel(3, 5, 1)
	got := level.GetCasterLevel()

	arcane, arcaneValid := got.GetArcane()
	divine, divineValid := got.GetDivine()
	primal, primalValid := got.GetPrimal()

	if arcane != 3 || !arcaneValid || divine != 5 || !divineValid || primal != 1 || !primalValid {
		t.Fatalf(
			"expected returned caster level (3, 5, 1), got arcane (%d, %t), divine (%d, %t), primal (%d, %t)",
			arcane, arcaneValid,
			divine, divineValid,
			primal, primalValid,
		)
	}
}

func TestCasterLevelSetters_RejectNegativeValues(t *testing.T) {
	level := NewCasterLevel(1, 2, 3)

	if ok := level.SetArcaneCasterLevel(-1); ok {
		t.Fatal("expected negative arcane caster level to be rejected")
	}

	if ok := level.SetDivineCasterLevel(-1); ok {
		t.Fatal("expected negative divine caster level to be rejected")
	}

	if ok := level.SetPrimalCasterLevel(-1); ok {
		t.Fatal("expected negative primal caster level to be rejected")
	}

	arcane, _ := level.GetArcane()
	divine, _ := level.GetDivine()
	primal, _ := level.GetPrimal()

	if arcane != 1 || divine != 2 || primal != 3 {
		t.Fatalf("expected pillars to remain unchanged, got (%d, %d, %d)", arcane, divine, primal)
	}
}

func TestCasterLevelDisablePillar_MakesItUnavailable(t *testing.T) {
	level := NewCasterLevel(1, 4, 2)

	level.DisableDivineCasterLevel()

	if _, ok := level.GetDivine(); ok {
		t.Fatal("expected divine pillar to be unavailable after disabling it")
	}
}

func TestCasterLevelSetCasterLevel_ReplacesWholeValueObject(t *testing.T) {
	level := NewCasterLevel(0, 0, 0)
	replacement := NewImpossibleCasterLevel()
	replacement.SetDivineCasterLevel(4)

	if ok := level.SetCasterLevel(replacement); !ok {
		t.Fatal("expected caster level to be replaced")
	}

	if _, ok := level.GetArcane(); ok {
		t.Fatal("expected arcane pillar to become unavailable")
	}

	divine, ok := level.GetDivine()
	if !ok || divine != 4 {
		t.Fatalf("expected divine pillar (4, true), got (%d, %t)", divine, ok)
	}

	if _, ok := level.GetPrimal(); ok {
		t.Fatal("expected primal pillar to become unavailable")
	}
}

func TestCasterLevelAddCasterLevel_EnablesUnavailablePillarsWhenNeeded(t *testing.T) {
	level := NewImpossibleCasterLevel()
	addition := NewImpossibleCasterLevel()
	addition.SetArcaneCasterLevel(3)
	addition.SetPrimalCasterLevel(2)

	if ok := level.AddCasterLevel(addition); !ok {
		t.Fatal("expected caster level contribution to be accepted")
	}

	arcane, arcaneValid := level.GetArcane()
	primal, primalValid := level.GetPrimal()

	if arcane != 3 || !arcaneValid || primal != 2 || !primalValid {
		t.Fatalf(
			"expected added pillars arcane (3, true) and primal (2, true), got arcane (%d, %t), primal (%d, %t)",
			arcane, arcaneValid,
			primal, primalValid,
		)
	}
}

func TestCasterLevelAddSpecificPillars_KeepPillarsIndependent(t *testing.T) {
	level := NewImpossibleCasterLevel()

	if ok := level.AddDivineCasterLevel(3); !ok {
		t.Fatal("expected divine contribution to be accepted")
	}

	if ok := level.AddPrimalCasterLevel(2); !ok {
		t.Fatal("expected primal contribution to be accepted")
	}

	if _, ok := level.GetArcane(); ok {
		t.Fatal("expected arcane pillar to remain unavailable")
	}

	divine, divineValid := level.GetDivine()
	primal, primalValid := level.GetPrimal()
	if divine != 3 || !divineValid || primal != 2 || !primalValid {
		t.Fatalf(
			"expected pillars divine (3, true) and primal (2, true), got divine (%d, %t), primal (%d, %t)",
			divine, divineValid,
			primal, primalValid,
		)
	}
}

func TestCasterLevelSetCasterLevel_RejectsNegativeWholeValue(t *testing.T) {
	level := NewCasterLevel(1, 1, 1)

	invalid := CasterLevel{
		arcane:      -1,
		arcaneValid: true,
	}

	if ok := level.SetCasterLevel(invalid); ok {
		t.Fatal("expected negative caster level to be rejected")
	}

	arcane, _ := level.GetArcane()
	divine, _ := level.GetDivine()
	primal, _ := level.GetPrimal()
	if arcane != 1 || divine != 1 || primal != 1 {
		t.Fatalf("expected pillars to remain unchanged, got (%d, %d, %d)", arcane, divine, primal)
	}
}

func TestCasterLevelAddSpecificPillars_RejectNegativeValues(t *testing.T) {
	level := NewCasterLevel(2, 2, 2)

	if ok := level.AddArcaneCasterLevel(-1); ok {
		t.Fatal("expected negative caster level contribution to be rejected")
	}

	arcane, _ := level.GetArcane()
	divine, _ := level.GetDivine()
	primal, _ := level.GetPrimal()
	if arcane != 2 || divine != 2 || primal != 2 {
		t.Fatalf("expected pillars to remain unchanged, got (%d, %d, %d)", arcane, divine, primal)
	}
}
