package feat

import "testing"

func TestCoreItemCreationFeats_SeedsEightCoreEntries(t *testing.T) {
	const expectedCount = 8

	if len(coreItemCreationFeats) != expectedCount {
		t.Fatalf("expected %d core item creation feats, got %d", expectedCount, len(coreItemCreationFeats))
	}

	if len(coreItemCreationFeatOrder) != expectedCount {
		t.Fatalf("expected %d ordered core item creation feats, got %d", expectedCount, len(coreItemCreationFeatOrder))
	}

	seen := make(map[FeatID]struct{}, len(coreItemCreationFeatOrder))
	for _, id := range coreItemCreationFeatOrder {
		if _, ok := seen[id]; ok {
			t.Fatalf("expected core item creation feat order not to duplicate %q", id)
		}
		seen[id] = struct{}{}

		value, ok := coreItemCreationFeats[id]
		if !ok {
			t.Fatalf("expected core item creation feat %q to be seeded", id)
		}

		if value.GetID() != id {
			t.Fatalf("expected core item creation feat id %q, got %q", id, value.GetID())
		}

		if value.GetCategory() != ItemCreationFeatCategory {
			t.Fatalf("expected core item creation feat %q category %q, got %q", id, ItemCreationFeatCategory, value.GetCategory())
		}

		if value.IsFighterBonusFeat() || !value.IsItemCreation() || value.IsMetamagic() {
			t.Fatalf("expected core item creation feat %q to carry only the item-creation flag", id)
		}

		prerequisites, ok := NewPrerequisiteList(value.GetPrerequisites())
		if !ok {
			t.Fatalf("expected core item creation feat %q prerequisites to remain valid", id)
		}

		if _, ok := NewFeat(
			value.GetID(),
			value.GetCategory(),
			prerequisites,
			value.IsFighterBonusFeat(),
			value.IsMetamagic(),
			value.IsItemCreation(),
		); !ok {
			t.Fatalf("expected core item creation feat %q to reconstruct from seeded metadata", id)
		}
	}
}

func TestCoreItemCreationFeats_DoNotSeedOtherCategories(t *testing.T) {
	nonItemCreationFeatIDs := []FeatID{
		AcrobaticFeatID,
		PowerAttackFeatID,
		BleedingCriticalFeatID,
		FeatID("Silent Spell"),
	}

	for _, id := range nonItemCreationFeatIDs {
		if _, ok := coreItemCreationFeats[id]; ok {
			t.Fatalf("expected non-item creation feat %q not to be in core item creation feat seeds", id)
		}
	}
}

func TestCoreItemCreationFeats_SeedKnownCorePrerequisites(t *testing.T) {
	expectedCasterLevels := map[FeatID]int{
		BrewPotionFeatID:             3,
		CraftMagicArmsAndArmorFeatID: 5,
		CraftRodFeatID:               9,
		CraftStaffFeatID:             11,
		CraftWandFeatID:              5,
		CraftWondrousItemFeatID:      3,
		ForgeRingFeatID:              7,
		ScribeScrollFeatID:           1,
	}

	for _, id := range coreItemCreationFeatOrder {
		assertItemCreationHasCasterLevelPrerequisite(t, id, expectedCasterLevels[id])
	}
}

func mustCoreItemCreationFeat(t *testing.T, id FeatID) Feat {
	t.Helper()

	value, ok := coreItemCreationFeats[id]
	if !ok {
		t.Fatalf("expected core item creation feat %q to be seeded", id)
	}

	return value
}

func assertItemCreationHasCasterLevelPrerequisite(t *testing.T, featID FeatID, minimumLevel int) {
	t.Helper()

	for _, prerequisite := range mustCoreItemCreationFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(CasterLevelPrerequisite)
		if ok && value.GetMinimumLevel() == minimumLevel {
			return
		}
	}

	t.Fatalf("expected core item creation feat %q to require caster level %d", featID, minimumLevel)
}
