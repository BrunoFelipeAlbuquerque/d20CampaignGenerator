package class

const (
	SimpleWeaponsWeaponProficiencyID  WeaponProficiencyID = "Simple Weapons"
	MartialWeaponsWeaponProficiencyID WeaponProficiencyID = "Martial Weapons"

	ClubWeaponProficiencyID                    WeaponProficiencyID = "Club"
	CrossbowHeavyWeaponProficiencyID           WeaponProficiencyID = "Crossbow, Heavy"
	CrossbowLightWeaponProficiencyID           WeaponProficiencyID = "Crossbow, Light"
	DaggerWeaponProficiencyID                  WeaponProficiencyID = "Dagger"
	DartWeaponProficiencyID                    WeaponProficiencyID = "Dart"
	HandCrossbowWeaponProficiencyID            WeaponProficiencyID = "Crossbow, Hand"
	HandaxeWeaponProficiencyID                 WeaponProficiencyID = "Handaxe"
	JavelinWeaponProficiencyID                 WeaponProficiencyID = "Javelin"
	KamaWeaponProficiencyID                    WeaponProficiencyID = "Kama"
	LongswordWeaponProficiencyID               WeaponProficiencyID = "Longsword"
	NunchakuWeaponProficiencyID                WeaponProficiencyID = "Nunchaku"
	QuarterstaffWeaponProficiencyID            WeaponProficiencyID = "Quarterstaff"
	RapierWeaponProficiencyID                  WeaponProficiencyID = "Rapier"
	ScytheWeaponProficiencyID                  WeaponProficiencyID = "Scythe"
	SaiWeaponProficiencyID                     WeaponProficiencyID = "Sai"
	SapWeaponProficiencyID                     WeaponProficiencyID = "Sap"
	ScimitarWeaponProficiencyID                WeaponProficiencyID = "Scimitar"
	ShortbowWeaponProficiencyID                WeaponProficiencyID = "Shortbow"
	ShortspearWeaponProficiencyID              WeaponProficiencyID = "Shortspear"
	ShortSwordWeaponProficiencyID              WeaponProficiencyID = "Short Sword"
	ShurikenWeaponProficiencyID                WeaponProficiencyID = "Shuriken"
	SianghamWeaponProficiencyID                WeaponProficiencyID = "Siangham"
	SickleWeaponProficiencyID                  WeaponProficiencyID = "Sickle"
	SlingWeaponProficiencyID                   WeaponProficiencyID = "Sling"
	SpearWeaponProficiencyID                   WeaponProficiencyID = "Spear"
	WhipWeaponProficiencyID                    WeaponProficiencyID = "Whip"
	FavoredWeaponOfDeityWeaponProficiencyID    WeaponProficiencyID = "Favored Weapon of Deity"
	WildShapeNaturalAttacksWeaponProficiencyID WeaponProficiencyID = "Wild Shape Natural Attacks"
)

const (
	LightArmorProficiencyID       ArmorProficiencyID = "Light Armor"
	MediumArmorProficiencyID      ArmorProficiencyID = "Medium Armor"
	HeavyArmorProficiencyID       ArmorProficiencyID = "Heavy Armor"
	ShieldArmorProficiencyID      ArmorProficiencyID = "Shields"
	TowerShieldArmorProficiencyID ArmorProficiencyID = "Tower Shields"
)

var validWeaponProficiencyIDs = map[WeaponProficiencyID]struct{}{
	SimpleWeaponsWeaponProficiencyID:           {},
	MartialWeaponsWeaponProficiencyID:          {},
	ClubWeaponProficiencyID:                    {},
	CrossbowHeavyWeaponProficiencyID:           {},
	CrossbowLightWeaponProficiencyID:           {},
	DaggerWeaponProficiencyID:                  {},
	DartWeaponProficiencyID:                    {},
	HandCrossbowWeaponProficiencyID:            {},
	HandaxeWeaponProficiencyID:                 {},
	JavelinWeaponProficiencyID:                 {},
	KamaWeaponProficiencyID:                    {},
	LongswordWeaponProficiencyID:               {},
	NunchakuWeaponProficiencyID:                {},
	QuarterstaffWeaponProficiencyID:            {},
	RapierWeaponProficiencyID:                  {},
	ScytheWeaponProficiencyID:                  {},
	SaiWeaponProficiencyID:                     {},
	SapWeaponProficiencyID:                     {},
	ScimitarWeaponProficiencyID:                {},
	ShortbowWeaponProficiencyID:                {},
	ShortspearWeaponProficiencyID:              {},
	ShortSwordWeaponProficiencyID:              {},
	ShurikenWeaponProficiencyID:                {},
	SianghamWeaponProficiencyID:                {},
	SickleWeaponProficiencyID:                  {},
	SlingWeaponProficiencyID:                   {},
	SpearWeaponProficiencyID:                   {},
	WhipWeaponProficiencyID:                    {},
	FavoredWeaponOfDeityWeaponProficiencyID:    {},
	WildShapeNaturalAttacksWeaponProficiencyID: {},
}

var validArmorProficiencyIDs = map[ArmorProficiencyID]struct{}{
	LightArmorProficiencyID:       {},
	MediumArmorProficiencyID:      {},
	HeavyArmorProficiencyID:       {},
	ShieldArmorProficiencyID:      {},
	TowerShieldArmorProficiencyID: {},
}

func isValidWeaponProficiencyID(id WeaponProficiencyID) bool {
	_, ok := validWeaponProficiencyIDs[id]
	return ok
}

func isValidArmorProficiencyID(id ArmorProficiencyID) bool {
	_, ok := validArmorProficiencyIDs[id]
	return ok
}
