package feat_test

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/feat"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

func TestPrerequisiteList_PublicConstructorsCompose(t *testing.T) {
	scorePrerequisite, ok := feat.NewAbilityScorePrerequisite(ability.DexterityScore, 15)
	if !ok {
		t.Fatal("expected ability score prerequisite to be valid")
	}

	classFeaturePrerequisite, ok := feat.NewClassFeaturePrerequisite(characterclass.FighterBonusFeatsClassFeatureID)
	if !ok {
		t.Fatal("expected fighter bonus feats prerequisite access term to be valid")
	}

	anySkillRanksPrerequisite, ok := feat.NewAnySkillRanksPrerequisite(
		[]skill.SkillID{skill.CraftSkillID, skill.ProfessionSkillID},
		5,
	)
	if !ok {
		t.Fatal("expected any-skill ranks prerequisite to be valid")
	}

	selectedFamiliarEligibilityPrerequisite := feat.NewSelectedFamiliarEligibilityPrerequisite()
	anyFeatPrerequisite, ok := feat.NewAnyFeatPrerequisite([]feat.FeatID{
		feat.FeatID("Catch Off-Guard"),
		feat.FeatID("Throw Anything"),
	})
	if !ok {
		t.Fatal("expected any-feat prerequisite to be valid")
	}

	featCategoryCountPrerequisite, ok := feat.NewFeatCategoryCountPrerequisite(feat.CriticalFeatCategory, 2)
	if !ok {
		t.Fatal("expected feat-category count prerequisite to be valid")
	}

	prerequisites, ok := feat.NewPrerequisiteList([]feat.Prerequisite{
		scorePrerequisite,
		classFeaturePrerequisite,
		anySkillRanksPrerequisite,
		selectedFamiliarEligibilityPrerequisite,
		anyFeatPrerequisite,
		featCategoryCountPrerequisite,
	})
	if !ok {
		t.Fatal("expected public constructor prerequisites to compose")
	}

	got := prerequisites.GetPrerequisites()
	if len(got) != 6 {
		t.Fatalf("expected 6 prerequisites, got %d", len(got))
	}
}
