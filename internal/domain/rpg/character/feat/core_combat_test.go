package feat

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

func TestCoreCombatFeats_SeedsNinetySixCoreEntries(t *testing.T) {
	const expectedCount = 96

	if len(coreCombatFeats) != expectedCount {
		t.Fatalf("expected %d core combat feats, got %d", expectedCount, len(coreCombatFeats))
	}

	if len(coreCombatFeatOrder) != expectedCount {
		t.Fatalf("expected %d ordered core combat feats, got %d", expectedCount, len(coreCombatFeatOrder))
	}

	seen := make(map[FeatID]struct{}, len(coreCombatFeatOrder))
	for _, id := range coreCombatFeatOrder {
		if _, ok := seen[id]; ok {
			t.Fatalf("expected core combat feat order not to duplicate %q", id)
		}
		seen[id] = struct{}{}

		value, ok := coreCombatFeats[id]
		if !ok {
			t.Fatalf("expected core combat feat %q to be seeded", id)
		}

		if value.GetID() != id {
			t.Fatalf("expected core combat feat id %q, got %q", id, value.GetID())
		}

		if value.GetCategory() != CombatFeatCategory {
			t.Fatalf("expected core combat feat %q category %q, got %q", id, CombatFeatCategory, value.GetCategory())
		}

		if !value.IsFighterBonusFeat() || value.IsItemCreation() || value.IsMetamagic() {
			t.Fatalf("expected core combat feat %q to carry only the fighter-bonus flag", id)
		}

		prerequisites, ok := NewPrerequisiteList(value.GetPrerequisites())
		if !ok {
			t.Fatalf("expected core combat feat %q prerequisites to remain valid", id)
		}

		if _, ok := NewFeat(
			value.GetID(),
			value.GetCategory(),
			prerequisites,
			value.IsFighterBonusFeat(),
			value.IsMetamagic(),
			value.IsItemCreation(),
		); !ok {
			t.Fatalf("expected core combat feat %q to reconstruct from seeded metadata", id)
		}
	}
}

func TestCoreCombatFeats_DoNotSeedOtherCategories(t *testing.T) {
	nonCombatFeatIDs := []FeatID{
		AcrobaticFeatID,
		FeatID("Bleeding Critical"),
		FeatID("Brew Potion"),
		FeatID("Silent Spell"),
	}

	for _, id := range nonCombatFeatIDs {
		if _, ok := coreCombatFeats[id]; ok {
			t.Fatalf("expected non-combat feat %q not to be in core combat feat seeds", id)
		}
	}
}

func TestCoreCombatFeats_SeedKnownCorePrerequisites(t *testing.T) {
	assertCombatHasFeatPrerequisite(t, ArcaneArmorMasteryFeatID, ArcaneArmorTrainingFeatID)
	assertCombatHasFeatPrerequisite(t, ArcaneArmorMasteryFeatID, ArmorProficiencyMediumFeatID)
	assertCombatHasCasterLevelPrerequisite(t, ArcaneArmorMasteryFeatID, 7)
	assertCombatHasSpellcastingPrerequisite(t, ArcaneStrikeFeatID, ArcaneSpellcastingAccess)

	assertCombatHasAbilityScorePrerequisite(t, WhirlwindAttackFeatID, ability.DexterityScore, 13)
	assertCombatHasAbilityScorePrerequisite(t, WhirlwindAttackFeatID, ability.IntelligenceScore, 13)
	assertCombatHasFeatPrerequisite(t, WhirlwindAttackFeatID, CombatExpertiseFeatID)
	assertCombatHasFeatPrerequisite(t, WhirlwindAttackFeatID, DodgeFeatID)
	assertCombatHasFeatPrerequisite(t, WhirlwindAttackFeatID, MobilityFeatID)
	assertCombatHasFeatPrerequisite(t, WhirlwindAttackFeatID, SpringAttackFeatID)
	assertCombatHasBaseAttackBonusPrerequisite(t, WhirlwindAttackFeatID, 4)

	assertCombatHasFeatPrerequisite(t, CriticalMasteryFeatID, CriticalFocusFeatID)
	assertCombatHasFeatCategoryCountPrerequisite(t, CriticalMasteryFeatID, CriticalFeatCategory, 2)
	assertCombatHasClassLevelPrerequisite(t, CriticalMasteryFeatID, characterclass.FighterClassID, 14)

	assertCombatHasAnyFeatPrerequisite(t, ImprovisedWeaponMasteryFeatID, []FeatID{CatchOffGuardFeatID, ThrowAnythingFeatID})
	assertCombatHasBaseAttackBonusPrerequisite(t, ImprovisedWeaponMasteryFeatID, 8)

	assertCombatHasSkillRanksPrerequisite(t, MountedCombatFeatID, skill.RideSkillID, 1)
	assertCombatHasAbilityScorePrerequisite(t, UnseatFeatID, ability.StrengthScore, 13)
	assertCombatHasSkillRanksPrerequisite(t, UnseatFeatID, skill.RideSkillID, 1)
	assertCombatHasFeatPrerequisite(t, UnseatFeatID, MountedCombatFeatID)
	assertCombatHasFeatPrerequisite(t, UnseatFeatID, PowerAttackFeatID)
	assertCombatHasFeatPrerequisite(t, UnseatFeatID, ImprovedBullRushFeatID)
	assertCombatHasBaseAttackBonusPrerequisite(t, UnseatFeatID, 1)

	assertCombatHasSelectedWeaponProficiencyPrerequisite(t, WeaponFocusFeatID)
	assertCombatHasBaseAttackBonusPrerequisite(t, WeaponFocusFeatID, 1)
	assertCombatHasSameSelectionFeatPrerequisite(t, GreaterWeaponSpecializationFeatID, GreaterWeaponFocusFeatID)
	assertCombatHasSameSelectionFeatPrerequisite(t, GreaterWeaponSpecializationFeatID, WeaponFocusFeatID)
	assertCombatHasSameSelectionFeatPrerequisite(t, GreaterWeaponSpecializationFeatID, WeaponSpecializationFeatID)
	assertCombatHasClassLevelPrerequisite(t, GreaterWeaponSpecializationFeatID, characterclass.FighterClassID, 12)
}

func mustCoreCombatFeat(t *testing.T, id FeatID) Feat {
	t.Helper()

	value, ok := coreCombatFeats[id]
	if !ok {
		t.Fatalf("expected core combat feat %q to be seeded", id)
	}

	return value
}

func assertCombatHasFeatPrerequisite(t *testing.T, featID FeatID, prerequisiteID FeatID) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(FeatPrerequisite)
		if ok && value.GetFeatID() == prerequisiteID {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require feat %q", featID, prerequisiteID)
}

func assertCombatHasAnyFeatPrerequisite(t *testing.T, featID FeatID, prerequisiteIDs []FeatID) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(AnyFeatPrerequisite)
		if !ok {
			continue
		}

		actualIDs := value.GetFeatIDs()
		if len(actualIDs) != len(prerequisiteIDs) {
			continue
		}

		matches := true
		for i, prerequisiteID := range prerequisiteIDs {
			if actualIDs[i] != prerequisiteID {
				matches = false
				break
			}
		}

		if matches {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require any feat from %v", featID, prerequisiteIDs)
}

func assertCombatHasFeatCategoryCountPrerequisite(
	t *testing.T,
	featID FeatID,
	category FeatCategory,
	minimumCount int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(FeatCategoryCountPrerequisite)
		if ok &&
			value.GetCategory() == category &&
			value.GetMinimumCount() == minimumCount {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require %d feats from category %q", featID, minimumCount, category)
}

func assertCombatHasSameSelectionFeatPrerequisite(
	t *testing.T,
	featID FeatID,
	prerequisiteID FeatID,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(SameSelectionFeatPrerequisite)
		if ok && value.GetFeatID() == prerequisiteID {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require same-selection feat %q", featID, prerequisiteID)
}

func assertCombatHasAbilityScorePrerequisite(
	t *testing.T,
	featID FeatID,
	abilityScoreID ability.AbilityScoreID,
	minimumScore int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(AbilityScorePrerequisite)
		if ok &&
			value.GetAbilityScoreID() == abilityScoreID &&
			value.GetMinimumScore() == minimumScore {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require %q %d", featID, abilityScoreID, minimumScore)
}

func assertCombatHasBaseAttackBonusPrerequisite(t *testing.T, featID FeatID, minimumBonus int) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(BaseAttackBonusPrerequisite)
		if ok && value.GetMinimumBonus() == minimumBonus {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require base attack bonus %d", featID, minimumBonus)
}

func assertCombatHasSkillRanksPrerequisite(
	t *testing.T,
	featID FeatID,
	skillID skill.SkillID,
	minimumRanks int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(SkillRanksPrerequisite)
		if ok &&
			value.GetSkillID() == skillID &&
			value.GetMinimumRanks() == minimumRanks {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require %q %d ranks", featID, skillID, minimumRanks)
}

func assertCombatHasSpellcastingPrerequisite(
	t *testing.T,
	featID FeatID,
	access SpellcastingAccess,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(SpellcastingPrerequisite)
		if ok && value.GetAccess() == access {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require spellcasting access %q", featID, access)
}

func assertCombatHasCasterLevelPrerequisite(t *testing.T, featID FeatID, minimumLevel int) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(CasterLevelPrerequisite)
		if ok && value.GetMinimumLevel() == minimumLevel {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require caster level %d", featID, minimumLevel)
}

func assertCombatHasClassLevelPrerequisite(
	t *testing.T,
	featID FeatID,
	classID characterclass.ClassID,
	minimumLevel int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(ClassLevelPrerequisite)
		if ok &&
			value.GetClassID() == classID &&
			value.GetMinimumLevel() == minimumLevel {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require class %q level %d", featID, classID, minimumLevel)
}

func assertCombatHasSelectedWeaponProficiencyPrerequisite(t *testing.T, featID FeatID) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		if _, ok := prerequisite.(SelectedWeaponProficiencyPrerequisite); ok {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require selected weapon proficiency", featID)
}
