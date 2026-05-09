package equipment

const (
	BucklerArmorID           ArmorID = "buckler"
	ChainShirtArmorID        ArmorID = "chain-shirt"
	LeatherArmorID           ArmorID = "leather"
	PaddedArmorID            ArmorID = "padded"
	ShieldHeavySteelArmorID  ArmorID = "shield-heavy-steel"
	ShieldHeavyWoodenArmorID ArmorID = "shield-heavy-wooden"
	ShieldLightSteelArmorID  ArmorID = "shield-light-steel"
	ShieldLightWoodenArmorID ArmorID = "shield-light-wooden"
	ShieldTowerArmorID       ArmorID = "shield-tower"
	StuddedLeatherArmorID    ArmorID = "studded-leather"
)

var coreArmor = mustBuildCoreArmor()

var coreArmorOrder = []ArmorID{
	PaddedArmorID,
	LeatherArmorID,
	StuddedLeatherArmorID,
	ChainShirtArmorID,
	BucklerArmorID,
	ShieldLightWoodenArmorID,
	ShieldLightSteelArmorID,
	ShieldHeavyWoodenArmorID,
	ShieldHeavySteelArmorID,
	ShieldTowerArmorID,
}

func GetArmorByID(id ArmorID) (Armor, bool) {
	value, ok := coreArmor[id]
	if !ok {
		return armor{}, false
	}

	return cloneArmor(value), true
}

func GetArmor() []Armor {
	values := make([]Armor, 0, len(coreArmorOrder))

	for _, id := range coreArmorOrder {
		values = append(values, cloneArmor(coreArmor[id]))
	}

	return values
}

func mustBuildCoreArmor() map[ArmorID]Armor {
	return map[ArmorID]Armor{
		PaddedArmorID: mustNewCoreLightArmor(
			PaddedArmorID,
			"Padded",
			1,
			8,
			0,
			5,
			500,
			160,
		),
		LeatherArmorID: mustNewCoreLightArmor(
			LeatherArmorID,
			"Leather",
			2,
			6,
			0,
			10,
			1000,
			240,
		),
		StuddedLeatherArmorID: mustNewCoreLightArmor(
			StuddedLeatherArmorID,
			"Studded leather",
			3,
			5,
			-1,
			15,
			2500,
			320,
		),
		ChainShirtArmorID: mustNewCoreLightArmor(
			ChainShirtArmorID,
			"Chain shirt",
			4,
			4,
			-2,
			20,
			10000,
			400,
		),
		BucklerArmorID: mustNewCoreShield(
			BucklerArmorID,
			"Buckler",
			1,
			-1,
			5,
			1500,
			80,
		),
		ShieldLightWoodenArmorID: mustNewCoreShield(
			ShieldLightWoodenArmorID,
			"Shield, light wooden",
			1,
			-1,
			5,
			300,
			80,
		),
		ShieldLightSteelArmorID: mustNewCoreShield(
			ShieldLightSteelArmorID,
			"Shield, light steel",
			1,
			-1,
			5,
			900,
			96,
		),
		ShieldHeavyWoodenArmorID: mustNewCoreShield(
			ShieldHeavyWoodenArmorID,
			"Shield, heavy wooden",
			2,
			-2,
			15,
			700,
			160,
		),
		ShieldHeavySteelArmorID: mustNewCoreShield(
			ShieldHeavySteelArmorID,
			"Shield, heavy steel",
			2,
			-2,
			15,
			2000,
			240,
		),
		ShieldTowerArmorID: mustNewCoreTowerShield(
			ShieldTowerArmorID,
			"Shield, tower",
			4,
			2,
			-10,
			50,
			3000,
			720,
		),
	}
}

func mustNewCoreLightArmor(
	id ArmorID,
	displayName string,
	armorClassBonus int,
	maximumDexterityBonus int,
	armorCheckPenalty int,
	arcaneSpellFailureChance int,
	copperPieces int,
	ounces int,
) Armor {
	return mustNewCoreArmor(
		id,
		displayName,
		LightArmorCategory,
		mustNewCoreArmorClassBonus(armorClassBonus),
		mustNewCoreArmorMaximumDexterityBonus(maximumDexterityBonus),
		mustNewCoreArmorCheckPenalty(armorCheckPenalty),
		mustNewCoreArmorArcaneSpellFailureChance(arcaneSpellFailureChance),
		NewNoArmorSpeedImpact(),
		copperPieces,
		ounces,
	)
}

func mustNewCoreShield(
	id ArmorID,
	displayName string,
	armorClassBonus int,
	armorCheckPenalty int,
	arcaneSpellFailureChance int,
	copperPieces int,
	ounces int,
) Armor {
	return mustNewCoreArmor(
		id,
		displayName,
		ShieldArmorCategory,
		mustNewCoreArmorClassBonus(armorClassBonus),
		NewNoArmorMaximumDexterityBonus(),
		mustNewCoreArmorCheckPenalty(armorCheckPenalty),
		mustNewCoreArmorArcaneSpellFailureChance(arcaneSpellFailureChance),
		NewNoArmorSpeedImpact(),
		copperPieces,
		ounces,
	)
}

func mustNewCoreTowerShield(
	id ArmorID,
	displayName string,
	armorClassBonus int,
	maximumDexterityBonus int,
	armorCheckPenalty int,
	arcaneSpellFailureChance int,
	copperPieces int,
	ounces int,
) Armor {
	return mustNewCoreArmor(
		id,
		displayName,
		TowerShieldArmorCategory,
		mustNewCoreArmorClassBonus(armorClassBonus),
		mustNewCoreArmorMaximumDexterityBonus(maximumDexterityBonus),
		mustNewCoreArmorCheckPenalty(armorCheckPenalty),
		mustNewCoreArmorArcaneSpellFailureChance(arcaneSpellFailureChance),
		NewNoArmorSpeedImpact(),
		copperPieces,
		ounces,
	)
}

func mustNewCoreArmor(
	id ArmorID,
	displayName string,
	category ArmorCategory,
	armorClassBonus ArmorClassBonus,
	maximumDexterityBonus ArmorMaximumDexterityBonus,
	armorCheckPenalty ArmorCheckPenalty,
	arcaneSpellFailureChance ArmorArcaneSpellFailureChance,
	speedImpact ArmorSpeedImpact,
	copperPieces int,
	ounces int,
) Armor {
	cost, ok := NewEquipmentCost(copperPieces)
	if !ok {
		panic("invalid core armor cost seed")
	}

	weight, ok := NewEquipmentWeightOunces(ounces)
	if !ok {
		panic("invalid core armor weight seed")
	}

	armor, ok := NewArmor(
		id,
		displayName,
		category,
		armorClassBonus,
		maximumDexterityBonus,
		armorCheckPenalty,
		arcaneSpellFailureChance,
		speedImpact,
		cost,
		weight,
	)
	if !ok {
		panic("invalid core armor seed")
	}

	return armor
}

func mustNewCoreArmorClassBonus(points int) ArmorClassBonus {
	bonus, ok := NewArmorClassBonus(points)
	if !ok {
		panic("invalid core armor class bonus seed")
	}

	return bonus
}

func mustNewCoreArmorMaximumDexterityBonus(points int) ArmorMaximumDexterityBonus {
	bonus, ok := NewArmorMaximumDexterityBonus(points)
	if !ok {
		panic("invalid core armor maximum Dexterity bonus seed")
	}

	return bonus
}

func mustNewCoreArmorCheckPenalty(penalty int) ArmorCheckPenalty {
	value, ok := NewArmorCheckPenalty(penalty)
	if !ok {
		panic("invalid core armor check penalty seed")
	}

	return value
}

func mustNewCoreArmorArcaneSpellFailureChance(percent int) ArmorArcaneSpellFailureChance {
	chance, ok := NewArmorArcaneSpellFailureChance(percent)
	if !ok {
		panic("invalid core armor arcane spell failure chance seed")
	}

	return chance
}

func cloneArmor(value Armor) Armor {
	return armor{
		id:                       value.id,
		displayName:              value.displayName,
		category:                 value.category,
		armorClassBonus:          value.armorClassBonus,
		maximumDexterityBonus:    value.maximumDexterityBonus,
		armorCheckPenalty:        value.armorCheckPenalty,
		arcaneSpellFailureChance: value.arcaneSpellFailureChance,
		speedImpact:              value.speedImpact,
		cost:                     value.cost,
		weight:                   value.weight,
	}
}
