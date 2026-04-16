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
	case "":
		return resolveHighestValue(modifiers)

	case ModifierUntyped:
		return sumModifiers(modifiers)

	case ModifierDodge:
		return sumModifiers(modifiers)

	case ModifierCircumstance:
		bySource := make(map[ModifierSource][]Modifier)

		for _, circumstanceModifier := range modifiers {
			bySource[circumstanceModifier.Source] = append(bySource[circumstanceModifier.Source], circumstanceModifier)
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
		sum += modifier.Value
	}
	return sum
}

func resolveHighestBonusAndStackingPenalties(modifiers []Modifier) int {
	highestBonus := 0
	penalties := 0

	for _, modifier := range modifiers {
		switch {
		case modifier.Value > highestBonus:
			highestBonus = modifier.Value
		case modifier.Value < 0:
			penalties += modifier.Value
		}
	}

	return highestBonus + penalties
}

func resolveHighestBonusAndWorstPenalty(modifiers []Modifier) int {
	highestBonus := 0
	worstPenalty := 0

	for _, modifier := range modifiers {
		switch {
		case modifier.Value > highestBonus:
			highestBonus = modifier.Value
		case modifier.Value < worstPenalty:
			worstPenalty = modifier.Value
		}
	}

	return highestBonus + worstPenalty
}

func resolveHighestValue(modifiers []Modifier) int {
	max := 0
	for i, modifier := range modifiers {
		if i == 0 || modifier.Value > max {
			max = modifier.Value
		}
	}

	return max
}
