package character

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/ability"
	characterrace "d20campaigngenerator/internal/domain/rpg/character/race"
)

func TestNewCharacterRace_ComposesCoreRaceThroughCharacterBoundary(t *testing.T) {
	selectedRace, ok := NewCharacterRace(characterrace.ElfRaceID)
	if !ok {
		t.Fatal("expected core race to compose through character boundary")
	}

	if selectedRace.GetRaceID() != characterrace.ElfRaceID {
		t.Fatalf("expected selected race id %q, got %q", characterrace.ElfRaceID, selectedRace.GetRaceID())
	}

	race, ok := selectedRace.GetRace()
	if !ok {
		t.Fatal("expected selected core race to resolve")
	}

	if race.GetID() != characterrace.ElfRaceID {
		t.Fatalf("expected resolved race id %q, got %q", characterrace.ElfRaceID, race.GetID())
	}

	if race.GetSize() != ability.MediumSize {
		t.Fatalf("expected resolved race size %q, got %q", ability.MediumSize, race.GetSize())
	}

	if race.GetBaseSpeed() != 30 {
		t.Fatalf("expected resolved race base speed 30, got %d", race.GetBaseSpeed())
	}

	if !race.HasFeature(characterrace.ElvenImmunitiesFeatureID) {
		t.Fatalf("expected resolved race to have feature %q", characterrace.ElvenImmunitiesFeatureID)
	}
}

func TestNewCharacterRace_RejectsUnknownRace(t *testing.T) {
	if _, ok := NewCharacterRace(characterrace.RaceID("android")); ok {
		t.Fatal("expected unknown race to be rejected")
	}
}

func TestNewCharacterRace_RejectsMalformedRaceID(t *testing.T) {
	if _, ok := NewCharacterRace(characterrace.RaceID(" human")); ok {
		t.Fatal("expected malformed race id to be rejected")
	}
}

func TestCharacterRace_GetRaceReturnsDetachedCatalogRace(t *testing.T) {
	selectedRace, ok := NewCharacterRace(characterrace.DwarfRaceID)
	if !ok {
		t.Fatal("expected core race to compose through character boundary")
	}

	first, ok := selectedRace.GetRace()
	if !ok {
		t.Fatal("expected selected core race to resolve")
	}

	languages := first.GetAutomaticLanguages()
	if len(languages) == 0 {
		t.Fatal("expected dwarf to have automatic languages")
	}
	languages[0] = characterrace.ElvenLanguageID

	second, ok := selectedRace.GetRace()
	if !ok {
		t.Fatal("expected selected core race to resolve again")
	}

	if second.GetAutomaticLanguages()[0] != characterrace.CommonLanguageID {
		t.Fatalf("expected resolved race language to remain %q, got %q", characterrace.CommonLanguageID, second.GetAutomaticLanguages()[0])
	}
}

func TestCharacterRace_ZeroValueDoesNotResolve(t *testing.T) {
	var selectedRace CharacterRace

	if _, ok := selectedRace.GetRace(); ok {
		t.Fatal("expected zero-value character race not to resolve")
	}
}
