package feat

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
	"d20campaigngenerator/internal/domain/rpg/character/spell"
)

func TestCoreGeneralFeats_SeedsFiftyFiveCoreEntries(t *testing.T) {
	const expectedCount = 55

	if len(coreGeneralFeats) != expectedCount {
		t.Fatalf("expected %d core general feats, got %d", expectedCount, len(coreGeneralFeats))
	}

	if len(coreGeneralFeatOrder) != expectedCount {
		t.Fatalf("expected %d ordered core general feats, got %d", expectedCount, len(coreGeneralFeatOrder))
	}

	seen := make(map[FeatID]struct{}, len(coreGeneralFeatOrder))
	for _, id := range coreGeneralFeatOrder {
		if _, ok := seen[id]; ok {
			t.Fatalf("expected core general feat order not to duplicate %q", id)
		}
		seen[id] = struct{}{}

		value, ok := coreGeneralFeats[id]
		if !ok {
			t.Fatalf("expected core general feat %q to be seeded", id)
		}

		if value.GetID() != id {
			t.Fatalf("expected core general feat id %q, got %q", id, value.GetID())
		}

		if value.GetCategory() != GeneralFeatCategory {
			t.Fatalf("expected core general feat %q category %q, got %q", id, GeneralFeatCategory, value.GetCategory())
		}

		if value.IsFighterBonusFeat() || value.IsItemCreation() || value.IsMetamagic() {
			t.Fatalf("expected core general feat %q not to carry non-general flags", id)
		}

		prerequisites, ok := NewPrerequisiteList(value.GetPrerequisites())
		if !ok {
			t.Fatalf("expected core general feat %q prerequisites to remain valid", id)
		}

		if _, ok := NewFeat(
			value.GetID(),
			value.GetCategory(),
			prerequisites,
			value.IsFighterBonusFeat(),
			value.IsMetamagic(),
			value.IsItemCreation(),
		); !ok {
			t.Fatalf("expected core general feat %q to reconstruct from seeded metadata", id)
		}
	}
}

func TestCoreGeneralFeats_DoNotSeedOtherCategories(t *testing.T) {
	nonGeneralFeatIDs := []FeatID{
		"Power Attack",
		"Bleeding Critical",
		"Brew Potion",
		"Silent Spell",
	}

	for _, id := range nonGeneralFeatIDs {
		if _, ok := coreGeneralFeats[id]; ok {
			t.Fatalf("expected non-general feat %q not to be in core general feat seeds", id)
		}
	}
}

func TestCoreGeneralFeats_SeedKnownCorePrerequisites(t *testing.T) {
	expectedPrerequisiteCounts := map[FeatID]int{
		AlignmentChannelFeatID:          1,
		ArmorProficiencyMediumFeatID:    1,
		ArmorProficiencyHeavyFeatID:     2,
		AugmentSummoningFeatID:          1,
		CommandUndeadFeatID:             1,
		ElementalChannelFeatID:          1,
		DiehardFeatID:                   1,
		ExtraChannelFeatID:              1,
		ExtraKiFeatID:                   1,
		ExtraLayOnHandsFeatID:           1,
		ExtraMercyFeatID:                2,
		ExtraPerformanceFeatID:          1,
		ExtraRageFeatID:                 1,
		ImprovedGreatFortitudeFeatID:    1,
		ImprovedChannelFeatID:           1,
		ImprovedFamiliarFeatID:          2,
		ImprovedIronWillFeatID:          1,
		LeadershipFeatID:                1,
		ImprovedLightningReflexesFeatID: 1,
		MasterCraftsmanFeatID:           1,
		NaturalSpellFeatID:              2,
		NimbleMovesFeatID:               1,
		AcrobaticStepsFeatID:            2,
		SelectiveChannelingFeatID:       2,
		GreaterSpellFocusFeatID:         1,
		SpellMasteryFeatID:              1,
		GreaterSpellPenetrationFeatID:   1,
		TurnUndeadFeatID:                1,
	}

	for _, id := range coreGeneralFeatOrder {
		value := mustCoreGeneralFeat(t, id)
		expected := expectedPrerequisiteCounts[id]
		if len(value.GetPrerequisites()) != expected {
			t.Fatalf("expected core general feat %q to have %d prerequisites, got %d", id, expected, len(value.GetPrerequisites()))
		}
	}

	assertHasClassFeaturePrerequisite(t, AlignmentChannelFeatID, characterclass.ChannelEnergyClassFeatureID)
	assertHasClassFeaturePrerequisite(t, CommandUndeadFeatID, characterclass.ChannelNegativeEnergyClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ElementalChannelFeatID, characterclass.ChannelEnergyClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ExtraChannelFeatID, characterclass.ChannelEnergyClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ExtraKiFeatID, characterclass.KiPoolClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ExtraLayOnHandsFeatID, characterclass.LayOnHandsClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ExtraMercyFeatID, characterclass.LayOnHandsClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ExtraMercyFeatID, characterclass.MercyClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ExtraPerformanceFeatID, characterclass.BardicPerformanceClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ExtraRageFeatID, characterclass.RageClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ImprovedChannelFeatID, characterclass.ChannelEnergyClassFeatureID)
	assertHasClassFeaturePrerequisite(t, ImprovedFamiliarFeatID, characterclass.FamiliarAccessClassFeatureID)
	assertHasClassFeaturePrerequisite(t, NaturalSpellFeatID, characterclass.WildShapeClassFeatureID)
	assertHasClassFeaturePrerequisite(t, SelectiveChannelingFeatID, characterclass.ChannelEnergyClassFeatureID)
	assertHasClassFeaturePrerequisite(t, TurnUndeadFeatID, characterclass.ChannelPositiveEnergyClassFeatureID)

	assertHasFeatPrerequisite(t, ArmorProficiencyMediumFeatID, ArmorProficiencyLightFeatID)
	assertHasFeatPrerequisite(t, ArmorProficiencyHeavyFeatID, ArmorProficiencyLightFeatID)
	assertHasFeatPrerequisite(t, ArmorProficiencyHeavyFeatID, ArmorProficiencyMediumFeatID)
	assertHasFeatPrerequisite(t, DiehardFeatID, EnduranceFeatID)
	assertHasFeatPrerequisite(t, ImprovedGreatFortitudeFeatID, GreatFortitudeFeatID)
	assertHasFeatPrerequisite(t, ImprovedIronWillFeatID, IronWillFeatID)
	assertHasFeatPrerequisite(t, ImprovedLightningReflexesFeatID, LightningReflexesFeatID)
	assertHasFeatPrerequisite(t, AcrobaticStepsFeatID, NimbleMovesFeatID)
	assertHasFeatPrerequisite(t, GreaterSpellPenetrationFeatID, SpellPenetrationFeatID)

	assertHasAbilityScorePrerequisite(t, NaturalSpellFeatID, ability.WisdomScore, 13)
	assertHasAbilityScorePrerequisite(t, NimbleMovesFeatID, ability.DexterityScore, 13)
	assertHasAbilityScorePrerequisite(t, AcrobaticStepsFeatID, ability.DexterityScore, 15)
	assertHasAbilityScorePrerequisite(t, SelectiveChannelingFeatID, ability.CharismaScore, 13)

	assertHasCharacterLevelPrerequisite(t, LeadershipFeatID, 7)
	assertHasClassLevelPrerequisite(t, SpellMasteryFeatID, characterclass.WizardClassID, 1)
	assertHasAnySkillRanksPrerequisite(t, MasterCraftsmanFeatID, []skill.SkillID{skill.CraftSkillID, skill.ProfessionSkillID}, 5)
	assertHasSelectedFamiliarEligibilityPrerequisite(t, ImprovedFamiliarFeatID)
	assertHasSameSelectionFeatPrerequisite(t, GreaterSpellFocusFeatID, SpellFocusFeatID)
	assertHasSpellSchoolFeatPrerequisite(t, AugmentSummoningFeatID, SpellFocusFeatID, spell.ConjurationSchoolID)
}

func mustCoreGeneralFeat(t *testing.T, id FeatID) Feat {
	t.Helper()

	value, ok := coreGeneralFeats[id]
	if !ok {
		t.Fatalf("expected core general feat %q to be seeded", id)
	}

	return value
}

func assertHasClassFeaturePrerequisite(
	t *testing.T,
	featID FeatID,
	featureID characterclass.ClassFeatureID,
) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(ClassFeaturePrerequisite)
		if ok && value.GetFeatureID() == featureID {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require class feature %q", featID, featureID)
}

func assertHasFeatPrerequisite(t *testing.T, featID FeatID, prerequisiteID FeatID) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(FeatPrerequisite)
		if ok && value.GetFeatID() == prerequisiteID {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require feat %q", featID, prerequisiteID)
}

func assertHasAbilityScorePrerequisite(
	t *testing.T,
	featID FeatID,
	abilityScoreID ability.AbilityScoreID,
	minimumScore int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(AbilityScorePrerequisite)
		if ok &&
			value.GetAbilityScoreID() == abilityScoreID &&
			value.GetMinimumScore() == minimumScore {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require %q %d", featID, abilityScoreID, minimumScore)
}

func assertHasCharacterLevelPrerequisite(t *testing.T, featID FeatID, minimumLevel int) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(CharacterLevelPrerequisite)
		if ok && value.GetMinimumLevel() == minimumLevel {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require character level %d", featID, minimumLevel)
}

func assertHasClassLevelPrerequisite(
	t *testing.T,
	featID FeatID,
	classID characterclass.ClassID,
	minimumLevel int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(ClassLevelPrerequisite)
		if ok &&
			value.GetClassID() == classID &&
			value.GetMinimumLevel() == minimumLevel {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require class %q level %d", featID, classID, minimumLevel)
}

func assertHasAnySkillRanksPrerequisite(
	t *testing.T,
	featID FeatID,
	skillIDs []skill.SkillID,
	minimumRanks int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(AnySkillRanksPrerequisite)
		if !ok || value.GetMinimumRanks() != minimumRanks {
			continue
		}

		actualSkillIDs := value.GetSkillIDs()
		if len(actualSkillIDs) != len(skillIDs) {
			continue
		}

		matches := true
		for i, skillID := range skillIDs {
			if actualSkillIDs[i] != skillID {
				matches = false
				break
			}
		}

		if matches {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require any of %v at %d ranks", featID, skillIDs, minimumRanks)
}

func assertHasSelectedFamiliarEligibilityPrerequisite(t *testing.T, featID FeatID) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		if _, ok := prerequisite.(SelectedFamiliarEligibilityPrerequisite); ok {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require selected familiar compatibility", featID)
}

func assertHasSameSelectionFeatPrerequisite(
	t *testing.T,
	featID FeatID,
	prerequisiteID FeatID,
) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(SameSelectionFeatPrerequisite)
		if ok && value.GetFeatID() == prerequisiteID {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require same-selection feat %q", featID, prerequisiteID)
}

func assertHasSpellSchoolFeatPrerequisite(
	t *testing.T,
	featID FeatID,
	prerequisiteID FeatID,
	schoolID spell.SchoolID,
) {
	t.Helper()

	for _, prerequisite := range mustCoreGeneralFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(SpellSchoolFeatPrerequisite)
		if ok &&
			value.GetFeatID() == prerequisiteID &&
			value.GetSchoolID() == schoolID {
			return
		}
	}

	t.Fatalf("expected core general feat %q to require feat %q for school %q", featID, prerequisiteID, schoolID)
}
