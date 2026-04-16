package ability

import "testing"

func almostEqual(a, b float64) bool {
	const epsilon = 0.01

	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	return diff < epsilon
}

func mustNewAbilityScore(t *testing.T, id AbilityScoreID, value AbilityScoreValue) AbilityScore {
	t.Helper()

	score, ok := NewAbilityScore(id, value)
	if !ok {
		t.Fatalf("expected ability score %q to be constructed", id)
	}

	return score
}

// ==============================
// CONSTRUCTOR
// ==============================

func TestNewAbilityScore_DerivesCanonicalMetadata(t *testing.T) {
	score := mustNewAbilityScore(t,
		StrengthScore,
		abilityScoreValue{
			value: 18,
			valid: true,
		},
	)

	if score.GetID() != StrengthScore {
		t.Errorf("expected id %q, got %q", StrengthScore, score.GetID())
	}

	if score.GetName() != "Strength" {
		t.Errorf("expected name Strength, got %q", score.GetName())
	}

	value, valid := score.GetValue().GetValue()
	if !valid {
		t.Fatal("expected constructed value to be valid")
	}

	if value != 18 {
		t.Errorf("expected value 18, got %d", value)
	}
}

func TestNewAbilityScore_PreservesSuppressedStoredValue(t *testing.T) {
	score := mustNewAbilityScore(t,
		ConstitutionScore,
		abilityScoreValue{
			value: 14,
			valid: false,
		},
	)

	if score.GetName() != "Constitution" {
		t.Errorf("expected name Constitution, got %q", score.GetName())
	}

	value, valid := score.GetValue().GetValue()
	if valid {
		t.Fatal("expected constitution score to be suppressed")
	}

	if value != 14 {
		t.Errorf("expected stored value 14 to be preserved, got %d", value)
	}
}

func TestAbilityScoreIDGetName_AllCanonicalScores(t *testing.T) {
	tests := map[abilityScoreID]string{
		StrengthScore:     "Strength",
		DexterityScore:    "Dexterity",
		ConstitutionScore: "Constitution",
		IntelligenceScore: "Intelligence",
		WisdomScore:       "Wisdom",
		CharismaScore:     "Charisma",
	}

	for id, want := range tests {
		got := id.GetName()
		if got != want {
			t.Errorf("expected %q for %q, got %q", want, id, got)
		}
	}
}

func TestAbilityScoreIDGetName_UnknownID(t *testing.T) {
	got := abilityScoreID("???").GetName()
	if got != "" {
		t.Errorf("expected empty name for unknown id, got %q", got)
	}
}

func TestNewAbilityScore_RejectsUnknownID(t *testing.T) {
	if _, ok := NewAbilityScore(AbilityScoreID("???"), abilityScoreValue{value: 10, valid: true}); ok {
		t.Fatal("expected unknown ability score id to be rejected")
	}
}

// ==============================
// VALUE OBJECT
// ==============================

func TestAbilityScoreValue_GettersAndSetters(t *testing.T) {
	value := abilityScoreValue{
		value: 12,
		valid: true,
	}

	gotValue, gotValid := value.GetValue()
	if gotValue != 12 || !gotValid {
		t.Fatalf("expected (12, true), got (%d, %t)", gotValue, gotValid)
	}

	if !value.IsValid() {
		t.Fatal("expected value to report valid")
	}

	var ok bool
	value, ok = value.WithValue(8)
	if !ok {
		t.Fatal("expected ability score value update to succeed")
	}
	value = value.WithValid(false)

	gotValue, gotValid = value.GetValue()
	if gotValue != 8 || gotValid {
		t.Fatalf("expected (8, false), got (%d, %t)", gotValue, gotValid)
	}
}

func TestNewAbilityScoreValue_RejectsNegativeValues(t *testing.T) {
	if _, ok := NewAbilityScoreValue(-1, true); ok {
		t.Fatal("expected negative ability score value to be rejected")
	}
}

// ==============================
// ABILITY SCORE MUTATION
// ==============================

func TestAbilityScore_SetValue_ReplacesWholeValueObject(t *testing.T) {
	score := mustNewAbilityScore(t,
		DexterityScore,
		abilityScoreValue{
			value: 16,
			valid: true,
		},
	)

	score.SetValue(abilityScoreValue{
		value: 10,
		valid: false,
	})

	value, valid := score.GetValue().GetValue()
	if value != 10 || valid {
		t.Fatalf("expected (10, false), got (%d, %t)", value, valid)
	}
}

func TestAbilityScore_SetScoreValue_UpdatesStoredValueWithoutChangingState(t *testing.T) {
	score := mustNewAbilityScore(t,
		WisdomScore,
		abilityScoreValue{
			value: 12,
			valid: false,
		},
	)

	score.SetScoreValue(18)

	value, valid := score.GetValue().GetValue()
	if value != 18 {
		t.Errorf("expected stored value 18, got %d", value)
	}

	if valid {
		t.Error("expected validity flag to remain false")
	}
}

func TestAbilityScore_SetValueValidity_TogglesAvailabilityWithoutErasingStoredValue(t *testing.T) {
	score := mustNewAbilityScore(t,
		ConstitutionScore,
		abilityScoreValue{
			value: 14,
			valid: true,
		},
	)

	score.SetValueValidity(false)

	value, valid := score.GetValue().GetValue()
	if value != 14 {
		t.Errorf("expected stored value 14, got %d", value)
	}

	if valid {
		t.Fatal("expected value to become invalid")
	}

	score.SetValueValidity(true)

	value, valid = score.GetValue().GetValue()
	if value != 14 || !valid {
		t.Fatalf("expected (14, true), got (%d, %t)", value, valid)
	}
}

// ==============================
// MODIFIER
// ==============================

func TestAbilityScore_GetModifier_WhenScoreIsSuppressed(t *testing.T) {
	score := mustNewAbilityScore(t,
		ConstitutionScore,
		abilityScoreValue{
			value: 18,
			valid: false,
		},
	)

	got, ok := score.GetModifier()
	if ok {
		t.Fatal("expected no modifier for suppressed score")
	}

	if got != 0 {
		t.Errorf("expected zero fallback modifier, got %d", got)
	}
}

func TestAbilityScore_GetModifier_ZeroScoreStillExists(t *testing.T) {
	score := mustNewAbilityScore(t,
		ConstitutionScore,
		abilityScoreValue{
			value: 0,
			valid: true,
		},
	)

	got, ok := score.GetModifier()
	if !ok {
		t.Fatal("expected zero score to still produce a modifier")
	}

	if got != -5 {
		t.Errorf("expected modifier -5, got %d", got)
	}
}

func TestAbilityScore_GetModifier_RestoresAfterSuppressionEnds(t *testing.T) {
	score := mustNewAbilityScore(t,
		ConstitutionScore,
		abilityScoreValue{
			value: 14,
			valid: true,
		},
	)

	score.SetValueValidity(false)

	if _, ok := score.GetModifier(); ok {
		t.Fatal("expected no modifier while score is suppressed")
	}

	score.SetValueValidity(true)

	got, ok := score.GetModifier()
	if !ok {
		t.Fatal("expected modifier after score is restored")
	}

	if got != 2 {
		t.Errorf("expected modifier 2, got %d", got)
	}
}

func TestCalculateAbilityModifier_Table(t *testing.T) {
	tests := []struct {
		score int
		want  int
	}{
		{0, -5},
		{1, -5},
		{2, -4},
		{3, -4},
		{8, -1},
		{9, -1},
		{10, 0},
		{11, 0},
		{12, 1},
		{13, 1},
		{18, 4},
		{19, 4},
		{20, 5},
		{21, 5},
	}

	for _, tt := range tests {
		got := calculateAbilityModifier(tt.score)
		if got != tt.want {
			t.Errorf("score %d: expected %d, got %d", tt.score, tt.want, got)
		}
	}
}

// ==============================
// STRENGTH CARRYING CAPACITY
// ==============================

func TestAbilityScore_GetCarryingCapacity_StrengthOnly(t *testing.T) {
	score := mustNewAbilityScore(t,
		DexterityScore,
		abilityScoreValue{
			value: 18,
			valid: true,
		},
	)

	if _, ok := score.GetCarryingCapacity(); ok {
		t.Fatal("expected carrying capacity to be unavailable for non-strength scores")
	}
}

func TestAbilityScore_GetCarryingCapacity_SuppressedStrength(t *testing.T) {
	score := mustNewAbilityScore(t,
		StrengthScore,
		abilityScoreValue{
			value: 18,
			valid: false,
		},
	)

	if _, ok := score.GetCarryingCapacity(); ok {
		t.Fatal("expected carrying capacity to be unavailable for suppressed strength")
	}
}

func TestAbilityScore_GetCarryingCapacity_ZeroStrength(t *testing.T) {
	score := mustNewAbilityScore(t,
		StrengthScore,
		abilityScoreValue{
			value: 0,
			valid: true,
		},
	)

	capacity, ok := score.GetCarryingCapacity()
	if !ok {
		t.Fatal("expected carrying capacity to resolve for zero strength")
	}

	light := capacity.GetLightLoadMax()
	if light.GetKilograms() != 0 {
		t.Errorf("expected light load 0kg, got %.1fkg", light.GetKilograms())
	}

	if light.GetPounds() != 0 {
		t.Errorf("expected light load 0lb, got %.2flb", light.GetPounds())
	}

	medium := capacity.GetMediumLoad()
	if medium.GetMin().GetKilograms() != 0 || medium.GetMax().GetKilograms() != 0 {
		t.Errorf(
			"expected medium load 0-0kg, got %.1f-%.1fkg",
			medium.GetMin().GetKilograms(),
			medium.GetMax().GetKilograms(),
		)
	}

	heavy := capacity.GetHeavyLoad()
	if heavy.GetMin().GetKilograms() != 0 || heavy.GetMax().GetKilograms() != 0 {
		t.Errorf(
			"expected heavy load 0-0kg, got %.1f-%.1fkg",
			heavy.GetMin().GetKilograms(),
			heavy.GetMax().GetKilograms(),
		)
	}
}

func TestAbilityScore_GetCarryingCapacity_UsesMetricStrengthTable(t *testing.T) {
	score := mustNewAbilityScore(t,
		StrengthScore,
		abilityScoreValue{
			value: 18,
			valid: true,
		},
	)

	capacity, ok := score.GetCarryingCapacity()
	if !ok {
		t.Fatal("expected carrying capacity for strength")
	}

	light := capacity.GetLightLoadMax()
	if light.GetKilograms() != 50 {
		t.Errorf("expected light load 50kg, got %.1fkg", light.GetKilograms())
	}

	if !almostEqual(light.GetPounds(), 110.23) {
		t.Errorf("expected light load about 110.23lb, got %.2flb", light.GetPounds())
	}

	medium := capacity.GetMediumLoad()
	if medium.GetMin().GetKilograms() != 50.5 || medium.GetMax().GetKilograms() != 100 {
		t.Errorf(
			"expected medium load 50.5-100kg, got %.1f-%.1fkg",
			medium.GetMin().GetKilograms(),
			medium.GetMax().GetKilograms(),
		)
	}

	if !almostEqual(medium.GetMin().GetPounds(), 111.33) || !almostEqual(medium.GetMax().GetPounds(), 220.46) {
		t.Errorf(
			"expected medium load about 111.33-220.46lb, got %.2f-%.2flb",
			medium.GetMin().GetPounds(),
			medium.GetMax().GetPounds(),
		)
	}

	heavy := capacity.GetHeavyLoad()
	if heavy.GetMin().GetKilograms() != 100.5 || heavy.GetMax().GetKilograms() != 150 {
		t.Errorf(
			"expected heavy load 100.5-150kg, got %.1f-%.1fkg",
			heavy.GetMin().GetKilograms(),
			heavy.GetMax().GetKilograms(),
		)
	}

	if !almostEqual(heavy.GetMin().GetPounds(), 221.56) || !almostEqual(heavy.GetMax().GetPounds(), 330.69) {
		t.Errorf(
			"expected heavy load about 221.56-330.69lb, got %.2f-%.2flb",
			heavy.GetMin().GetPounds(),
			heavy.GetMax().GetPounds(),
		)
	}
}

func TestAbilityScore_GetCarryingCapacity_PlusTenMultipliesByFour(t *testing.T) {
	score := mustNewAbilityScore(t,
		StrengthScore,
		abilityScoreValue{
			value: 39,
			valid: true,
		},
	)

	capacity, ok := score.GetCarryingCapacity()
	if !ok {
		t.Fatal("expected carrying capacity for strength 39")
	}

	light := capacity.GetLightLoadMax()
	if light.GetKilograms() != 932 {
		t.Errorf("expected light load 932kg, got %.1fkg", light.GetKilograms())
	}

	if !almostEqual(light.GetPounds(), 2054.71) {
		t.Errorf("expected light load about 2054.71lb, got %.2flb", light.GetPounds())
	}

	medium := capacity.GetMediumLoad()
	if medium.GetMin().GetKilograms() != 954 || medium.GetMax().GetKilograms() != 1866 {
		t.Errorf(
			"expected medium load 954-1866kg, got %.1f-%.1fkg",
			medium.GetMin().GetKilograms(),
			medium.GetMax().GetKilograms(),
		)
	}

	if !almostEqual(medium.GetMin().GetPounds(), 2103.21) || !almostEqual(medium.GetMax().GetPounds(), 4113.83) {
		t.Errorf(
			"expected medium load about 2103.21-4113.83lb, got %.2f-%.2flb",
			medium.GetMin().GetPounds(),
			medium.GetMax().GetPounds(),
		)
	}

	heavy := capacity.GetHeavyLoad()
	if heavy.GetMin().GetKilograms() != 1868 || heavy.GetMax().GetKilograms() != 2800 {
		t.Errorf(
			"expected heavy load 1868-2800kg, got %.1f-%.1fkg",
			heavy.GetMin().GetKilograms(),
			heavy.GetMax().GetKilograms(),
		)
	}

	if !almostEqual(heavy.GetMin().GetPounds(), 4118.24) || !almostEqual(heavy.GetMax().GetPounds(), 6172.94) {
		t.Errorf(
			"expected heavy load about 4118.24-6172.94lb, got %.2f-%.2flb",
			heavy.GetMin().GetPounds(),
			heavy.GetMax().GetPounds(),
		)
	}
}

// ==============================
// SPELLCASTING PROFILE
// ==============================

func TestAbilityScore_GetSpellcastingProfile_SuppressedScore(t *testing.T) {
	score := mustNewAbilityScore(t,
		IntelligenceScore,
		abilityScoreValue{
			value: 18,
			valid: false,
		},
	)

	if _, ok := score.GetSpellcastingProfile(); ok {
		t.Fatal("expected suppressed score to have no spellcasting profile")
	}
}

func TestAbilityScore_GetSpellcastingProfile_LowCastingScore(t *testing.T) {
	score := mustNewAbilityScore(t,
		WisdomScore,
		abilityScoreValue{
			value: 10,
			valid: true,
		},
	)

	profile, ok := score.GetSpellcastingProfile()
	if !ok {
		t.Fatal("expected spellcasting profile")
	}

	if profile.GetMaxSpellLevel() != 0 {
		t.Errorf("expected max spell level 0, got %d", profile.GetMaxSpellLevel())
	}

	if profile.GetBonusSpells(1) != 0 {
		t.Errorf("expected no bonus 1st-level spells, got %d", profile.GetBonusSpells(1))
	}
}

func TestAbilityScore_GetSpellcastingProfile_TracksMaxSpellLevelFromScore(t *testing.T) {
	score := mustNewAbilityScore(t,
		CharismaScore,
		abilityScoreValue{
			value: 12,
			valid: true,
		},
	)

	profile, ok := score.GetSpellcastingProfile()
	if !ok {
		t.Fatal("expected spellcasting profile")
	}

	if profile.GetMaxSpellLevel() != 2 {
		t.Errorf("expected max spell level 2, got %d", profile.GetMaxSpellLevel())
	}

	if profile.GetBonusSpells(1) != 1 {
		t.Errorf("expected one bonus 1st-level spell, got %d", profile.GetBonusSpells(1))
	}

	if profile.GetBonusSpells(2) != 0 {
		t.Errorf("expected no bonus 2nd-level spells, got %d", profile.GetBonusSpells(2))
	}
}

func TestSpellcastingAbilityProfile_GetBonusSpells_TableProgression(t *testing.T) {
	score := mustNewAbilityScore(t,
		IntelligenceScore,
		abilityScoreValue{
			value: 20,
			valid: true,
		},
	)

	profile, ok := score.GetSpellcastingProfile()
	if !ok {
		t.Fatal("expected spellcasting profile")
	}

	tests := []struct {
		level int
		want  int
	}{
		{0, 0},
		{1, 2},
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 0},
	}

	for _, tt := range tests {
		got := profile.GetBonusSpells(tt.level)
		if got != tt.want {
			t.Errorf("spell level %d: expected %d bonus spells, got %d", tt.level, tt.want, got)
		}
	}
}
