package character

import characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"

type characterEquipment struct {
	id       characterequipment.EquipmentID
	quantity int
}
type CharacterEquipment = characterEquipment

func NewCharacterEquipment(
	id characterequipment.EquipmentID,
	quantity int,
) (CharacterEquipment, bool) {
	if quantity <= 0 {
		return characterEquipment{}, false
	}

	if _, ok := characterequipment.GetEquipmentByID(id); !ok {
		return characterEquipment{}, false
	}

	return characterEquipment{
		id:       id,
		quantity: quantity,
	}, true
}

func (e characterEquipment) GetEquipmentID() characterequipment.EquipmentID {
	return e.id
}

func (e characterEquipment) GetQuantity() int {
	return e.quantity
}

func (e characterEquipment) GetEquipment() (characterequipment.Equipment, bool) {
	return characterequipment.GetEquipmentByID(e.id)
}
