package feat_test

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	"d20campaigngenerator/internal/domain/rpg/character/feat"
)

func TestFeat_PublicConstructorComposesPrerequisites(t *testing.T) {
	strengthPrerequisite, ok := feat.NewAbilityScorePrerequisite(ability.StrengthScore, 13)
	if !ok {
		t.Fatal("expected strength prerequisite to be valid")
	}

	prerequisites, ok := feat.NewPrerequisiteList([]feat.Prerequisite{
		strengthPrerequisite,
	})
	if !ok {
		t.Fatal("expected prerequisite list to be valid")
	}

	value, ok := feat.NewFeat(
		feat.FeatID("Power Attack"),
		feat.CombatFeatCategory,
		prerequisites,
		true,
		false,
		false,
	)
	if !ok {
		t.Fatal("expected public feat constructor to succeed")
	}

	if value.GetID().GetName() != "Power Attack" {
		t.Fatalf("expected public feat name %q, got %q", "Power Attack", value.GetID().GetName())
	}

	if value.GetCategory() != feat.CombatFeatCategory || !value.IsFighterBonusFeat() {
		t.Fatal("expected public feat metadata to expose combat fighter-bonus category")
	}

	if len(value.GetPrerequisites()) != 1 {
		t.Fatalf("expected one public feat prerequisite, got %d", len(value.GetPrerequisites()))
	}
}
