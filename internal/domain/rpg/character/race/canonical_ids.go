package race

const (
	CommonLanguageID   LanguageID = "Common"
	DwarvenLanguageID  LanguageID = "Dwarven"
	ElvenLanguageID    LanguageID = "Elven"
	GnomeLanguageID    LanguageID = "Gnome"
	SylvanLanguageID   LanguageID = "Sylvan"
	OrcLanguageID      LanguageID = "Orc"
	HalflingLanguageID LanguageID = "Halfling"
)

const (
	SlowAndSteadyFeatureID      RacialFeatureID = "Slow and Steady"
	DarkvisionFeatureID         RacialFeatureID = "Darkvision"
	DefensiveTrainingFeatureID  RacialFeatureID = "Defensive Training"
	HardyFeatureID              RacialFeatureID = "Hardy"
	StabilityFeatureID          RacialFeatureID = "Stability"
	GreedFeatureID              RacialFeatureID = "Greed"
	StonecunningFeatureID       RacialFeatureID = "Stonecunning"
	HatredFeatureID             RacialFeatureID = "Hatred"
	WeaponFamiliarityFeatureID  RacialFeatureID = "Weapon Familiarity"
	LowLightVisionFeatureID     RacialFeatureID = "Low-Light Vision"
	ElvenImmunitiesFeatureID    RacialFeatureID = "Elven Immunities"
	KeenSensesFeatureID         RacialFeatureID = "Keen Senses"
	ElvenMagicFeatureID         RacialFeatureID = "Elven Magic"
	MagicFeatureID              RacialFeatureID = ElvenMagicFeatureID
	IllusionResistanceFeatureID RacialFeatureID = "Illusion Resistance"
	ObsessiveFeatureID          RacialFeatureID = "Obsessive"
	GnomeMagicFeatureID         RacialFeatureID = "Gnome Magic"
	AdaptabilityFeatureID       RacialFeatureID = "Adaptability"
	ElfBloodFeatureID           RacialFeatureID = "Elf Blood"
	MultitalentedFeatureID      RacialFeatureID = "Multitalented"
	OrcBloodFeatureID           RacialFeatureID = "Orc Blood"
	OrcFerocityFeatureID        RacialFeatureID = "Orc Ferocity"
	IntimidatingFeatureID       RacialFeatureID = "Intimidating"
	FearlessFeatureID           RacialFeatureID = "Fearless"
	HalflingLuckFeatureID       RacialFeatureID = "Halfling Luck"
	SureFootedFeatureID         RacialFeatureID = "Sure-Footed"
	BonusFeatFeatureID          RacialFeatureID = "Bonus Feat"
	SkilledFeatureID            RacialFeatureID = "Skilled"
)

var validLanguageIDs = map[LanguageID]struct{}{
	CommonLanguageID:   {},
	DwarvenLanguageID:  {},
	ElvenLanguageID:    {},
	GnomeLanguageID:    {},
	SylvanLanguageID:   {},
	OrcLanguageID:      {},
	HalflingLanguageID: {},
}

var validRacialFeatureIDs = map[RacialFeatureID]struct{}{
	SlowAndSteadyFeatureID:      {},
	DarkvisionFeatureID:         {},
	DefensiveTrainingFeatureID:  {},
	HardyFeatureID:              {},
	StabilityFeatureID:          {},
	GreedFeatureID:              {},
	StonecunningFeatureID:       {},
	HatredFeatureID:             {},
	WeaponFamiliarityFeatureID:  {},
	LowLightVisionFeatureID:     {},
	ElvenImmunitiesFeatureID:    {},
	KeenSensesFeatureID:         {},
	ElvenMagicFeatureID:         {},
	IllusionResistanceFeatureID: {},
	ObsessiveFeatureID:          {},
	GnomeMagicFeatureID:         {},
	AdaptabilityFeatureID:       {},
	ElfBloodFeatureID:           {},
	MultitalentedFeatureID:      {},
	OrcBloodFeatureID:           {},
	OrcFerocityFeatureID:        {},
	IntimidatingFeatureID:       {},
	FearlessFeatureID:           {},
	HalflingLuckFeatureID:       {},
	SureFootedFeatureID:         {},
	BonusFeatFeatureID:          {},
	SkilledFeatureID:            {},
}
