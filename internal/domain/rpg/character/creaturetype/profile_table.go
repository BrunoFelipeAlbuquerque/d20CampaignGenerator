package creaturetype

import ability "d20campaigngenerator/internal/domain/rpg/character/ability"

var creatureTypeProfiles = map[CreatureTypeID]CreatureTypeProfile{
	AberrationType: mustNewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		[]ability.SavingThrowID{ability.WillSave},
		0,
		4,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{Darkvision60Trait, BreatheEatSleepTrait},
	),
	AnimalType: mustNewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		[]ability.SavingThrowID{ability.FortitudeSave, ability.ReflexSave},
		0,
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{LowLightVisionTrait, BreatheEatSleepTrait},
	),
	ConstructType: mustNewCreatureTypeProfile(
		ability.D10HitDie,
		ability.BaseAttackBonusFull,
		nil,
		0,
		2,
		ability.ConstructHitPoints,
		[]CreatureTypeTraitID{
			Darkvision60Trait,
			LowLightVisionTrait,
			NoConstitutionTrait,
			ImmunityMindAffectingTrait,
			ImmunityPoisonTrait,
			ImmunitySleepTrait,
			ImmunityParalysisTrait,
			ImmunityStunTrait,
			ImmunityDiseaseTrait,
			ImmunityDeathEffectsTrait,
			NotSubjectToNonlethalTrait,
			DestroyedAtZeroHPTrait,
		},
	),
	DragonType: mustNewCreatureTypeProfile(
		ability.D12HitDie,
		ability.BaseAttackBonusFull,
		[]ability.SavingThrowID{ability.FortitudeSave, ability.ReflexSave, ability.WillSave},
		0,
		6,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{
			Darkvision60Trait,
			LowLightVisionTrait,
			ImmunitySleepTrait,
			ImmunityParalysisTrait,
		},
	),
	FeyType: mustNewCreatureTypeProfile(
		ability.D6HitDie,
		ability.BaseAttackBonusHalf,
		[]ability.SavingThrowID{ability.ReflexSave, ability.WillSave},
		0,
		6,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{LowLightVisionTrait},
	),
	HumanoidType: mustNewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		nil,
		1,
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{BreatheEatSleepTrait},
	),
	MagicalBeastType: mustNewCreatureTypeProfile(
		ability.D10HitDie,
		ability.BaseAttackBonusFull,
		[]ability.SavingThrowID{ability.FortitudeSave, ability.ReflexSave},
		0,
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{Darkvision60Trait, LowLightVisionTrait},
	),
	MonstrousHumanoidType: mustNewCreatureTypeProfile(
		ability.D10HitDie,
		ability.BaseAttackBonusFull,
		[]ability.SavingThrowID{ability.ReflexSave, ability.WillSave},
		0,
		4,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{Darkvision60Trait},
	),
	OozeType: mustNewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		nil,
		0,
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{
			BlindTrait,
			Blindsight60Trait,
			NoIntelligenceTrait,
			MindlessTrait,
			ImmunityMindAffectingTrait,
			ImmunityPoisonTrait,
			ImmunitySleepTrait,
			ImmunityParalysisTrait,
			ImmunityStunTrait,
			ImmunityPolymorphTrait,
			NotSubjectToCriticalHitsTrait,
		},
	),
	OutsiderType: mustNewCreatureTypeProfile(
		ability.D10HitDie,
		ability.BaseAttackBonusFull,
		nil,
		2,
		6,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{
			Darkvision60Trait,
			NoNeedToEatSleepBreatheTrait,
		},
	),
	PlantType: mustNewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		[]ability.SavingThrowID{ability.FortitudeSave},
		0,
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{
			LowLightVisionTrait,
			ImmunityMindAffectingTrait,
			ImmunityPoisonTrait,
			ImmunitySleepTrait,
			ImmunityParalysisTrait,
			ImmunityStunTrait,
			ImmunityPolymorphTrait,
			NotSubjectToCriticalHitsTrait,
		},
	),
	UndeadType: mustNewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		[]ability.SavingThrowID{ability.WillSave},
		0,
		4,
		ability.UndeadHitPoints,
		[]CreatureTypeTraitID{
			Darkvision60Trait,
			NoConstitutionTrait,
			ImmunityMindAffectingTrait,
			ImmunityPoisonTrait,
			ImmunitySleepTrait,
			ImmunityParalysisTrait,
			ImmunityStunTrait,
			ImmunityDiseaseTrait,
			ImmunityDeathEffectsTrait,
			NotSubjectToNonlethalTrait,
			DestroyedAtZeroHPTrait,
		},
	),
	VerminType: mustNewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		[]ability.SavingThrowID{ability.FortitudeSave},
		0,
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{
			Darkvision60Trait,
			NoIntelligenceTrait,
			MindlessTrait,
		},
	),
}

func GetCreatureTypeProfile(typeID CreatureTypeID) (CreatureTypeProfile, bool) {
	profile, ok := creatureTypeProfiles[typeID]
	return profile, ok
}

func mustNewCreatureTypeProfile(
	hitDieType ability.HitDieType,
	babProgression ability.BaseAttackBonusProgression,
	fixedGoodSaves []ability.SavingThrowID,
	selectableGoodSaveCount int,
	skillPointsPerHD int,
	hitPointKind ability.HitPointKind,
	traitIDs []CreatureTypeTraitID,
) CreatureTypeProfile {
	profile, ok := NewCreatureTypeProfile(
		hitDieType,
		babProgression,
		fixedGoodSaves,
		selectableGoodSaveCount,
		skillPointsPerHD,
		hitPointKind,
		traitIDs,
	)
	if !ok {
		panic("invalid creature type profile")
	}

	return profile
}
