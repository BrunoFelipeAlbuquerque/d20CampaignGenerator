package modifier

const (
	// MODIFIER SOURCE: COMBAT
	SourceFlanking     ModifierCircumstanceSource = "combat.flanking"
	SourceHigherGround ModifierCircumstanceSource = "combat.higher_ground"
	SourceProneTarget  ModifierCircumstanceSource = "combat.target_prone"
	SourceCharging     ModifierCircumstanceSource = "combat.charging"
	SourceCover        ModifierCircumstanceSource = "combat.cover"
	SourceSoftCover    ModifierCircumstanceSource = "combat.soft_cover"

	// MODIFIER SOURCE: PERCEPTION
	SourceConcealment      ModifierCircumstanceSource = "perception.concealment"
	SourceTotalConcealment ModifierCircumstanceSource = "perception.total_concealment"
	SourceInvisibleTarget  ModifierCircumstanceSource = "perception.invisible_target"
	SourcePoorLighting     ModifierCircumstanceSource = "perception.poor_lighting"

	// MODIFIER SOURCE: CONDITON
	SourceTargetHelpless   ModifierCircumstanceSource = "condition.helpless"
	SourceTargetBlinded    ModifierCircumstanceSource = "condition.blinded"
	SourceTargetStunned    ModifierCircumstanceSource = "condition.stunned"
	SourceTargetFlatFooted ModifierCircumstanceSource = "condition.flat_footed"

	// MODIFIER SOURCE: ENVIRONMENT
	SourceSlipperySurface ModifierCircumstanceSource = "environment.slippery"
	SourceStrongWind      ModifierCircumstanceSource = "environment.strong_wind"
	SourceExtremeHeat     ModifierCircumstanceSource = "environment.extreme_heat"
	SourceExtremeCold     ModifierCircumstanceSource = "environment.extreme_cold"

	// MODIFIER SOURCE: SKILL
	SourceFavorableConditions   ModifierCircumstanceSource = "skill.favorable_conditions"
	SourceUnfavorableConditions ModifierCircumstanceSource = "skill.unfavorable_conditions"
	SourceDistraction           ModifierCircumstanceSource = "skill.distraction"

	// MODIFIER TYPES
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
