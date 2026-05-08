package equipment

import "strings"

type armorID string
type ArmorID = armorID

type armorCategory string
type ArmorCategory = armorCategory

const (
	LightArmorCategory       ArmorCategory = "Light Armor"
	MediumArmorCategory      ArmorCategory = "Medium Armor"
	HeavyArmorCategory       ArmorCategory = "Heavy Armor"
	ShieldArmorCategory      ArmorCategory = "Shield"
	TowerShieldArmorCategory ArmorCategory = "Tower Shield"
)

type armorClassBonus struct {
	points int
	valid  bool
}
type ArmorClassBonus = armorClassBonus

type armorMaximumDexterityBonus struct {
	points     int
	hasMaximum bool
	valid      bool
}
type ArmorMaximumDexterityBonus = armorMaximumDexterityBonus

type armorCheckPenalty struct {
	penalty int
	valid   bool
}
type ArmorCheckPenalty = armorCheckPenalty

type armorArcaneSpellFailureChance struct {
	percent int
	valid   bool
}
type ArmorArcaneSpellFailureChance = armorArcaneSpellFailureChance

type armorSpeedImpact struct {
	speedFor30FeetBase int
	speedFor20FeetBase int
	limitsRunning      bool
	hasImpact          bool
	valid              bool
}
type ArmorSpeedImpact = armorSpeedImpact

type armor struct {
	id                       armorID
	displayName              string
	category                 armorCategory
	armorClassBonus          armorClassBonus
	maximumDexterityBonus    armorMaximumDexterityBonus
	armorCheckPenalty        armorCheckPenalty
	arcaneSpellFailureChance armorArcaneSpellFailureChance
	speedImpact              armorSpeedImpact
	cost                     equipmentCost
	weight                   equipmentWeight
}
type Armor = armor

func NewArmorClassBonus(points int) (ArmorClassBonus, bool) {
	if points <= 0 {
		return armorClassBonus{}, false
	}

	return armorClassBonus{
		points: points,
		valid:  true,
	}, true
}

func NewArmorMaximumDexterityBonus(points int) (ArmorMaximumDexterityBonus, bool) {
	if points < 0 {
		return armorMaximumDexterityBonus{}, false
	}

	return armorMaximumDexterityBonus{
		points:     points,
		hasMaximum: true,
		valid:      true,
	}, true
}

func NewNoArmorMaximumDexterityBonus() ArmorMaximumDexterityBonus {
	return armorMaximumDexterityBonus{
		valid: true,
	}
}

func NewArmorCheckPenalty(penalty int) (ArmorCheckPenalty, bool) {
	if penalty > 0 {
		return armorCheckPenalty{}, false
	}

	return armorCheckPenalty{
		penalty: penalty,
		valid:   true,
	}, true
}

func NewArmorArcaneSpellFailureChance(percent int) (ArmorArcaneSpellFailureChance, bool) {
	if percent < 0 || percent > 100 {
		return armorArcaneSpellFailureChance{}, false
	}

	return armorArcaneSpellFailureChance{
		percent: percent,
		valid:   true,
	}, true
}

func NewArmorSpeedImpact(
	speedFor30FeetBase int,
	speedFor20FeetBase int,
	limitsRunning bool,
) (ArmorSpeedImpact, bool) {
	if speedFor30FeetBase <= 0 ||
		speedFor30FeetBase >= 30 ||
		speedFor20FeetBase <= 0 ||
		speedFor20FeetBase >= 20 {
		return armorSpeedImpact{}, false
	}

	return armorSpeedImpact{
		speedFor30FeetBase: speedFor30FeetBase,
		speedFor20FeetBase: speedFor20FeetBase,
		limitsRunning:      limitsRunning,
		hasImpact:          true,
		valid:              true,
	}, true
}

func NewNoArmorSpeedImpact() ArmorSpeedImpact {
	return armorSpeedImpact{
		valid: true,
	}
}

func NewArmor(
	id ArmorID,
	displayName string,
	category ArmorCategory,
	armorClassBonus ArmorClassBonus,
	maximumDexterityBonus ArmorMaximumDexterityBonus,
	armorCheckPenalty ArmorCheckPenalty,
	arcaneSpellFailureChance ArmorArcaneSpellFailureChance,
	speedImpact ArmorSpeedImpact,
	cost EquipmentCost,
	weight EquipmentWeight,
) (Armor, bool) {
	if !isValidArmorID(id) ||
		!isValidDisplayName(displayName) ||
		!isValidArmorCategory(category) ||
		!isValidArmorClassBonus(armorClassBonus) ||
		!isValidArmorMaximumDexterityBonus(maximumDexterityBonus) ||
		!isValidArmorCheckPenalty(armorCheckPenalty) ||
		!isValidArmorArcaneSpellFailureChance(arcaneSpellFailureChance) ||
		!isValidArmorSpeedImpact(speedImpact) ||
		!isValidArmorCost(cost) ||
		!isValidArmorWeight(weight) ||
		!isValidArmorProfileCombination(category, maximumDexterityBonus, speedImpact) {
		return armor{}, false
	}

	return armor{
		id:                       id,
		displayName:              displayName,
		category:                 category,
		armorClassBonus:          armorClassBonus,
		maximumDexterityBonus:    maximumDexterityBonus,
		armorCheckPenalty:        armorCheckPenalty,
		arcaneSpellFailureChance: arcaneSpellFailureChance,
		speedImpact:              speedImpact,
		cost:                     cost,
		weight:                   weight,
	}, true
}

func (b armorClassBonus) GetPoints() int {
	return b.points
}

func (b armorMaximumDexterityBonus) HasMaximum() bool {
	return b.valid && b.hasMaximum
}

func (b armorMaximumDexterityBonus) GetPoints() int {
	return b.points
}

func (p armorCheckPenalty) GetPenalty() int {
	return p.penalty
}

func (c armorArcaneSpellFailureChance) GetPercent() int {
	return c.percent
}

func (s armorSpeedImpact) HasImpact() bool {
	return s.valid && s.hasImpact
}

func (s armorSpeedImpact) GetSpeedFor30FeetBase() int {
	return s.speedFor30FeetBase
}

func (s armorSpeedImpact) GetSpeedFor20FeetBase() int {
	return s.speedFor20FeetBase
}

func (s armorSpeedImpact) LimitsRunning() bool {
	return s.limitsRunning
}

func (a armor) GetID() ArmorID {
	return a.id
}

func (a armor) GetDisplayName() string {
	return a.displayName
}

func (a armor) GetCategory() ArmorCategory {
	return a.category
}

func (a armor) GetArmorClassBonus() ArmorClassBonus {
	return a.armorClassBonus
}

func (a armor) GetMaximumDexterityBonus() ArmorMaximumDexterityBonus {
	return a.maximumDexterityBonus
}

func (a armor) GetArmorCheckPenalty() ArmorCheckPenalty {
	return a.armorCheckPenalty
}

func (a armor) GetArcaneSpellFailureChance() ArmorArcaneSpellFailureChance {
	return a.arcaneSpellFailureChance
}

func (a armor) GetSpeedImpact() ArmorSpeedImpact {
	return a.speedImpact
}

func (a armor) GetCost() EquipmentCost {
	return a.cost
}

func (a armor) GetWeight() EquipmentWeight {
	return a.weight
}

func isValidArmorID(id ArmorID) bool {
	value := string(id)
	return value != "" && strings.TrimSpace(value) == value
}

func isValidArmorCategory(category ArmorCategory) bool {
	switch category {
	case LightArmorCategory,
		MediumArmorCategory,
		HeavyArmorCategory,
		ShieldArmorCategory,
		TowerShieldArmorCategory:
		return true
	default:
		return false
	}
}

func isValidArmorClassBonus(bonus ArmorClassBonus) bool {
	return bonus.valid && bonus.points > 0
}

func isValidArmorMaximumDexterityBonus(bonus ArmorMaximumDexterityBonus) bool {
	if !bonus.valid {
		return false
	}

	if !bonus.hasMaximum {
		return bonus.points == 0
	}

	return bonus.points >= 0
}

func isValidArmorCheckPenalty(penalty ArmorCheckPenalty) bool {
	return penalty.valid && penalty.penalty <= 0
}

func isValidArmorArcaneSpellFailureChance(chance ArmorArcaneSpellFailureChance) bool {
	return chance.valid && chance.percent >= 0 && chance.percent <= 100
}

func isValidArmorSpeedImpact(speedImpact ArmorSpeedImpact) bool {
	if !speedImpact.valid {
		return false
	}

	if !speedImpact.hasImpact {
		return speedImpact.speedFor30FeetBase == 0 &&
			speedImpact.speedFor20FeetBase == 0 &&
			!speedImpact.limitsRunning
	}

	return speedImpact.speedFor30FeetBase > 0 &&
		speedImpact.speedFor30FeetBase < 30 &&
		speedImpact.speedFor20FeetBase > 0 &&
		speedImpact.speedFor20FeetBase < 20
}

func isValidArmorCost(cost EquipmentCost) bool {
	return isValidEquipmentCost(cost)
}

func isValidArmorWeight(weight EquipmentWeight) bool {
	return isValidEquipmentWeight(weight)
}

func isValidArmorProfileCombination(
	category ArmorCategory,
	maximumDexterityBonus ArmorMaximumDexterityBonus,
	speedImpact ArmorSpeedImpact,
) bool {
	switch category {
	case LightArmorCategory:
		return maximumDexterityBonus.HasMaximum() && !speedImpact.HasImpact()
	case MediumArmorCategory:
		return maximumDexterityBonus.HasMaximum() &&
			speedImpact.HasImpact() &&
			!speedImpact.LimitsRunning()
	case HeavyArmorCategory:
		return maximumDexterityBonus.HasMaximum() &&
			speedImpact.HasImpact() &&
			speedImpact.LimitsRunning()
	case ShieldArmorCategory:
		return !maximumDexterityBonus.HasMaximum() && !speedImpact.HasImpact()
	case TowerShieldArmorCategory:
		return maximumDexterityBonus.HasMaximum() && !speedImpact.HasImpact()
	default:
		return false
	}
}
