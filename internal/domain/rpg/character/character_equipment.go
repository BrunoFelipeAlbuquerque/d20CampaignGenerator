package character

import characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"

type characterEquipment struct {
	ref      characterequipment.CarryableItemRef
	quantity int
}
type CharacterEquipment = characterEquipment

func NewCharacterEquipment(
	ref characterequipment.CarryableItemRef,
	quantity int,
) (CharacterEquipment, bool) {
	if quantity <= 0 {
		return characterEquipment{}, false
	}

	if _, ok := characterequipment.GetCarryableItemByRef(ref); !ok {
		return characterEquipment{}, false
	}

	return characterEquipment{
		ref:      ref,
		quantity: quantity,
	}, true
}

func NewCharacterAdventuringGear(
	id characterequipment.EquipmentID,
	quantity int,
) (CharacterEquipment, bool) {
	ref, ok := characterequipment.NewEquipmentCarryableItemRef(id)
	if !ok {
		return characterEquipment{}, false
	}

	return NewCharacterEquipment(ref, quantity)
}

func (e characterEquipment) GetCarryableItemRef() characterequipment.CarryableItemRef {
	return e.ref
}

func (e characterEquipment) GetEquipmentID() characterequipment.EquipmentID {
	if e.ref.GetKind() != characterequipment.EquipmentCarryableItemKind {
		return ""
	}

	return characterequipment.EquipmentID(e.ref.GetID())
}

func (e characterEquipment) GetQuantity() int {
	return e.quantity
}

func (e characterEquipment) GetCarryableItem() (characterequipment.CarryableItem, bool) {
	return characterequipment.GetCarryableItemByRef(e.ref)
}

func (e characterEquipment) GetEquipment() (characterequipment.Equipment, bool) {
	if e.ref.GetKind() != characterequipment.EquipmentCarryableItemKind {
		var value characterequipment.Equipment
		return value, false
	}

	return characterequipment.GetEquipmentByID(characterequipment.EquipmentID(e.ref.GetID()))
}
