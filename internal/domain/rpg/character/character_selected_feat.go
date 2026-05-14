package character

import characterfeat "d20campaigngenerator/internal/domain/rpg/character/feat"

type characterSelectedWeaponFeat struct {
	featID         characterfeat.FeatID
	selectedWeapon characterSelectedWeapon
	valid          bool
}
type CharacterSelectedWeaponFeat = characterSelectedWeaponFeat

func NewCharacterSelectedWeaponFeat(
	featID characterfeat.FeatID,
	selectedWeapon CharacterSelectedWeapon,
) (CharacterSelectedWeaponFeat, bool) {
	feat, ok := characterfeat.GetFeatByID(featID)
	if !ok || !isWeaponSelectionFeat(feat) {
		return characterSelectedWeaponFeat{}, false
	}

	selectedWeaponValue, ok := buildCharacterSelectedWeapon(selectedWeapon)
	if !ok || !selectedWeaponValue.valid {
		return characterSelectedWeaponFeat{}, false
	}

	return characterSelectedWeaponFeat{
		featID:         featID,
		selectedWeapon: selectedWeaponValue,
		valid:          true,
	}, true
}

func (f characterSelectedWeaponFeat) GetFeatID() characterfeat.FeatID {
	if !f.valid {
		return ""
	}

	return f.featID
}

func (f characterSelectedWeaponFeat) GetSelectedWeapon() CharacterSelectedWeapon {
	if !f.valid {
		return characterSelectedWeapon{}
	}

	return f.selectedWeapon
}

func isWeaponSelectionFeat(feat characterfeat.Feat) bool {
	if feat.GetCategory() != characterfeat.CombatFeatCategory {
		return false
	}

	for _, prerequisite := range feat.GetPrerequisites() {
		switch prerequisite.(type) {
		case characterfeat.SelectedWeaponProficiencyPrerequisite,
			characterfeat.SameSelectionFeatPrerequisite:
			return true
		}
	}

	return false
}
