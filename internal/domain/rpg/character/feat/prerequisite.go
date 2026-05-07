package feat

import (
	"strings"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
	"d20campaigngenerator/internal/domain/rpg/character/spell"
)

type featID string
type FeatID = featID

type prerequisiteKind string
type PrerequisiteKind = prerequisiteKind

type spellcastingAccess string
type SpellcastingAccess = spellcastingAccess

const (
	AbilityScorePrerequisiteKind                PrerequisiteKind = "Ability Score"
	BaseAttackBonusPrerequisiteKind             PrerequisiteKind = "Base Attack Bonus"
	SkillRanksPrerequisiteKind                  PrerequisiteKind = "Skill Ranks"
	AnySkillRanksPrerequisiteKind               PrerequisiteKind = "Any Skill Ranks"
	SpellcastingPrerequisiteKind                PrerequisiteKind = "Spellcasting"
	CasterLevelPrerequisiteKind                 PrerequisiteKind = "Caster Level"
	CharacterLevelPrerequisiteKind              PrerequisiteKind = "Character Level"
	ClassLevelPrerequisiteKind                  PrerequisiteKind = "Class Level"
	ClassFeaturePrerequisiteKind                PrerequisiteKind = "Class Feature"
	SelectedWeaponProficiencyPrerequisiteKind   PrerequisiteKind = "Selected Weapon Proficiency"
	SelectedFamiliarEligibilityPrerequisiteKind PrerequisiteKind = "Selected Familiar Eligibility"
	FeatPrerequisiteKind                        PrerequisiteKind = "Feat"
	AnyFeatPrerequisiteKind                     PrerequisiteKind = "Any Feat"
	SameSelectionFeatPrerequisiteKind           PrerequisiteKind = "Same Selection Feat"
	SpellSchoolFeatPrerequisiteKind             PrerequisiteKind = "Spell School Feat"
)

const (
	AnySpellcastingAccess    SpellcastingAccess = "Any"
	ArcaneSpellcastingAccess SpellcastingAccess = "Arcane"
	DivineSpellcastingAccess SpellcastingAccess = "Divine"
)

type prerequisite interface {
	isPrerequisite()
	isValid() bool
	GetKind() PrerequisiteKind
}
type Prerequisite = prerequisite

type prerequisiteList struct {
	prerequisites []Prerequisite
}
type PrerequisiteList = prerequisiteList

type abilityScorePrerequisite struct {
	abilityScoreID ability.AbilityScoreID
	minimumScore   int
}
type AbilityScorePrerequisite = abilityScorePrerequisite

type baseAttackBonusPrerequisite struct {
	minimumBonus int
}
type BaseAttackBonusPrerequisite = baseAttackBonusPrerequisite

type skillRanksPrerequisite struct {
	skillID      skill.SkillID
	minimumRanks int
}
type SkillRanksPrerequisite = skillRanksPrerequisite

type anySkillRanksPrerequisite struct {
	skillIDs     []skill.SkillID
	minimumRanks int
}
type AnySkillRanksPrerequisite = anySkillRanksPrerequisite

type spellcastingPrerequisite struct {
	access spellcastingAccess
}
type SpellcastingPrerequisite = spellcastingPrerequisite

type casterLevelPrerequisite struct {
	minimumLevel int
}
type CasterLevelPrerequisite = casterLevelPrerequisite

type characterLevelPrerequisite struct {
	minimumLevel int
}
type CharacterLevelPrerequisite = characterLevelPrerequisite

type classLevelPrerequisite struct {
	classID      characterclass.ClassID
	minimumLevel int
}
type ClassLevelPrerequisite = classLevelPrerequisite

type classFeaturePrerequisite struct {
	featureID characterclass.ClassFeatureID
}
type ClassFeaturePrerequisite = classFeaturePrerequisite

type selectedWeaponProficiencyPrerequisite struct {
	valid bool
}
type SelectedWeaponProficiencyPrerequisite = selectedWeaponProficiencyPrerequisite

type selectedFamiliarEligibilityPrerequisite struct {
	valid bool
}
type SelectedFamiliarEligibilityPrerequisite = selectedFamiliarEligibilityPrerequisite

type featPrerequisite struct {
	featID featID
}
type FeatPrerequisite = featPrerequisite

type anyFeatPrerequisite struct {
	featIDs []featID
}
type AnyFeatPrerequisite = anyFeatPrerequisite

type sameSelectionFeatPrerequisite struct {
	featID featID
}
type SameSelectionFeatPrerequisite = sameSelectionFeatPrerequisite

type spellSchoolFeatPrerequisite struct {
	featID   featID
	schoolID spell.SchoolID
}
type SpellSchoolFeatPrerequisite = spellSchoolFeatPrerequisite

func NewPrerequisiteList(prerequisites []Prerequisite) (PrerequisiteList, bool) {
	copied := make([]Prerequisite, 0, len(prerequisites))
	for _, prerequisite := range prerequisites {
		if !isValidPrerequisite(prerequisite) {
			return prerequisiteList{}, false
		}

		copied = append(copied, prerequisite)
	}

	return prerequisiteList{prerequisites: copied}, true
}

func NewAbilityScorePrerequisite(
	id ability.AbilityScoreID,
	minimumScore int,
) (AbilityScorePrerequisite, bool) {
	value := abilityScorePrerequisite{
		abilityScoreID: id,
		minimumScore:   minimumScore,
	}
	if !value.isValid() {
		return abilityScorePrerequisite{}, false
	}

	return value, true
}

func NewBaseAttackBonusPrerequisite(minimumBonus int) (BaseAttackBonusPrerequisite, bool) {
	value := baseAttackBonusPrerequisite{minimumBonus: minimumBonus}
	if !value.isValid() {
		return baseAttackBonusPrerequisite{}, false
	}

	return value, true
}

func NewSkillRanksPrerequisite(
	id skill.SkillID,
	minimumRanks int,
) (SkillRanksPrerequisite, bool) {
	value := skillRanksPrerequisite{
		skillID:      id,
		minimumRanks: minimumRanks,
	}
	if !value.isValid() {
		return skillRanksPrerequisite{}, false
	}

	return value, true
}

func NewAnySkillRanksPrerequisite(
	ids []skill.SkillID,
	minimumRanks int,
) (AnySkillRanksPrerequisite, bool) {
	value := anySkillRanksPrerequisite{
		skillIDs:     append([]skill.SkillID(nil), ids...),
		minimumRanks: minimumRanks,
	}
	if !value.isValid() {
		return anySkillRanksPrerequisite{}, false
	}

	return value, true
}

func NewSpellcastingPrerequisite(access SpellcastingAccess) (SpellcastingPrerequisite, bool) {
	value := spellcastingPrerequisite{access: access}
	if !value.isValid() {
		return spellcastingPrerequisite{}, false
	}

	return value, true
}

func NewCasterLevelPrerequisite(minimumLevel int) (CasterLevelPrerequisite, bool) {
	value := casterLevelPrerequisite{minimumLevel: minimumLevel}
	if !value.isValid() {
		return casterLevelPrerequisite{}, false
	}

	return value, true
}

func NewCharacterLevelPrerequisite(minimumLevel int) (CharacterLevelPrerequisite, bool) {
	value := characterLevelPrerequisite{minimumLevel: minimumLevel}
	if !value.isValid() {
		return characterLevelPrerequisite{}, false
	}

	return value, true
}

func NewClassLevelPrerequisite(
	id characterclass.ClassID,
	minimumLevel int,
) (ClassLevelPrerequisite, bool) {
	value := classLevelPrerequisite{
		classID:      id,
		minimumLevel: minimumLevel,
	}
	if !value.isValid() {
		return classLevelPrerequisite{}, false
	}

	return value, true
}

func NewClassFeaturePrerequisite(
	id characterclass.ClassFeatureID,
) (ClassFeaturePrerequisite, bool) {
	value := classFeaturePrerequisite{featureID: id}
	if !value.isValid() {
		return classFeaturePrerequisite{}, false
	}

	return value, true
}

func NewSelectedWeaponProficiencyPrerequisite() SelectedWeaponProficiencyPrerequisite {
	return selectedWeaponProficiencyPrerequisite{valid: true}
}

func NewSelectedFamiliarEligibilityPrerequisite() SelectedFamiliarEligibilityPrerequisite {
	return selectedFamiliarEligibilityPrerequisite{valid: true}
}

func NewFeatPrerequisite(id FeatID) (FeatPrerequisite, bool) {
	value := featPrerequisite{featID: id}
	if !value.isValid() {
		return featPrerequisite{}, false
	}

	return value, true
}

func NewAnyFeatPrerequisite(ids []FeatID) (AnyFeatPrerequisite, bool) {
	copied := make([]featID, 0, len(ids))
	for _, id := range ids {
		copied = append(copied, featID(id))
	}

	value := anyFeatPrerequisite{featIDs: copied}
	if !value.isValid() {
		return anyFeatPrerequisite{}, false
	}

	return value, true
}

func NewSameSelectionFeatPrerequisite(id FeatID) (SameSelectionFeatPrerequisite, bool) {
	value := sameSelectionFeatPrerequisite{featID: id}
	if !value.isValid() {
		return sameSelectionFeatPrerequisite{}, false
	}

	return value, true
}

func NewSpellSchoolFeatPrerequisite(
	id FeatID,
	schoolID spell.SchoolID,
) (SpellSchoolFeatPrerequisite, bool) {
	value := spellSchoolFeatPrerequisite{
		featID:   id,
		schoolID: schoolID,
	}
	if !value.isValid() {
		return spellSchoolFeatPrerequisite{}, false
	}

	return value, true
}

func (p prerequisiteList) GetPrerequisites() []Prerequisite {
	return append([]Prerequisite(nil), p.prerequisites...)
}

func (p abilityScorePrerequisite) GetKind() PrerequisiteKind {
	return AbilityScorePrerequisiteKind
}

func (p abilityScorePrerequisite) GetAbilityScoreID() ability.AbilityScoreID {
	return p.abilityScoreID
}

func (p abilityScorePrerequisite) GetMinimumScore() int {
	return p.minimumScore
}

func (p baseAttackBonusPrerequisite) GetKind() PrerequisiteKind {
	return BaseAttackBonusPrerequisiteKind
}

func (p baseAttackBonusPrerequisite) GetMinimumBonus() int {
	return p.minimumBonus
}

func (p skillRanksPrerequisite) GetKind() PrerequisiteKind {
	return SkillRanksPrerequisiteKind
}

func (p skillRanksPrerequisite) GetSkillID() skill.SkillID {
	return p.skillID
}

func (p skillRanksPrerequisite) GetMinimumRanks() int {
	return p.minimumRanks
}

func (p anySkillRanksPrerequisite) GetKind() PrerequisiteKind {
	return AnySkillRanksPrerequisiteKind
}

func (p anySkillRanksPrerequisite) GetSkillIDs() []skill.SkillID {
	return append([]skill.SkillID(nil), p.skillIDs...)
}

func (p anySkillRanksPrerequisite) GetMinimumRanks() int {
	return p.minimumRanks
}

func (p spellcastingPrerequisite) GetKind() PrerequisiteKind {
	return SpellcastingPrerequisiteKind
}

func (p spellcastingPrerequisite) GetAccess() SpellcastingAccess {
	return p.access
}

func (p casterLevelPrerequisite) GetKind() PrerequisiteKind {
	return CasterLevelPrerequisiteKind
}

func (p casterLevelPrerequisite) GetMinimumLevel() int {
	return p.minimumLevel
}

func (p characterLevelPrerequisite) GetKind() PrerequisiteKind {
	return CharacterLevelPrerequisiteKind
}

func (p characterLevelPrerequisite) GetMinimumLevel() int {
	return p.minimumLevel
}

func (p classLevelPrerequisite) GetKind() PrerequisiteKind {
	return ClassLevelPrerequisiteKind
}

func (p classLevelPrerequisite) GetClassID() characterclass.ClassID {
	return p.classID
}

func (p classLevelPrerequisite) GetMinimumLevel() int {
	return p.minimumLevel
}

func (p classFeaturePrerequisite) GetKind() PrerequisiteKind {
	return ClassFeaturePrerequisiteKind
}

func (p classFeaturePrerequisite) GetFeatureID() characterclass.ClassFeatureID {
	return p.featureID
}

func (p selectedWeaponProficiencyPrerequisite) GetKind() PrerequisiteKind {
	return SelectedWeaponProficiencyPrerequisiteKind
}

func (p selectedFamiliarEligibilityPrerequisite) GetKind() PrerequisiteKind {
	return SelectedFamiliarEligibilityPrerequisiteKind
}

func (p featPrerequisite) GetKind() PrerequisiteKind {
	return FeatPrerequisiteKind
}

func (p featPrerequisite) GetFeatID() FeatID {
	return p.featID
}

func (p anyFeatPrerequisite) GetKind() PrerequisiteKind {
	return AnyFeatPrerequisiteKind
}

func (p anyFeatPrerequisite) GetFeatIDs() []FeatID {
	ids := make([]FeatID, 0, len(p.featIDs))
	for _, id := range p.featIDs {
		ids = append(ids, FeatID(id))
	}

	return ids
}

func (p sameSelectionFeatPrerequisite) GetKind() PrerequisiteKind {
	return SameSelectionFeatPrerequisiteKind
}

func (p sameSelectionFeatPrerequisite) GetFeatID() FeatID {
	return p.featID
}

func (p spellSchoolFeatPrerequisite) GetKind() PrerequisiteKind {
	return SpellSchoolFeatPrerequisiteKind
}

func (p spellSchoolFeatPrerequisite) GetFeatID() FeatID {
	return p.featID
}

func (p spellSchoolFeatPrerequisite) GetSchoolID() spell.SchoolID {
	return p.schoolID
}

func (p abilityScorePrerequisite) isPrerequisite() {}

func (p abilityScorePrerequisite) isValid() bool {
	return p.abilityScoreID.GetName() != "" && p.minimumScore > 0
}

func (p baseAttackBonusPrerequisite) isPrerequisite() {}

func (p baseAttackBonusPrerequisite) isValid() bool {
	return p.minimumBonus > 0
}

func (p skillRanksPrerequisite) isPrerequisite() {}

func (p skillRanksPrerequisite) isValid() bool {
	return p.skillID.GetName() != "" && p.minimumRanks > 0
}

func (p anySkillRanksPrerequisite) isPrerequisite() {}

func (p anySkillRanksPrerequisite) isValid() bool {
	if len(p.skillIDs) == 0 || p.minimumRanks <= 0 {
		return false
	}

	seen := make(map[skill.SkillID]struct{}, len(p.skillIDs))
	for _, id := range p.skillIDs {
		if id.GetName() == "" {
			return false
		}

		if _, ok := seen[id]; ok {
			return false
		}

		seen[id] = struct{}{}
	}

	return true
}

func (p spellcastingPrerequisite) isPrerequisite() {}

func (p spellcastingPrerequisite) isValid() bool {
	return isValidSpellcastingAccess(p.access)
}

func (p casterLevelPrerequisite) isPrerequisite() {}

func (p casterLevelPrerequisite) isValid() bool {
	return p.minimumLevel > 0
}

func (p characterLevelPrerequisite) isPrerequisite() {}

func (p characterLevelPrerequisite) isValid() bool {
	return p.minimumLevel > 0
}

func (p classLevelPrerequisite) isPrerequisite() {}

func (p classLevelPrerequisite) isValid() bool {
	_, ok := characterclass.GetClassByID(p.classID)
	return ok && p.minimumLevel > 0
}

func (p classFeaturePrerequisite) isPrerequisite() {}

func (p classFeaturePrerequisite) isValid() bool {
	return p.featureID.GetName() != ""
}

func (p selectedWeaponProficiencyPrerequisite) isPrerequisite() {}

func (p selectedWeaponProficiencyPrerequisite) isValid() bool {
	return p.valid
}

func (p selectedFamiliarEligibilityPrerequisite) isPrerequisite() {}

func (p selectedFamiliarEligibilityPrerequisite) isValid() bool {
	return p.valid
}

func (p featPrerequisite) isPrerequisite() {}

func (p featPrerequisite) isValid() bool {
	return isValidFeatID(p.featID)
}

func (p anyFeatPrerequisite) isPrerequisite() {}

func (p anyFeatPrerequisite) isValid() bool {
	if len(p.featIDs) == 0 {
		return false
	}

	seen := make(map[featID]struct{}, len(p.featIDs))
	for _, id := range p.featIDs {
		if !isValidFeatID(FeatID(id)) {
			return false
		}

		if _, ok := seen[id]; ok {
			return false
		}

		seen[id] = struct{}{}
	}

	return true
}

func (p sameSelectionFeatPrerequisite) isPrerequisite() {}

func (p sameSelectionFeatPrerequisite) isValid() bool {
	return isValidFeatID(p.featID)
}

func (p spellSchoolFeatPrerequisite) isPrerequisite() {}

func (p spellSchoolFeatPrerequisite) isValid() bool {
	return isValidSpellSchoolID(p.schoolID) && isValidFeatID(p.featID)
}

func isValidPrerequisite(value Prerequisite) bool {
	return value != nil && value.isValid()
}

func isValidSpellcastingAccess(value SpellcastingAccess) bool {
	switch value {
	case AnySpellcastingAccess,
		ArcaneSpellcastingAccess,
		DivineSpellcastingAccess:
		return true
	default:
		return false
	}
}

func isValidSpellSchoolID(id spell.SchoolID) bool {
	switch id {
	case spell.AbjurationSchoolID,
		spell.ConjurationSchoolID,
		spell.DivinationSchoolID,
		spell.EnchantmentSchoolID,
		spell.EvocationSchoolID,
		spell.IllusionSchoolID,
		spell.NecromancySchoolID,
		spell.TransmutationSchoolID,
		spell.UniversalSchoolID:
		return true
	default:
		return false
	}
}

func isValidFeatID(id FeatID) bool {
	value := string(id)
	return value != "" && strings.TrimSpace(value) == value
}
