package skill

import "strings"

type skillID string
type SkillID = skillID

const (
	CraftSkillID      SkillID = "Craft"
	KnowledgeSkillID  SkillID = "Knowledge"
	PerformSkillID    SkillID = "Perform"
	ProfessionSkillID SkillID = "Profession"
)

type skill struct {
	id                       skillID
	trainedOnly              bool
	armorCheckPenaltyApplies bool
	grouped                  bool
}
type Skill = skill

func NewSkill(id SkillID, trainedOnly bool, armorCheckPenaltyApplies bool, grouped bool) (Skill, bool) {
	if !isValidSkillID(id) || grouped != isGroupedSkillID(id) {
		return skill{}, false
	}

	return skill{
		id:                       id,
		trainedOnly:              trainedOnly,
		armorCheckPenaltyApplies: armorCheckPenaltyApplies,
		grouped:                  grouped,
	}, true
}

func (id skillID) GetName() string {
	name := strings.TrimSpace(string(id))
	if name == "" || name != string(id) {
		return ""
	}

	return name
}

func (s skill) GetID() SkillID {
	return s.id
}

func (s skill) IsTrainedOnly() bool {
	return s.trainedOnly
}

func (s skill) AppliesArmorCheckPenalty() bool {
	return s.armorCheckPenaltyApplies
}

func (s skill) IsGrouped() bool {
	return s.grouped
}

func isValidSkillID(id SkillID) bool {
	return id.GetName() != ""
}

func isGroupedSkillID(id SkillID) bool {
	switch id {
	case CraftSkillID, KnowledgeSkillID, PerformSkillID, ProfessionSkillID:
		return true
	default:
		return false
	}
}
