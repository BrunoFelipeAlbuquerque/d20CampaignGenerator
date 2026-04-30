package feat

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
)

func TestNewFeat_ConstructsValidatedFeatChassis(t *testing.T) {
	baseAttackBonusPrerequisite := mustNewBaseAttackBonusPrerequisite(t, 1)
	prerequisites := mustNewPrerequisiteList(t, []Prerequisite{baseAttackBonusPrerequisite})

	feat, ok := NewFeat(
		FeatID("Power Attack"),
		CombatFeatCategory,
		prerequisites,
		true,
		false,
		false,
	)
	if !ok {
		t.Fatal("expected feat chassis to be constructed")
	}

	if feat.GetID() != FeatID("Power Attack") {
		t.Fatalf("expected feat id %q, got %q", FeatID("Power Attack"), feat.GetID())
	}

	if feat.GetID().GetName() != "Power Attack" {
		t.Fatalf("expected feat name %q, got %q", "Power Attack", feat.GetID().GetName())
	}

	if feat.GetCategory() != CombatFeatCategory {
		t.Fatalf("expected feat category %q, got %q", CombatFeatCategory, feat.GetCategory())
	}

	if feat.GetCategory().GetName() != "Combat" {
		t.Fatalf("expected category name %q, got %q", "Combat", feat.GetCategory().GetName())
	}

	if !feat.IsFighterBonusFeat() {
		t.Fatal("expected combat feat to be marked as fighter bonus feat")
	}

	if feat.IsMetamagic() {
		t.Fatal("expected combat feat not to be marked as metamagic")
	}

	if feat.IsItemCreation() {
		t.Fatal("expected combat feat not to be marked as item creation")
	}

	got := feat.GetPrerequisites()
	if len(got) != 1 || got[0].GetKind() != BaseAttackBonusPrerequisiteKind {
		t.Fatalf("expected base attack bonus prerequisite, got %v", got)
	}
}

func TestNewFeat_AcceptsCoreFeatCategoriesAndFlags(t *testing.T) {
	testCases := []struct {
		id               FeatID
		category         FeatCategory
		fighterBonusFeat bool
		metamagic        bool
		itemCreation     bool
	}{
		{FeatID("Acrobatic"), GeneralFeatCategory, false, false, false},
		{FeatID("Weapon Focus"), CombatFeatCategory, true, false, false},
		{FeatID("Staggering Critical"), CriticalFeatCategory, true, false, false},
		{FeatID("Craft Wand"), ItemCreationFeatCategory, false, false, true},
		{FeatID("Still Spell"), MetamagicFeatCategory, false, true, false},
	}

	prerequisites := mustNewPrerequisiteList(t, nil)

	for _, tc := range testCases {
		if _, ok := NewFeat(
			tc.id,
			tc.category,
			prerequisites,
			tc.fighterBonusFeat,
			tc.metamagic,
			tc.itemCreation,
		); !ok {
			t.Fatalf("expected feat %q category %q to be valid", tc.id, tc.category)
		}
	}
}

func TestNewFeat_RejectsInvalidInputs(t *testing.T) {
	validPrerequisites := mustNewPrerequisiteList(t, nil)

	if _, ok := NewFeat("", GeneralFeatCategory, validPrerequisites, false, false, false); ok {
		t.Fatal("expected empty feat id to be rejected")
	}

	if _, ok := NewFeat(FeatID(" Power Attack"), CombatFeatCategory, validPrerequisites, true, false, false); ok {
		t.Fatal("expected unnormalized feat id to be rejected")
	}

	if _, ok := NewFeat(FeatID("Power Attack"), FeatCategory("Teamwork"), validPrerequisites, false, false, false); ok {
		t.Fatal("expected unknown feat category to be rejected")
	}

	if _, ok := NewFeat(
		FeatID("Power Attack"),
		CombatFeatCategory,
		validPrerequisites,
		false,
		false,
		false,
	); ok {
		t.Fatal("expected combat feat without fighter bonus flag to be rejected")
	}

	if _, ok := NewFeat(
		FeatID("Acrobatic"),
		GeneralFeatCategory,
		validPrerequisites,
		true,
		false,
		false,
	); ok {
		t.Fatal("expected general feat with fighter bonus flag to be rejected")
	}

	if _, ok := NewFeat(
		FeatID("Still Spell"),
		MetamagicFeatCategory,
		validPrerequisites,
		false,
		false,
		false,
	); ok {
		t.Fatal("expected metamagic feat without metamagic flag to be rejected")
	}

	if _, ok := NewFeat(
		FeatID("Craft Wand"),
		ItemCreationFeatCategory,
		validPrerequisites,
		false,
		false,
		false,
	); ok {
		t.Fatal("expected item creation feat without item creation flag to be rejected")
	}

	if _, ok := NewFeat(
		FeatID("Craft Wand"),
		ItemCreationFeatCategory,
		validPrerequisites,
		false,
		true,
		true,
	); ok {
		t.Fatal("expected conflicting item creation and metamagic flags to be rejected")
	}

	invalidPrerequisites := prerequisiteList{prerequisites: []Prerequisite{nil}}
	if _, ok := NewFeat(
		FeatID("Acrobatic"),
		GeneralFeatCategory,
		invalidPrerequisites,
		false,
		false,
		false,
	); ok {
		t.Fatal("expected invalid prerequisite list to be rejected")
	}
}

func TestFeat_GetPrerequisitesReturnsDefensiveCopy(t *testing.T) {
	strengthPrerequisite := mustNewAbilityScorePrerequisite(t, ability.StrengthScore, 13)
	baseAttackBonusPrerequisite := mustNewBaseAttackBonusPrerequisite(t, 1)
	prerequisites := mustNewPrerequisiteList(t, []Prerequisite{strengthPrerequisite})

	feat, ok := NewFeat(
		FeatID("Power Attack"),
		CombatFeatCategory,
		prerequisites,
		true,
		false,
		false,
	)
	if !ok {
		t.Fatal("expected feat chassis to be constructed")
	}

	first := feat.GetPrerequisites()
	first[0] = baseAttackBonusPrerequisite

	second := feat.GetPrerequisites()
	if second[0].GetKind() != AbilityScorePrerequisiteKind {
		t.Fatal("expected feat prerequisites getter to return a defensive copy")
	}
}

func TestFeatIDAndCategoryNamesRejectInvalidValues(t *testing.T) {
	if FeatID(" Power Attack").GetName() != "" {
		t.Fatal("expected invalid feat id to return empty name")
	}

	if FeatCategory("Teamwork").GetName() != "" {
		t.Fatal("expected invalid feat category to return empty name")
	}
}

func mustNewPrerequisiteList(
	t *testing.T,
	prerequisites []Prerequisite,
) PrerequisiteList {
	t.Helper()

	value, ok := NewPrerequisiteList(prerequisites)
	if !ok {
		t.Fatal("expected prerequisite list to be valid")
	}

	return value
}
