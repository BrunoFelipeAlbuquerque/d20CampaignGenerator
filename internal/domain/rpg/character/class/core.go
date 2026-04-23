package class

import (
	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

const (
	BarbarianClassID ClassID = "barbarian"
	BardClassID      ClassID = "bard"
	ClericClassID    ClassID = "cleric"
	DruidClassID     ClassID = "druid"
	FighterClassID   ClassID = "fighter"
	MonkClassID      ClassID = "monk"
	PaladinClassID   ClassID = "paladin"
	RangerClassID    ClassID = "ranger"
	RogueClassID     ClassID = "rogue"
	SorcererClassID  ClassID = "sorcerer"
	WizardClassID    ClassID = "wizard"
)

const (
	knowledgeArcanaSkillID        skill.SkillID = "Knowledge (arcana)"
	knowledgeDungeoneeringSkillID skill.SkillID = "Knowledge (dungeoneering)"
	knowledgeEngineeringSkillID   skill.SkillID = "Knowledge (engineering)"
	knowledgeGeographySkillID     skill.SkillID = "Knowledge (geography)"
	knowledgeHistorySkillID       skill.SkillID = "Knowledge (history)"
	knowledgeLocalSkillID         skill.SkillID = "Knowledge (local)"
	knowledgeNatureSkillID        skill.SkillID = "Knowledge (nature)"
	knowledgeNobilitySkillID      skill.SkillID = "Knowledge (nobility)"
	knowledgePlanesSkillID        skill.SkillID = "Knowledge (planes)"
	knowledgeReligionSkillID      skill.SkillID = "Knowledge (religion)"
)

var coreClasses = mustBuildCoreClasses()

var coreClassOrder = []ClassID{
	BarbarianClassID,
	BardClassID,
	ClericClassID,
	DruidClassID,
	FighterClassID,
	MonkClassID,
	PaladinClassID,
	RangerClassID,
	RogueClassID,
	SorcererClassID,
	WizardClassID,
}

func mustBuildCoreClasses() map[ClassID]Class {
	return map[ClassID]Class{
		BarbarianClassID: mustNewClass(
			BarbarianClassID,
			ability.D12HitDie,
			ability.BaseAttackBonusFull,
			mustNewSaveProgressions(ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowPoor),
			4,
			[]skill.SkillID{
				skill.AcrobaticsSkillID,
				skill.ClimbSkillID,
				skill.CraftSkillID,
				skill.HandleAnimalSkillID,
				skill.IntimidateSkillID,
				knowledgeNatureSkillID,
				skill.PerceptionSkillID,
				skill.RideSkillID,
				skill.SurvivalSkillID,
				skill.SwimSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
				MartialWeaponsWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
				MediumArmorProficiencyID,
				ShieldArmorProficiencyID,
			},
			NewNonSpellcastingProfile(),
		),
		BardClassID: mustNewClass(
			BardClassID,
			ability.D8HitDie,
			ability.BaseAttackBonusThreeQuarters,
			mustNewSaveProgressions(ability.SavingThrowPoor, ability.SavingThrowGood, ability.SavingThrowGood),
			6,
			[]skill.SkillID{
				skill.AcrobaticsSkillID,
				skill.AppraiseSkillID,
				skill.BluffSkillID,
				skill.ClimbSkillID,
				skill.CraftSkillID,
				skill.DiplomacySkillID,
				skill.DisguiseSkillID,
				skill.EscapeArtistSkillID,
				skill.IntimidateSkillID,
				skill.KnowledgeSkillID,
				skill.LinguisticsSkillID,
				skill.PerceptionSkillID,
				skill.PerformSkillID,
				skill.ProfessionSkillID,
				skill.SenseMotiveSkillID,
				skill.SleightOfHandSkillID,
				skill.SpellcraftSkillID,
				skill.StealthSkillID,
				skill.UseMagicDeviceSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
				LongswordWeaponProficiencyID,
				RapierWeaponProficiencyID,
				SapWeaponProficiencyID,
				ShortSwordWeaponProficiencyID,
				ShortbowWeaponProficiencyID,
				WhipWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
				ShieldArmorProficiencyID,
			},
			mustNewSpellcastingProfile(ArcaneSpontaneousSpellcastingKind, ability.CharismaScore),
		),
		ClericClassID: mustNewClass(
			ClericClassID,
			ability.D8HitDie,
			ability.BaseAttackBonusThreeQuarters,
			mustNewSaveProgressions(ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowGood),
			2,
			[]skill.SkillID{
				skill.AppraiseSkillID,
				skill.CraftSkillID,
				skill.DiplomacySkillID,
				skill.HealSkillID,
				knowledgeArcanaSkillID,
				knowledgeHistorySkillID,
				knowledgeNobilitySkillID,
				knowledgePlanesSkillID,
				knowledgeReligionSkillID,
				skill.LinguisticsSkillID,
				skill.ProfessionSkillID,
				skill.SenseMotiveSkillID,
				skill.SpellcraftSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
				FavoredWeaponOfDeityWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
				MediumArmorProficiencyID,
				ShieldArmorProficiencyID,
			},
			mustNewSpellcastingProfile(DivinePreparedSpellcastingKind, ability.WisdomScore),
		),
		DruidClassID: mustNewClass(
			DruidClassID,
			ability.D8HitDie,
			ability.BaseAttackBonusThreeQuarters,
			mustNewSaveProgressions(ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowGood),
			4,
			[]skill.SkillID{
				skill.ClimbSkillID,
				skill.CraftSkillID,
				skill.FlySkillID,
				skill.HandleAnimalSkillID,
				skill.HealSkillID,
				knowledgeGeographySkillID,
				knowledgeNatureSkillID,
				skill.PerceptionSkillID,
				skill.ProfessionSkillID,
				skill.RideSkillID,
				skill.SpellcraftSkillID,
				skill.SurvivalSkillID,
				skill.SwimSkillID,
			},
			[]WeaponProficiencyID{
				ClubWeaponProficiencyID,
				DaggerWeaponProficiencyID,
				DartWeaponProficiencyID,
				QuarterstaffWeaponProficiencyID,
				ScimitarWeaponProficiencyID,
				ScytheWeaponProficiencyID,
				SickleWeaponProficiencyID,
				ShortspearWeaponProficiencyID,
				SlingWeaponProficiencyID,
				SpearWeaponProficiencyID,
				WildShapeNaturalAttacksWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
				MediumArmorProficiencyID,
				ShieldArmorProficiencyID,
			},
			mustNewSpellcastingProfile(DivinePreparedSpellcastingKind, ability.WisdomScore),
		),
		FighterClassID: mustNewClass(
			FighterClassID,
			ability.D10HitDie,
			ability.BaseAttackBonusFull,
			mustNewSaveProgressions(ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowPoor),
			2,
			[]skill.SkillID{
				skill.ClimbSkillID,
				skill.CraftSkillID,
				skill.HandleAnimalSkillID,
				skill.IntimidateSkillID,
				knowledgeDungeoneeringSkillID,
				knowledgeEngineeringSkillID,
				skill.ProfessionSkillID,
				skill.RideSkillID,
				skill.SurvivalSkillID,
				skill.SwimSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
				MartialWeaponsWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
				MediumArmorProficiencyID,
				HeavyArmorProficiencyID,
				ShieldArmorProficiencyID,
				TowerShieldArmorProficiencyID,
			},
			NewNonSpellcastingProfile(),
		),
		MonkClassID: mustNewClass(
			MonkClassID,
			ability.D8HitDie,
			ability.BaseAttackBonusThreeQuarters,
			mustNewSaveProgressions(ability.SavingThrowGood, ability.SavingThrowGood, ability.SavingThrowGood),
			4,
			[]skill.SkillID{
				skill.AcrobaticsSkillID,
				skill.ClimbSkillID,
				skill.CraftSkillID,
				skill.EscapeArtistSkillID,
				skill.IntimidateSkillID,
				knowledgeHistorySkillID,
				knowledgeReligionSkillID,
				skill.PerceptionSkillID,
				skill.PerformSkillID,
				skill.ProfessionSkillID,
				skill.RideSkillID,
				skill.SenseMotiveSkillID,
				skill.StealthSkillID,
				skill.SwimSkillID,
			},
			[]WeaponProficiencyID{
				ClubWeaponProficiencyID,
				CrossbowHeavyWeaponProficiencyID,
				CrossbowLightWeaponProficiencyID,
				DaggerWeaponProficiencyID,
				HandaxeWeaponProficiencyID,
				JavelinWeaponProficiencyID,
				KamaWeaponProficiencyID,
				NunchakuWeaponProficiencyID,
				QuarterstaffWeaponProficiencyID,
				SaiWeaponProficiencyID,
				ShortspearWeaponProficiencyID,
				ShortSwordWeaponProficiencyID,
				ShurikenWeaponProficiencyID,
				SianghamWeaponProficiencyID,
				SlingWeaponProficiencyID,
				SpearWeaponProficiencyID,
			},
			nil,
			NewNonSpellcastingProfile(),
		),
		PaladinClassID: mustNewClass(
			PaladinClassID,
			ability.D10HitDie,
			ability.BaseAttackBonusFull,
			mustNewSaveProgressions(ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowGood),
			2,
			[]skill.SkillID{
				skill.CraftSkillID,
				skill.DiplomacySkillID,
				skill.HandleAnimalSkillID,
				skill.HealSkillID,
				knowledgeNobilitySkillID,
				knowledgeReligionSkillID,
				skill.ProfessionSkillID,
				skill.RideSkillID,
				skill.SenseMotiveSkillID,
				skill.SpellcraftSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
				MartialWeaponsWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
				MediumArmorProficiencyID,
				HeavyArmorProficiencyID,
				ShieldArmorProficiencyID,
			},
			mustNewSpellcastingProfile(DivinePreparedSpellcastingKind, ability.CharismaScore),
		),
		RangerClassID: mustNewClass(
			RangerClassID,
			ability.D10HitDie,
			ability.BaseAttackBonusFull,
			mustNewSaveProgressions(ability.SavingThrowGood, ability.SavingThrowGood, ability.SavingThrowPoor),
			6,
			[]skill.SkillID{
				skill.ClimbSkillID,
				skill.CraftSkillID,
				skill.HandleAnimalSkillID,
				skill.HealSkillID,
				skill.IntimidateSkillID,
				knowledgeDungeoneeringSkillID,
				knowledgeGeographySkillID,
				knowledgeNatureSkillID,
				skill.PerceptionSkillID,
				skill.ProfessionSkillID,
				skill.RideSkillID,
				skill.SpellcraftSkillID,
				skill.StealthSkillID,
				skill.SurvivalSkillID,
				skill.SwimSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
				MartialWeaponsWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
				MediumArmorProficiencyID,
				ShieldArmorProficiencyID,
			},
			mustNewSpellcastingProfile(DivinePreparedSpellcastingKind, ability.WisdomScore),
		),
		RogueClassID: mustNewClass(
			RogueClassID,
			ability.D8HitDie,
			ability.BaseAttackBonusThreeQuarters,
			mustNewSaveProgressions(ability.SavingThrowPoor, ability.SavingThrowGood, ability.SavingThrowPoor),
			8,
			[]skill.SkillID{
				skill.AcrobaticsSkillID,
				skill.AppraiseSkillID,
				skill.BluffSkillID,
				skill.ClimbSkillID,
				skill.CraftSkillID,
				skill.DiplomacySkillID,
				skill.DisableDeviceSkillID,
				skill.DisguiseSkillID,
				skill.EscapeArtistSkillID,
				skill.IntimidateSkillID,
				knowledgeDungeoneeringSkillID,
				knowledgeLocalSkillID,
				skill.LinguisticsSkillID,
				skill.PerceptionSkillID,
				skill.PerformSkillID,
				skill.ProfessionSkillID,
				skill.SenseMotiveSkillID,
				skill.SleightOfHandSkillID,
				skill.StealthSkillID,
				skill.SwimSkillID,
				skill.UseMagicDeviceSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
				HandCrossbowWeaponProficiencyID,
				RapierWeaponProficiencyID,
				SapWeaponProficiencyID,
				ShortbowWeaponProficiencyID,
				ShortSwordWeaponProficiencyID,
			},
			[]ArmorProficiencyID{
				LightArmorProficiencyID,
			},
			NewNonSpellcastingProfile(),
		),
		SorcererClassID: mustNewClass(
			SorcererClassID,
			ability.D6HitDie,
			ability.BaseAttackBonusHalf,
			mustNewSaveProgressions(ability.SavingThrowPoor, ability.SavingThrowPoor, ability.SavingThrowGood),
			2,
			[]skill.SkillID{
				skill.AppraiseSkillID,
				skill.BluffSkillID,
				skill.CraftSkillID,
				skill.FlySkillID,
				skill.IntimidateSkillID,
				knowledgeArcanaSkillID,
				skill.ProfessionSkillID,
				skill.SpellcraftSkillID,
				skill.UseMagicDeviceSkillID,
			},
			[]WeaponProficiencyID{
				SimpleWeaponsWeaponProficiencyID,
			},
			nil,
			mustNewSpellcastingProfile(ArcaneSpontaneousSpellcastingKind, ability.CharismaScore),
		),
		WizardClassID: mustNewClass(
			WizardClassID,
			ability.D6HitDie,
			ability.BaseAttackBonusHalf,
			mustNewSaveProgressions(ability.SavingThrowPoor, ability.SavingThrowPoor, ability.SavingThrowGood),
			2,
			[]skill.SkillID{
				skill.AppraiseSkillID,
				skill.CraftSkillID,
				skill.FlySkillID,
				skill.KnowledgeSkillID,
				skill.LinguisticsSkillID,
				skill.ProfessionSkillID,
				skill.SpellcraftSkillID,
			},
			[]WeaponProficiencyID{
				ClubWeaponProficiencyID,
				DaggerWeaponProficiencyID,
				CrossbowHeavyWeaponProficiencyID,
				CrossbowLightWeaponProficiencyID,
				QuarterstaffWeaponProficiencyID,
			},
			nil,
			mustNewSpellcastingProfile(ArcanePreparedSpellcastingKind, ability.IntelligenceScore),
		),
	}
}

func mustNewClass(
	id ClassID,
	hitDieType ability.HitDieType,
	baseAttackBonus ability.BaseAttackBonusProgression,
	saveProgressions SaveProgressions,
	skillRanksPerLevel int,
	classSkills []skill.SkillID,
	weaponProficiencies []WeaponProficiencyID,
	armorProficiencies []ArmorProficiencyID,
	spellcasting SpellcastingProfile,
) Class {
	class, ok := NewClass(
		id,
		hitDieType,
		baseAttackBonus,
		saveProgressions,
		skillRanksPerLevel,
		classSkills,
		weaponProficiencies,
		armorProficiencies,
		spellcasting,
	)
	if !ok {
		panic("invalid core class seed")
	}

	return class
}

func mustNewSaveProgressions(
	fortitude ability.SavingThrowProgression,
	reflex ability.SavingThrowProgression,
	will ability.SavingThrowProgression,
) SaveProgressions {
	value, ok := NewSaveProgressions(fortitude, reflex, will)
	if !ok {
		panic("invalid core class save progressions")
	}

	return value
}

func mustNewSpellcastingProfile(
	kind SpellcastingKind,
	keyAbility ability.AbilityScoreID,
) SpellcastingProfile {
	value, ok := NewSpellcastingProfile(kind, keyAbility)
	if !ok {
		panic("invalid core class spellcasting profile")
	}

	return value
}
