package skill

var coreSkills = mustBuildCoreSkills()

var coreSkillOrder = []SkillID{
	AcrobaticsSkillID,
	AppraiseSkillID,
	BluffSkillID,
	ClimbSkillID,
	CraftSkillID,
	DiplomacySkillID,
	DisableDeviceSkillID,
	DisguiseSkillID,
	EscapeArtistSkillID,
	FlySkillID,
	HandleAnimalSkillID,
	HealSkillID,
	IntimidateSkillID,
	KnowledgeSkillID,
	LinguisticsSkillID,
	PerceptionSkillID,
	PerformSkillID,
	ProfessionSkillID,
	RideSkillID,
	SenseMotiveSkillID,
	SleightOfHandSkillID,
	SpellcraftSkillID,
	StealthSkillID,
	SurvivalSkillID,
	SwimSkillID,
	UseMagicDeviceSkillID,
}

func GetSkillByID(id SkillID) (Skill, bool) {
	value, ok := coreSkills[id]
	if !ok {
		return skill{}, false
	}

	return value, true
}

func GetSkills() []Skill {
	skills := make([]Skill, 0, len(coreSkillOrder))

	for _, id := range coreSkillOrder {
		skills = append(skills, coreSkills[id])
	}

	return skills
}

func mustBuildCoreSkills() map[SkillID]Skill {
	return map[SkillID]Skill{
		AcrobaticsSkillID:     mustNewSkill(AcrobaticsSkillID, false, true, false),
		AppraiseSkillID:       mustNewSkill(AppraiseSkillID, false, false, false),
		BluffSkillID:          mustNewSkill(BluffSkillID, false, false, false),
		ClimbSkillID:          mustNewSkill(ClimbSkillID, false, true, false),
		CraftSkillID:          mustNewSkill(CraftSkillID, false, false, true),
		DiplomacySkillID:      mustNewSkill(DiplomacySkillID, false, false, false),
		DisableDeviceSkillID:  mustNewSkill(DisableDeviceSkillID, true, true, false),
		DisguiseSkillID:       mustNewSkill(DisguiseSkillID, false, false, false),
		EscapeArtistSkillID:   mustNewSkill(EscapeArtistSkillID, false, true, false),
		FlySkillID:            mustNewSkill(FlySkillID, false, true, false),
		HandleAnimalSkillID:   mustNewSkill(HandleAnimalSkillID, true, false, false),
		HealSkillID:           mustNewSkill(HealSkillID, false, false, false),
		IntimidateSkillID:     mustNewSkill(IntimidateSkillID, false, false, false),
		KnowledgeSkillID:      mustNewSkill(KnowledgeSkillID, true, false, true),
		LinguisticsSkillID:    mustNewSkill(LinguisticsSkillID, true, false, false),
		PerceptionSkillID:     mustNewSkill(PerceptionSkillID, false, false, false),
		PerformSkillID:        mustNewSkill(PerformSkillID, false, false, true),
		ProfessionSkillID:     mustNewSkill(ProfessionSkillID, true, false, true),
		RideSkillID:           mustNewSkill(RideSkillID, false, true, false),
		SenseMotiveSkillID:    mustNewSkill(SenseMotiveSkillID, false, false, false),
		SleightOfHandSkillID:  mustNewSkill(SleightOfHandSkillID, true, true, false),
		SpellcraftSkillID:     mustNewSkill(SpellcraftSkillID, true, false, false),
		StealthSkillID:        mustNewSkill(StealthSkillID, false, true, false),
		SurvivalSkillID:       mustNewSkill(SurvivalSkillID, false, false, false),
		SwimSkillID:           mustNewSkill(SwimSkillID, false, true, false),
		UseMagicDeviceSkillID: mustNewSkill(UseMagicDeviceSkillID, true, false, false),
	}
}

func mustNewSkill(id SkillID, trainedOnly bool, armorCheckPenaltyApplies bool, grouped bool) Skill {
	skill, ok := NewSkill(id, trainedOnly, armorCheckPenaltyApplies, grouped)
	if !ok {
		panic("invalid core skill seed")
	}

	return skill
}
