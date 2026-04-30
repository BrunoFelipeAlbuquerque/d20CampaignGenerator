package feat

type featCategory string
type FeatCategory = featCategory

const (
	GeneralFeatCategory      FeatCategory = "General"
	CombatFeatCategory       FeatCategory = "Combat"
	CriticalFeatCategory     FeatCategory = "Critical"
	ItemCreationFeatCategory FeatCategory = "Item Creation"
	MetamagicFeatCategory    FeatCategory = "Metamagic"
)

type feat struct {
	id               featID
	category         featCategory
	prerequisites    prerequisiteList
	fighterBonusFeat bool
	metamagic        bool
	itemCreation     bool
}
type Feat = feat

func NewFeat(
	id FeatID,
	category FeatCategory,
	prerequisites PrerequisiteList,
	fighterBonusFeat bool,
	metamagic bool,
	itemCreation bool,
) (Feat, bool) {
	if !isValidFeatID(id) ||
		!isValidFeatCategory(category) ||
		!isValidPrerequisiteList(prerequisites) ||
		!isValidFeatFlags(category, fighterBonusFeat, metamagic, itemCreation) {
		return feat{}, false
	}

	return feat{
		id:               id,
		category:         category,
		prerequisites:    clonePrerequisiteList(prerequisites),
		fighterBonusFeat: fighterBonusFeat,
		metamagic:        metamagic,
		itemCreation:     itemCreation,
	}, true
}

func (id featID) GetName() string {
	if !isValidFeatID(FeatID(id)) {
		return ""
	}

	return string(id)
}

func (c featCategory) GetName() string {
	if !isValidFeatCategory(FeatCategory(c)) {
		return ""
	}

	return string(c)
}

func (f feat) GetID() FeatID {
	return f.id
}

func (f feat) GetCategory() FeatCategory {
	return f.category
}

func (f feat) GetPrerequisites() []Prerequisite {
	return f.prerequisites.GetPrerequisites()
}

func (f feat) IsFighterBonusFeat() bool {
	return f.fighterBonusFeat
}

func (f feat) IsMetamagic() bool {
	return f.metamagic
}

func (f feat) IsItemCreation() bool {
	return f.itemCreation
}

func isValidFeatCategory(category FeatCategory) bool {
	switch category {
	case GeneralFeatCategory,
		CombatFeatCategory,
		CriticalFeatCategory,
		ItemCreationFeatCategory,
		MetamagicFeatCategory:
		return true
	default:
		return false
	}
}

func isValidPrerequisiteList(prerequisites PrerequisiteList) bool {
	for _, prerequisite := range prerequisites.GetPrerequisites() {
		if !isValidPrerequisite(prerequisite) {
			return false
		}
	}

	return true
}

func isValidFeatFlags(
	category FeatCategory,
	fighterBonusFeat bool,
	metamagic bool,
	itemCreation bool,
) bool {
	switch category {
	case GeneralFeatCategory:
		return !fighterBonusFeat && !metamagic && !itemCreation
	case CombatFeatCategory, CriticalFeatCategory:
		return fighterBonusFeat && !metamagic && !itemCreation
	case ItemCreationFeatCategory:
		return !fighterBonusFeat && !metamagic && itemCreation
	case MetamagicFeatCategory:
		return !fighterBonusFeat && metamagic && !itemCreation
	default:
		return false
	}
}

func clonePrerequisiteList(prerequisites PrerequisiteList) PrerequisiteList {
	return prerequisiteList{
		prerequisites: prerequisites.GetPrerequisites(),
	}
}
