package spell

import "strings"

type spellID string
type SpellID = spellID

type schoolID string
type SchoolID = schoolID

type descriptorID string
type DescriptorID = descriptorID

type componentID string
type ComponentID = componentID

type spell struct {
	id              spellID
	school          schoolID
	descriptors     []descriptorID
	components      []componentID
	castingTime     string
	spellRange      string
	targetEffect    string
	duration        string
	savingThrow     string
	spellResistance string
}
type Spell = spell

const (
	AbjurationSchoolID    SchoolID = "Abjuration"
	ConjurationSchoolID   SchoolID = "Conjuration"
	DivinationSchoolID    SchoolID = "Divination"
	EnchantmentSchoolID   SchoolID = "Enchantment"
	EvocationSchoolID     SchoolID = "Evocation"
	IllusionSchoolID      SchoolID = "Illusion"
	NecromancySchoolID    SchoolID = "Necromancy"
	TransmutationSchoolID SchoolID = "Transmutation"
	UniversalSchoolID     SchoolID = "Universal"
)

const (
	VerbalComponentID      ComponentID = "V"
	SomaticComponentID     ComponentID = "S"
	MaterialComponentID    ComponentID = "M"
	FocusComponentID       ComponentID = "F"
	DivineFocusComponentID ComponentID = "DF"
)

func NewSpell(
	id SpellID,
	school SchoolID,
	descriptors []DescriptorID,
	components []ComponentID,
	castingTime string,
	spellRange string,
	targetEffect string,
	duration string,
	savingThrow string,
	spellResistance string,
) (Spell, bool) {
	if !isValidSpellID(id) ||
		!isValidSchoolID(school) ||
		!isValidTextClause(castingTime) ||
		!isValidTextClause(spellRange) ||
		!isValidTextClause(targetEffect) ||
		!isValidTextClause(duration) ||
		!isValidTextClause(savingThrow) ||
		!isValidTextClause(spellResistance) {
		return spell{}, false
	}

	dedupedDescriptors, ok := dedupeDescriptorIDs(descriptors)
	if !ok {
		return spell{}, false
	}

	dedupedComponents, ok := dedupeComponentIDs(components)
	if !ok {
		return spell{}, false
	}

	return spell{
		id:              id,
		school:          school,
		descriptors:     dedupedDescriptors,
		components:      dedupedComponents,
		castingTime:     castingTime,
		spellRange:      spellRange,
		targetEffect:    targetEffect,
		duration:        duration,
		savingThrow:     savingThrow,
		spellResistance: spellResistance,
	}, true
}

func (s spell) GetID() SpellID {
	return s.id
}

func (s spell) GetSchool() SchoolID {
	return s.school
}

func (s spell) GetDescriptors() []DescriptorID {
	return append([]DescriptorID(nil), s.descriptors...)
}

func (s spell) GetComponents() []ComponentID {
	return append([]ComponentID(nil), s.components...)
}

func (s spell) GetCastingTime() string {
	return s.castingTime
}

func (s spell) GetRange() string {
	return s.spellRange
}

func (s spell) GetTargetEffect() string {
	return s.targetEffect
}

func (s spell) GetDuration() string {
	return s.duration
}

func (s spell) GetSavingThrow() string {
	return s.savingThrow
}

func (s spell) GetSpellResistance() string {
	return s.spellResistance
}

func isValidSpellID(id SpellID) bool {
	value := string(id)
	return value != "" && strings.TrimSpace(value) == value
}

func isValidSchoolID(id SchoolID) bool {
	switch id {
	case AbjurationSchoolID,
		ConjurationSchoolID,
		DivinationSchoolID,
		EnchantmentSchoolID,
		EvocationSchoolID,
		IllusionSchoolID,
		NecromancySchoolID,
		TransmutationSchoolID,
		UniversalSchoolID:
		return true
	default:
		return false
	}
}

func isValidDescriptorID(id DescriptorID) bool {
	value := string(id)
	return value != "" && strings.TrimSpace(value) == value
}

func isValidComponentID(id ComponentID) bool {
	switch id {
	case VerbalComponentID,
		SomaticComponentID,
		MaterialComponentID,
		FocusComponentID,
		DivineFocusComponentID:
		return true
	default:
		return false
	}
}

func isValidTextClause(value string) bool {
	return value != "" && strings.TrimSpace(value) == value
}

func dedupeDescriptorIDs(descriptors []DescriptorID) ([]DescriptorID, bool) {
	if len(descriptors) == 0 {
		return nil, true
	}

	seen := make(map[DescriptorID]struct{}, len(descriptors))
	deduped := make([]DescriptorID, 0, len(descriptors))

	for _, descriptor := range descriptors {
		if !isValidDescriptorID(descriptor) {
			return nil, false
		}

		if _, ok := seen[descriptor]; ok {
			continue
		}

		seen[descriptor] = struct{}{}
		deduped = append(deduped, descriptor)
	}

	return deduped, true
}

func dedupeComponentIDs(components []ComponentID) ([]ComponentID, bool) {
	if len(components) == 0 {
		return nil, false
	}

	seen := make(map[ComponentID]struct{}, len(components))
	deduped := make([]ComponentID, 0, len(components))

	for _, component := range components {
		if !isValidComponentID(component) {
			return nil, false
		}

		if _, ok := seen[component]; ok {
			continue
		}

		seen[component] = struct{}{}
		deduped = append(deduped, component)
	}

	return deduped, true
}
