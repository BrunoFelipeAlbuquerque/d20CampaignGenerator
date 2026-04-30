package feat_test

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/feat"
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

	prerequisites, ok := feat.NewPrerequisiteList([]feat.Prerequisite{
		scorePrerequisite,
		classFeaturePrerequisite,
	})
	if !ok {
		t.Fatal("expected public constructor prerequisites to compose")
	}

	got := prerequisites.GetPrerequisites()
	if len(got) != 2 {
		t.Fatalf("expected 2 prerequisites, got %d", len(got))
	}
}
