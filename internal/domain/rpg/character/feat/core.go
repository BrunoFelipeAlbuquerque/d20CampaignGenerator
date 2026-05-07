package feat

import (
	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
	"d20campaigngenerator/internal/domain/rpg/character/spell"
)

const (
	AcrobaticFeatID                 FeatID = "Acrobatic"
	AlertnessFeatID                 FeatID = "Alertness"
	AlignmentChannelFeatID          FeatID = "Alignment Channel"
	AnimalAffinityFeatID            FeatID = "Animal Affinity"
	ArmorProficiencyLightFeatID     FeatID = "Armor Proficiency, Light"
	ArmorProficiencyMediumFeatID    FeatID = "Armor Proficiency, Medium"
	ArmorProficiencyHeavyFeatID     FeatID = "Armor Proficiency, Heavy"
	AthleticFeatID                  FeatID = "Athletic"
	AugmentSummoningFeatID          FeatID = "Augment Summoning"
	CombatCastingFeatID             FeatID = "Combat Casting"
	CommandUndeadFeatID             FeatID = "Command Undead"
	DeceitfulFeatID                 FeatID = "Deceitful"
	DeftHandsFeatID                 FeatID = "Deft Hands"
	ElementalChannelFeatID          FeatID = "Elemental Channel"
	EnduranceFeatID                 FeatID = "Endurance"
	DiehardFeatID                   FeatID = "Diehard"
	EschewMaterialsFeatID           FeatID = "Eschew Materials"
	ExtraChannelFeatID              FeatID = "Extra Channel"
	ExtraKiFeatID                   FeatID = "Extra Ki"
	ExtraLayOnHandsFeatID           FeatID = "Extra Lay On Hands"
	ExtraMercyFeatID                FeatID = "Extra Mercy"
	ExtraPerformanceFeatID          FeatID = "Extra Performance"
	ExtraRageFeatID                 FeatID = "Extra Rage"
	FleetFeatID                     FeatID = "Fleet"
	GreatFortitudeFeatID            FeatID = "Great Fortitude"
	ImprovedGreatFortitudeFeatID    FeatID = "Improved Great Fortitude"
	ImprovedChannelFeatID           FeatID = "Improved Channel"
	ImprovedCounterspellFeatID      FeatID = "Improved Counterspell"
	ImprovedFamiliarFeatID          FeatID = "Improved Familiar"
	IronWillFeatID                  FeatID = "Iron Will"
	ImprovedIronWillFeatID          FeatID = "Improved Iron Will"
	LeadershipFeatID                FeatID = "Leadership"
	LightningReflexesFeatID         FeatID = "Lightning Reflexes"
	ImprovedLightningReflexesFeatID FeatID = "Improved Lightning Reflexes"
	MagicalAptitudeFeatID           FeatID = "Magical Aptitude"
	MartialWeaponProficiencyFeatID  FeatID = "Martial Weapon Proficiency"
	MasterCraftsmanFeatID           FeatID = "Master Craftsman"
	NaturalSpellFeatID              FeatID = "Natural Spell"
	NimbleMovesFeatID               FeatID = "Nimble Moves"
	AcrobaticStepsFeatID            FeatID = "Acrobatic Steps"
	PersuasiveFeatID                FeatID = "Persuasive"
	RunFeatID                       FeatID = "Run"
	SelectiveChannelingFeatID       FeatID = "Selective Channeling"
	SelfSufficientFeatID            FeatID = "Self-Sufficient"
	ShieldProficiencyFeatID         FeatID = "Shield Proficiency"
	SimpleWeaponProficiencyFeatID   FeatID = "Simple Weapon Proficiency"
	SkillFocusFeatID                FeatID = "Skill Focus"
	SpellFocusFeatID                FeatID = "Spell Focus"
	GreaterSpellFocusFeatID         FeatID = "Greater Spell Focus"
	SpellMasteryFeatID              FeatID = "Spell Mastery"
	SpellPenetrationFeatID          FeatID = "Spell Penetration"
	GreaterSpellPenetrationFeatID   FeatID = "Greater Spell Penetration"
	StealthyFeatID                  FeatID = "Stealthy"
	ToughnessFeatID                 FeatID = "Toughness"
	TurnUndeadFeatID                FeatID = "Turn Undead"
)

var coreGeneralFeats = mustBuildCoreGeneralFeats()

var coreGeneralFeatOrder = []FeatID{
	AcrobaticFeatID,
	AlertnessFeatID,
	AlignmentChannelFeatID,
	AnimalAffinityFeatID,
	ArmorProficiencyLightFeatID,
	ArmorProficiencyMediumFeatID,
	ArmorProficiencyHeavyFeatID,
	AthleticFeatID,
	AugmentSummoningFeatID,
	CombatCastingFeatID,
	CommandUndeadFeatID,
	DeceitfulFeatID,
	DeftHandsFeatID,
	ElementalChannelFeatID,
	EnduranceFeatID,
	DiehardFeatID,
	EschewMaterialsFeatID,
	ExtraChannelFeatID,
	ExtraKiFeatID,
	ExtraLayOnHandsFeatID,
	ExtraMercyFeatID,
	ExtraPerformanceFeatID,
	ExtraRageFeatID,
	FleetFeatID,
	GreatFortitudeFeatID,
	ImprovedGreatFortitudeFeatID,
	ImprovedChannelFeatID,
	ImprovedCounterspellFeatID,
	ImprovedFamiliarFeatID,
	IronWillFeatID,
	ImprovedIronWillFeatID,
	LeadershipFeatID,
	LightningReflexesFeatID,
	ImprovedLightningReflexesFeatID,
	MagicalAptitudeFeatID,
	MartialWeaponProficiencyFeatID,
	MasterCraftsmanFeatID,
	NaturalSpellFeatID,
	NimbleMovesFeatID,
	AcrobaticStepsFeatID,
	PersuasiveFeatID,
	RunFeatID,
	SelectiveChannelingFeatID,
	SelfSufficientFeatID,
	ShieldProficiencyFeatID,
	SimpleWeaponProficiencyFeatID,
	SkillFocusFeatID,
	SpellFocusFeatID,
	GreaterSpellFocusFeatID,
	SpellMasteryFeatID,
	SpellPenetrationFeatID,
	GreaterSpellPenetrationFeatID,
	StealthyFeatID,
	ToughnessFeatID,
	TurnUndeadFeatID,
}

func mustBuildCoreGeneralFeats() map[FeatID]Feat {
	return map[FeatID]Feat{
		AcrobaticFeatID:              mustNewCoreGeneralFeat(AcrobaticFeatID),
		AlertnessFeatID:              mustNewCoreGeneralFeat(AlertnessFeatID),
		AlignmentChannelFeatID:       mustNewCoreGeneralFeat(AlignmentChannelFeatID, mustClassFeaturePrerequisite(characterclass.ChannelEnergyClassFeatureID)),
		AnimalAffinityFeatID:         mustNewCoreGeneralFeat(AnimalAffinityFeatID),
		ArmorProficiencyLightFeatID:  mustNewCoreGeneralFeat(ArmorProficiencyLightFeatID),
		ArmorProficiencyMediumFeatID: mustNewCoreGeneralFeat(ArmorProficiencyMediumFeatID, mustFeatPrerequisite(ArmorProficiencyLightFeatID)),
		ArmorProficiencyHeavyFeatID: mustNewCoreGeneralFeat(
			ArmorProficiencyHeavyFeatID,
			mustFeatPrerequisite(ArmorProficiencyLightFeatID),
			mustFeatPrerequisite(ArmorProficiencyMediumFeatID),
		),
		AthleticFeatID:         mustNewCoreGeneralFeat(AthleticFeatID),
		AugmentSummoningFeatID: mustNewCoreGeneralFeat(AugmentSummoningFeatID, mustSpellSchoolFeatPrerequisite(SpellFocusFeatID, spell.ConjurationSchoolID)),
		CombatCastingFeatID:    mustNewCoreGeneralFeat(CombatCastingFeatID),
		CommandUndeadFeatID:    mustNewCoreGeneralFeat(CommandUndeadFeatID, mustClassFeaturePrerequisite(characterclass.ChannelNegativeEnergyClassFeatureID)),
		DeceitfulFeatID:        mustNewCoreGeneralFeat(DeceitfulFeatID),
		DeftHandsFeatID:        mustNewCoreGeneralFeat(DeftHandsFeatID),
		ElementalChannelFeatID: mustNewCoreGeneralFeat(ElementalChannelFeatID, mustClassFeaturePrerequisite(characterclass.ChannelEnergyClassFeatureID)),
		EnduranceFeatID:        mustNewCoreGeneralFeat(EnduranceFeatID),
		DiehardFeatID:          mustNewCoreGeneralFeat(DiehardFeatID, mustFeatPrerequisite(EnduranceFeatID)),
		EschewMaterialsFeatID:  mustNewCoreGeneralFeat(EschewMaterialsFeatID),
		ExtraChannelFeatID:     mustNewCoreGeneralFeat(ExtraChannelFeatID, mustClassFeaturePrerequisite(characterclass.ChannelEnergyClassFeatureID)),
		ExtraKiFeatID:          mustNewCoreGeneralFeat(ExtraKiFeatID, mustClassFeaturePrerequisite(characterclass.KiPoolClassFeatureID)),
		ExtraLayOnHandsFeatID:  mustNewCoreGeneralFeat(ExtraLayOnHandsFeatID, mustClassFeaturePrerequisite(characterclass.LayOnHandsClassFeatureID)),
		ExtraMercyFeatID: mustNewCoreGeneralFeat(
			ExtraMercyFeatID,
			mustClassFeaturePrerequisite(characterclass.LayOnHandsClassFeatureID),
			mustClassFeaturePrerequisite(characterclass.MercyClassFeatureID),
		),
		ExtraPerformanceFeatID:          mustNewCoreGeneralFeat(ExtraPerformanceFeatID, mustClassFeaturePrerequisite(characterclass.BardicPerformanceClassFeatureID)),
		ExtraRageFeatID:                 mustNewCoreGeneralFeat(ExtraRageFeatID, mustClassFeaturePrerequisite(characterclass.RageClassFeatureID)),
		FleetFeatID:                     mustNewCoreGeneralFeat(FleetFeatID),
		GreatFortitudeFeatID:            mustNewCoreGeneralFeat(GreatFortitudeFeatID),
		ImprovedGreatFortitudeFeatID:    mustNewCoreGeneralFeat(ImprovedGreatFortitudeFeatID, mustFeatPrerequisite(GreatFortitudeFeatID)),
		ImprovedChannelFeatID:           mustNewCoreGeneralFeat(ImprovedChannelFeatID, mustClassFeaturePrerequisite(characterclass.ChannelEnergyClassFeatureID)),
		ImprovedCounterspellFeatID:      mustNewCoreGeneralFeat(ImprovedCounterspellFeatID),
		ImprovedFamiliarFeatID:          mustNewCoreGeneralFeat(ImprovedFamiliarFeatID, mustClassFeaturePrerequisite(characterclass.FamiliarAccessClassFeatureID), NewSelectedFamiliarEligibilityPrerequisite()),
		IronWillFeatID:                  mustNewCoreGeneralFeat(IronWillFeatID),
		ImprovedIronWillFeatID:          mustNewCoreGeneralFeat(ImprovedIronWillFeatID, mustFeatPrerequisite(IronWillFeatID)),
		LeadershipFeatID:                mustNewCoreGeneralFeat(LeadershipFeatID, mustCharacterLevelPrerequisite(7)),
		LightningReflexesFeatID:         mustNewCoreGeneralFeat(LightningReflexesFeatID),
		ImprovedLightningReflexesFeatID: mustNewCoreGeneralFeat(ImprovedLightningReflexesFeatID, mustFeatPrerequisite(LightningReflexesFeatID)),
		MagicalAptitudeFeatID:           mustNewCoreGeneralFeat(MagicalAptitudeFeatID),
		MartialWeaponProficiencyFeatID:  mustNewCoreGeneralFeat(MartialWeaponProficiencyFeatID),
		MasterCraftsmanFeatID: mustNewCoreGeneralFeat(
			MasterCraftsmanFeatID,
			mustAnySkillRanksPrerequisite([]skill.SkillID{skill.CraftSkillID, skill.ProfessionSkillID}, 5),
		),
		NaturalSpellFeatID: mustNewCoreGeneralFeat(
			NaturalSpellFeatID,
			mustAbilityScorePrerequisite(ability.WisdomScore, 13),
			mustClassFeaturePrerequisite(characterclass.WildShapeClassFeatureID),
		),
		NimbleMovesFeatID:         mustNewCoreGeneralFeat(NimbleMovesFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 13)),
		AcrobaticStepsFeatID:      mustNewCoreGeneralFeat(AcrobaticStepsFeatID, mustAbilityScorePrerequisite(ability.DexterityScore, 15), mustFeatPrerequisite(NimbleMovesFeatID)),
		PersuasiveFeatID:          mustNewCoreGeneralFeat(PersuasiveFeatID),
		RunFeatID:                 mustNewCoreGeneralFeat(RunFeatID),
		SelectiveChannelingFeatID: mustNewCoreGeneralFeat(SelectiveChannelingFeatID, mustAbilityScorePrerequisite(ability.CharismaScore, 13), mustClassFeaturePrerequisite(characterclass.ChannelEnergyClassFeatureID)),
		SelfSufficientFeatID:      mustNewCoreGeneralFeat(SelfSufficientFeatID),
		ShieldProficiencyFeatID:   mustNewCoreGeneralFeat(ShieldProficiencyFeatID),
		SimpleWeaponProficiencyFeatID: mustNewCoreGeneralFeat(
			SimpleWeaponProficiencyFeatID,
		),
		SkillFocusFeatID:              mustNewCoreGeneralFeat(SkillFocusFeatID),
		SpellFocusFeatID:              mustNewCoreGeneralFeat(SpellFocusFeatID),
		GreaterSpellFocusFeatID:       mustNewCoreGeneralFeat(GreaterSpellFocusFeatID, mustSameSelectionFeatPrerequisite(SpellFocusFeatID)),
		SpellMasteryFeatID:            mustNewCoreGeneralFeat(SpellMasteryFeatID, mustClassLevelPrerequisite(characterclass.WizardClassID, 1)),
		SpellPenetrationFeatID:        mustNewCoreGeneralFeat(SpellPenetrationFeatID),
		GreaterSpellPenetrationFeatID: mustNewCoreGeneralFeat(GreaterSpellPenetrationFeatID, mustFeatPrerequisite(SpellPenetrationFeatID)),
		StealthyFeatID:                mustNewCoreGeneralFeat(StealthyFeatID),
		ToughnessFeatID:               mustNewCoreGeneralFeat(ToughnessFeatID),
		TurnUndeadFeatID:              mustNewCoreGeneralFeat(TurnUndeadFeatID, mustClassFeaturePrerequisite(characterclass.ChannelPositiveEnergyClassFeatureID)),
	}
}

func mustNewCoreGeneralFeat(id FeatID, prerequisites ...Prerequisite) Feat {
	prerequisiteList, ok := NewPrerequisiteList(prerequisites)
	if !ok {
		panic("invalid core general feat prerequisite seed")
	}

	value, ok := NewFeat(id, GeneralFeatCategory, prerequisiteList, false, false, false)
	if !ok {
		panic("invalid core general feat seed")
	}

	return value
}

func mustAbilityScorePrerequisite(id ability.AbilityScoreID, minimumScore int) AbilityScorePrerequisite {
	value, ok := NewAbilityScorePrerequisite(id, minimumScore)
	if !ok {
		panic("invalid core ability score prerequisite seed")
	}

	return value
}

func mustFeatPrerequisite(id FeatID) FeatPrerequisite {
	value, ok := NewFeatPrerequisite(id)
	if !ok {
		panic("invalid core feat prerequisite seed")
	}

	return value
}

func mustSameSelectionFeatPrerequisite(id FeatID) SameSelectionFeatPrerequisite {
	value, ok := NewSameSelectionFeatPrerequisite(id)
	if !ok {
		panic("invalid core same-selection feat prerequisite seed")
	}

	return value
}

func mustSpellSchoolFeatPrerequisite(id FeatID, schoolID spell.SchoolID) SpellSchoolFeatPrerequisite {
	value, ok := NewSpellSchoolFeatPrerequisite(id, schoolID)
	if !ok {
		panic("invalid core spell-school feat prerequisite seed")
	}

	return value
}

func mustCharacterLevelPrerequisite(minimumLevel int) CharacterLevelPrerequisite {
	value, ok := NewCharacterLevelPrerequisite(minimumLevel)
	if !ok {
		panic("invalid core character level prerequisite seed")
	}

	return value
}

func mustClassLevelPrerequisite(
	id characterclass.ClassID,
	minimumLevel int,
) ClassLevelPrerequisite {
	value, ok := NewClassLevelPrerequisite(id, minimumLevel)
	if !ok {
		panic("invalid core class level prerequisite seed")
	}

	return value
}

func mustClassFeaturePrerequisite(
	id characterclass.ClassFeatureID,
) ClassFeaturePrerequisite {
	value, ok := NewClassFeaturePrerequisite(id)
	if !ok {
		panic("invalid core class feature prerequisite seed")
	}

	return value
}

func mustAnySkillRanksPrerequisite(
	ids []skill.SkillID,
	minimumRanks int,
) AnySkillRanksPrerequisite {
	value, ok := NewAnySkillRanksPrerequisite(ids, minimumRanks)
	if !ok {
		panic("invalid core any-skill ranks prerequisite seed")
	}

	return value
}
