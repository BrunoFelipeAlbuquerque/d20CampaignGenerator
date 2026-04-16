package creaturetype

type creatureTypeID string
type CreatureTypeID = creatureTypeID

const (
	AberrationType        CreatureTypeID = "Aberration"
	AnimalType            CreatureTypeID = "Animal"
	ConstructType         CreatureTypeID = "Construct"
	DragonType            CreatureTypeID = "Dragon"
	FeyType               CreatureTypeID = "Fey"
	HumanoidType          CreatureTypeID = "Humanoid"
	MagicalBeastType      CreatureTypeID = "Magical Beast"
	MonstrousHumanoidType CreatureTypeID = "Monstrous Humanoid"
	OozeType              CreatureTypeID = "Ooze"
	OutsiderType          CreatureTypeID = "Outsider"
	PlantType             CreatureTypeID = "Plant"
	UndeadType            CreatureTypeID = "Undead"
	VerminType            CreatureTypeID = "Vermin"
)

func isValidCreatureTypeID(value CreatureTypeID) bool {
	switch value {
	case AberrationType,
		AnimalType,
		ConstructType,
		DragonType,
		FeyType,
		HumanoidType,
		MagicalBeastType,
		MonstrousHumanoidType,
		OozeType,
		OutsiderType,
		PlantType,
		UndeadType,
		VerminType:
		return true
	default:
		return false
	}
}
