package modifier

type ModifierType string

type ModifierSource string

type ModifierCircumstanceSource = ModifierSource

type ModifierList []Modifier

type ModifierCondition []Condition

type Target interface {
	isTarget()
}

type Condition interface {
	isCondition()
}

type Modifier struct {
	Type      ModifierType
	Value     int
	Source    ModifierSource
	Target    Target
	Condition ModifierCondition
}

func (mods ModifierList) ModifierResolve() int {
	total := 0

	grouped := make(map[ModifierType][]Modifier)

	for _, m := range mods {
		grouped[m.Type] = append(grouped[m.Type], m)
	}

	for t, group := range grouped {
		total += resolveByType(t, group)
	}

	return total
}

func resolveByType(modifierType ModifierType, modifiers []Modifier) int {
	switch modifierType {
	case ModifierUntyped:
		sum := 0
		for _, untypedModifier := range modifiers {
			sum += untypedModifier.Value
		}
		return sum

	case ModifierDodge:
		sum := 0
		for _, dodgeModifier := range modifiers {
			sum += dodgeModifier.Value
		}
		return sum

	case ModifierCircumstance:
		bySource := make(map[ModifierSource]int)

		for _, circumstanceModifier := range modifiers {
			if current, ok := bySource[circumstanceModifier.Source]; !ok || circumstanceModifier.Value > current {
				bySource[circumstanceModifier.Source] = circumstanceModifier.Value
			}
		}

		sum := 0
		for _, value := range bySource {
			sum += value
		}
		return sum

	default:
		max := 0
		for i, modifier := range modifiers {
			if i == 0 || modifier.Value > max {
				max = modifier.Value
			}
		}
		return max
	}
}
