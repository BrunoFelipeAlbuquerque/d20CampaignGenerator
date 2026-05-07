package feat

import "testing"

func TestCoreCriticalFeats_SeedsEightCoreEntries(t *testing.T) {
	const expectedCount = 8

	if len(coreCriticalFeats) != expectedCount {
		t.Fatalf("expected %d core critical feats, got %d", expectedCount, len(coreCriticalFeats))
	}

	if len(coreCriticalFeatOrder) != expectedCount {
		t.Fatalf("expected %d ordered core critical feats, got %d", expectedCount, len(coreCriticalFeatOrder))
	}

	seen := make(map[FeatID]struct{}, len(coreCriticalFeatOrder))
	for _, id := range coreCriticalFeatOrder {
		if _, ok := seen[id]; ok {
			t.Fatalf("expected core critical feat order not to duplicate %q", id)
		}
		seen[id] = struct{}{}

		value, ok := coreCriticalFeats[id]
		if !ok {
			t.Fatalf("expected core critical feat %q to be seeded", id)
		}

		if value.GetID() != id {
			t.Fatalf("expected core critical feat id %q, got %q", id, value.GetID())
		}

		if value.GetCategory() != CriticalFeatCategory {
			t.Fatalf("expected core critical feat %q category %q, got %q", id, CriticalFeatCategory, value.GetCategory())
		}

		if !value.IsFighterBonusFeat() || value.IsItemCreation() || value.IsMetamagic() {
			t.Fatalf("expected core critical feat %q to carry only the fighter-bonus flag", id)
		}

		prerequisites, ok := NewPrerequisiteList(value.GetPrerequisites())
		if !ok {
			t.Fatalf("expected core critical feat %q prerequisites to remain valid", id)
		}

		if _, ok := NewFeat(
			value.GetID(),
			value.GetCategory(),
			prerequisites,
			value.IsFighterBonusFeat(),
			value.IsMetamagic(),
			value.IsItemCreation(),
		); !ok {
			t.Fatalf("expected core critical feat %q to reconstruct from seeded metadata", id)
		}
	}
}

func TestCoreCriticalFeats_DoNotSeedOtherCategories(t *testing.T) {
	nonCriticalFeatIDs := []FeatID{
		CriticalFocusFeatID,
		CriticalMasteryFeatID,
		PowerAttackFeatID,
		FeatID("Brew Potion"),
		FeatID("Silent Spell"),
	}

	for _, id := range nonCriticalFeatIDs {
		if _, ok := coreCriticalFeats[id]; ok {
			t.Fatalf("expected non-critical feat %q not to be in core critical feat seeds", id)
		}
	}
}

func TestCoreCriticalFeats_SeedKnownCorePrerequisites(t *testing.T) {
	expectedBaseAttackBonus := map[FeatID]int{
		BleedingCriticalFeatID:   11,
		BlindingCriticalFeatID:   15,
		DeafeningCriticalFeatID:  13,
		ExhaustingCriticalFeatID: 15,
		SickeningCriticalFeatID:  11,
		StaggeringCriticalFeatID: 13,
		StunningCriticalFeatID:   17,
		TiringCriticalFeatID:     13,
	}

	for _, id := range coreCriticalFeatOrder {
		assertCriticalHasFeatPrerequisite(t, id, CriticalFocusFeatID)
		assertCriticalHasBaseAttackBonusPrerequisite(t, id, expectedBaseAttackBonus[id])
	}

	assertCriticalHasFeatPrerequisite(t, ExhaustingCriticalFeatID, TiringCriticalFeatID)
	assertCriticalHasFeatPrerequisite(t, StunningCriticalFeatID, StaggeringCriticalFeatID)
}

func mustCoreCriticalFeat(t *testing.T, id FeatID) Feat {
	t.Helper()

	value, ok := coreCriticalFeats[id]
	if !ok {
		t.Fatalf("expected core critical feat %q to be seeded", id)
	}

	return value
}

func assertCriticalHasFeatPrerequisite(t *testing.T, featID FeatID, prerequisiteID FeatID) {
	t.Helper()

	for _, prerequisite := range mustCoreCriticalFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(FeatPrerequisite)
		if ok && value.GetFeatID() == prerequisiteID {
			return
		}
	}

	t.Fatalf("expected core critical feat %q to require feat %q", featID, prerequisiteID)
}

func assertCriticalHasBaseAttackBonusPrerequisite(t *testing.T, featID FeatID, minimumBonus int) {
	t.Helper()

	for _, prerequisite := range mustCoreCriticalFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(BaseAttackBonusPrerequisite)
		if ok && value.GetMinimumBonus() == minimumBonus {
			return
		}
	}

	t.Fatalf("expected core critical feat %q to require base attack bonus %d", featID, minimumBonus)
}
