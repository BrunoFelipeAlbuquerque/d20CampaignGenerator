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
		bonusChoice BonusLanguageChoice
		featureName RacialFeatureID
	}{
		{DwarfRaceID, ability.MediumSize, 20, []LanguageID{CommonLanguageID, DwarvenLanguageID}, mustCoreBonusLanguageChoice(t, []LanguageID{GiantLanguageID, GnomeLanguageID, GoblinLanguageID, OrcLanguageID, TerranLanguageID, UndercommonLanguageID}, false), StonecunningFeatureID},
		{ElfRaceID, ability.MediumSize, 30, []LanguageID{CommonLanguageID, ElvenLanguageID}, mustCoreBonusLanguageChoice(t, []LanguageID{CelestialLanguageID, DraconicLanguageID, GnollLanguageID, GnomeLanguageID, GoblinLanguageID, OrcLanguageID, SylvanLanguageID}, false), ElvenImmunitiesFeatureID},
		{GnomeRaceID, ability.SmallSize, 20, []LanguageID{CommonLanguageID, GnomeLanguageID, SylvanLanguageID}, mustCoreBonusLanguageChoice(t, []LanguageID{DraconicLanguageID, DwarvenLanguageID, ElvenLanguageID, GiantLanguageID, GoblinLanguageID, OrcLanguageID}, false), GnomeMagicFeatureID},
		{HalfElfRaceID, ability.MediumSize, 30, []LanguageID{CommonLanguageID, ElvenLanguageID}, mustCoreBonusLanguageChoice(t, nil, true), MultitalentedFeatureID},
		{HalfOrcRaceID, ability.MediumSize, 30, []LanguageID{CommonLanguageID, OrcLanguageID}, mustCoreBonusLanguageChoice(t, []LanguageID{AbyssalLanguageID, DraconicLanguageID, GiantLanguageID, GnollLanguageID, GoblinLanguageID}, false), IntimidatingFeatureID},
		{HalflingRaceID, ability.SmallSize, 20, []LanguageID{CommonLanguageID, HalflingLanguageID}, mustCoreBonusLanguageChoice(t, []LanguageID{DwarvenLanguageID, ElvenLanguageID, GnomeLanguageID, GoblinLanguageID, OrcLanguageID}, false), HalflingLuckFeatureID},
		{HumanRaceID, ability.MediumSize, 30, []LanguageID{CommonLanguageID}, mustCoreBonusLanguageChoice(t, nil, true), BonusFeatFeatureID},
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

		languages := race.GetAutomaticLanguages()
		if len(languages) != len(tc.languages) {
			t.Fatalf("expected race %q to have %d automatic languages, got %d", tc.id, len(tc.languages), len(languages))
		}

		for i, language := range tc.languages {
			if languages[i] != language {
				t.Fatalf("expected race %q automatic language at %d to be %q, got %q", tc.id, i, language, languages[i])
			}
		}

		bonusChoice, ok := race.GetBonusLanguageChoice()
		if !ok {
			t.Fatalf("expected race %q bonus language choice metadata", tc.id)
		}

		if bonusChoice.AllowsAnyNonSecret() != tc.bonusChoice.AllowsAnyNonSecret() {
			t.Fatalf("expected race %q any-non-secret=%t, got %t", tc.id, tc.bonusChoice.AllowsAnyNonSecret(), bonusChoice.AllowsAnyNonSecret())
		}

		actualBonusLanguages := bonusChoice.GetLanguageIDs()
		expectedBonusLanguages := tc.bonusChoice.GetLanguageIDs()
		if len(actualBonusLanguages) != len(expectedBonusLanguages) {
			t.Fatalf("expected race %q to have %d bonus languages, got %d", tc.id, len(expectedBonusLanguages), len(actualBonusLanguages))
		}

		for i, language := range expectedBonusLanguages {
			if actualBonusLanguages[i] != language {
				t.Fatalf("expected race %q bonus language at %d to be %q, got %q", tc.id, i, language, actualBonusLanguages[i])
			}
		}

		if !race.HasFeature(tc.featureName) {
			t.Fatalf("expected race %q to have feature %q", tc.id, tc.featureName)
		}
	}
}

func TestCoreRaces_SeedsCorrectedCoreFeaturePresence(t *testing.T) {
	testCases := []struct {
		id       RaceID
		features []RacialFeatureID
	}{
		{ElfRaceID, []RacialFeatureID{ElvenImmunitiesFeatureID, ElvenMagicFeatureID}},
		{GnomeRaceID, []RacialFeatureID{GnomeMagicFeatureID, KeenSensesFeatureID}},
		{HalfElfRaceID, []RacialFeatureID{ElvenImmunitiesFeatureID, KeenSensesFeatureID, MultitalentedFeatureID}},
		{HalfOrcRaceID, []RacialFeatureID{OrcFerocityFeatureID, IntimidatingFeatureID}},
	}

	for _, tc := range testCases {
		race := coreRaces[tc.id]

		for _, featureID := range tc.features {
			if !race.HasFeature(featureID) {
				t.Fatalf("expected race %q to have feature %q", tc.id, featureID)
			}
		}
	}
}

func TestCoreRaces_DoesNotDuplicateStructuralFactsAsFeatures(t *testing.T) {
	structuralFeatureMarkers := []RacialFeatureID{
		"Medium",
		"Small",
		"Normal Speed",
		"Slow Speed",
	}

	for raceID, race := range coreRaces {
		for _, featureID := range structuralFeatureMarkers {
			if race.HasFeature(featureID) {
				t.Fatalf("expected race %q not to encode structural fact %q as a racial feature", raceID, featureID)
			}
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

func TestCoreRaces_SeedsSelectableAbilityScoreModifierForVariableBonuses(t *testing.T) {
	testCases := []RaceID{HalfElfRaceID, HalfOrcRaceID, HumanRaceID}

	for _, raceID := range testCases {
		race := coreRaces[raceID]

		if len(race.GetAbilityScoreModifiers()) != 0 {
			t.Fatalf("expected race %q to have no fixed ability score modifiers in this seed, got %d", raceID, len(race.GetAbilityScoreModifiers()))
		}

		selectableModifier, ok := race.GetSelectableAbilityScoreModifier()
		if !ok {
			t.Fatalf("expected race %q to expose selectable ability score modifier metadata", raceID)
		}

		if selectableModifier != 2 {
			t.Fatalf("expected race %q selectable ability score modifier to be 2, got %d", raceID, selectableModifier)
		}

		if race.HasFeature(RacialFeatureID("Flexible Ability Bonus")) {
			t.Fatalf("expected race %q variable ability bonus not to be encoded as a feature marker", raceID)
		}
	}
}

func TestGetRaceByID_ReturnsSeededCoreRace(t *testing.T) {
	race, ok := GetRaceByID(ElfRaceID)
	if !ok {
		t.Fatal("expected elf to be returned from core race lookup")
	}

	if race.GetID() != ElfRaceID {
		t.Fatalf("expected race id %q, got %q", ElfRaceID, race.GetID())
	}

	if !race.HasFeature(ElvenImmunitiesFeatureID) {
		t.Fatal("expected looked up elf to expose feature queries")
	}
}

func TestGetRaceByID_ReturnsSelectableAbilityScoreModifierMetadata(t *testing.T) {
	race, ok := GetRaceByID(HumanRaceID)
	if !ok {
		t.Fatal("expected human to be returned from core race lookup")
	}

	selectableModifier, ok := race.GetSelectableAbilityScoreModifier()
	if !ok {
		t.Fatal("expected looked up human to expose selectable ability score modifier metadata")
	}

	if selectableModifier != 2 {
		t.Fatalf("expected looked up human selectable ability score modifier to be 2, got %d", selectableModifier)
	}
}

func TestGetRaceByID_ReturnsDetachedCopy(t *testing.T) {
	first, ok := GetRaceByID(DwarfRaceID)
	if !ok {
		t.Fatal("expected dwarf to be returned from core race lookup")
	}

	first.abilityScoreModifiers[0].modifier = 99
	first.automaticLanguages[0] = "Changed"
	first.bonusLanguageChoice.languageIDs[0] = "Changed"
	first.racialFeatures[0] = "Changed"

	second, ok := GetRaceByID(DwarfRaceID)
	if !ok {
		t.Fatal("expected dwarf to be returned from core race lookup")
	}

	if second.abilityScoreModifiers[0].modifier != 2 {
		t.Fatalf("expected stored dwarf constitution modifier to remain 2, got %d", second.abilityScoreModifiers[0].modifier)
	}

	if second.automaticLanguages[0] != CommonLanguageID {
		t.Fatalf("expected stored dwarf language to remain Common, got %q", second.automaticLanguages[0])
	}

	if second.bonusLanguageChoice.languageIDs[0] != GiantLanguageID {
		t.Fatalf("expected stored dwarf bonus language to remain Giant, got %q", second.bonusLanguageChoice.languageIDs[0])
	}

	if second.racialFeatures[0] != SlowAndSteadyFeatureID {
		t.Fatalf("expected stored dwarf feature to remain Slow and Steady, got %q", second.racialFeatures[0])
	}
}

func TestGetRaceByID_RejectsUnknownRace(t *testing.T) {
	if _, ok := GetRaceByID(RaceID("android")); ok {
		t.Fatal("expected unknown race lookup to fail")
	}
}

func TestGetRaces_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	races := GetRaces()
	if len(races) != len(coreRaceOrder) {
		t.Fatalf("expected %d queried races, got %d", len(coreRaceOrder), len(races))
	}

	for i, expectedID := range coreRaceOrder {
		if races[i].GetID() != expectedID {
			t.Fatalf("expected race at index %d to be %q, got %q", i, expectedID, races[i].GetID())
		}
	}
}

func TestGetRaces_ReturnsDetachedCopies(t *testing.T) {
	first := GetRaces()
	second := GetRaces()

	first[0].automaticLanguages[0] = "Changed"
	first[0].bonusLanguageChoice.languageIDs[0] = "Changed"
	first[0].racialFeatures[0] = "Changed"
	first[0].abilityScoreModifiers[0].modifier = 99

	if second[0].automaticLanguages[0] != CommonLanguageID {
		t.Fatalf("expected stored race language to remain %q, got %q", CommonLanguageID, second[0].automaticLanguages[0])
	}

	if second[0].bonusLanguageChoice.languageIDs[0] != GiantLanguageID {
		t.Fatalf("expected stored race bonus language to remain %q, got %q", GiantLanguageID, second[0].bonusLanguageChoice.languageIDs[0])
	}

	if second[0].racialFeatures[0] != SlowAndSteadyFeatureID {
		t.Fatalf("expected stored race feature to remain %q, got %q", SlowAndSteadyFeatureID, second[0].racialFeatures[0])
	}

	if second[0].abilityScoreModifiers[0].modifier != 2 {
		t.Fatalf("expected stored race modifier to remain 2, got %d", second[0].abilityScoreModifiers[0].modifier)
	}
}

func mustCoreBonusLanguageChoice(t *testing.T, languageIDs []LanguageID, anyNonSecret bool) BonusLanguageChoice {
	t.Helper()

	value, ok := NewBonusLanguageChoice(languageIDs, anyNonSecret)
	if !ok {
		t.Fatal("expected core bonus language choice to be constructed")
	}

	return value
}
