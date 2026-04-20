package ability

import "testing"

func TestIsCoreSize_ExcludesProjectHouseRuleSize(t *testing.T) {
	if !IsCoreSize(MediumSize) {
		t.Fatal("expected medium to be core size")
	}

	if IsCoreSize(TitanicSize) {
		t.Fatal("expected titanic to be non-core size")
	}
}

func TestIsProjectConstructBonusHPTableSize_SeparatesCustomTable(t *testing.T) {
	if !IsProjectConstructBonusHPTableSize(ColossalSize) {
		t.Fatal("expected colossal to use custom construct bonus HP table")
	}

	if !IsProjectConstructBonusHPTableSize(TitanicSize) {
		t.Fatal("expected titanic to use custom construct bonus HP table")
	}

	if IsProjectConstructBonusHPTableSize(MediumSize) {
		t.Fatal("expected medium to not use custom construct bonus HP table")
	}
}

func TestIsCoreCasterSource_ExcludesProjectHouseRuleSource(t *testing.T) {
	if !IsCoreCasterSource(ArcaneCasterSource) {
		t.Fatal("expected arcane source to be core")
	}

	if !IsCoreCasterSource(DivineCasterSource) {
		t.Fatal("expected divine source to be core")
	}

	if IsCoreCasterSource(PrimalCasterSource) {
		t.Fatal("expected primal source to be non-core")
	}
}
