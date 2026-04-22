package ability

func IsCoreSize(value Size) bool {
	switch value {
	case FineSize,
		DiminutiveSize,
		TinySize,
		SmallSize,
		MediumSize,
		LargeSize,
		HugeSize,
		GargantuanSize,
		ColossalSize:
		return true
	default:
		return false
	}
}

func IsProjectHouseRuleSize(value Size) bool {
	return value == TitanicSize
}

func IsProjectConstructBonusHPTableSize(value Size) bool {
	return value == TitanicSize
}

func IsCoreCasterSource(source CasterSource) bool {
	switch source {
	case ArcaneCasterSource, DivineCasterSource:
		return true
	default:
		return false
	}
}

func IsProjectHouseRuleCasterSource(source CasterSource) bool {
	return source == PrimalCasterSource
}
