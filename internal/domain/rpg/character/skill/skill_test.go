package skill

import "testing"

func TestNewSkill_ConstructsValidatedSkillChassis(t *testing.T) {
	skill, ok := NewSkill(SkillID("Acrobatics"), false, true)
	if !ok {
		t.Fatal("expected skill chassis to be constructed")
	}

	if skill.GetID() != SkillID("Acrobatics") {
		t.Fatalf("expected skill id %q, got %q", SkillID("Acrobatics"), skill.GetID())
	}

	if skill.GetID().GetName() != "Acrobatics" {
		t.Fatalf("expected skill name %q, got %q", "Acrobatics", skill.GetID().GetName())
	}

	if skill.IsTrainedOnly() {
		t.Fatal("expected Acrobatics not to be trained only")
	}

	if !skill.AppliesArmorCheckPenalty() {
		t.Fatal("expected Acrobatics to apply armor check penalty metadata")
	}
}

func TestNewSkill_AllowsCoreMultiwordSkillIDs(t *testing.T) {
	skill, ok := NewSkill(SkillID("Sleight of Hand"), true, true)
	if !ok {
		t.Fatal("expected multiword skill id to be accepted")
	}

	if !skill.IsTrainedOnly() {
		t.Fatal("expected Sleight of Hand to preserve trained-only metadata")
	}

	if !skill.AppliesArmorCheckPenalty() {
		t.Fatal("expected Sleight of Hand to preserve armor check penalty metadata")
	}
}

func TestNewSkill_RejectsInvalidInputs(t *testing.T) {
	if _, ok := NewSkill("", false, false); ok {
		t.Fatal("expected empty skill id to be rejected")
	}

	if _, ok := NewSkill("   ", false, false); ok {
		t.Fatal("expected blank skill id to be rejected")
	}

	if _, ok := NewSkill(" Acrobatics", false, true); ok {
		t.Fatal("expected skill id with surrounding whitespace to be rejected")
	}

	if SkillID(" ").GetName() != "" {
		t.Fatal("expected invalid skill id name lookup to be empty")
	}
}
