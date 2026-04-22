package skill

import "strings"

type skillID string
type SkillID = skillID

type skill struct {
	id                       skillID
	trainedOnly              bool
	armorCheckPenaltyApplies bool
}
type Skill = skill

func NewSkill(id SkillID, trainedOnly bool, armorCheckPenaltyApplies bool) (Skill, bool) {
	if !isValidSkillID(id) {
		return skill{}, false
	}

	return skill{
		id:                       id,
		trainedOnly:              trainedOnly,
		armorCheckPenaltyApplies: armorCheckPenaltyApplies,
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

func isValidSkillID(id SkillID) bool {
	return id.GetName() != ""
}
