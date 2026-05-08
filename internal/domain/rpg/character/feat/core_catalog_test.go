package feat

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/spell"
)

func TestGetFeatByID_ReturnsSeededCoreFeatAcrossCategories(t *testing.T) {
	testCases := []struct {
		id       FeatID
		category FeatCategory
	}{
		{AcrobaticFeatID, GeneralFeatCategory},
		{PowerAttackFeatID, CombatFeatCategory},
		{BleedingCriticalFeatID, CriticalFeatCategory},
		{BrewPotionFeatID, ItemCreationFeatCategory},
		{SilentSpellFeatID, MetamagicFeatCategory},
	}

	for _, tc := range testCases {
		feat, ok := GetFeatByID(tc.id)
		if !ok {
			t.Fatalf("expected core feat lookup for %q to succeed", tc.id)
		}

		if feat.GetID() != tc.id {
			t.Fatalf("expected feat id %q, got %q", tc.id, feat.GetID())
		}

		if feat.GetCategory() != tc.category {
			t.Fatalf("expected feat %q category %q, got %q", tc.id, tc.category, feat.GetCategory())
		}
	}
}

func TestGetFeatByID_ReturnsDetachedCopy(t *testing.T) {
	first, ok := GetFeatByID(PowerAttackFeatID)
	if !ok {
		t.Fatal("expected power attack lookup to succeed")
	}

	if len(first.prerequisites.prerequisites) == 0 {
		t.Fatal("expected power attack to have seeded prerequisites")
	}

	first.prerequisites.prerequisites[0] = mustFeatPrerequisite(WeaponFocusFeatID)

	second, ok := GetFeatByID(PowerAttackFeatID)
	if !ok {
		t.Fatal("expected power attack lookup to succeed")
	}

	prerequisite, ok := second.prerequisites.prerequisites[0].(AbilityScorePrerequisite)
	if !ok {
		t.Fatalf("expected stored power attack prerequisite to remain ability score, got %T", second.prerequisites.prerequisites[0])
	}

	if prerequisite.GetAbilityScoreID().GetName() != "Strength" || prerequisite.GetMinimumScore() != 13 {
		t.Fatalf("expected stored power attack prerequisite to remain Strength 13, got %q %d", prerequisite.GetAbilityScoreID(), prerequisite.GetMinimumScore())
	}

	first, ok = GetFeatByID(ImprovisedWeaponMasteryFeatID)
	if !ok {
		t.Fatal("expected improvised weapon mastery lookup to succeed")
	}

	anyFeatPrerequisite, ok := first.prerequisites.prerequisites[0].(AnyFeatPrerequisite)
	if !ok {
		t.Fatalf("expected improvised weapon mastery prerequisite to be any-feat, got %T", first.prerequisites.prerequisites[0])
	}

	anyFeatPrerequisite.featIDs[0] = WeaponFocusFeatID

	second, ok = GetFeatByID(ImprovisedWeaponMasteryFeatID)
	if !ok {
		t.Fatal("expected improvised weapon mastery lookup to succeed")
	}

	anyFeatPrerequisite, ok = second.prerequisites.prerequisites[0].(AnyFeatPrerequisite)
	if !ok {
		t.Fatalf("expected stored improvised weapon mastery prerequisite to be any-feat, got %T", second.prerequisites.prerequisites[0])
	}

	if anyFeatPrerequisite.featIDs[0] != CatchOffGuardFeatID {
		t.Fatalf("expected stored improvised weapon mastery prerequisite to remain Catch Off-Guard, got %q", anyFeatPrerequisite.featIDs[0])
	}
}

func TestGetFeatByID_RejectsUnknownFeat(t *testing.T) {
	if _, ok := GetFeatByID(FeatID("Mythic Power Attack")); ok {
		t.Fatal("expected unknown feat lookup to fail")
	}
}

func TestCoreFeatPrerequisiteReferences_SeededCoreFeatsAllResolve(t *testing.T) {
	feats := GetFeats()
	if !validateCoreFeatPrerequisiteReferences(feats, seededCoreFeatIDs(feats)) {
		t.Fatal("expected all seeded core feat prerequisites to reference seeded core feats")
	}
}

func TestCoreFeatPrerequisiteReferences_RejectsMissingReferencedFeat(t *testing.T) {
	missingID := FeatID("Missing Core Feat")

	testCases := []struct {
		name         string
		prerequisite Prerequisite
	}{
		{
			name:         "feat prerequisite",
			prerequisite: mustFeatPrerequisite(missingID),
		},
		{
			name:         "any feat prerequisite",
			prerequisite: mustAnyFeatPrerequisite([]FeatID{missingID}),
		},
		{
			name:         "same-selection feat prerequisite",
			prerequisite: mustSameSelectionFeatPrerequisite(missingID),
		},
		{
			name:         "spell-school feat prerequisite",
			prerequisite: mustSpellSchoolFeatPrerequisite(missingID, spell.ConjurationSchoolID),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			feat := mustCoreReferenceValidationFeat(t, tc.prerequisite)
			seededIDs := map[FeatID]struct{}{
				feat.GetID(): {},
			}

			if validateCoreFeatPrerequisiteReferences([]Feat{feat}, seededIDs) {
				t.Fatalf("expected %s with missing referenced feat to be rejected", tc.name)
			}
		})
	}
}

func TestGetFeats_ReturnsSeededCatalogInCoreOrder(t *testing.T) {
	feats := GetFeats()
	if len(feats) != len(coreFeatOrder) {
		t.Fatalf("expected %d queried feats, got %d", len(coreFeatOrder), len(feats))
	}

	seen := make(map[FeatID]struct{}, len(feats))
	for i, expectedID := range coreFeatOrder {
		if feats[i].GetID() != expectedID {
			t.Fatalf("expected feat at index %d to be %q, got %q", i, expectedID, feats[i].GetID())
		}

		if _, ok := seen[feats[i].GetID()]; ok {
			t.Fatalf("expected queried core feats not to duplicate %q", feats[i].GetID())
		}
		seen[feats[i].GetID()] = struct{}{}
	}
}

func TestGetFeats_ReturnsDetachedCopies(t *testing.T) {
	first := GetFeats()
	second := GetFeats()

	if len(first[0].prerequisites.prerequisites) != 0 {
		t.Fatal("expected acrobatic not to have seeded prerequisites")
	}

	powerAttackIndex := indexOfCoreFeat(t, PowerAttackFeatID)
	first[0] = feat{}
	first[powerAttackIndex].prerequisites.prerequisites[0] = mustFeatPrerequisite(WeaponFocusFeatID)

	if second[0].GetID() != AcrobaticFeatID {
		t.Fatalf("expected stored first feat to remain Acrobatic, got %q", second[0].GetID())
	}

	prerequisite, ok := second[powerAttackIndex].prerequisites.prerequisites[0].(AbilityScorePrerequisite)
	if !ok {
		t.Fatalf("expected stored power attack prerequisite to remain ability score, got %T", second[powerAttackIndex].prerequisites.prerequisites[0])
	}

	if prerequisite.GetAbilityScoreID().GetName() != "Strength" || prerequisite.GetMinimumScore() != 13 {
		t.Fatalf("expected stored power attack prerequisite to remain Strength 13, got %q %d", prerequisite.GetAbilityScoreID(), prerequisite.GetMinimumScore())
	}
}

func indexOfCoreFeat(t *testing.T, id FeatID) int {
	t.Helper()

	for i, coreID := range coreFeatOrder {
		if coreID == id {
			return i
		}
	}

	t.Fatalf("expected core feat order to include %q", id)
	return -1
}

func mustCoreReferenceValidationFeat(t *testing.T, prerequisite Prerequisite) Feat {
	t.Helper()

	prerequisites, ok := NewPrerequisiteList([]Prerequisite{prerequisite})
	if !ok {
		t.Fatal("expected prerequisite list to be valid")
	}

	feat, ok := NewFeat(
		FeatID("Reference Validation Test Feat"),
		GeneralFeatCategory,
		prerequisites,
		false,
		false,
		false,
	)
	if !ok {
		t.Fatal("expected reference validation test feat to be valid")
	}

	return feat
}
