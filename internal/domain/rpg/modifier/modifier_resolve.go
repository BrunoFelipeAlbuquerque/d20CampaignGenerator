package modifier

import "strings"

type ModifierType string

type ModifierSource string

type targetID string
type TargetID = targetID

type conditionID string
type ConditionID = conditionID

type ModifierCircumstanceSource = ModifierSource

type ModifierList []Modifier

type ModifierCondition []Condition

type Target interface {
	isTarget()
}

type Condition interface {
	isCondition()
}

type targetRef struct {
	id targetID
}

type TargetRef = targetRef

type conditionRef struct {
	id conditionID
}

type ConditionRef = conditionRef

type Modifier struct {
	modifierType ModifierType
	value        int
	source       ModifierSource
	target       Target
	condition    ModifierCondition
}

func NewTargetRef(id string) (TargetRef, bool) {
	value := strings.TrimSpace(id)
	if value == "" || value != id {
		return targetRef{}, false
	}

	return targetRef{id: TargetID(value)}, true
}

func NewConditionRef(id string) (ConditionRef, bool) {
	value := strings.TrimSpace(id)
	if value == "" || value != id {
		return conditionRef{}, false
	}

	return conditionRef{id: ConditionID(value)}, true
}

func NewModifier(
	modifierType ModifierType,
	value int,
	source ModifierSource,
	target Target,
	condition ModifierCondition,
) (Modifier, bool) {
	if !isValidModifierType(modifierType) {
		return Modifier{}, false
	}

	if modifierType == ModifierCircumstance && validateModifierSource(source) != nil {
		return Modifier{}, false
	}

	return Modifier{
		modifierType: modifierType,
		value:        value,
		source:       source,
		target:       target,
		condition:    append(ModifierCondition(nil), condition...),
	}, true
}

func (m Modifier) GetType() ModifierType {
	return m.modifierType
}

func (t targetRef) isTarget() {}

func (t targetRef) GetID() TargetID {
	return t.id
}

func (c conditionRef) isCondition() {}

func (c conditionRef) GetID() ConditionID {
	return c.id
}

func (m Modifier) GetValue() int {
	return m.value
}

func (m Modifier) GetSource() ModifierSource {
	return m.source
}

func (m Modifier) GetTarget() Target {
	return m.target
}

func (m Modifier) GetCondition() ModifierCondition {
	return append(ModifierCondition(nil), m.condition...)
}

func (mods ModifierList) ModifierResolve() (int, bool) {
	total := 0

	grouped := make(map[ModifierType][]Modifier)

	for _, m := range mods {
		if !isValidModifier(m) {
			return 0, false
		}

		grouped[m.modifierType] = append(grouped[m.modifierType], m)
	}

	for t, group := range grouped {
		total += resolveByType(t, group)
	}

	return total, true
}

func isValidModifier(m Modifier) bool {
	if !isValidModifierType(m.modifierType) {
		return false
	}

	if m.modifierType == ModifierCircumstance && validateModifierSource(m.source) != nil {
		return false
	}

	return true
}

func resolveByType(modifierType ModifierType, modifiers []Modifier) int {
	switch modifierType {
	case ModifierUntyped:
		return sumModifiers(modifiers)

	case ModifierDodge:
		return sumModifiers(modifiers)

	case ModifierCircumstance:
		bySource := make(map[ModifierSource][]Modifier)

		for _, circumstanceModifier := range modifiers {
			bySource[circumstanceModifier.source] = append(bySource[circumstanceModifier.source], circumstanceModifier)
		}

		sum := 0
		for _, sourceModifiers := range bySource {
			sum += resolveHighestBonusAndWorstPenalty(sourceModifiers)
		}
		return sum

	default:
		return resolveHighestBonusAndStackingPenalties(modifiers)
	}
}

func sumModifiers(modifiers []Modifier) int {
	sum := 0
	for _, modifier := range modifiers {
		sum += modifier.value
	}
	return sum
}

func resolveHighestBonusAndStackingPenalties(modifiers []Modifier) int {
	highestBonus := 0
	penalties := 0

	for _, modifier := range modifiers {
		switch {
		case modifier.value > highestBonus:
			highestBonus = modifier.value
		case modifier.value < 0:
			penalties += modifier.value
		}
	}

	return highestBonus + penalties
}

func resolveHighestBonusAndWorstPenalty(modifiers []Modifier) int {
	highestBonus := 0
	worstPenalty := 0

	for _, modifier := range modifiers {
		switch {
		case modifier.value > highestBonus:
			highestBonus = modifier.value
		case modifier.value < worstPenalty:
			worstPenalty = modifier.value
		}
	}

	return highestBonus + worstPenalty
}

func resolveHighestValue(modifiers []Modifier) int {
	max := 0
	for i, modifier := range modifiers {
		if i == 0 || modifier.value > max {
			max = modifier.value
		}
	}

	return max
}
