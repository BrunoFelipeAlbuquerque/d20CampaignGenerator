package character

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/ability"
	characterfeat "d20campaigngenerator/internal/domain/rpg/character/feat"
	characterrace "d20campaigngenerator/internal/domain/rpg/character/race"
)

func TestNewFixedRacialCharacterAbilityScores_ComposesFixedRaceModifiers(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.DwarfRaceID)

	scores, ok := NewFixedRacialCharacterAbilityScores(
		selectedRace,
		mustNewBaseCharacterAbilityScoresForTest(t, 10),
	)
	if !ok {
		t.Fatal("expected dwarf fixed racial ability modifiers to compose")
	}

	assertCharacterAbilityScoresForTest(t, scores, map[ability.AbilityScoreID]int{
		ability.StrengthScore:     10,
		ability.DexterityScore:    10,
		ability.ConstitutionScore: 12,
		ability.IntelligenceScore: 10,
		ability.WisdomScore:       12,
		ability.CharismaScore:     8,
	})
}

func TestNewFixedRacialCharacterAbilityScores_ExposesCharacterAbilityScoreFacts(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.ElfRaceID)

	scores, ok := NewFixedRacialCharacterAbilityScores(
		selectedRace,
		[]CharacterAbilityScore{
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, 11),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, 11),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.CharismaScore, 10),
		},
	)
	if !ok {
		t.Fatal("expected elf fixed racial ability modifiers to compose")
	}

	state, ok := NewCharacterFeatPrerequisiteState(scores, 0, nil, nil, nil, nil, nil)
	if !ok {
		t.Fatal("expected composed ability scores to feed feat prerequisite state")
	}

	dexterityPrerequisite, ok := characterfeat.NewAbilityScorePrerequisite(ability.DexterityScore, 13)
	if !ok {
		t.Fatal("expected dexterity prerequisite to construct")
	}

	if !state.SatisfiesPrerequisite(dexterityPrerequisite) {
		t.Fatal("expected composed elf dexterity score to satisfy a character feat prerequisite")
	}
}

func TestNewFixedRacialCharacterAbilityScores_RejectsSelectableModifierRace(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.HumanRaceID)

	if _, ok := NewFixedRacialCharacterAbilityScores(
		selectedRace,
		mustNewBaseCharacterAbilityScoresForTest(t, 10),
	); ok {
		t.Fatal("expected fixed-only path to reject a selectable-modifier race")
	}
}

func TestNewFixedRacialCharacterAbilityScores_RejectsInvalidBaseScores(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.DwarfRaceID)

	tests := []struct {
		name   string
		scores []CharacterAbilityScore
	}{
		{
			name: "missing score",
			scores: []CharacterAbilityScore{
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, 10),
			},
		},
		{
			name: "duplicate score",
			scores: []CharacterAbilityScore{
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 11),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.CharismaScore, 10),
			},
		},
		{
			name: "malformed score",
			scores: []CharacterAbilityScore{
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, 10),
				mustNewCharacterAbilityScoreForAbilityTest(t, ability.CharismaScore, 10),
				{id: ability.AbilityScoreID("LCK"), score: 10},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, ok := NewFixedRacialCharacterAbilityScores(selectedRace, test.scores); ok {
				t.Fatalf("expected %s to fail", test.name)
			}
		})
	}
}

func TestNewFixedRacialCharacterAbilityScores_RejectsZeroValueRace(t *testing.T) {
	var selectedRace CharacterRace

	if _, ok := NewFixedRacialCharacterAbilityScores(
		selectedRace,
		mustNewBaseCharacterAbilityScoresForTest(t, 10),
	); ok {
		t.Fatal("expected zero-value race to fail")
	}
}

func TestNewFixedRacialCharacterAbilityScores_RejectsInvalidComposedScores(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.HalflingRaceID)

	if _, ok := NewFixedRacialCharacterAbilityScores(
		selectedRace,
		[]CharacterAbilityScore{
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 1),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.CharismaScore, 10),
		},
	); ok {
		t.Fatal("expected racial modifiers that produce invalid ability scores to fail")
	}
}

func TestNewSelectableRacialCharacterAbilityScores_ComposesSelectableRaceModifiers(t *testing.T) {
	tests := []struct {
		name            string
		raceID          characterrace.RaceID
		selectedAbility ability.AbilityScoreID
		expectedScores  map[ability.AbilityScoreID]int
	}{
		{
			name:            "human",
			raceID:          characterrace.HumanRaceID,
			selectedAbility: ability.StrengthScore,
			expectedScores: map[ability.AbilityScoreID]int{
				ability.StrengthScore:     12,
				ability.DexterityScore:    10,
				ability.ConstitutionScore: 10,
				ability.IntelligenceScore: 10,
				ability.WisdomScore:       10,
				ability.CharismaScore:     10,
			},
		},
		{
			name:            "half-elf",
			raceID:          characterrace.HalfElfRaceID,
			selectedAbility: ability.IntelligenceScore,
			expectedScores: map[ability.AbilityScoreID]int{
				ability.StrengthScore:     10,
				ability.DexterityScore:    10,
				ability.ConstitutionScore: 10,
				ability.IntelligenceScore: 12,
				ability.WisdomScore:       10,
				ability.CharismaScore:     10,
			},
		},
		{
			name:            "half-orc",
			raceID:          characterrace.HalfOrcRaceID,
			selectedAbility: ability.CharismaScore,
			expectedScores: map[ability.AbilityScoreID]int{
				ability.StrengthScore:     10,
				ability.DexterityScore:    10,
				ability.ConstitutionScore: 10,
				ability.IntelligenceScore: 10,
				ability.WisdomScore:       10,
				ability.CharismaScore:     12,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scores, ok := NewSelectableRacialCharacterAbilityScores(
				mustNewCharacterRaceForAbilityTest(t, test.raceID),
				mustNewBaseCharacterAbilityScoresForTest(t, 10),
				[]CharacterSelectedAbilityScore{
					mustNewCharacterSelectedAbilityScoreForTest(t, test.selectedAbility),
				},
			)
			if !ok {
				t.Fatalf("expected %s selectable racial ability modifier to compose", test.name)
			}

			assertCharacterAbilityScoresForTest(t, scores, test.expectedScores)
		})
	}
}

func TestNewSelectableRacialCharacterAbilityScores_ExposesCharacterAbilityScoreFacts(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.HumanRaceID)

	scores, ok := NewSelectableRacialCharacterAbilityScores(
		selectedRace,
		[]CharacterAbilityScore{
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 11),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.CharismaScore, 10),
		},
		[]CharacterSelectedAbilityScore{
			mustNewCharacterSelectedAbilityScoreForTest(t, ability.StrengthScore),
		},
	)
	if !ok {
		t.Fatal("expected human selectable racial ability modifier to compose")
	}

	state, ok := NewCharacterFeatPrerequisiteState(scores, 1, nil, nil, nil, nil, nil)
	if !ok {
		t.Fatal("expected composed selectable ability scores to feed feat prerequisite state")
	}

	powerAttack, ok := characterfeat.GetFeatByID(characterfeat.PowerAttackFeatID)
	if !ok {
		t.Fatal("expected power attack to resolve")
	}

	if !state.SatisfiesFeat(powerAttack) {
		t.Fatal("expected composed selectable strength score to satisfy power attack")
	}
}

func TestNewSelectableRacialCharacterAbilityScores_RejectsFixedModifierRace(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.DwarfRaceID)

	if _, ok := NewSelectableRacialCharacterAbilityScores(
		selectedRace,
		mustNewBaseCharacterAbilityScoresForTest(t, 10),
		[]CharacterSelectedAbilityScore{mustNewCharacterSelectedAbilityScoreForTest(t, ability.StrengthScore)},
	); ok {
		t.Fatal("expected selectable-only path to reject a fixed-modifier race")
	}
}

func TestNewSelectableRacialCharacterAbilityScores_RejectsInvalidSelectedAbilities(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.HumanRaceID)

	tests := []struct {
		name              string
		selectedAbilities []CharacterSelectedAbilityScore
	}{
		{
			name:              "missing",
			selectedAbilities: nil,
		},
		{
			name: "duplicate",
			selectedAbilities: []CharacterSelectedAbilityScore{
				mustNewCharacterSelectedAbilityScoreForTest(t, ability.StrengthScore),
				mustNewCharacterSelectedAbilityScoreForTest(t, ability.StrengthScore),
			},
		},
		{
			name: "unknown",
			selectedAbilities: []CharacterSelectedAbilityScore{
				{id: ability.AbilityScoreID("LCK"), valid: true},
			},
		},
		{
			name: "zero value",
			selectedAbilities: []CharacterSelectedAbilityScore{
				{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, ok := NewSelectableRacialCharacterAbilityScores(
				selectedRace,
				mustNewBaseCharacterAbilityScoresForTest(t, 10),
				test.selectedAbilities,
			); ok {
				t.Fatalf("expected %s selected ability input to fail", test.name)
			}
		})
	}
}

func TestNewSelectableRacialCharacterAbilityScores_RejectsInvalidBaseScores(t *testing.T) {
	selectedRace := mustNewCharacterRaceForAbilityTest(t, characterrace.HumanRaceID)

	if _, ok := NewSelectableRacialCharacterAbilityScores(
		selectedRace,
		[]CharacterAbilityScore{
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, 10),
			mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, 10),
			{id: ability.AbilityScoreID("LCK"), score: 10},
		},
		[]CharacterSelectedAbilityScore{
			mustNewCharacterSelectedAbilityScoreForTest(t, ability.StrengthScore),
		},
	); ok {
		t.Fatal("expected malformed base ability score input to fail")
	}
}

func TestNewSelectableRacialCharacterAbilityScores_RejectsZeroValueRace(t *testing.T) {
	var selectedRace CharacterRace

	if _, ok := NewSelectableRacialCharacterAbilityScores(
		selectedRace,
		mustNewBaseCharacterAbilityScoresForTest(t, 10),
		[]CharacterSelectedAbilityScore{mustNewCharacterSelectedAbilityScoreForTest(t, ability.StrengthScore)},
	); ok {
		t.Fatal("expected zero-value race to fail")
	}
}

func mustNewCharacterRaceForAbilityTest(
	t *testing.T,
	id characterrace.RaceID,
) CharacterRace {
	t.Helper()

	selectedRace, ok := NewCharacterRace(id)
	if !ok {
		t.Fatalf("expected character race %q to compose", id)
	}

	return selectedRace
}

func mustNewBaseCharacterAbilityScoresForTest(t *testing.T, score int) []CharacterAbilityScore {
	t.Helper()

	return []CharacterAbilityScore{
		mustNewCharacterAbilityScoreForAbilityTest(t, ability.StrengthScore, score),
		mustNewCharacterAbilityScoreForAbilityTest(t, ability.DexterityScore, score),
		mustNewCharacterAbilityScoreForAbilityTest(t, ability.ConstitutionScore, score),
		mustNewCharacterAbilityScoreForAbilityTest(t, ability.IntelligenceScore, score),
		mustNewCharacterAbilityScoreForAbilityTest(t, ability.WisdomScore, score),
		mustNewCharacterAbilityScoreForAbilityTest(t, ability.CharismaScore, score),
	}
}

func mustNewCharacterAbilityScoreForAbilityTest(
	t *testing.T,
	id ability.AbilityScoreID,
	score int,
) CharacterAbilityScore {
	t.Helper()

	value, ok := NewCharacterAbilityScore(id, score)
	if !ok {
		t.Fatalf("expected ability score %q %d to compose", id, score)
	}

	return value
}

func mustNewCharacterSelectedAbilityScoreForTest(
	t *testing.T,
	id ability.AbilityScoreID,
) CharacterSelectedAbilityScore {
	t.Helper()

	value, ok := NewCharacterSelectedAbilityScore(id)
	if !ok {
		t.Fatalf("expected selected ability score %q to compose", id)
	}

	return value
}

func assertCharacterAbilityScoresForTest(
	t *testing.T,
	scores []CharacterAbilityScore,
	expected map[ability.AbilityScoreID]int,
) {
	t.Helper()

	if len(scores) != len(expected) {
		t.Fatalf("expected %d ability scores, got %d", len(expected), len(scores))
	}

	actual := make(map[ability.AbilityScoreID]int, len(scores))
	for _, score := range scores {
		actual[score.GetAbilityScoreID()] = score.GetScore()
	}

	for scoreID, expectedScore := range expected {
		if actual[scoreID] != expectedScore {
			t.Fatalf("expected %q score %d, got %d", scoreID, expectedScore, actual[scoreID])
		}
	}
}
