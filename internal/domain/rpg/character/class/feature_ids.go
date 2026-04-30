package class

type classFeatureID string
type ClassFeatureID = classFeatureID

const (
	BardicPerformanceClassFeatureID     ClassFeatureID = "Bardic Performance"
	ChannelEnergyClassFeatureID         ClassFeatureID = "Channel Energy"
	ChannelNegativeEnergyClassFeatureID ClassFeatureID = "Channel Negative Energy"
	ChannelPositiveEnergyClassFeatureID ClassFeatureID = "Channel Positive Energy"
	FighterBonusFeatsClassFeatureID     ClassFeatureID = "Fighter Bonus Feats"
	KiPoolClassFeatureID                ClassFeatureID = "Ki Pool"
	LayOnHandsClassFeatureID            ClassFeatureID = "Lay On Hands"
	MercyClassFeatureID                 ClassFeatureID = "Mercy"
	RageClassFeatureID                  ClassFeatureID = "Rage"
	WildShapeClassFeatureID             ClassFeatureID = "Wild Shape"
)

func (id classFeatureID) GetName() string {
	if !isValidClassFeatureID(ClassFeatureID(id)) {
		return ""
	}

	return string(id)
}

func isValidClassFeatureID(id ClassFeatureID) bool {
	switch id {
	case BardicPerformanceClassFeatureID,
		ChannelEnergyClassFeatureID,
		ChannelNegativeEnergyClassFeatureID,
		ChannelPositiveEnergyClassFeatureID,
		FighterBonusFeatsClassFeatureID,
		KiPoolClassFeatureID,
		LayOnHandsClassFeatureID,
		MercyClassFeatureID,
		RageClassFeatureID,
		WildShapeClassFeatureID:
		return true
	default:
		return false
	}
}
