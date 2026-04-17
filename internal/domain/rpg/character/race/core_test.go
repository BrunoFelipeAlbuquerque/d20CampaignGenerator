package race

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
)

func TestCoreRaces_SeedsSevenCoreEntries(t *testing.T) {
	if len(coreRaces) != 7 {
		t.Fatalf("expected 7 core races, got %d", len(coreRaces))
	}

	testCases := []struct {
		id          RaceID
		size        ability.Size
		baseSpeed   int
		languages   []LanguageID
		featureName RacialFeatureID
	}{
		{DwarfRaceID, ability.MediumSize, 20, []LanguageID{"Common", "Dwarven"}, "Stonecunning"},
		{ElfRaceID, ability.MediumSize, 30, []LanguageID{"Common", "Elven"}, "Elven Immunities"},
		{GnomeRaceID, ability.SmallSize, 20, []LanguageID{"Common", "Gnome", "Sylvan"}, "Gnome Magic"},
		{HalfElfRaceID, ability.MediumSize, 30, []LanguageID{"Common", "Elven"}, "Multitalented"},
		{HalfOrcRaceID, ability.MediumSize, 30, []LanguageID{"Common", "Orc"}, "Intimidating"},
		{HalflingRaceID, ability.SmallSize, 20, []LanguageID{"Common", "Halfling"}, "Halfling Luck"},
		{HumanRaceID, ability.MediumSize, 30, []LanguageID{"Common"}, "Bonus Feat"},
	}

	for _, tc := range testCases {
		race, ok := coreRaces[tc.id]
		if !ok {
			t.Fatalf("expected core race %q to exist", tc.id)
		}

		if race.GetID() != tc.id {
			t.Fatalf("expected race id %q, got %q", tc.id, race.GetID())
		}

		if race.GetSize() != tc.size {
			t.Fatalf("expected race %q size %q, got %q", tc.id, tc.size, race.GetSize())
		}

		if race.GetBaseSpeed() != tc.baseSpeed {
			t.Fatalf("expected race %q base speed %d, got %d", tc.id, tc.baseSpeed, race.GetBaseSpeed())
		}

		languages := race.GetRacialLanguages()
		if len(languages) != len(tc.languages) {
			t.Fatalf("expected race %q to have %d languages, got %d", tc.id, len(tc.languages), len(languages))
		}

		for i, language := range tc.languages {
			if languages[i] != language {
				t.Fatalf("expected race %q language at %d to be %q, got %q", tc.id, i, language, languages[i])
			}
		}

		if !race.HasRacialFeature(tc.featureName) {
			t.Fatalf("expected race %q to have feature %q", tc.id, tc.featureName)
		}
	}
}

func TestCoreRaces_SeedsAbilityScoreModifiersWhereFixed(t *testing.T) {
	testCases := []struct {
		id        RaceID
		modifiers map[ability.AbilityScoreID]int
	}{
		{
			id: DwarfRaceID,
			modifiers: map[ability.AbilityScoreID]int{
				ability.ConstitutionScore: 2,
				ability.WisdomScore:       2,
				ability.CharismaScore:     -2,
			},
		},
		{
			id: ElfRaceID,
			modifiers: map[ability.AbilityScoreID]int{
				ability.DexterityScore:    2,
				ability.IntelligenceScore: 2,
				ability.ConstitutionScore: -2,
			},
		},
		{
			id: GnomeRaceID,
			modifiers: map[ability.AbilityScoreID]int{
				ability.ConstitutionScore: 2,
				ability.CharismaScore:     2,
				ability.StrengthScore:     -2,
			},
		},
		{
			id: HalflingRaceID,
			modifiers: map[ability.AbilityScoreID]int{
				ability.DexterityScore: 2,
				ability.CharismaScore:  2,
				ability.StrengthScore:  -2,
			},
		},
	}

	for _, tc := range testCases {
		race := coreRaces[tc.id]
		modifiers := race.GetAbilityScoreModifiers()

		if len(modifiers) != len(tc.modifiers) {
			t.Fatalf("expected race %q to have %d ability score modifiers, got %d", tc.id, len(tc.modifiers), len(modifiers))
		}

		actual := make(map[ability.AbilityScoreID]int, len(modifiers))
		for _, modifier := range modifiers {
			actual[modifier.GetScoreID()] = modifier.GetModifier()
		}

		for scoreID, expected := range tc.modifiers {
			if actual[scoreID] != expected {
				t.Fatalf("expected race %q modifier %q to be %d, got %d", tc.id, scoreID, expected, actual[scoreID])
			}
		}
	}
}

func TestCoreRaces_UsesFlexibleAbilityBonusFeatureForVariableBonuses(t *testing.T) {
	testCases := []RaceID{HalfElfRaceID, HalfOrcRaceID, HumanRaceID}

	for _, raceID := range testCases {
		race := coreRaces[raceID]

		if len(race.GetAbilityScoreModifiers()) != 0 {
			t.Fatalf("expected race %q to have no fixed ability score modifiers in this seed, got %d", raceID, len(race.GetAbilityScoreModifiers()))
		}

		if !race.HasRacialFeature("Flexible Ability Bonus") {
			t.Fatalf("expected race %q to carry Flexible Ability Bonus marker", raceID)
		}
	}
}
