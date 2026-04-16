package creaturetype

type creatureSubtypeID string
type CreatureSubtypeID = creatureSubtypeID

const (
	AquaticSubtype     CreatureSubtypeID = "Aquatic"
	AugmentedSubtype   CreatureSubtypeID = "Augmented"
	ElementalSubtype   CreatureSubtypeID = "Elemental"
	IncorporealSubtype CreatureSubtypeID = "Incorporeal"
	NativeSubtype      CreatureSubtypeID = "Native"
	SwarmSubtype       CreatureSubtypeID = "Swarm"
)

func isValidCreatureSubtypeID(value CreatureSubtypeID) bool {
	switch value {
	case AquaticSubtype,
		AugmentedSubtype,
		ElementalSubtype,
		IncorporealSubtype,
		NativeSubtype,
		SwarmSubtype:
		return true
	default:
		return false
	}
}
