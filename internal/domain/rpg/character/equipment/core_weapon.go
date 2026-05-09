package equipment

const (
	ClubWeaponID          WeaponID = "club"
	CrossbowHeavyWeaponID WeaponID = "crossbow-heavy"
	CrossbowLightWeaponID WeaponID = "crossbow-light"
	DaggerWeaponID        WeaponID = "dagger"
	DartWeaponID          WeaponID = "dart"
	GauntletWeaponID      WeaponID = "gauntlet"
	JavelinWeaponID       WeaponID = "javelin"
	LongspearWeaponID     WeaponID = "longspear"
	MaceHeavyWeaponID     WeaponID = "mace-heavy"
	MaceLightWeaponID     WeaponID = "mace-light"
	MorningstarWeaponID   WeaponID = "morningstar"
	QuarterstaffWeaponID  WeaponID = "quarterstaff"
	ShortspearWeaponID    WeaponID = "shortspear"
	SickleWeaponID        WeaponID = "sickle"
	SlingWeaponID         WeaponID = "sling"
	SpearWeaponID         WeaponID = "spear"
	UnarmedStrikeWeaponID WeaponID = "unarmed-strike"
)

var coreSimpleWeapons = mustBuildCoreSimpleWeapons()

var coreSimpleWeaponOrder = []WeaponID{
	GauntletWeaponID,
	UnarmedStrikeWeaponID,
	DaggerWeaponID,
	MaceLightWeaponID,
	SickleWeaponID,
	ClubWeaponID,
	MaceHeavyWeaponID,
	MorningstarWeaponID,
	ShortspearWeaponID,
	LongspearWeaponID,
	QuarterstaffWeaponID,
	SpearWeaponID,
	CrossbowHeavyWeaponID,
	CrossbowLightWeaponID,
	DartWeaponID,
	JavelinWeaponID,
	SlingWeaponID,
}

func mustBuildCoreSimpleWeapons() map[WeaponID]Weapon {
	return map[WeaponID]Weapon{
		GauntletWeaponID: mustNewCoreSimpleWeapon(
			GauntletWeaponID,
			"Gauntlet",
			UnarmedAttackWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 2, 1, 3),
			mustNewCoreWeaponCriticalProfile(20, 2),
			NewNoWeaponRangeIncrement(),
			200,
			16,
		),
		UnarmedStrikeWeaponID: mustNewCoreSimpleWeapon(
			UnarmedStrikeWeaponID,
			"Unarmed strike",
			UnarmedAttackWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 2, 1, 3),
			mustNewCoreWeaponCriticalProfile(20, 2),
			NewNoWeaponRangeIncrement(),
			0,
			0,
		),
		DaggerWeaponID: mustNewCoreSimpleWeapon(
			DaggerWeaponID,
			"Dagger",
			LightMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 3, 1, 4),
			mustNewCoreWeaponCriticalProfile(19, 2),
			mustNewCoreWeaponRangeIncrement(10),
			200,
			16,
		),
		MaceLightWeaponID: mustNewCoreSimpleWeapon(
			MaceLightWeaponID,
			"Mace, light",
			LightMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 4, 1, 6),
			mustNewCoreWeaponCriticalProfile(20, 2),
			NewNoWeaponRangeIncrement(),
			500,
			64,
		),
		SickleWeaponID: mustNewCoreSimpleWeapon(
			SickleWeaponID,
			"Sickle",
			LightMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 4, 1, 6),
			mustNewCoreWeaponCriticalProfile(20, 2),
			NewNoWeaponRangeIncrement(),
			600,
			32,
		),
		ClubWeaponID: mustNewCoreSimpleWeapon(
			ClubWeaponID,
			"Club",
			OneHandedMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 4, 1, 6),
			mustNewCoreWeaponCriticalProfile(20, 2),
			mustNewCoreWeaponRangeIncrement(10),
			0,
			48,
		),
		MaceHeavyWeaponID: mustNewCoreSimpleWeapon(
			MaceHeavyWeaponID,
			"Mace, heavy",
			OneHandedMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 6, 1, 8),
			mustNewCoreWeaponCriticalProfile(20, 2),
			NewNoWeaponRangeIncrement(),
			1200,
			128,
		),
		MorningstarWeaponID: mustNewCoreSimpleWeapon(
			MorningstarWeaponID,
			"Morningstar",
			OneHandedMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 6, 1, 8),
			mustNewCoreWeaponCriticalProfile(20, 2),
			NewNoWeaponRangeIncrement(),
			800,
			96,
		),
		ShortspearWeaponID: mustNewCoreSimpleWeapon(
			ShortspearWeaponID,
			"Shortspear",
			OneHandedMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 4, 1, 6),
			mustNewCoreWeaponCriticalProfile(20, 2),
			mustNewCoreWeaponRangeIncrement(20),
			100,
			48,
		),
		LongspearWeaponID: mustNewCoreSimpleWeapon(
			LongspearWeaponID,
			"Longspear",
			TwoHandedMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 6, 1, 8),
			mustNewCoreWeaponCriticalProfile(20, 3),
			NewNoWeaponRangeIncrement(),
			500,
			144,
		),
		QuarterstaffWeaponID: mustNewCoreSimpleWeapon(
			QuarterstaffWeaponID,
			"Quarterstaff",
			TwoHandedMeleeWeaponCategory,
			mustNewCoreDoubleWeaponDamageProfile(1, 4, 1, 6),
			mustNewCoreWeaponCriticalProfile(20, 2),
			NewNoWeaponRangeIncrement(),
			0,
			64,
		),
		SpearWeaponID: mustNewCoreSimpleWeapon(
			SpearWeaponID,
			"Spear",
			TwoHandedMeleeWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 6, 1, 8),
			mustNewCoreWeaponCriticalProfile(20, 3),
			mustNewCoreWeaponRangeIncrement(20),
			200,
			96,
		),
		CrossbowHeavyWeaponID: mustNewCoreSimpleWeapon(
			CrossbowHeavyWeaponID,
			"Crossbow, heavy",
			RangedWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 8, 1, 10),
			mustNewCoreWeaponCriticalProfile(19, 2),
			mustNewCoreWeaponRangeIncrement(120),
			5000,
			128,
		),
		CrossbowLightWeaponID: mustNewCoreSimpleWeapon(
			CrossbowLightWeaponID,
			"Crossbow, light",
			RangedWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 6, 1, 8),
			mustNewCoreWeaponCriticalProfile(19, 2),
			mustNewCoreWeaponRangeIncrement(80),
			3500,
			64,
		),
		DartWeaponID: mustNewCoreSimpleWeapon(
			DartWeaponID,
			"Dart",
			RangedWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 3, 1, 4),
			mustNewCoreWeaponCriticalProfile(20, 2),
			mustNewCoreWeaponRangeIncrement(20),
			50,
			8,
		),
		JavelinWeaponID: mustNewCoreSimpleWeapon(
			JavelinWeaponID,
			"Javelin",
			RangedWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 4, 1, 6),
			mustNewCoreWeaponCriticalProfile(20, 2),
			mustNewCoreWeaponRangeIncrement(30),
			100,
			32,
		),
		SlingWeaponID: mustNewCoreSimpleWeapon(
			SlingWeaponID,
			"Sling",
			RangedWeaponCategory,
			mustNewCoreWeaponDamageProfile(1, 3, 1, 4),
			mustNewCoreWeaponCriticalProfile(20, 2),
			mustNewCoreWeaponRangeIncrement(50),
			0,
			0,
		),
	}
}

func mustNewCoreSimpleWeapon(
	id WeaponID,
	displayName string,
	category WeaponCategory,
	damageProfile WeaponDamageProfile,
	criticalProfile WeaponCriticalProfile,
	rangeIncrement WeaponRangeIncrement,
	copperPieces int,
	ounces int,
) Weapon {
	cost, ok := NewEquipmentCost(copperPieces)
	if !ok {
		panic("invalid core simple weapon cost seed")
	}

	weight, ok := NewEquipmentWeightOunces(ounces)
	if !ok {
		panic("invalid core simple weapon weight seed")
	}

	weapon, ok := NewWeapon(
		id,
		displayName,
		SimpleWeaponProficiencyCategory,
		category,
		damageProfile,
		criticalProfile,
		rangeIncrement,
		cost,
		weight,
	)
	if !ok {
		panic("invalid core simple weapon seed")
	}

	return weapon
}

func mustNewCoreWeaponDamageProfile(
	smallDiceCount int,
	smallDieSides int,
	mediumDiceCount int,
	mediumDieSides int,
) WeaponDamageProfile {
	small := mustNewCoreWeaponDamageDice(smallDiceCount, smallDieSides)
	medium := mustNewCoreWeaponDamageDice(mediumDiceCount, mediumDieSides)

	damageProfile, ok := NewWeaponDamageProfile(small, medium)
	if !ok {
		panic("invalid core simple weapon damage profile seed")
	}

	return damageProfile
}

func mustNewCoreDoubleWeaponDamageProfile(
	smallDiceCount int,
	smallDieSides int,
	mediumDiceCount int,
	mediumDieSides int,
) WeaponDamageProfile {
	small := mustNewCoreWeaponDamageDice(smallDiceCount, smallDieSides)
	medium := mustNewCoreWeaponDamageDice(mediumDiceCount, mediumDieSides)

	damageProfile, ok := NewDoubleWeaponDamageProfile(small, medium, small, medium)
	if !ok {
		panic("invalid core simple double weapon damage profile seed")
	}

	return damageProfile
}

func mustNewCoreWeaponDamageDice(diceCount int, dieSides int) WeaponDamage {
	damage, ok := NewWeaponDamageDice(diceCount, dieSides)
	if !ok {
		panic("invalid core simple weapon damage seed")
	}

	return damage
}

func mustNewCoreWeaponCriticalProfile(
	threatMinimum int,
	multiplier int,
) WeaponCriticalProfile {
	criticalProfile, ok := NewWeaponCriticalProfile(threatMinimum, multiplier)
	if !ok {
		panic("invalid core simple weapon critical profile seed")
	}

	return criticalProfile
}

func mustNewCoreWeaponRangeIncrement(feet int) WeaponRangeIncrement {
	rangeIncrement, ok := NewWeaponRangeIncrementFeet(feet)
	if !ok {
		panic("invalid core simple weapon range increment seed")
	}

	return rangeIncrement
}
