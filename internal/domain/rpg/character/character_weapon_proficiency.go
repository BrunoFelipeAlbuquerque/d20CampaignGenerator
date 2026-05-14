package character

import (
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"
)

type characterSelectedWeapon struct {
	id                  characterequipment.WeaponID
	proficiencyCategory characterequipment.WeaponProficiencyCategory
	valid               bool
}
type CharacterSelectedWeapon = characterSelectedWeapon

func NewCharacterSelectedWeapon(
	id characterequipment.WeaponID,
) (CharacterSelectedWeapon, bool) {
	weapon, ok := characterequipment.GetWeaponByID(id)
	if !ok {
		return characterSelectedWeapon{}, false
	}

	return characterSelectedWeapon{
		id:                  weapon.GetID(),
		proficiencyCategory: weapon.GetProficiencyCategory(),
		valid:               true,
	}, true
}

func (w characterSelectedWeapon) GetWeaponID() characterequipment.WeaponID {
	if !w.valid {
		return ""
	}

	return w.id
}

func (w characterSelectedWeapon) GetProficiencyCategory() characterequipment.WeaponProficiencyCategory {
	if !w.valid {
		return ""
	}

	return w.proficiencyCategory
}

func (w characterSelectedWeapon) GetWeapon() (characterequipment.Weapon, bool) {
	if !w.valid {
		return characterequipment.Weapon{}, false
	}

	return characterequipment.GetWeaponByID(w.id)
}

func (w characterSelectedWeapon) IsProficientWith(
	proficiencies []characterclass.WeaponProficiencyID,
) bool {
	if !w.valid {
		return false
	}

	categoryProficiencyID, hasCategoryMapping := weaponCategoryProficiencyID(w.proficiencyCategory)
	individualProficiencyID, hasIndividualMapping := individualWeaponProficiencyID(w.id)
	if !hasCategoryMapping && !hasIndividualMapping {
		return false
	}

	for _, proficiencyID := range proficiencies {
		if hasCategoryMapping && proficiencyID == categoryProficiencyID {
			return true
		}

		if hasIndividualMapping && proficiencyID == individualProficiencyID {
			return true
		}
	}

	return false
}

func weaponCategoryProficiencyID(
	category characterequipment.WeaponProficiencyCategory,
) (characterclass.WeaponProficiencyID, bool) {
	switch category {
	case characterequipment.SimpleWeaponProficiencyCategory:
		return characterclass.SimpleWeaponsWeaponProficiencyID, true
	case characterequipment.MartialWeaponProficiencyCategory:
		return characterclass.MartialWeaponsWeaponProficiencyID, true
	default:
		return "", false
	}
}

func individualWeaponProficiencyID(
	id characterequipment.WeaponID,
) (characterclass.WeaponProficiencyID, bool) {
	proficiencyID, ok := individualWeaponProficiencyIDs[id]
	return proficiencyID, ok
}

var individualWeaponProficiencyIDs = map[characterequipment.WeaponID]characterclass.WeaponProficiencyID{
	characterequipment.ClubWeaponID:          characterclass.ClubWeaponProficiencyID,
	characterequipment.CrossbowHeavyWeaponID: characterclass.CrossbowHeavyWeaponProficiencyID,
	characterequipment.CrossbowLightWeaponID: characterclass.CrossbowLightWeaponProficiencyID,
	characterequipment.DaggerWeaponID:        characterclass.DaggerWeaponProficiencyID,
	characterequipment.DartWeaponID:          characterclass.DartWeaponProficiencyID,
	characterequipment.JavelinWeaponID:       characterclass.JavelinWeaponProficiencyID,
	characterequipment.QuarterstaffWeaponID:  characterclass.QuarterstaffWeaponProficiencyID,
	characterequipment.ShortspearWeaponID:    characterclass.ShortspearWeaponProficiencyID,
	characterequipment.SickleWeaponID:        characterclass.SickleWeaponProficiencyID,
	characterequipment.SlingWeaponID:         characterclass.SlingWeaponProficiencyID,
	characterequipment.SpearWeaponID:         characterclass.SpearWeaponProficiencyID,
}
