package skill

import "testing"

func TestCoreSkills_SeedsTwentySixCoreEntries(t *testing.T) {
	testCases := []struct {
		id                       SkillID
		trainedOnly              bool
		armorCheckPenaltyApplies bool
		grouped                  bool
	}{
		{AcrobaticsSkillID, false, true, false},
		{AppraiseSkillID, false, false, false},
		{BluffSkillID, false, false, false},
		{ClimbSkillID, false, true, false},
		{CraftSkillID, false, false, true},
		{DiplomacySkillID, false, false, false},
		{DisableDeviceSkillID, true, true, false},
		{DisguiseSkillID, false, false, false},
		{EscapeArtistSkillID, false, true, false},
		{FlySkillID, false, true, false},
		{HandleAnimalSkillID, true, false, false},
		{HealSkillID, false, false, false},
		{IntimidateSkillID, false, false, false},
		{KnowledgeSkillID, true, false, true},
		{LinguisticsSkillID, true, false, false},
		{PerceptionSkillID, false, false, false},
		{PerformSkillID, false, false, true},
		{ProfessionSkillID, true, false, true},
		{RideSkillID, false, true, false},
		{SenseMotiveSkillID, false, false, false},
		{SleightOfHandSkillID, true, true, false},
		{SpellcraftSkillID, true, false, false},
		{StealthSkillID, false, true, false},
		{SurvivalSkillID, false, false, false},
		{SwimSkillID, false, true, false},
		{UseMagicDeviceSkillID, true, false, false},
	}

	if len(coreSkills) != len(testCases) {
		t.Fatalf("expected %d core skills, got %d", len(testCases), len(coreSkills))
	}

	for _, tc := range testCases {
		skill, ok := coreSkills[tc.id]
		if !ok {
			t.Fatalf("expected core skill %q to be seeded", tc.id)
		}

		if skill.GetID() != tc.id {
			t.Fatalf("expected skill id %q, got %q", tc.id, skill.GetID())
		}

		if skill.IsTrainedOnly() != tc.trainedOnly {
			t.Fatalf("expected skill %q trained-only=%t, got %t", tc.id, tc.trainedOnly, skill.IsTrainedOnly())
		}

		if skill.AppliesArmorCheckPenalty() != tc.armorCheckPenaltyApplies {
			t.Fatalf("expected skill %q armor-check-penalty=%t, got %t", tc.id, tc.armorCheckPenaltyApplies, skill.AppliesArmorCheckPenalty())
		}

		if skill.IsGrouped() != tc.grouped {
			t.Fatalf("expected skill %q grouped=%t, got %t", tc.id, tc.grouped, skill.IsGrouped())
		}
	}
}

func TestNewSkill_AcceptsEverySeededCoreSkillID(t *testing.T) {
	for id, seeded := range coreSkills {
		skill, ok := NewSkill(id, seeded.IsTrainedOnly(), seeded.AppliesArmorCheckPenalty(), seeded.IsGrouped())
		if !ok {
			t.Fatalf("expected skill %q to be constructible from its seeded metadata", id)
		}

		if skill.GetID() != seeded.GetID() {
			t.Fatalf("expected constructed skill id %q, got %q", seeded.GetID(), skill.GetID())
		}
	}
}

func TestNewSkill_RejectsUnknownCoreLikeSkillIDs(t *testing.T) {
	invalidIDs := []SkillID{
		"Jump",
		"Open Lock",
		"knowledge",
		"Use magic device",
		"Knowledge(arcana)",
	}

	for _, id := range invalidIDs {
		if _, ok := NewSkill(id, false, false, false); ok {
			t.Fatalf("expected non-core or non-canonical skill id %q to be rejected", id)
		}
	}
}

func TestNewSkill_AcceptsGroupedSpecializationsWithoutChangingCoreCatalog(t *testing.T) {
	skill, ok := NewSkill("Perform (sing)", false, false, true)
	if !ok {
		t.Fatal("expected grouped specialization to be accepted")
	}

	if skill.GetFamilyID() != PerformSkillID {
		t.Fatalf("expected grouped specialization family %q, got %q", PerformSkillID, skill.GetFamilyID())
	}

	if len(coreSkills) != len(coreSkillOrder) {
		t.Fatalf("expected grouped specialization support not to change core catalog size, got %d skills and %d ordered ids", len(coreSkills), len(coreSkillOrder))
	}
}

func TestGetSkillByID_ReturnsSeededCoreSkill(t *testing.T) {
	skill, ok := GetSkillByID(PerceptionSkillID)
	if !ok {
		t.Fatal("expected perception skill lookup to succeed")
	}

	if skill.GetID() != PerceptionSkillID {
		t.Fatalf("expected looked up skill id %q, got %q", PerceptionSkillID, skill.GetID())
	}

	if skill.IsTrainedOnly() {
		t.Fatal("expected perception to not be trained-only")
	}
}

func TestGetSkillByID_RejectsUnknownSkill(t *testing.T) {
	if _, ok := GetSkillByID(SkillID("Jump")); ok {
		t.Fatal("expected unknown skill lookup to fail")
	}
}

func TestGetSkills_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	skills := GetSkills()
	if len(skills) != len(coreSkillOrder) {
		t.Fatalf("expected %d queried skills, got %d", len(coreSkillOrder), len(skills))
	}

	for i, expectedID := range coreSkillOrder {
		if skills[i].GetID() != expectedID {
			t.Fatalf("expected skill at index %d to be %q, got %q", i, expectedID, skills[i].GetID())
		}
	}
}

func TestGetSkills_ReturnsDetachedSlice(t *testing.T) {
	first := GetSkills()
	second := GetSkills()

	first[0] = skill{}

	if second[0].GetID() != AcrobaticsSkillID {
		t.Fatalf("expected detached skill slice to preserve %q, got %q", AcrobaticsSkillID, second[0].GetID())
	}
}
