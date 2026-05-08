package character

import (
	"testing"

	characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"
)

func TestNewCharacterEquipment_ComposesCoreEquipmentThroughCharacterBoundary(t *testing.T) {
	selectedEquipment, ok := NewCharacterEquipment(characterequipment.BackpackEmptyEquipmentID, 2)
	if !ok {
		t.Fatal("expected core equipment to compose through character boundary")
	}

	if selectedEquipment.GetEquipmentID() != characterequipment.BackpackEmptyEquipmentID {
		t.Fatalf("expected selected equipment id %q, got %q", characterequipment.BackpackEmptyEquipmentID, selectedEquipment.GetEquipmentID())
	}

	if selectedEquipment.GetQuantity() != 2 {
		t.Fatalf("expected selected equipment quantity 2, got %d", selectedEquipment.GetQuantity())
	}

	equipment, ok := selectedEquipment.GetEquipment()
	if !ok {
		t.Fatal("expected selected core equipment to resolve")
	}

	if equipment.GetID() != characterequipment.BackpackEmptyEquipmentID {
		t.Fatalf("expected resolved equipment id %q, got %q", characterequipment.BackpackEmptyEquipmentID, equipment.GetID())
	}

	if equipment.GetDisplayName() != "Backpack (empty)" {
		t.Fatalf("expected resolved equipment display name %q, got %q", "Backpack (empty)", equipment.GetDisplayName())
	}

	if equipment.GetCost().GetCopperPieces() != 200 {
		t.Fatalf("expected resolved equipment cost 200 cp, got %d cp", equipment.GetCost().GetCopperPieces())
	}

	if equipment.GetWeight().GetOunces() != 32 {
		t.Fatalf("expected resolved equipment weight 32 oz, got %d oz", equipment.GetWeight().GetOunces())
	}
}

func TestNewCharacterEquipment_RejectsUnknownEquipment(t *testing.T) {
	if _, ok := NewCharacterEquipment(characterequipment.EquipmentID("ten-foot-pole"), 1); ok {
		t.Fatal("expected unknown equipment to be rejected")
	}
}

func TestNewCharacterEquipment_RejectsMalformedEquipmentID(t *testing.T) {
	if _, ok := NewCharacterEquipment(characterequipment.EquipmentID(" backpack-empty"), 1); ok {
		t.Fatal("expected malformed equipment id to be rejected")
	}
}

func TestNewCharacterEquipment_RejectsNonPositiveQuantity(t *testing.T) {
	for _, quantity := range []int{0, -1} {
		if _, ok := NewCharacterEquipment(characterequipment.BackpackEmptyEquipmentID, quantity); ok {
			t.Fatalf("expected quantity %d to be rejected", quantity)
		}
	}
}

func TestCharacterEquipment_ZeroValueDoesNotResolve(t *testing.T) {
	var selectedEquipment CharacterEquipment

	if _, ok := selectedEquipment.GetEquipment(); ok {
		t.Fatal("expected zero-value character equipment not to resolve")
	}
}
