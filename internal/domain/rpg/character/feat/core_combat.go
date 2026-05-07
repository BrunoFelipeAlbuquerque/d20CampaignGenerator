package feat

import (
	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

const (
	AgileManeuversFeatID              FeatID = "Agile Maneuvers"
	ArcaneArmorTrainingFeatID         FeatID = "Arcane Armor Training"
	ArcaneArmorMasteryFeatID          FeatID = "Arcane Armor Mastery"
	ArcaneStrikeFeatID                FeatID = "Arcane Strike"
	BlindFightFeatID                  FeatID = "Blind-Fight"
	CatchOffGuardFeatID               FeatID = "Catch Off-Guard"
	ChannelSmiteFeatID                FeatID = "Channel Smite"
	CombatExpertiseFeatID             FeatID = "Combat Expertise"
	ImprovedDisarmFeatID              FeatID = "Improved Disarm"
	GreaterDisarmFeatID               FeatID = "Greater Disarm"
	ImprovedFeintFeatID               FeatID = "Improved Feint"
	GreaterFeintFeatID                FeatID = "Greater Feint"
	ImprovedTripFeatID                FeatID = "Improved Trip"
	GreaterTripFeatID                 FeatID = "Greater Trip"
	WhirlwindAttackFeatID             FeatID = "Whirlwind Attack"
	CombatReflexesFeatID              FeatID = "Combat Reflexes"
	StandStillFeatID                  FeatID = "Stand Still"
	CriticalFocusFeatID               FeatID = "Critical Focus"
	DeadlyAimFeatID                   FeatID = "Deadly Aim"
	DefensiveCombatTrainingFeatID     FeatID = "Defensive Combat Training"
	DisruptiveFeatID                  FeatID = "Disruptive"
	SpellbreakerFeatID                FeatID = "Spellbreaker"
	DodgeFeatID                       FeatID = "Dodge"
	MobilityFeatID                    FeatID = "Mobility"
	SpringAttackFeatID                FeatID = "Spring Attack"
	WindStanceFeatID                  FeatID = "Wind Stance"
	LightningStanceFeatID             FeatID = "Lightning Stance"
	ExoticWeaponProficiencyFeatID     FeatID = "Exotic Weapon Proficiency"
	ImprovedCriticalFeatID            FeatID = "Improved Critical"
	ImprovedInitiativeFeatID          FeatID = "Improved Initiative"
	ImprovedUnarmedStrikeFeatID       FeatID = "Improved Unarmed Strike"
	DeflectArrowsFeatID               FeatID = "Deflect Arrows"
	SnatchArrowsFeatID                FeatID = "Snatch Arrows"
	ImprovedGrappleFeatID             FeatID = "Improved Grapple"
	GreaterGrappleFeatID              FeatID = "Greater Grapple"
	ScorpionStyleFeatID               FeatID = "Scorpion Style"
	GorgonsFistFeatID                 FeatID = "Gorgon's Fist"
	MedusasWrathFeatID                FeatID = "Medusa's Wrath"
	StunningFistFeatID                FeatID = "Stunning Fist"
	ImprovisedWeaponMasteryFeatID     FeatID = "Improvised Weapon Mastery"
	IntimidatingProwessFeatID         FeatID = "Intimidating Prowess"
	LungeFeatID                       FeatID = "Lunge"
	MountedCombatFeatID               FeatID = "Mounted Combat"
	MountedArcheryFeatID              FeatID = "Mounted Archery"
	RideByAttackFeatID                FeatID = "Ride-By Attack"
	SpiritedChargeFeatID              FeatID = "Spirited Charge"
	TrampleFeatID                     FeatID = "Trample"
	UnseatFeatID                      FeatID = "Unseat"
	PointBlankShotFeatID              FeatID = "Point-Blank Shot"
	FarShotFeatID                     FeatID = "Far Shot"
	PreciseShotFeatID                 FeatID = "Precise Shot"
	ImprovedPreciseShotFeatID         FeatID = "Improved Precise Shot"
	PinpointTargetingFeatID           FeatID = "Pinpoint Targeting"
	ShotOnTheRunFeatID                FeatID = "Shot on the Run"
	RapidShotFeatID                   FeatID = "Rapid Shot"
	ManyshotFeatID                    FeatID = "Manyshot"
	PowerAttackFeatID                 FeatID = "Power Attack"
	CleaveFeatID                      FeatID = "Cleave"
	GreatCleaveFeatID                 FeatID = "Great Cleave"
	ImprovedBullRushFeatID            FeatID = "Improved Bull Rush"
	GreaterBullRushFeatID             FeatID = "Greater Bull Rush"
	ImprovedOverrunFeatID             FeatID = "Improved Overrun"
	GreaterOverrunFeatID              FeatID = "Greater Overrun"
	ImprovedSunderFeatID              FeatID = "Improved Sunder"
	GreaterSunderFeatID               FeatID = "Greater Sunder"
	QuickDrawFeatID                   FeatID = "Quick Draw"
	RapidReloadFeatID                 FeatID = "Rapid Reload"
	ImprovedShieldBashFeatID          FeatID = "Improved Shield Bash"
	ShieldSlamFeatID                  FeatID = "Shield Slam"
	ShieldMasterFeatID                FeatID = "Shield Master"
	ShieldFocusFeatID                 FeatID = "Shield Focus"
	GreaterShieldFocusFeatID          FeatID = "Greater Shield Focus"
	TowerShieldProficiencyFeatID      FeatID = "Tower Shield Proficiency"
	StepUpFeatID                      FeatID = "Step Up"
	StrikeBackFeatID                  FeatID = "Strike Back"
	ThrowAnythingFeatID               FeatID = "Throw Anything"
	TwoWeaponFightingFeatID           FeatID = "Two-Weapon Fighting"
	DoubleSliceFeatID                 FeatID = "Double Slice"
	TwoWeaponRendFeatID               FeatID = "Two-Weapon Rend"
	ImprovedTwoWeaponFightingFeatID   FeatID = "Improved Two-Weapon Fighting"
	GreaterTwoWeaponFightingFeatID    FeatID = "Greater Two-Weapon Fighting"
	TwoWeaponDefenseFeatID            FeatID = "Two-Weapon Defense"
	VitalStrikeFeatID                 FeatID = "Vital Strike"
	ImprovedVitalStrikeFeatID         FeatID = "Improved Vital Strike"
	GreaterVitalStrikeFeatID          FeatID = "Greater Vital Strike"
	WeaponFinesseFeatID               FeatID = "Weapon Finesse"
	WeaponFocusFeatID                 FeatID = "Weapon Focus"
	DazzlingDisplayFeatID             FeatID = "Dazzling Display"
	ShatterDefensesFeatID             FeatID = "Shatter Defenses"
	DeadlyStrokeFeatID                FeatID = "Deadly Stroke"
	GreaterWeaponFocusFeatID          FeatID = "Greater Weapon Focus"
	PenetratingStrikeFeatID           FeatID = "Penetrating Strike"
	GreaterPenetratingStrikeFeatID    FeatID = "Greater Penetrating Strike"
	WeaponSpecializationFeatID        FeatID = "Weapon Specialization"
	GreaterWeaponSpecializationFeatID FeatID = "Greater Weapon Specialization"
)

var coreCombatFeats = mustBuildCoreCombatFeats()

var coreCombatFeatOrder = []FeatID{
	AgileManeuversFeatID,
	ArcaneArmorTrainingFeatID,
	ArcaneArmorMasteryFeatID,
	ArcaneStrikeFeatID,
	BlindFightFeatID,
	CatchOffGuardFeatID,
	ChannelSmiteFeatID,
	CombatExpertiseFeatID,
	ImprovedDisarmFeatID,
	GreaterDisarmFeatID,
	ImprovedFeintFeatID,
	GreaterFeintFeatID,
	ImprovedTripFeatID,
	GreaterTripFeatID,
	WhirlwindAttackFeatID,
	CombatReflexesFeatID,
	StandStillFeatID,
	CriticalFocusFeatID,
	DeadlyAimFeatID,
	DefensiveCombatTrainingFeatID,
	DisruptiveFeatID,
	SpellbreakerFeatID,
	DodgeFeatID,
	MobilityFeatID,
	SpringAttackFeatID,
	WindStanceFeatID,
	LightningStanceFeatID,
	ExoticWeaponProficiencyFeatID,
	ImprovedCriticalFeatID,
	ImprovedInitiativeFeatID,
	ImprovedUnarmedStrikeFeatID,
	DeflectArrowsFeatID,
	SnatchArrowsFeatID,
	ImprovedGrappleFeatID,
	GreaterGrappleFeatID,
	ScorpionStyleFeatID,
	GorgonsFistFeatID,
	MedusasWrathFeatID,
	StunningFistFeatID,
	ImprovisedWeaponMasteryFeatID,
	IntimidatingProwessFeatID,
	LungeFeatID,
	MountedCombatFeatID,
	MountedArcheryFeatID,
	RideByAttackFeatID,
	SpiritedChargeFeatID,
	TrampleFeatID,
	UnseatFeatID,
	PointBlankShotFeatID,
	FarShotFeatID,
	PreciseShotFeatID,
	ImprovedPreciseShotFeatID,
	PinpointTargetingFeatID,
	ShotOnTheRunFeatID,
	RapidShotFeatID,
	ManyshotFeatID,
	PowerAttackFeatID,
	CleaveFeatID,
	GreatCleaveFeatID,
	ImprovedBullRushFeatID,
	GreaterBullRushFeatID,
	ImprovedOverrunFeatID,
	GreaterOverrunFeatID,
	ImprovedSunderFeatID,
	GreaterSunderFeatID,
	QuickDrawFeatID,
	RapidReloadFeatID,
	ImprovedShieldBashFeatID,
	ShieldSlamFeatID,
	ShieldMasterFeatID,
	ShieldFocusFeatID,
	GreaterShieldFocusFeatID,
	TowerShieldProficiencyFeatID,
	StepUpFeatID,
	StrikeBackFeatID,
	ThrowAnythingFeatID,
	TwoWeaponFightingFeatID,
	DoubleSliceFeatID,
	TwoWeaponRendFeatID,
	ImprovedTwoWeaponFightingFeatID,
	GreaterTwoWeaponFightingFeatID,
	TwoWeaponDefenseFeatID,
	VitalStrikeFeatID,
	ImprovedVitalStrikeFeatID,
	GreaterVitalStrikeFeatID,
	WeaponFinesseFeatID,
	WeaponFocusFeatID,
	DazzlingDisplayFeatID,
	ShatterDefensesFeatID,
	DeadlyStrokeFeatID,
	GreaterWeaponFocusFeatID,
	PenetratingStrikeFeatID,
	GreaterPenetratingStrikeFeatID,
	WeaponSpecializationFeatID,
	GreaterWeaponSpecializationFeatID,
}

func mustBuildCoreCombatFeats() map[FeatID]Feat {
	return map[FeatID]Feat{
		AgileManeuversFeatID:      mustNewCoreCombatFeat(AgileManeuversFeatID),
		ArcaneArmorTrainingFeatID: mustNewCoreCombatFeat(ArcaneArmorTrainingFeatID, mustFeatPrerequisite(ArmorProficiencyLightFeatID), mustCasterLevelPrerequisite(3)),
		ArcaneArmorMasteryFeatID: mustNewCoreCombatFeat(
			ArcaneArmorMasteryFeatID,
			mustFeatPrerequisite(ArcaneArmorTrainingFeatID),
			mustFeatPrerequisite(ArmorProficiencyMediumFeatID),
			mustCasterLevelPrerequisite(7),
		),
		ArcaneStrikeFeatID:    mustNewCoreCombatFeat(ArcaneStrikeFeatID, mustSpellcastingPrerequisite(ArcaneSpellcastingAccess)),
		BlindFightFeatID:      mustNewCoreCombatFeat(BlindFightFeatID),
		CatchOffGuardFeatID:   mustNewCoreCombatFeat(CatchOffGuardFeatID),
		ChannelSmiteFeatID:    mustNewCoreCombatFeat(ChannelSmiteFeatID, mustClassFeaturePrerequisite(characterclass.ChannelEnergyClassFeatureID)),
		CombatExpertiseFeatID: mustNewCoreCombatFeat(CombatExpertiseFeatID, mustAbilityScorePrerequisite(ability.IntelligenceScore, 13)),
		ImprovedDisarmFeatID:  mustNewCoreCombatFeat(ImprovedDisarmFeatID, mustAbilityScorePrerequisite(ability.IntelligenceScore, 13), mustFeatPrerequisite(CombatExpertiseFeatID)),
		GreaterDisarmFeatID:   mustNewCoreCombatFeat(GreaterDisarmFeatID, mustAbilityScorePrerequisite(ability.IntelligenceScore, 13), mustFeatPrerequisite(CombatExpertiseFeatID), mustFeatPrerequisite(ImprovedDisarmFeatID), mustBaseAttackBonusPrerequisite(6)),
		ImprovedFeintFeatID:   mustNewCoreCombatFeat(ImprovedFeintFeatID, mustAbilityScorePrerequisite(ability.IntelligenceScore, 13), mustFeatPrerequisite(CombatExpertiseFeatID)),
		GreaterFeintFeatID:    mustNewCoreCombatFeat(GreaterFeintFeatID, mustAbilityScorePrerequisite(ability.IntelligenceScore, 13), mustFeatPrerequisite(CombatExpertiseFeatID), mustFeatPrerequisite(ImprovedFeintFeatID), mustBaseAttackBonusPrerequisite(6)),
		ImprovedTripFeatID:    mustNewCoreCombatFeat(ImprovedTripFeatID, mustAbilityScorePrerequisite(ability.IntelligenceScore, 13), mustFeatPrerequisite(CombatExpertiseFeatID)),
		GreaterTripFeatID:     mustNewCoreCombatFeat(GreaterTripFeatID, mustAbilityScorePrerequisite(ability.IntelligenceScore, 13), mustFeatPrerequisite(CombatExpertiseFeatID), mustFeatPrerequisite(ImprovedTripFeatID), mustBaseAttackBonusPrerequisite(6)),
		WhirlwindAttackFeatID: mustNewCoreCombatFeat(WhirlwindAttackFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustAbilityScorePrerequisite(ability.IntelligenceScore, 13), mustFeatPrerequisite(CombatExpertiseFeatID), mustFeatPrerequisite(DodgeFeatID), mustFeatPrerequisite(MobilityFeatID), mustFeatPrerequisite(SpringAttackFeatID), mustBaseAttackBonusPrerequisite(4)),
		CombatReflexesFeatID:  mustNewCoreCombatFeat(CombatReflexesFeatID),
		StandStillFeatID:      mustNewCoreCombatFeat(StandStillFeatID, mustFeatPrerequisite(CombatReflexesFeatID)),
		CriticalFocusFeatID:   mustNewCoreCombatFeat(CriticalFocusFeatID, mustBaseAttackBonusPrerequisite(9)),
		DeadlyAimFeatID:       mustNewCoreCombatFeat(DeadlyAimFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustBaseAttackBonusPrerequisite(1)),
		DefensiveCombatTrainingFeatID: mustNewCoreCombatFeat(
			DefensiveCombatTrainingFeatID,
		),
		DisruptiveFeatID:      mustNewCoreCombatFeat(DisruptiveFeatID, mustClassLevelPrerequisite(characterclass.FighterClassID, 6)),
		SpellbreakerFeatID:    mustNewCoreCombatFeat(SpellbreakerFeatID, mustFeatPrerequisite(DisruptiveFeatID), mustClassLevelPrerequisite(characterclass.FighterClassID, 10)),
		DodgeFeatID:           mustNewCoreCombatFeat(DodgeFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13)),
		MobilityFeatID:        mustNewCoreCombatFeat(MobilityFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustFeatPrerequisite(DodgeFeatID)),
		SpringAttackFeatID:    mustNewCoreCombatFeat(SpringAttackFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustFeatPrerequisite(DodgeFeatID), mustFeatPrerequisite(MobilityFeatID), mustBaseAttackBonusPrerequisite(4)),
		WindStanceFeatID:      mustNewCoreCombatFeat(WindStanceFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 15), mustFeatPrerequisite(DodgeFeatID), mustBaseAttackBonusPrerequisite(6)),
		LightningStanceFeatID: mustNewCoreCombatFeat(LightningStanceFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 17), mustFeatPrerequisite(DodgeFeatID), mustFeatPrerequisite(WindStanceFeatID), mustBaseAttackBonusPrerequisite(11)),
		ExoticWeaponProficiencyFeatID: mustNewCoreCombatFeat(
			ExoticWeaponProficiencyFeatID,
			mustBaseAttackBonusPrerequisite(1),
		),
		ImprovedCriticalFeatID:      mustNewCoreCombatFeat(ImprovedCriticalFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustBaseAttackBonusPrerequisite(8)),
		ImprovedInitiativeFeatID:    mustNewCoreCombatFeat(ImprovedInitiativeFeatID),
		ImprovedUnarmedStrikeFeatID: mustNewCoreCombatFeat(ImprovedUnarmedStrikeFeatID),
		DeflectArrowsFeatID:         mustNewCoreCombatFeat(DeflectArrowsFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID)),
		SnatchArrowsFeatID:          mustNewCoreCombatFeat(SnatchArrowsFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 15), mustFeatPrerequisite(DeflectArrowsFeatID), mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID)),
		ImprovedGrappleFeatID:       mustNewCoreCombatFeat(ImprovedGrappleFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID)),
		GreaterGrappleFeatID:        mustNewCoreCombatFeat(GreaterGrappleFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustFeatPrerequisite(ImprovedGrappleFeatID), mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID), mustBaseAttackBonusPrerequisite(6)),
		ScorpionStyleFeatID:         mustNewCoreCombatFeat(ScorpionStyleFeatID, mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID)),
		GorgonsFistFeatID:           mustNewCoreCombatFeat(GorgonsFistFeatID, mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID), mustFeatPrerequisite(ScorpionStyleFeatID), mustBaseAttackBonusPrerequisite(6)),
		MedusasWrathFeatID:          mustNewCoreCombatFeat(MedusasWrathFeatID, mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID), mustFeatPrerequisite(GorgonsFistFeatID), mustFeatPrerequisite(ScorpionStyleFeatID), mustBaseAttackBonusPrerequisite(11)),
		StunningFistFeatID:          mustNewCoreCombatFeat(StunningFistFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustAbilityScorePrerequisite(ability.WisdomScore, 13), mustFeatPrerequisite(ImprovedUnarmedStrikeFeatID), mustBaseAttackBonusPrerequisite(8)),
		ImprovisedWeaponMasteryFeatID: mustNewCoreCombatFeat(
			ImprovisedWeaponMasteryFeatID,
			mustAnyFeatPrerequisite([]FeatID{CatchOffGuardFeatID, ThrowAnythingFeatID}),
			mustBaseAttackBonusPrerequisite(8),
		),
		IntimidatingProwessFeatID: mustNewCoreCombatFeat(IntimidatingProwessFeatID),
		LungeFeatID:               mustNewCoreCombatFeat(LungeFeatID, mustBaseAttackBonusPrerequisite(6)),
		MountedCombatFeatID:       mustNewCoreCombatFeat(MountedCombatFeatID, mustSkillRanksPrerequisite(skill.RideSkillID, 1)),
		MountedArcheryFeatID:      mustNewCoreCombatFeat(MountedArcheryFeatID, mustSkillRanksPrerequisite(skill.RideSkillID, 1), mustFeatPrerequisite(MountedCombatFeatID)),
		RideByAttackFeatID:        mustNewCoreCombatFeat(RideByAttackFeatID, mustSkillRanksPrerequisite(skill.RideSkillID, 1), mustFeatPrerequisite(MountedCombatFeatID)),
		SpiritedChargeFeatID:      mustNewCoreCombatFeat(SpiritedChargeFeatID, mustSkillRanksPrerequisite(skill.RideSkillID, 1), mustFeatPrerequisite(MountedCombatFeatID), mustFeatPrerequisite(RideByAttackFeatID)),
		TrampleFeatID:             mustNewCoreCombatFeat(TrampleFeatID, mustSkillRanksPrerequisite(skill.RideSkillID, 1), mustFeatPrerequisite(MountedCombatFeatID)),
		UnseatFeatID:              mustNewCoreCombatFeat(UnseatFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustSkillRanksPrerequisite(skill.RideSkillID, 1), mustFeatPrerequisite(MountedCombatFeatID), mustFeatPrerequisite(PowerAttackFeatID), mustFeatPrerequisite(ImprovedBullRushFeatID), mustBaseAttackBonusPrerequisite(1)),
		PointBlankShotFeatID:      mustNewCoreCombatFeat(PointBlankShotFeatID),
		FarShotFeatID:             mustNewCoreCombatFeat(FarShotFeatID, mustFeatPrerequisite(PointBlankShotFeatID)),
		PreciseShotFeatID:         mustNewCoreCombatFeat(PreciseShotFeatID, mustFeatPrerequisite(PointBlankShotFeatID)),
		ImprovedPreciseShotFeatID: mustNewCoreCombatFeat(ImprovedPreciseShotFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 19), mustFeatPrerequisite(PointBlankShotFeatID), mustFeatPrerequisite(PreciseShotFeatID), mustBaseAttackBonusPrerequisite(11)),
		PinpointTargetingFeatID:   mustNewCoreCombatFeat(PinpointTargetingFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 19), mustFeatPrerequisite(ImprovedPreciseShotFeatID), mustFeatPrerequisite(PointBlankShotFeatID), mustFeatPrerequisite(PreciseShotFeatID), mustBaseAttackBonusPrerequisite(16)),
		ShotOnTheRunFeatID:        mustNewCoreCombatFeat(ShotOnTheRunFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustFeatPrerequisite(DodgeFeatID), mustFeatPrerequisite(MobilityFeatID), mustFeatPrerequisite(PointBlankShotFeatID), mustBaseAttackBonusPrerequisite(4)),
		RapidShotFeatID:           mustNewCoreCombatFeat(RapidShotFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13), mustFeatPrerequisite(PointBlankShotFeatID)),
		ManyshotFeatID:            mustNewCoreCombatFeat(ManyshotFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 17), mustFeatPrerequisite(PointBlankShotFeatID), mustFeatPrerequisite(RapidShotFeatID), mustBaseAttackBonusPrerequisite(6)),
		PowerAttackFeatID:         mustNewCoreCombatFeat(PowerAttackFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustBaseAttackBonusPrerequisite(1)),
		CleaveFeatID:              mustNewCoreCombatFeat(CleaveFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(1)),
		GreatCleaveFeatID:         mustNewCoreCombatFeat(GreatCleaveFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(CleaveFeatID), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(4)),
		ImprovedBullRushFeatID:    mustNewCoreCombatFeat(ImprovedBullRushFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(1)),
		GreaterBullRushFeatID:     mustNewCoreCombatFeat(GreaterBullRushFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(ImprovedBullRushFeatID), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(6)),
		ImprovedOverrunFeatID:     mustNewCoreCombatFeat(ImprovedOverrunFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(1)),
		GreaterOverrunFeatID:      mustNewCoreCombatFeat(GreaterOverrunFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(ImprovedOverrunFeatID), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(6)),
		ImprovedSunderFeatID:      mustNewCoreCombatFeat(ImprovedSunderFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(1)),
		GreaterSunderFeatID:       mustNewCoreCombatFeat(GreaterSunderFeatID, mustAbilityScorePrerequisite(ability.StrengthScore, 13), mustFeatPrerequisite(ImprovedSunderFeatID), mustFeatPrerequisite(PowerAttackFeatID), mustBaseAttackBonusPrerequisite(6)),
		QuickDrawFeatID:           mustNewCoreCombatFeat(QuickDrawFeatID, mustBaseAttackBonusPrerequisite(1)),
		RapidReloadFeatID:         mustNewCoreCombatFeat(RapidReloadFeatID, NewSelectedWeaponProficiencyPrerequisite()),
		ImprovedShieldBashFeatID:  mustNewCoreCombatFeat(ImprovedShieldBashFeatID, mustFeatPrerequisite(ShieldProficiencyFeatID)),
		ShieldSlamFeatID:          mustNewCoreCombatFeat(ShieldSlamFeatID, mustFeatPrerequisite(ImprovedShieldBashFeatID), mustFeatPrerequisite(ShieldProficiencyFeatID), mustFeatPrerequisite(TwoWeaponFightingFeatID), mustBaseAttackBonusPrerequisite(6)),
		ShieldMasterFeatID:        mustNewCoreCombatFeat(ShieldMasterFeatID, mustFeatPrerequisite(ImprovedShieldBashFeatID), mustFeatPrerequisite(ShieldProficiencyFeatID), mustFeatPrerequisite(ShieldSlamFeatID), mustFeatPrerequisite(TwoWeaponFightingFeatID), mustBaseAttackBonusPrerequisite(11)),
		ShieldFocusFeatID:         mustNewCoreCombatFeat(ShieldFocusFeatID, mustFeatPrerequisite(ShieldProficiencyFeatID), mustBaseAttackBonusPrerequisite(1)),
		GreaterShieldFocusFeatID:  mustNewCoreCombatFeat(GreaterShieldFocusFeatID, mustFeatPrerequisite(ShieldFocusFeatID), mustFeatPrerequisite(ShieldProficiencyFeatID), mustBaseAttackBonusPrerequisite(1), mustClassLevelPrerequisite(characterclass.FighterClassID, 8)),
		TowerShieldProficiencyFeatID: mustNewCoreCombatFeat(
			TowerShieldProficiencyFeatID,
			mustFeatPrerequisite(ShieldProficiencyFeatID),
		),
		StepUpFeatID:            mustNewCoreCombatFeat(StepUpFeatID, mustBaseAttackBonusPrerequisite(1)),
		StrikeBackFeatID:        mustNewCoreCombatFeat(StrikeBackFeatID, mustBaseAttackBonusPrerequisite(11)),
		ThrowAnythingFeatID:     mustNewCoreCombatFeat(ThrowAnythingFeatID),
		TwoWeaponFightingFeatID: mustNewCoreCombatFeat(TwoWeaponFightingFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 15)),
		DoubleSliceFeatID:       mustNewCoreCombatFeat(DoubleSliceFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 15), mustFeatPrerequisite(TwoWeaponFightingFeatID)),
		TwoWeaponRendFeatID:     mustNewCoreCombatFeat(TwoWeaponRendFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 17), mustFeatPrerequisite(DoubleSliceFeatID), mustFeatPrerequisite(ImprovedTwoWeaponFightingFeatID), mustFeatPrerequisite(TwoWeaponFightingFeatID), mustBaseAttackBonusPrerequisite(11)),
		ImprovedTwoWeaponFightingFeatID: mustNewCoreCombatFeat(
			ImprovedTwoWeaponFightingFeatID,
			mustAbilityScorePrerequisite(ability.DexterityScore, 17),
			mustFeatPrerequisite(TwoWeaponFightingFeatID),
			mustBaseAttackBonusPrerequisite(6),
		),
		GreaterTwoWeaponFightingFeatID: mustNewCoreCombatFeat(
			GreaterTwoWeaponFightingFeatID,
			mustAbilityScorePrerequisite(ability.DexterityScore, 19),
			mustFeatPrerequisite(ImprovedTwoWeaponFightingFeatID),
			mustFeatPrerequisite(TwoWeaponFightingFeatID),
			mustBaseAttackBonusPrerequisite(11),
		),
		TwoWeaponDefenseFeatID:    mustNewCoreCombatFeat(TwoWeaponDefenseFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 15), mustFeatPrerequisite(TwoWeaponFightingFeatID)),
		VitalStrikeFeatID:         mustNewCoreCombatFeat(VitalStrikeFeatID, mustBaseAttackBonusPrerequisite(6)),
		ImprovedVitalStrikeFeatID: mustNewCoreCombatFeat(ImprovedVitalStrikeFeatID, mustFeatPrerequisite(VitalStrikeFeatID), mustBaseAttackBonusPrerequisite(11)),
		GreaterVitalStrikeFeatID:  mustNewCoreCombatFeat(GreaterVitalStrikeFeatID, mustFeatPrerequisite(ImprovedVitalStrikeFeatID), mustFeatPrerequisite(VitalStrikeFeatID), mustBaseAttackBonusPrerequisite(16)),
		WeaponFinesseFeatID:       mustNewCoreCombatFeat(WeaponFinesseFeatID),
		WeaponFocusFeatID:         mustNewCoreCombatFeat(WeaponFocusFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustBaseAttackBonusPrerequisite(1)),
		DazzlingDisplayFeatID:     mustNewCoreCombatFeat(DazzlingDisplayFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustSameSelectionFeatPrerequisite(WeaponFocusFeatID)),
		ShatterDefensesFeatID:     mustNewCoreCombatFeat(ShatterDefensesFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustSameSelectionFeatPrerequisite(WeaponFocusFeatID), mustSameSelectionFeatPrerequisite(DazzlingDisplayFeatID), mustBaseAttackBonusPrerequisite(6)),
		DeadlyStrokeFeatID:        mustNewCoreCombatFeat(DeadlyStrokeFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustSameSelectionFeatPrerequisite(DazzlingDisplayFeatID), mustSameSelectionFeatPrerequisite(GreaterWeaponFocusFeatID), mustSameSelectionFeatPrerequisite(ShatterDefensesFeatID), mustSameSelectionFeatPrerequisite(WeaponFocusFeatID), mustBaseAttackBonusPrerequisite(11)),
		GreaterWeaponFocusFeatID:  mustNewCoreCombatFeat(GreaterWeaponFocusFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustSameSelectionFeatPrerequisite(WeaponFocusFeatID), mustBaseAttackBonusPrerequisite(1), mustClassLevelPrerequisite(characterclass.FighterClassID, 8)),
		PenetratingStrikeFeatID:   mustNewCoreCombatFeat(PenetratingStrikeFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustSameSelectionFeatPrerequisite(WeaponFocusFeatID), mustBaseAttackBonusPrerequisite(1), mustClassLevelPrerequisite(characterclass.FighterClassID, 12)),
		GreaterPenetratingStrikeFeatID: mustNewCoreCombatFeat(
			GreaterPenetratingStrikeFeatID,
			mustSameSelectionFeatPrerequisite(PenetratingStrikeFeatID),
			mustSameSelectionFeatPrerequisite(WeaponFocusFeatID),
			mustClassLevelPrerequisite(characterclass.FighterClassID, 16),
		),
		WeaponSpecializationFeatID: mustNewCoreCombatFeat(WeaponSpecializationFeatID, NewSelectedWeaponProficiencyPrerequisite(), mustSameSelectionFeatPrerequisite(WeaponFocusFeatID), mustClassLevelPrerequisite(characterclass.FighterClassID, 4)),
		GreaterWeaponSpecializationFeatID: mustNewCoreCombatFeat(
			GreaterWeaponSpecializationFeatID,
			NewSelectedWeaponProficiencyPrerequisite(),
			mustSameSelectionFeatPrerequisite(GreaterWeaponFocusFeatID),
			mustSameSelectionFeatPrerequisite(WeaponFocusFeatID),
			mustSameSelectionFeatPrerequisite(WeaponSpecializationFeatID),
			mustClassLevelPrerequisite(characterclass.FighterClassID, 12),
		),
	}
}

func mustNewCoreCombatFeat(id FeatID, prerequisites ...Prerequisite) Feat {
	prerequisiteList, ok := NewPrerequisiteList(prerequisites)
	if !ok {
		panic("invalid core combat feat prerequisite seed")
	}

	value, ok := NewFeat(id, CombatFeatCategory, prerequisiteList, true, false, false)
	if !ok {
		panic("invalid core combat feat seed")
	}

	return value
}

func mustBaseAttackBonusPrerequisite(minimumBonus int) BaseAttackBonusPrerequisite {
	value, ok := NewBaseAttackBonusPrerequisite(minimumBonus)
	if !ok {
		panic("invalid core base attack bonus prerequisite seed")
	}

	return value
}

func mustCasterLevelPrerequisite(minimumLevel int) CasterLevelPrerequisite {
	value, ok := NewCasterLevelPrerequisite(minimumLevel)
	if !ok {
		panic("invalid core caster level prerequisite seed")
	}

	return value
}

func mustSkillRanksPrerequisite(id skill.SkillID, minimumRanks int) SkillRanksPrerequisite {
	value, ok := NewSkillRanksPrerequisite(id, minimumRanks)
	if !ok {
		panic("invalid core skill ranks prerequisite seed")
	}

	return value
}

func mustSpellcastingPrerequisite(access SpellcastingAccess) SpellcastingPrerequisite {
	value, ok := NewSpellcastingPrerequisite(access)
	if !ok {
		panic("invalid core spellcasting prerequisite seed")
	}

	return value
}

func mustAnyFeatPrerequisite(ids []FeatID) AnyFeatPrerequisite {
	value, ok := NewAnyFeatPrerequisite(ids)
	if !ok {
		panic("invalid core any-feat prerequisite seed")
	}

	return value
}
