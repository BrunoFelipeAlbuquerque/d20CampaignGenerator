package equipment

type carryableItemKind string
type CarryableItemKind = carryableItemKind

const (
	EquipmentCarryableItemKind CarryableItemKind = "Equipment"
	WeaponCarryableItemKind    CarryableItemKind = "Weapon"
	ArmorCarryableItemKind     CarryableItemKind = "Armor"
)

type carryableItemRef struct {
	kind  carryableItemKind
	id    string
	valid bool
}
type CarryableItemRef = carryableItemRef

type carryableItem struct {
	ref         carryableItemRef
	displayName string
	cost        equipmentCost
	weight      equipmentWeight
}
type CarryableItem = carryableItem

func NewEquipmentCarryableItemRef(id EquipmentID) (CarryableItemRef, bool) {
	if !isValidEquipmentID(id) {
		return carryableItemRef{}, false
	}

	return carryableItemRef{
		kind:  EquipmentCarryableItemKind,
		id:    string(id),
		valid: true,
	}, true
}

func NewWeaponCarryableItemRef(id WeaponID) (CarryableItemRef, bool) {
	if !isValidWeaponID(id) {
		return carryableItemRef{}, false
	}

	return carryableItemRef{
		kind:  WeaponCarryableItemKind,
		id:    string(id),
		valid: true,
	}, true
}

func NewArmorCarryableItemRef(id ArmorID) (CarryableItemRef, bool) {
	if !isValidArmorID(id) {
		return carryableItemRef{}, false
	}

	return carryableItemRef{
		kind:  ArmorCarryableItemKind,
		id:    string(id),
		valid: true,
	}, true
}

func NewCarryableItemFromEquipment(value Equipment) (CarryableItem, bool) {
	if _, ok := NewEquipment(value.id, value.displayName, value.category, value.cost, value.weight); !ok {
		return carryableItem{}, false
	}

	ref, ok := NewEquipmentCarryableItemRef(value.id)
	if !ok {
		return carryableItem{}, false
	}

	return carryableItem{
		ref:         ref,
		displayName: value.displayName,
		cost:        value.cost,
		weight:      value.weight,
	}, true
}

func NewCarryableItemFromWeapon(value Weapon) (CarryableItem, bool) {
	if _, ok := NewWeapon(
		value.id,
		value.displayName,
		value.proficiencyCategory,
		value.category,
		value.damageProfile,
		value.criticalProfile,
		value.rangeIncrement,
		value.cost,
		value.weight,
	); !ok {
		return carryableItem{}, false
	}

	ref, ok := NewWeaponCarryableItemRef(value.id)
	if !ok {
		return carryableItem{}, false
	}

	return carryableItem{
		ref:         ref,
		displayName: value.displayName,
		cost:        value.cost,
		weight:      value.weight,
	}, true
}

func NewCarryableItemFromArmor(value Armor) (CarryableItem, bool) {
	if _, ok := NewArmor(
		value.id,
		value.displayName,
		value.category,
		value.armorClassBonus,
		value.maximumDexterityBonus,
		value.armorCheckPenalty,
		value.arcaneSpellFailureChance,
		value.speedImpact,
		value.cost,
		value.weight,
	); !ok {
		return carryableItem{}, false
	}

	ref, ok := NewArmorCarryableItemRef(value.id)
	if !ok {
		return carryableItem{}, false
	}

	return carryableItem{
		ref:         ref,
		displayName: value.displayName,
		cost:        value.cost,
		weight:      value.weight,
	}, true
}

func GetCarryableItemByRef(ref CarryableItemRef) (CarryableItem, bool) {
	if !isValidCarryableItemRef(ref) {
		return carryableItem{}, false
	}

	switch ref.kind {
	case EquipmentCarryableItemKind:
		value, ok := GetEquipmentByID(EquipmentID(ref.id))
		if !ok {
			return carryableItem{}, false
		}

		return NewCarryableItemFromEquipment(value)
	case WeaponCarryableItemKind:
		value, ok := GetWeaponByID(WeaponID(ref.id))
		if !ok {
			return carryableItem{}, false
		}

		return NewCarryableItemFromWeapon(value)
	case ArmorCarryableItemKind:
		return carryableItem{}, false
	default:
		return carryableItem{}, false
	}
}

func GetCarryableItems() []CarryableItem {
	values := make([]CarryableItem, 0, len(coreEquipmentOrder))

	for _, id := range coreEquipmentOrder {
		ref, ok := NewEquipmentCarryableItemRef(id)
		if !ok {
			continue
		}

		value, ok := GetCarryableItemByRef(ref)
		if ok {
			values = append(values, value)
		}
	}

	return values
}

func (r carryableItemRef) GetKind() CarryableItemKind {
	if !isValidCarryableItemRef(r) {
		return ""
	}

	return r.kind
}

func (r carryableItemRef) GetID() string {
	if !isValidCarryableItemRef(r) {
		return ""
	}

	return r.id
}

func (i carryableItem) GetRef() CarryableItemRef {
	return i.ref
}

func (i carryableItem) GetDisplayName() string {
	return i.displayName
}

func (i carryableItem) GetCost() EquipmentCost {
	return i.cost
}

func (i carryableItem) GetWeight() EquipmentWeight {
	return i.weight
}

func isValidCarryableItemRef(ref CarryableItemRef) bool {
	if !ref.valid {
		return false
	}

	switch ref.kind {
	case EquipmentCarryableItemKind:
		return isValidEquipmentID(EquipmentID(ref.id))
	case WeaponCarryableItemKind:
		return isValidWeaponID(WeaponID(ref.id))
	case ArmorCarryableItemKind:
		return isValidArmorID(ArmorID(ref.id))
	default:
		return false
	}
}
