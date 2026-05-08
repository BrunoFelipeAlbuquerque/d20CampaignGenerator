package character

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

func TestNewCharacterClass_ComposesCoreClassThroughCharacterBoundary(t *testing.T) {
	selectedClass, ok := NewCharacterClass(characterclass.WizardClassID)
	if !ok {
		t.Fatal("expected core class to compose through character boundary")
	}

	if selectedClass.GetClassID() != characterclass.WizardClassID {
		t.Fatalf("expected selected class id %q, got %q", characterclass.WizardClassID, selectedClass.GetClassID())
	}

	class, ok := selectedClass.GetClass()
	if !ok {
		t.Fatal("expected selected core class to resolve")
	}

	if class.GetID() != characterclass.WizardClassID {
		t.Fatalf("expected resolved class id %q, got %q", characterclass.WizardClassID, class.GetID())
	}

	if class.GetHitDieType() != ability.D6HitDie {
		t.Fatalf("expected resolved class hit die %q, got %q", ability.D6HitDie, class.GetHitDieType())
	}

	if class.GetBaseAttackBonusProgression() != ability.BaseAttackBonusHalf {
		t.Fatalf("expected resolved class BAB progression %q, got %q", ability.BaseAttackBonusHalf, class.GetBaseAttackBonusProgression())
	}

	if class.GetSkillRanksPerLevel() != 2 {
		t.Fatalf("expected resolved class skill ranks per level 2, got %d", class.GetSkillRanksPerLevel())
	}

	if class.GetSpellcasting().GetKind() != characterclass.ArcanePreparedSpellcastingKind {
		t.Fatalf("expected resolved class spellcasting %q, got %q", characterclass.ArcanePreparedSpellcastingKind, class.GetSpellcasting().GetKind())
	}
}

func TestNewCharacterClass_RejectsUnknownClass(t *testing.T) {
	if _, ok := NewCharacterClass(characterclass.ClassID("alchemist")); ok {
		t.Fatal("expected unknown class to be rejected")
	}
}

func TestNewCharacterClass_RejectsMalformedClassID(t *testing.T) {
	if _, ok := NewCharacterClass(characterclass.ClassID(" fighter")); ok {
		t.Fatal("expected malformed class id to be rejected")
	}
}

func TestCharacterClass_GetClassReturnsDetachedCatalogClass(t *testing.T) {
	selectedClass, ok := NewCharacterClass(characterclass.BarbarianClassID)
	if !ok {
		t.Fatal("expected core class to compose through character boundary")
	}

	first, ok := selectedClass.GetClass()
	if !ok {
		t.Fatal("expected selected core class to resolve")
	}

	classSkills := first.GetClassSkills()
	if len(classSkills) == 0 {
		t.Fatal("expected barbarian to have class skills")
	}
	classSkills[0] = skill.AppraiseSkillID

	second, ok := selectedClass.GetClass()
	if !ok {
		t.Fatal("expected selected core class to resolve again")
	}

	if second.GetClassSkills()[0] != skill.AcrobaticsSkillID {
		t.Fatalf("expected resolved class skill to remain %q, got %q", skill.AcrobaticsSkillID, second.GetClassSkills()[0])
	}
}

func TestCharacterClass_ZeroValueDoesNotResolve(t *testing.T) {
	var selectedClass CharacterClass

	if _, ok := selectedClass.GetClass(); ok {
		t.Fatal("expected zero-value character class not to resolve")
	}
}
