package character

import (
	"testing"

	characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"
)

func TestNewCharacterEquipment_ComposesCoreEquipmentThroughCharacterBoundary(t *testing.T) {
	ref := mustNewCharacterEquipmentRefForTest(t, characterequipment.BackpackEmptyEquipmentID)

	selectedEquipment, ok := NewCharacterEquipment(ref, 2)
	if !ok {
		t.Fatal("expected core equipment to compose through character boundary")
	}

	if selectedEquipment.GetCarryableItemRef().GetKind() != characterequipment.EquipmentCarryableItemKind {
		t.Fatalf("expected selected carryable kind %q, got %q", characterequipment.EquipmentCarryableItemKind, selectedEquipment.GetCarryableItemRef().GetKind())
	}

	if selectedEquipment.GetEquipmentID() != characterequipment.BackpackEmptyEquipmentID {
		t.Fatalf("expected selected equipment id %q, got %q", characterequipment.BackpackEmptyEquipmentID, selectedEquipment.GetEquipmentID())
	}

	if selectedEquipment.GetQuantity() != 2 {
		t.Fatalf("expected selected equipment quantity 2, got %d", selectedEquipment.GetQuantity())
	}

	carryableItem, ok := selectedEquipment.GetCarryableItem()
	if !ok {
		t.Fatal("expected selected carryable item to resolve")
	}

	if carryableItem.GetWeight().GetOunces() != 32 {
		t.Fatalf("expected carryable item weight 32 oz, got %d oz", carryableItem.GetWeight().GetOunces())
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
	ref := mustNewCharacterEquipmentRefForTest(t, characterequipment.EquipmentID("ten-foot-pole"))

	if _, ok := NewCharacterEquipment(ref, 1); ok {
		t.Fatal("expected unknown equipment to be rejected")
	}
}

func TestNewCharacterEquipment_RejectsMalformedEquipmentID(t *testing.T) {
	if _, ok := characterequipment.NewEquipmentCarryableItemRef(characterequipment.EquipmentID(" backpack-empty")); ok {
		t.Fatal("expected malformed equipment id to be rejected")
	}

	if _, ok := NewCharacterEquipment(characterequipment.CarryableItemRef{}, 1); ok {
		t.Fatal("expected zero-value carryable item ref to be rejected")
	}
}

func TestNewCharacterEquipment_RejectsNonPositiveQuantity(t *testing.T) {
	ref := mustNewCharacterEquipmentRefForTest(t, characterequipment.BackpackEmptyEquipmentID)

	for _, quantity := range []int{0, -1} {
		if _, ok := NewCharacterEquipment(ref, quantity); ok {
			t.Fatalf("expected quantity %d to be rejected", quantity)
		}
	}
}

func TestCharacterEquipment_ZeroValueDoesNotResolve(t *testing.T) {
	var selectedEquipment CharacterEquipment

	if _, ok := selectedEquipment.GetCarryableItem(); ok {
		t.Fatal("expected zero-value character equipment not to resolve as a carryable item")
	}

	if _, ok := selectedEquipment.GetEquipment(); ok {
		t.Fatal("expected zero-value character equipment not to resolve")
	}
}

func TestNewCharacterAdventuringGear_ComposesEquipmentRefConvenience(t *testing.T) {
	selectedEquipment, ok := NewCharacterAdventuringGear(characterequipment.BackpackEmptyEquipmentID, 1)
	if !ok {
		t.Fatal("expected character adventuring gear convenience constructor to compose")
	}

	if selectedEquipment.GetCarryableItemRef().GetKind() != characterequipment.EquipmentCarryableItemKind {
		t.Fatalf("expected selected carryable kind %q, got %q", characterequipment.EquipmentCarryableItemKind, selectedEquipment.GetCarryableItemRef().GetKind())
	}
}

func mustNewCharacterEquipmentRefForTest(
	t *testing.T,
	id characterequipment.EquipmentID,
) characterequipment.CarryableItemRef {
	t.Helper()

	ref, ok := characterequipment.NewEquipmentCarryableItemRef(id)
	if !ok {
		t.Fatalf("expected equipment carryable ref %q to compose", id)
	}

	return ref
}
