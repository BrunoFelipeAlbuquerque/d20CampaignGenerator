package equipment

import "strings"

type weaponID string
type WeaponID = weaponID

type weaponProficiencyCategory string
type WeaponProficiencyCategory = weaponProficiencyCategory

const (
	SimpleWeaponProficiencyCategory  WeaponProficiencyCategory = "Simple"
	MartialWeaponProficiencyCategory WeaponProficiencyCategory = "Martial"
	ExoticWeaponProficiencyCategory  WeaponProficiencyCategory = "Exotic"
)

type weaponCategory string
type WeaponCategory = weaponCategory

const (
	UnarmedAttackWeaponCategory  WeaponCategory = "Unarmed Attack"
	LightMeleeWeaponCategory     WeaponCategory = "Light Melee"
	OneHandedMeleeWeaponCategory WeaponCategory = "One-Handed Melee"
	TwoHandedMeleeWeaponCategory WeaponCategory = "Two-Handed Melee"
	RangedWeaponCategory         WeaponCategory = "Ranged"
)

type weaponDamageKind string
type WeaponDamageKind = weaponDamageKind

const (
	NoWeaponDamageKind   WeaponDamageKind = "None"
	DiceWeaponDamageKind WeaponDamageKind = "Dice"
	FlatWeaponDamageKind WeaponDamageKind = "Flat"
)

type weaponDamage struct {
	kind       weaponDamageKind
	diceCount  int
	dieSides   int
	flatPoints int
	valid      bool
}
type WeaponDamage = weaponDamage

type weaponDamageProfile struct {
	smallPrimary    weaponDamage
	mediumPrimary   weaponDamage
	smallSecondary  weaponDamage
	mediumSecondary weaponDamage
	hasSecondary    bool
	valid           bool
}
type WeaponDamageProfile = weaponDamageProfile

type weaponCriticalProfile struct {
	threatMinimum       int
	primaryMultiplier   int
	secondaryMultiplier int
	hasSecondary        bool
	hasCritical         bool
	valid               bool
}
type WeaponCriticalProfile = weaponCriticalProfile

type weaponRangeIncrement struct {
	feet     int
	meters   float64
	hasRange bool
	valid    bool
}
type WeaponRangeIncrement = weaponRangeIncrement

type weapon struct {
	id                  weaponID
	displayName         string
	proficiencyCategory weaponProficiencyCategory
	category            weaponCategory
	damageProfile       weaponDamageProfile
	criticalProfile     weaponCriticalProfile
	rangeIncrement      weaponRangeIncrement
	cost                equipmentCost
	weight              equipmentWeight
}
type Weapon = weapon

func NewWeaponDamageDice(diceCount int, dieSides int) (WeaponDamage, bool) {
	if diceCount <= 0 || !isValidWeaponDamageDieSides(dieSides) {
		return weaponDamage{}, false
	}

	return weaponDamage{
		kind:      DiceWeaponDamageKind,
		diceCount: diceCount,
		dieSides:  dieSides,
		valid:     true,
	}, true
}

func NewWeaponFlatDamage(points int) (WeaponDamage, bool) {
	if points <= 0 {
		return weaponDamage{}, false
	}

	return weaponDamage{
		kind:       FlatWeaponDamageKind,
		flatPoints: points,
		valid:      true,
	}, true
}

func NewNoWeaponDamage() WeaponDamage {
	return weaponDamage{
		kind:  NoWeaponDamageKind,
		valid: true,
	}
}

func NewWeaponDamageProfile(small WeaponDamage, medium WeaponDamage) (WeaponDamageProfile, bool) {
	if !isValidWeaponDamage(small) ||
		!isValidWeaponDamage(medium) ||
		small.HasDamage() != medium.HasDamage() {
		return weaponDamageProfile{}, false
	}

	return weaponDamageProfile{
		smallPrimary:  small,
		mediumPrimary: medium,
		valid:         true,
	}, true
}

func NewDoubleWeaponDamageProfile(
	smallPrimary WeaponDamage,
	mediumPrimary WeaponDamage,
	smallSecondary WeaponDamage,
	mediumSecondary WeaponDamage,
) (WeaponDamageProfile, bool) {
	if !isValidWeaponDamage(smallPrimary) ||
		!isValidWeaponDamage(mediumPrimary) ||
		!isValidWeaponDamage(smallSecondary) ||
		!isValidWeaponDamage(mediumSecondary) ||
		!smallPrimary.HasDamage() ||
		!mediumPrimary.HasDamage() ||
		!smallSecondary.HasDamage() ||
		!mediumSecondary.HasDamage() {
		return weaponDamageProfile{}, false
	}

	return weaponDamageProfile{
		smallPrimary:    smallPrimary,
		mediumPrimary:   mediumPrimary,
		smallSecondary:  smallSecondary,
		mediumSecondary: mediumSecondary,
		hasSecondary:    true,
		valid:           true,
	}, true
}

func NewWeaponCriticalProfile(threatMinimum int, multiplier int) (WeaponCriticalProfile, bool) {
	if !isValidWeaponCriticalThreatMinimum(threatMinimum) || !isValidWeaponCriticalMultiplier(multiplier) {
		return weaponCriticalProfile{}, false
	}

	return weaponCriticalProfile{
		threatMinimum:     threatMinimum,
		primaryMultiplier: multiplier,
		hasCritical:       true,
		valid:             true,
	}, true
}

func NewDoubleWeaponCriticalProfile(
	threatMinimum int,
	primaryMultiplier int,
	secondaryMultiplier int,
) (WeaponCriticalProfile, bool) {
	if !isValidWeaponCriticalThreatMinimum(threatMinimum) ||
		!isValidWeaponCriticalMultiplier(primaryMultiplier) ||
		!isValidWeaponCriticalMultiplier(secondaryMultiplier) {
		return weaponCriticalProfile{}, false
	}

	return weaponCriticalProfile{
		threatMinimum:       threatMinimum,
		primaryMultiplier:   primaryMultiplier,
		secondaryMultiplier: secondaryMultiplier,
		hasSecondary:        true,
		hasCritical:         true,
		valid:               true,
	}, true
}

func NewNoWeaponCriticalProfile() WeaponCriticalProfile {
	return weaponCriticalProfile{
		valid: true,
	}
}

func NewWeaponRangeIncrementFeet(feet int) (WeaponRangeIncrement, bool) {
	if feet <= 0 {
		return weaponRangeIncrement{}, false
	}

	return newWeaponRangeIncrement(feet, feetToMeters(feet))
}

func NewWeaponRangeIncrementMeters(meters float64) (WeaponRangeIncrement, bool) {
	if !isValidMetricValue(meters, false) {
		return weaponRangeIncrement{}, false
	}

	return newWeaponRangeIncrement(metersToFeet(meters), meters)
}

func newWeaponRangeIncrement(feet int, meters float64) (WeaponRangeIncrement, bool) {
	if feet <= 0 ||
		!isValidMetricValue(meters, false) ||
		metersToFeet(meters) != feet {
		return weaponRangeIncrement{}, false
	}

	return weaponRangeIncrement{
		feet:     feet,
		meters:   meters,
		hasRange: true,
		valid:    true,
	}, true
}

func NewNoWeaponRangeIncrement() WeaponRangeIncrement {
	return weaponRangeIncrement{
		valid: true,
	}
}

func NewWeapon(
	id WeaponID,
	displayName string,
	proficiencyCategory WeaponProficiencyCategory,
	category WeaponCategory,
	damageProfile WeaponDamageProfile,
	criticalProfile WeaponCriticalProfile,
	rangeIncrement WeaponRangeIncrement,
	cost EquipmentCost,
	weight EquipmentWeight,
) (Weapon, bool) {
	if !isValidWeaponID(id) ||
		!isValidDisplayName(displayName) ||
		!isValidWeaponProficiencyCategory(proficiencyCategory) ||
		!isValidWeaponCategory(category) ||
		!isValidWeaponDamageProfile(damageProfile) ||
		!isValidWeaponCriticalProfile(criticalProfile) ||
		!isValidWeaponRangeIncrement(rangeIncrement) ||
		!isValidWeaponCost(cost) ||
		!isValidWeaponWeight(weight) ||
		!isValidWeaponProfileCombination(category, damageProfile, criticalProfile, rangeIncrement) {
		return weapon{}, false
	}

	return weapon{
		id:                  id,
		displayName:         displayName,
		proficiencyCategory: proficiencyCategory,
		category:            category,
		damageProfile:       damageProfile,
		criticalProfile:     criticalProfile,
		rangeIncrement:      rangeIncrement,
		cost:                cost,
		weight:              weight,
	}, true
}

func (d weaponDamage) GetKind() WeaponDamageKind {
	if !isValidWeaponDamage(WeaponDamage(d)) {
		return ""
	}

	return d.kind
}

func (d weaponDamage) HasDamage() bool {
	return d.valid && d.kind != NoWeaponDamageKind
}

func (d weaponDamage) GetDiceCount() int {
	return d.diceCount
}

func (d weaponDamage) GetDieSides() int {
	return d.dieSides
}

func (d weaponDamage) GetFlatPoints() int {
	return d.flatPoints
}

func (p weaponDamageProfile) GetSmallDamage() WeaponDamage {
	return p.smallPrimary
}

func (p weaponDamageProfile) GetMediumDamage() WeaponDamage {
	return p.mediumPrimary
}

func (p weaponDamageProfile) HasSecondaryDamage() bool {
	return p.hasSecondary
}

func (p weaponDamageProfile) GetSecondarySmallDamage() WeaponDamage {
	return p.smallSecondary
}

func (p weaponDamageProfile) GetSecondaryMediumDamage() WeaponDamage {
	return p.mediumSecondary
}

func (p weaponDamageProfile) HasDamage() bool {
	return p.valid && p.smallPrimary.HasDamage()
}

func (p weaponCriticalProfile) HasCritical() bool {
	return p.valid && p.hasCritical
}

func (p weaponCriticalProfile) GetThreatMinimum() int {
	return p.threatMinimum
}

func (p weaponCriticalProfile) GetThreatMaximum() int {
	if !p.HasCritical() {
		return 0
	}

	return 20
}

func (p weaponCriticalProfile) GetPrimaryMultiplier() int {
	return p.primaryMultiplier
}

func (p weaponCriticalProfile) HasSecondaryMultiplier() bool {
	return p.hasSecondary
}

func (p weaponCriticalProfile) GetSecondaryMultiplier() int {
	return p.secondaryMultiplier
}

func (r weaponRangeIncrement) HasRangeIncrement() bool {
	return r.valid && r.hasRange
}

func (r weaponRangeIncrement) GetFeet() int {
	return r.feet
}

func (r weaponRangeIncrement) GetMeters() float64 {
	if !r.HasRangeIncrement() {
		return 0
	}

	return r.meters
}

func (w weapon) GetID() WeaponID {
	return w.id
}

func (w weapon) GetDisplayName() string {
	return w.displayName
}

func (w weapon) GetProficiencyCategory() WeaponProficiencyCategory {
	return w.proficiencyCategory
}

func (w weapon) GetCategory() WeaponCategory {
	return w.category
}

func (w weapon) GetDamageProfile() WeaponDamageProfile {
	return w.damageProfile
}

func (w weapon) GetCriticalProfile() WeaponCriticalProfile {
	return w.criticalProfile
}

func (w weapon) GetRangeIncrement() WeaponRangeIncrement {
	return w.rangeIncrement
}

func (w weapon) GetCost() EquipmentCost {
	return w.cost
}

func (w weapon) GetWeight() EquipmentWeight {
	return w.weight
}

func isValidWeaponID(id WeaponID) bool {
	value := string(id)
	return value != "" && strings.TrimSpace(value) == value
}

func isValidWeaponProficiencyCategory(category WeaponProficiencyCategory) bool {
	switch category {
	case SimpleWeaponProficiencyCategory,
		MartialWeaponProficiencyCategory,
		ExoticWeaponProficiencyCategory:
		return true
	default:
		return false
	}
}

func isValidWeaponCategory(category WeaponCategory) bool {
	switch category {
	case UnarmedAttackWeaponCategory,
		LightMeleeWeaponCategory,
		OneHandedMeleeWeaponCategory,
		TwoHandedMeleeWeaponCategory,
		RangedWeaponCategory:
		return true
	default:
		return false
	}
}

func isValidWeaponDamage(damage WeaponDamage) bool {
	if !damage.valid {
		return false
	}

	switch damage.kind {
	case NoWeaponDamageKind:
		return damage.diceCount == 0 && damage.dieSides == 0 && damage.flatPoints == 0
	case DiceWeaponDamageKind:
		return damage.diceCount > 0 && isValidWeaponDamageDieSides(damage.dieSides) && damage.flatPoints == 0
	case FlatWeaponDamageKind:
		return damage.diceCount == 0 && damage.dieSides == 0 && damage.flatPoints > 0
	default:
		return false
	}
}

func isValidWeaponDamageDieSides(dieSides int) bool {
	switch dieSides {
	case 2, 3, 4, 6, 8, 10, 12:
		return true
	default:
		return false
	}
}

func isValidWeaponDamageProfile(profile WeaponDamageProfile) bool {
	if !profile.valid ||
		!isValidWeaponDamage(profile.smallPrimary) ||
		!isValidWeaponDamage(profile.mediumPrimary) ||
		profile.smallPrimary.HasDamage() != profile.mediumPrimary.HasDamage() {
		return false
	}

	if !profile.hasSecondary {
		return !profile.smallSecondary.valid && !profile.mediumSecondary.valid
	}

	return isValidWeaponDamage(profile.smallSecondary) &&
		isValidWeaponDamage(profile.mediumSecondary) &&
		profile.smallSecondary.HasDamage() &&
		profile.mediumSecondary.HasDamage()
}

func isValidWeaponCriticalProfile(profile WeaponCriticalProfile) bool {
	if !profile.valid {
		return false
	}

	if !profile.hasCritical {
		return profile.threatMinimum == 0 &&
			profile.primaryMultiplier == 0 &&
			profile.secondaryMultiplier == 0 &&
			!profile.hasSecondary
	}

	if !isValidWeaponCriticalThreatMinimum(profile.threatMinimum) ||
		!isValidWeaponCriticalMultiplier(profile.primaryMultiplier) {
		return false
	}

	if !profile.hasSecondary {
		return profile.secondaryMultiplier == 0
	}

	return isValidWeaponCriticalMultiplier(profile.secondaryMultiplier)
}

func isValidWeaponCriticalThreatMinimum(threatMinimum int) bool {
	switch threatMinimum {
	case 18, 19, 20:
		return true
	default:
		return false
	}
}

func isValidWeaponCriticalMultiplier(multiplier int) bool {
	switch multiplier {
	case 2, 3, 4:
		return true
	default:
		return false
	}
}

func isValidWeaponRangeIncrement(rangeIncrement WeaponRangeIncrement) bool {
	if !rangeIncrement.valid {
		return false
	}

	if rangeIncrement.hasRange {
		return rangeIncrement.feet > 0 &&
			isValidMetricValue(rangeIncrement.meters, false) &&
			metersToFeet(rangeIncrement.meters) == rangeIncrement.feet
	}

	return rangeIncrement.feet == 0 && rangeIncrement.meters == 0
}

func isValidWeaponCost(cost EquipmentCost) bool {
	return isValidEquipmentCost(cost)
}

func isValidWeaponWeight(weight EquipmentWeight) bool {
	return isValidEquipmentWeight(weight)
}

func isValidWeaponProfileCombination(
	category WeaponCategory,
	damageProfile WeaponDamageProfile,
	criticalProfile WeaponCriticalProfile,
	rangeIncrement WeaponRangeIncrement,
) bool {
	if damageProfile.HasDamage() != criticalProfile.HasCritical() {
		return false
	}

	if !damageProfile.HasSecondaryDamage() && criticalProfile.HasSecondaryMultiplier() {
		return false
	}

	if category == RangedWeaponCategory && !rangeIncrement.HasRangeIncrement() {
		return false
	}

	return true
}
