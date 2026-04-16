package ability

import "testing"

func TestNewCasterLevel_PreservesThreePresentSources(t *testing.T) {
	level, ok := NewCasterLevel(4, 2, 7)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}

	arcane, arcaneValid := level.GetSourceLevel(ArcaneCasterSource)
	divine, divineValid := level.GetSourceLevel(DivineCasterSource)
	primal, primalValid := level.GetSourceLevel(PrimalCasterSource)

	if arcane != 4 || !arcaneValid || divine != 2 || !divineValid || primal != 7 || !primalValid {
		t.Fatalf(
			"expected present sources (4, 2, 7), got arcane (%d, %t), divine (%d, %t), primal (%d, %t)",
			arcane, arcaneValid,
			divine, divineValid,
			primal, primalValid,
		)
	}
}

func TestNewImpossibleCasterLevel_MarksAllSourcesAsUnavailable(t *testing.T) {
	level := NewImpossibleCasterLevel()

	if _, ok := level.GetSourceLevel(ArcaneCasterSource); ok {
		t.Fatal("expected arcane source to be unavailable")
	}

	if _, ok := level.GetSourceLevel(DivineCasterSource); ok {
		t.Fatal("expected divine source to be unavailable")
	}

	if _, ok := level.GetSourceLevel(PrimalCasterSource); ok {
		t.Fatal("expected primal source to be unavailable")
	}
}

func TestCasterLevelSetSourceLevel_RejectsNegativeValues(t *testing.T) {
	level, ok := NewCasterLevel(1, 2, 3)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}

	if ok := level.SetSourceLevel(ArcaneCasterSource, -1); ok {
		t.Fatal("expected negative arcane source level to be rejected")
	}

	arcane, _ := level.GetSourceLevel(ArcaneCasterSource)
	divine, _ := level.GetSourceLevel(DivineCasterSource)
	primal, _ := level.GetSourceLevel(PrimalCasterSource)
	if arcane != 1 || divine != 2 || primal != 3 {
		t.Fatalf("expected sources to remain unchanged, got (%d, %d, %d)", arcane, divine, primal)
	}
}

func TestCasterLevelDisableSourceLevel_MakesSourceUnavailable(t *testing.T) {
	level, ok := NewCasterLevel(1, 4, 2)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}

	if ok := level.DisableSourceLevel(DivineCasterSource); !ok {
		t.Fatal("expected divine source to be disabled")
	}

	if _, ok := level.GetSourceLevel(DivineCasterSource); ok {
		t.Fatal("expected divine source to be unavailable after disabling it")
	}
}

func TestCasterLevelSetCasterLevel_ReplacesWholeValueObject(t *testing.T) {
	level, ok := NewCasterLevel(0, 0, 0)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}

	replacement := NewImpossibleCasterLevel()
	if ok := replacement.SetSourceLevel(DivineCasterSource, 4); !ok {
		t.Fatal("expected replacement caster level to accept divine source")
	}

	if ok := level.SetCasterLevel(replacement); !ok {
		t.Fatal("expected caster level to be replaced")
	}

	if _, ok := level.GetSourceLevel(ArcaneCasterSource); ok {
		t.Fatal("expected arcane source to become unavailable")
	}

	divine, ok := level.GetSourceLevel(DivineCasterSource)
	if !ok || divine != 4 {
		t.Fatalf("expected divine source (4, true), got (%d, %t)", divine, ok)
	}

	if _, ok := level.GetSourceLevel(PrimalCasterSource); ok {
		t.Fatal("expected primal source to become unavailable")
	}
}

func TestCasterLevelAddCasterLevel_EnablesUnavailableSourcesWhenNeeded(t *testing.T) {
	level := NewImpossibleCasterLevel()
	addition := NewImpossibleCasterLevel()

	if ok := addition.SetSourceLevel(ArcaneCasterSource, 3); !ok {
		t.Fatal("expected arcane source level to be set")
	}

	if ok := addition.SetSourceLevel(PrimalCasterSource, 2); !ok {
		t.Fatal("expected primal source level to be set")
	}

	if ok := level.AddCasterLevel(addition); !ok {
		t.Fatal("expected caster level contribution to be accepted")
	}

	arcane, arcaneValid := level.GetSourceLevel(ArcaneCasterSource)
	primal, primalValid := level.GetSourceLevel(PrimalCasterSource)

	if arcane != 3 || !arcaneValid || primal != 2 || !primalValid {
		t.Fatalf(
			"expected added sources arcane (3, true) and primal (2, true), got arcane (%d, %t), primal (%d, %t)",
			arcane, arcaneValid,
			primal, primalValid,
		)
	}
}

func TestCasterLevelAddSourceLevel_KeepsSourcesIndependent(t *testing.T) {
	level := NewImpossibleCasterLevel()

	if ok := level.AddSourceLevel(DivineCasterSource, 3); !ok {
		t.Fatal("expected divine contribution to be accepted")
	}

	if ok := level.AddSourceLevel(PrimalCasterSource, 2); !ok {
		t.Fatal("expected primal contribution to be accepted")
	}

	if _, ok := level.GetSourceLevel(ArcaneCasterSource); ok {
		t.Fatal("expected arcane source to remain unavailable")
	}

	divine, divineValid := level.GetSourceLevel(DivineCasterSource)
	primal, primalValid := level.GetSourceLevel(PrimalCasterSource)
	if divine != 3 || !divineValid || primal != 2 || !primalValid {
		t.Fatalf(
			"expected sources divine (3, true) and primal (2, true), got divine (%d, %t), primal (%d, %t)",
			divine, divineValid,
			primal, primalValid,
		)
	}
}

func TestCasterLevelSetCasterLevel_RejectsNegativeWholeValue(t *testing.T) {
	level, ok := NewCasterLevel(1, 1, 1)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}

	invalid := CasterLevel{
		arcane: nullableInt{
			value: -1,
			valid: true,
		},
	}

	if ok := level.SetCasterLevel(invalid); ok {
		t.Fatal("expected negative caster level to be rejected")
	}

	arcane, _ := level.GetSourceLevel(ArcaneCasterSource)
	divine, _ := level.GetSourceLevel(DivineCasterSource)
	primal, _ := level.GetSourceLevel(PrimalCasterSource)
	if arcane != 1 || divine != 1 || primal != 1 {
		t.Fatalf("expected sources to remain unchanged, got (%d, %d, %d)", arcane, divine, primal)
	}
}

func TestCasterLevelAddSourceLevel_RejectsNegativeValues(t *testing.T) {
	level, ok := NewCasterLevel(2, 2, 2)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}

	if ok := level.AddSourceLevel(ArcaneCasterSource, -1); ok {
		t.Fatal("expected negative caster level contribution to be rejected")
	}

	arcane, _ := level.GetSourceLevel(ArcaneCasterSource)
	divine, _ := level.GetSourceLevel(DivineCasterSource)
	primal, _ := level.GetSourceLevel(PrimalCasterSource)
	if arcane != 2 || divine != 2 || primal != 2 {
		t.Fatalf("expected sources to remain unchanged, got (%d, %d, %d)", arcane, divine, primal)
	}
}

func TestCasterLevelSourceAPI_RejectsUnknownSource(t *testing.T) {
	level, ok := NewCasterLevel(1, 1, 1)
	if !ok {
		t.Fatal("expected caster level to be constructed")
	}

	if _, ok := level.GetSourceLevel(CasterSource("Mystic")); ok {
		t.Fatal("expected unknown source lookup to be rejected")
	}

	if ok := level.SetSourceLevel(CasterSource("Mystic"), 1); ok {
		t.Fatal("expected unknown source set to be rejected")
	}

	if ok := level.AddSourceLevel(CasterSource("Mystic"), 1); ok {
		t.Fatal("expected unknown source add to be rejected")
	}

	if ok := level.DisableSourceLevel(CasterSource("Mystic")); ok {
		t.Fatal("expected unknown source disable to be rejected")
	}
}

func TestNewCasterLevel_RejectsNegativeConstructorValues(t *testing.T) {
	if _, ok := NewCasterLevel(-1, 0, 0); ok {
		t.Fatal("expected invalid caster level to be rejected")
	}
}
