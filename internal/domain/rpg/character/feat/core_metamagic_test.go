package feat

import "testing"

func TestCoreMetamagicFeats_SeedsNineCoreEntries(t *testing.T) {
	const expectedCount = 9

	if len(coreMetamagicFeats) != expectedCount {
		t.Fatalf("expected %d core metamagic feats, got %d", expectedCount, len(coreMetamagicFeats))
	}

	if len(coreMetamagicFeatOrder) != expectedCount {
		t.Fatalf("expected %d ordered core metamagic feats, got %d", expectedCount, len(coreMetamagicFeatOrder))
	}

	seen := make(map[FeatID]struct{}, len(coreMetamagicFeatOrder))
	for _, id := range coreMetamagicFeatOrder {
		if _, ok := seen[id]; ok {
			t.Fatalf("expected core metamagic feat order not to duplicate %q", id)
		}
		seen[id] = struct{}{}

		value, ok := coreMetamagicFeats[id]
		if !ok {
			t.Fatalf("expected core metamagic feat %q to be seeded", id)
		}

		if value.GetID() != id {
			t.Fatalf("expected core metamagic feat id %q, got %q", id, value.GetID())
		}

		if value.GetCategory() != MetamagicFeatCategory {
			t.Fatalf("expected core metamagic feat %q category %q, got %q", id, MetamagicFeatCategory, value.GetCategory())
		}

		if value.IsFighterBonusFeat() || value.IsItemCreation() || !value.IsMetamagic() {
			t.Fatalf("expected core metamagic feat %q to carry only the metamagic flag", id)
		}

		prerequisites, ok := NewPrerequisiteList(value.GetPrerequisites())
		if !ok {
			t.Fatalf("expected core metamagic feat %q prerequisites to remain valid", id)
		}

		if _, ok := NewFeat(
			value.GetID(),
			value.GetCategory(),
			prerequisites,
			value.IsFighterBonusFeat(),
			value.IsMetamagic(),
			value.IsItemCreation(),
		); !ok {
			t.Fatalf("expected core metamagic feat %q to reconstruct from seeded metadata", id)
		}
	}
}

func TestCoreMetamagicFeats_DoNotSeedOtherCategories(t *testing.T) {
	nonMetamagicFeatIDs := []FeatID{
		AcrobaticFeatID,
		PowerAttackFeatID,
		BleedingCriticalFeatID,
		BrewPotionFeatID,
	}

	for _, id := range nonMetamagicFeatIDs {
		if _, ok := coreMetamagicFeats[id]; ok {
			t.Fatalf("expected non-metamagic feat %q not to be in core metamagic feat seeds", id)
		}
	}
}

func TestCoreMetamagicFeats_SeedNoPrerequisites(t *testing.T) {
	for _, id := range coreMetamagicFeatOrder {
		if prerequisites := mustCoreMetamagicFeat(t, id).GetPrerequisites(); len(prerequisites) != 0 {
			t.Fatalf("expected core metamagic feat %q to have no prerequisites, got %d", id, len(prerequisites))
		}
	}
}

func mustCoreMetamagicFeat(t *testing.T, id FeatID) Feat {
	t.Helper()

	value, ok := coreMetamagicFeats[id]
	if !ok {
		t.Fatalf("expected core metamagic feat %q to be seeded", id)
	}

	return value
}
