package race

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

const (
	DwarfRaceID    RaceID = "dwarf"
	ElfRaceID      RaceID = "elf"
	GnomeRaceID    RaceID = "gnome"
	HalfElfRaceID  RaceID = "half-elf"
	HalfOrcRaceID  RaceID = "half-orc"
	HalflingRaceID RaceID = "halfling"
	HumanRaceID    RaceID = "human"
)

var coreRaces = mustBuildCoreRaces()

func GetRaceByID(id RaceID) (Race, bool) {
	value, ok := coreRaces[id]
	if !ok {
		return race{}, false
	}

	return cloneRace(value), true
}

func mustBuildCoreRaces() map[RaceID]Race {
	return map[RaceID]Race{
		DwarfRaceID: mustNewRace(
			DwarfRaceID,
			ability.MediumSize,
			20,
			mustAbilityScoreModifiers(
				mustNewAbilityScoreModifier(ability.ConstitutionScore, 2),
				mustNewAbilityScoreModifier(ability.WisdomScore, 2),
				mustNewAbilityScoreModifier(ability.CharismaScore, -2),
			),
			0,
			[]LanguageID{"Common", "Dwarven"},
			[]RacialFeatureID{
				"Slow and Steady",
				"Darkvision",
				"Defensive Training",
				"Hardy",
				"Stability",
				"Greed",
				"Stonecunning",
				"Hatred",
				"Weapon Familiarity",
			},
		),
		ElfRaceID: mustNewRace(
			ElfRaceID,
			ability.MediumSize,
			30,
			mustAbilityScoreModifiers(
				mustNewAbilityScoreModifier(ability.DexterityScore, 2),
				mustNewAbilityScoreModifier(ability.IntelligenceScore, 2),
				mustNewAbilityScoreModifier(ability.ConstitutionScore, -2),
			),
			0,
			[]LanguageID{"Common", "Elven"},
			[]RacialFeatureID{
				"Medium",
				"Normal Speed",
				"Low-Light Vision",
				"Elven Immunities",
				"Keen Senses",
				"Weapon Familiarity",
				"Magic",
			},
		),
		GnomeRaceID: mustNewRace(
			GnomeRaceID,
			ability.SmallSize,
			20,
			mustAbilityScoreModifiers(
				mustNewAbilityScoreModifier(ability.ConstitutionScore, 2),
				mustNewAbilityScoreModifier(ability.CharismaScore, 2),
				mustNewAbilityScoreModifier(ability.StrengthScore, -2),
			),
			0,
			[]LanguageID{"Common", "Gnome", "Sylvan"},
			[]RacialFeatureID{
				"Small",
				"Slow Speed",
				"Low-Light Vision",
				"Defensive Training",
				"Illusion Resistance",
				"Hatred",
				"Weapon Familiarity",
				"Obsessive",
				"Gnome Magic",
				"Keen Senses",
			},
		),
		HalfElfRaceID: mustNewRace(
			HalfElfRaceID,
			ability.MediumSize,
			30,
			nil,
			2,
			[]LanguageID{"Common", "Elven"},
			[]RacialFeatureID{
				"Adaptability",
				"Elf Blood",
				"Low-Light Vision",
				"Elven Immunities",
				"Keen Senses",
				"Multitalented",
			},
		),
		HalfOrcRaceID: mustNewRace(
			HalfOrcRaceID,
			ability.MediumSize,
			30,
			nil,
			2,
			[]LanguageID{"Common", "Orc"},
			[]RacialFeatureID{
				"Orc Blood",
				"Darkvision",
				"Orc Ferocity",
				"Weapon Familiarity",
				"Intimidating",
			},
		),
		HalflingRaceID: mustNewRace(
			HalflingRaceID,
			ability.SmallSize,
			20,
			mustAbilityScoreModifiers(
				mustNewAbilityScoreModifier(ability.DexterityScore, 2),
				mustNewAbilityScoreModifier(ability.CharismaScore, 2),
				mustNewAbilityScoreModifier(ability.StrengthScore, -2),
			),
			0,
			[]LanguageID{"Common", "Halfling"},
			[]RacialFeatureID{
				"Small",
				"Slow Speed",
				"Fearless",
				"Halfling Luck",
				"Keen Senses",
				"Sure-Footed",
				"Weapon Familiarity",
			},
		),
		HumanRaceID: mustNewRace(
			HumanRaceID,
			ability.MediumSize,
			30,
			nil,
			2,
			[]LanguageID{"Common"},
			[]RacialFeatureID{
				"Bonus Feat",
				"Skilled",
			},
		),
	}
}

func cloneRace(value Race) Race {
	return race{
		id:                             value.id,
		size:                           value.size,
		baseSpeed:                      value.baseSpeed,
		abilityScoreModifiers:          append([]AbilityScoreModifier(nil), value.abilityScoreModifiers...),
		selectableAbilityScoreModifier: value.selectableAbilityScoreModifier,
		racialLanguages:                append([]LanguageID(nil), value.racialLanguages...),
		racialFeatures:                 append([]RacialFeatureID(nil), value.racialFeatures...),
	}
}

func mustNewRace(
	id RaceID,
	size ability.Size,
	baseSpeed int,
	abilityScoreModifiers []AbilityScoreModifier,
	selectableAbilityScoreModifier int,
	racialLanguages []LanguageID,
	racialFeatures []RacialFeatureID,
) Race {
	race, ok := NewRace(
		id,
		size,
		baseSpeed,
		abilityScoreModifiers,
		selectableAbilityScoreModifier,
		racialLanguages,
		racialFeatures,
	)
	if !ok {
		panic("invalid core race seed")
	}

	return race
}

func mustAbilityScoreModifiers(modifiers ...AbilityScoreModifier) []AbilityScoreModifier {
	return modifiers
}

func mustNewAbilityScoreModifier(scoreID ability.AbilityScoreID, modifier int) AbilityScoreModifier {
	value, ok := NewAbilityScoreModifier(scoreID, modifier)
	if !ok {
		panic("invalid core race ability score modifier")
	}

	return value
}
