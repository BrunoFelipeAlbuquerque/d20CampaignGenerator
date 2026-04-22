package modifier

const (
	// MODIFIER SOURCE: COMBAT
	SourceFlanking     ModifierSource = "combat.flanking"
	SourceHigherGround ModifierSource = "combat.higher_ground"
	SourceProneTarget  ModifierSource = "combat.target_prone"
	SourceCharging     ModifierSource = "combat.charging"
	SourceCover        ModifierSource = "combat.cover"
	SourceSoftCover    ModifierSource = "combat.soft_cover"

	// MODIFIER SOURCE: PERCEPTION
	SourceConcealment      ModifierSource = "perception.concealment"
	SourceTotalConcealment ModifierSource = "perception.total_concealment"
	SourceInvisibleTarget  ModifierSource = "perception.invisible_target"
	SourcePoorLighting     ModifierSource = "perception.poor_lighting"

	// MODIFIER SOURCE: CONDITON
	SourceTargetHelpless   ModifierSource = "condition.helpless"
	SourceTargetBlinded    ModifierSource = "condition.blinded"
	SourceTargetStunned    ModifierSource = "condition.stunned"
	SourceTargetFlatFooted ModifierSource = "condition.flat_footed"

	// MODIFIER SOURCE: ENVIRONMENT
	SourceSlipperySurface ModifierSource = "environment.slippery"
	SourceStrongWind      ModifierSource = "environment.strong_wind"
	SourceExtremeHeat     ModifierSource = "environment.extreme_heat"
	SourceExtremeCold     ModifierSource = "environment.extreme_cold"

	// MODIFIER SOURCE: SKILL
	SourceFavorableConditions   ModifierSource = "skill.favorable_conditions"
	SourceUnfavorableConditions ModifierSource = "skill.unfavorable_conditions"
	SourceDistraction           ModifierSource = "skill.distraction"

	// MODIFIER TYPES
	ModifierUntyped      ModifierType = "untyped"
	ModifierAlchemical   ModifierType = "alchemical"
	ModifierArmor        ModifierType = "armor"
	ModifierCircumstance ModifierType = "circumstance"
	ModifierCompetence   ModifierType = "competence"
	ModifierDeflection   ModifierType = "deflection"
	ModifierDodge        ModifierType = "dodge"
	ModifierEnhancement  ModifierType = "enhancement"
	ModifierInherent     ModifierType = "inherent"
	ModifierInsight      ModifierType = "insight"
	ModifierLuck         ModifierType = "luck"
	ModifierMorale       ModifierType = "morale"
	ModifierNaturalArmor ModifierType = "natural_armor"
	ModifierProfane      ModifierType = "profane"
	ModifierRacial       ModifierType = "racial"
	ModifierResistance   ModifierType = "resistance"
	ModifierSacred       ModifierType = "sacred"
	ModifierShield       ModifierType = "shield"
	ModifierSize         ModifierType = "size"
	ModifierTrait        ModifierType = "trait"
)

func isValidModifierType(value ModifierType) bool {
	switch value {
	case ModifierUntyped,
		ModifierAlchemical,
		ModifierArmor,
		ModifierCircumstance,
		ModifierCompetence,
		ModifierDeflection,
		ModifierDodge,
		ModifierEnhancement,
		ModifierInherent,
		ModifierInsight,
		ModifierLuck,
		ModifierMorale,
		ModifierNaturalArmor,
		ModifierProfane,
		ModifierRacial,
		ModifierResistance,
		ModifierSacred,
		ModifierShield,
		ModifierSize,
		ModifierTrait:
		return true
	default:
		return false
	}
}
