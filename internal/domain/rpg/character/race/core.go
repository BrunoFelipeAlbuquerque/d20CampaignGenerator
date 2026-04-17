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
			},
		),
		HalfElfRaceID: mustNewRace(
			HalfElfRaceID,
			ability.MediumSize,
			30,
			nil,
			[]LanguageID{"Common", "Elven"},
			[]RacialFeatureID{
				"Adaptability",
				"Elf Blood",
				"Multitalented",
				"Low-Light Vision",
				"Flexible Ability Bonus",
			},
		),
		HalfOrcRaceID: mustNewRace(
			HalfOrcRaceID,
			ability.MediumSize,
			30,
			nil,
			[]LanguageID{"Common", "Orc"},
			[]RacialFeatureID{
				"Orc Blood",
				"Darkvision",
				"Weapon Familiarity",
				"Intimidating",
				"Flexible Ability Bonus",
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
			[]LanguageID{"Common"},
			[]RacialFeatureID{
				"Bonus Feat",
				"Skilled",
				"Flexible Ability Bonus",
			},
		),
	}
}

func mustNewRace(
	id RaceID,
	size ability.Size,
	baseSpeed int,
	abilityScoreModifiers []AbilityScoreModifier,
	racialLanguages []LanguageID,
	racialFeatures []RacialFeatureID,
) Race {
	race, ok := NewRace(id, size, baseSpeed, abilityScoreModifiers, racialLanguages, racialFeatures)
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
