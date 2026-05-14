package character

import (
	"testing"

	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"
)

func TestNewCharacterSelectedWeapon_ComposesSeededWeaponThroughCharacterBoundary(t *testing.T) {
	selectedWeapon := mustNewCharacterSelectedWeaponForTest(t, characterequipment.DaggerWeaponID)

	if selectedWeapon.GetWeaponID() != characterequipment.DaggerWeaponID {
		t.Fatalf("expected selected weapon id %q, got %q", characterequipment.DaggerWeaponID, selectedWeapon.GetWeaponID())
	}

	if selectedWeapon.GetProficiencyCategory() != characterequipment.SimpleWeaponProficiencyCategory {
		t.Fatalf(
			"expected selected weapon proficiency category %q, got %q",
			characterequipment.SimpleWeaponProficiencyCategory,
			selectedWeapon.GetProficiencyCategory(),
		)
	}

	weapon, ok := selectedWeapon.GetWeapon()
	if !ok {
		t.Fatal("expected selected weapon to resolve")
	}

	if weapon.GetDisplayName() != "Dagger" {
		t.Fatalf("expected selected weapon display name %q, got %q", "Dagger", weapon.GetDisplayName())
	}
}

func TestCharacterSelectedWeapon_IsProficientWithCategoryProficiency(t *testing.T) {
	selectedWeapon := mustNewCharacterSelectedWeaponForTest(t, characterequipment.GauntletWeaponID)

	if !selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.SimpleWeaponsWeaponProficiencyID,
	}) {
		t.Fatal("expected simple weapon category proficiency to satisfy selected gauntlet proficiency")
	}

	if selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.MartialWeaponsWeaponProficiencyID,
	}) {
		t.Fatal("expected martial weapon category proficiency not to satisfy selected gauntlet proficiency")
	}
}

func TestCharacterSelectedWeapon_IsProficientWithIndividualProficiency(t *testing.T) {
	selectedWeapon := mustNewCharacterSelectedWeaponForTest(t, characterequipment.DaggerWeaponID)

	if !selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.DaggerWeaponProficiencyID,
	}) {
		t.Fatal("expected dagger weapon proficiency to satisfy selected dagger proficiency")
	}

	if selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.ClubWeaponProficiencyID,
	}) {
		t.Fatal("expected club weapon proficiency not to satisfy selected dagger proficiency")
	}
}

func TestCharacterSelectedWeapon_DoesNotUseRawWeaponIDAsProficiency(t *testing.T) {
	selectedWeapon := mustNewCharacterSelectedWeaponForTest(t, characterequipment.DaggerWeaponID)

	if selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.WeaponProficiencyID(characterequipment.DaggerWeaponID),
	}) {
		t.Fatal("expected raw equipment weapon id not to satisfy selected dagger proficiency")
	}
}

func TestNewCharacterSelectedWeapon_RejectsUnknownAndMalformedWeapons(t *testing.T) {
	if _, ok := NewCharacterSelectedWeapon(characterequipment.WeaponID("longsword")); ok {
		t.Fatal("expected unseeded selected weapon to be rejected")
	}

	if _, ok := NewCharacterSelectedWeapon(characterequipment.WeaponID(" dagger")); ok {
		t.Fatal("expected malformed selected weapon id to be rejected")
	}
}

func TestCharacterSelectedWeapon_UnsupportedMappingsFailClosed(t *testing.T) {
	selectedWeapon := characterSelectedWeapon{
		id:                  characterequipment.WeaponID("net"),
		proficiencyCategory: characterequipment.ExoticWeaponProficiencyCategory,
		valid:               true,
	}

	if selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.SimpleWeaponsWeaponProficiencyID,
		characterclass.MartialWeaponsWeaponProficiencyID,
	}) {
		t.Fatal("expected unsupported exotic selected weapon mapping to fail closed")
	}
}

func TestCharacterSelectedWeapon_UnsupportedCategoryCanStillUseIndividualMapping(t *testing.T) {
	selectedWeapon := characterSelectedWeapon{
		id:                  characterequipment.DaggerWeaponID,
		proficiencyCategory: characterequipment.ExoticWeaponProficiencyCategory,
		valid:               true,
	}

	if !selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.DaggerWeaponProficiencyID,
	}) {
		t.Fatal("expected individual selected weapon mapping to satisfy proficiency when category mapping is unsupported")
	}
}

func TestCharacterSelectedWeapon_ZeroValueFailsClosed(t *testing.T) {
	var selectedWeapon CharacterSelectedWeapon

	if selectedWeapon.GetWeaponID() != "" {
		t.Fatalf("expected zero-value selected weapon id to be empty, got %q", selectedWeapon.GetWeaponID())
	}

	if selectedWeapon.GetProficiencyCategory() != "" {
		t.Fatalf(
			"expected zero-value selected weapon proficiency category to be empty, got %q",
			selectedWeapon.GetProficiencyCategory(),
		)
	}

	if _, ok := selectedWeapon.GetWeapon(); ok {
		t.Fatal("expected zero-value selected weapon not to resolve")
	}

	if selectedWeapon.IsProficientWith([]characterclass.WeaponProficiencyID{
		characterclass.SimpleWeaponsWeaponProficiencyID,
	}) {
		t.Fatal("expected zero-value selected weapon proficiency check to fail closed")
	}
}

func mustNewCharacterSelectedWeaponForTest(
	t *testing.T,
	id characterequipment.WeaponID,
) CharacterSelectedWeapon {
	t.Helper()

	selectedWeapon, ok := NewCharacterSelectedWeapon(id)
	if !ok {
		t.Fatalf("expected selected weapon %q to compose", id)
	}

	return selectedWeapon
}
