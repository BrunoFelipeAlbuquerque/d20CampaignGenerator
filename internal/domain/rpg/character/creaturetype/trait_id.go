package creaturetype

type creatureTypeTraitID string
type CreatureTypeTraitID = creatureTypeTraitID

const (
	Darkvision60Trait             CreatureTypeTraitID = "Darkvision60"
	LowLightVisionTrait           CreatureTypeTraitID = "LowLightVision"
	BreatheEatSleepTrait          CreatureTypeTraitID = "BreatheEatSleep"
	NoNeedToEatSleepBreatheTrait  CreatureTypeTraitID = "NoNeedToEatSleepBreathe"
	NoConstitutionTrait           CreatureTypeTraitID = "NoConstitution"
	NoIntelligenceTrait           CreatureTypeTraitID = "NoIntelligence"
	MindlessTrait                 CreatureTypeTraitID = "Mindless"
	BlindTrait                    CreatureTypeTraitID = "Blind"
	Blindsight60Trait             CreatureTypeTraitID = "Blindsight60"
	DestroyedAtZeroHPTrait        CreatureTypeTraitID = "DestroyedAtZeroHP"
	NotSubjectToCriticalHitsTrait CreatureTypeTraitID = "NotSubjectToCriticalHits"
	NotSubjectToNonlethalTrait    CreatureTypeTraitID = "NotSubjectToNonlethal"
	ImmunityMindAffectingTrait    CreatureTypeTraitID = "ImmunityMindAffecting"
	ImmunityPoisonTrait           CreatureTypeTraitID = "ImmunityPoison"
	ImmunitySleepTrait            CreatureTypeTraitID = "ImmunitySleep"
	ImmunityParalysisTrait        CreatureTypeTraitID = "ImmunityParalysis"
	ImmunityStunTrait             CreatureTypeTraitID = "ImmunityStun"
	ImmunityPolymorphTrait        CreatureTypeTraitID = "ImmunityPolymorph"
	ImmunityDiseaseTrait          CreatureTypeTraitID = "ImmunityDisease"
	ImmunityDeathEffectsTrait     CreatureTypeTraitID = "ImmunityDeathEffects"
	OneGoodSaveChoiceTrait        CreatureTypeTraitID = "OneGoodSaveChoice"
	TwoGoodSaveChoicesTrait       CreatureTypeTraitID = "TwoGoodSaveChoices"
)

func isValidCreatureTypeTraitID(value CreatureTypeTraitID) bool {
	switch value {
	case Darkvision60Trait,
		LowLightVisionTrait,
		BreatheEatSleepTrait,
		NoNeedToEatSleepBreatheTrait,
		NoConstitutionTrait,
		NoIntelligenceTrait,
		MindlessTrait,
		BlindTrait,
		Blindsight60Trait,
		DestroyedAtZeroHPTrait,
		NotSubjectToCriticalHitsTrait,
		NotSubjectToNonlethalTrait,
		ImmunityMindAffectingTrait,
		ImmunityPoisonTrait,
		ImmunitySleepTrait,
		ImmunityParalysisTrait,
		ImmunityStunTrait,
		ImmunityPolymorphTrait,
		ImmunityDiseaseTrait,
		ImmunityDeathEffectsTrait,
		OneGoodSaveChoiceTrait,
		TwoGoodSaveChoicesTrait:
		return true
	default:
		return false
	}
}
