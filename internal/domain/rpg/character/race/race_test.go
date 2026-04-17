package race

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
)

func TestNewRace_ConstructsValidatedRaceChassis(t *testing.T) {
	strengthModifier, ok := NewAbilityScoreModifier(ability.StrengthScore, 2)
	if !ok {
		t.Fatal("expected strength modifier to be constructed")
	}

	constitutionModifier, ok := NewAbilityScoreModifier(ability.ConstitutionScore, -2)
	if !ok {
		t.Fatal("expected constitution modifier to be constructed")
	}

	race, ok := NewRace(
		RaceID("elf"),
		ability.MediumSize,
		30,
		[]AbilityScoreModifier{strengthModifier, constitutionModifier},
		[]LanguageID{"Common", "Elven"},
		[]RacialFeatureID{"Keen Senses", "Low-Light Vision"},
	)
	if !ok {
		t.Fatal("expected race chassis to be constructed")
	}

	if race.GetID() != RaceID("elf") {
		t.Fatalf("expected race id %q, got %q", RaceID("elf"), race.GetID())
	}

	if race.GetSize() != ability.MediumSize {
		t.Fatalf("expected race size %q, got %q", ability.MediumSize, race.GetSize())
	}

	if race.GetBaseSpeed() != 30 {
		t.Fatalf("expected base speed 30, got %d", race.GetBaseSpeed())
	}

	modifiers := race.GetAbilityScoreModifiers()
	if len(modifiers) != 2 {
		t.Fatalf("expected 2 ability score modifiers, got %d", len(modifiers))
	}

	if modifiers[0].GetScoreID() != ability.StrengthScore || modifiers[0].GetModifier() != 2 {
		t.Fatalf("expected first modifier to be STR +2, got (%q, %d)", modifiers[0].GetScoreID(), modifiers[0].GetModifier())
	}

	languages := race.GetRacialLanguages()
	if len(languages) != 2 || languages[0] != "Common" || languages[1] != "Elven" {
		t.Fatalf("expected racial languages [Common Elven], got %v", languages)
	}

	if !race.HasFeature("Keen Senses") {
		t.Fatal("expected Keen Senses feature to be present")
	}
}

func TestNewRace_DedupesModifiersLanguagesAndFeatures(t *testing.T) {
	intelligenceModifier, ok := NewAbilityScoreModifier(ability.IntelligenceScore, 2)
	if !ok {
		t.Fatal("expected intelligence modifier to be constructed")
	}

	race, ok := NewRace(
		RaceID("gnome"),
		ability.SmallSize,
		20,
		[]AbilityScoreModifier{intelligenceModifier, intelligenceModifier},
		[]LanguageID{"Common", "Gnome", "Common"},
		[]RacialFeatureID{"Defensive Training", "Defensive Training", "Keen Senses"},
	)
	if !ok {
		t.Fatal("expected race chassis to be constructed")
	}

	if len(race.GetAbilityScoreModifiers()) != 1 {
		t.Fatalf("expected deduped ability score modifier length 1, got %d", len(race.GetAbilityScoreModifiers()))
	}

	if len(race.GetRacialLanguages()) != 2 {
		t.Fatalf("expected deduped racial languages length 2, got %d", len(race.GetRacialLanguages()))
	}

	if len(race.GetRacialFeatures()) != 2 {
		t.Fatalf("expected deduped racial features length 2, got %d", len(race.GetRacialFeatures()))
	}
}

func TestNewRace_RejectsInvalidInputs(t *testing.T) {
	if _, ok := NewAbilityScoreModifier(ability.AbilityScoreID("LCK"), 2); ok {
		t.Fatal("expected invalid ability score id to be rejected")
	}

	if _, ok := NewAbilityScoreModifier(ability.StrengthScore, 0); ok {
		t.Fatal("expected zero ability score modifier to be rejected")
	}

	validModifier, ok := NewAbilityScoreModifier(ability.DexterityScore, 2)
	if !ok {
		t.Fatal("expected dexterity modifier to be constructed")
	}

	if _, ok := NewRace("", ability.MediumSize, 30, nil, nil, nil); ok {
		t.Fatal("expected empty race id to be rejected")
	}

	if _, ok := NewRace(RaceID("human"), ability.Size("Gigantic"), 30, nil, nil, nil); ok {
		t.Fatal("expected invalid size to be rejected")
	}

	if _, ok := NewRace(RaceID("human"), ability.MediumSize, 0, nil, nil, nil); ok {
		t.Fatal("expected non-positive base speed to be rejected")
	}

	if _, ok := NewRace(
		RaceID("human"),
		ability.MediumSize,
		30,
		[]AbilityScoreModifier{{scoreID: ability.AbilityScoreID("LCK"), modifier: 2}},
		nil,
		nil,
	); ok {
		t.Fatal("expected invalid ability score modifier entry to be rejected")
	}

	if _, ok := NewRace(
		RaceID("human"),
		ability.MediumSize,
		30,
		[]AbilityScoreModifier{validModifier},
		[]LanguageID{""},
		nil,
	); ok {
		t.Fatal("expected empty racial language to be rejected")
	}

	if _, ok := NewRace(
		RaceID("human"),
		ability.MediumSize,
		30,
		[]AbilityScoreModifier{validModifier},
		nil,
		[]RacialFeatureID{""},
	); ok {
		t.Fatal("expected empty racial feature to be rejected")
	}
}

func TestRace_GettersReturnDefensiveCopies(t *testing.T) {
	dexterityModifier, ok := NewAbilityScoreModifier(ability.DexterityScore, 2)
	if !ok {
		t.Fatal("expected dexterity modifier to be constructed")
	}

	race, ok := NewRace(
		RaceID("halfling"),
		ability.SmallSize,
		20,
		[]AbilityScoreModifier{dexterityModifier},
		[]LanguageID{"Common", "Halfling"},
		[]RacialFeatureID{"Sure-Footed"},
	)
	if !ok {
		t.Fatal("expected race chassis to be constructed")
	}

	modifiers := race.GetAbilityScoreModifiers()
	languages := race.GetRacialLanguages()
	features := race.GetRacialFeatures()

	modifiers[0] = AbilityScoreModifier{}
	languages[0] = "Changed"
	features[0] = "Changed"

	if race.GetAbilityScoreModifiers()[0].GetScoreID() != ability.DexterityScore {
		t.Fatal("expected ability score modifiers getter to return a defensive copy")
	}

	if race.GetRacialLanguages()[0] != "Common" {
		t.Fatal("expected racial languages getter to return a defensive copy")
	}

	if race.GetRacialFeatures()[0] != "Sure-Footed" {
		t.Fatal("expected racial features getter to return a defensive copy")
	}
}
