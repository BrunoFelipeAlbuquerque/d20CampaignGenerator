package feat

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

func TestCoreCombatFeats_SeedsNinetyFiveCoreEntries(t *testing.T) {
	const expectedCount = 95

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
		ShieldProficiencyFeatID,
		FeatID("Staggering Critical"),
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
	assertCombatHasAbilityScorePrerequisite(t, PowerAttackFeatID, ability.StrengthScore, 13)
	assertCombatHasBaseAttackBonusPrerequisite(t, PowerAttackFeatID, 1)

	assertCombatHasFeatPrerequisite(t, ArcaneArmorTrainingFeatID, ArmorProficiencyLightFeatID)
	assertCombatHasCasterLevelPrerequisite(t, ArcaneArmorTrainingFeatID, 3)
	assertCombatHasFeatPrerequisite(t, ArcaneArmorMasteryFeatID, ArcaneArmorTrainingFeatID)
	assertCombatHasFeatPrerequisite(t, ArcaneArmorMasteryFeatID, ArmorProficiencyMediumFeatID)
	assertCombatHasCasterLevelPrerequisite(t, ArcaneArmorMasteryFeatID, 7)
	assertCombatHasSpellcastingPrerequisite(t, ArcaneStrikeFeatID, ArcaneSpellcastingAccess)

	assertCombatHasClassFeaturePrerequisite(t, ChannelSmiteFeatID, characterclass.ChannelEnergyClassFeatureID)
	assertCombatHasClassLevelPrerequisite(t, DisruptiveFeatID, characterclass.FighterClassID, 6)
	assertCombatHasFeatPrerequisite(t, SpellbreakerFeatID, DisruptiveFeatID)
	assertCombatHasClassLevelPrerequisite(t, SpellbreakerFeatID, characterclass.FighterClassID, 10)

	assertCombatHasSkillRanksPrerequisite(t, MountedCombatFeatID, skill.RideSkillID, 1)
	assertCombatHasSkillRanksPrerequisite(t, MountedArcheryFeatID, skill.RideSkillID, 1)
	assertCombatHasFeatPrerequisite(t, MountedArcheryFeatID, MountedCombatFeatID)
	assertCombatHasAnyFeatPrerequisite(t, ImprovisedWeaponMasteryFeatID, []FeatID{CatchOffGuardFeatID, ThrowAnythingFeatID})

	assertCombatHasSelectedWeaponProficiencyPrerequisite(t, WeaponFocusFeatID)
	assertCombatHasBaseAttackBonusPrerequisite(t, WeaponFocusFeatID, 1)
	assertCombatHasSameSelectionFeatPrerequisite(t, GreaterWeaponFocusFeatID, WeaponFocusFeatID)
	assertCombatHasClassLevelPrerequisite(t, GreaterWeaponFocusFeatID, characterclass.FighterClassID, 8)
	assertCombatHasSameSelectionFeatPrerequisite(t, GreaterWeaponSpecializationFeatID, WeaponSpecializationFeatID)
	assertCombatHasClassLevelPrerequisite(t, GreaterWeaponSpecializationFeatID, characterclass.FighterClassID, 12)

	assertCombatHasFeatPrerequisite(t, ShieldSlamFeatID, ImprovedShieldBashFeatID)
	assertCombatHasFeatPrerequisite(t, ShieldSlamFeatID, ShieldProficiencyFeatID)
	assertCombatHasBaseAttackBonusPrerequisite(t, ShieldSlamFeatID, 6)

	assertCombatHasAbilityScorePrerequisite(t, WhirlwindAttackFeatID, ability.DexterityScore, 13)
	assertCombatHasAbilityScorePrerequisite(t, WhirlwindAttackFeatID, ability.IntelligenceScore, 13)
	assertCombatHasFeatPrerequisite(t, WhirlwindAttackFeatID, SpringAttackFeatID)
	assertCombatHasBaseAttackBonusPrerequisite(t, WhirlwindAttackFeatID, 4)

	assertCombatHasBaseAttackBonusPrerequisite(t, CriticalFocusFeatID, 9)
}

func mustCoreCombatFeat(t *testing.T, id FeatID) Feat {
	t.Helper()

	value, ok := coreCombatFeats[id]
	if !ok {
		t.Fatalf("expected core combat feat %q to be seeded", id)
	}

	return value
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

func assertCombatHasBaseAttackBonusPrerequisite(
	t *testing.T,
	featID FeatID,
	minimumBonus int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(BaseAttackBonusPrerequisite)
		if ok && value.GetMinimumBonus() == minimumBonus {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require BAB %d", featID, minimumBonus)
}

func assertCombatHasCasterLevelPrerequisite(
	t *testing.T,
	featID FeatID,
	minimumLevel int,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(CasterLevelPrerequisite)
		if ok && value.GetMinimumLevel() == minimumLevel {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require caster level %d", featID, minimumLevel)
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

func assertCombatHasClassFeaturePrerequisite(
	t *testing.T,
	featID FeatID,
	featureID characterclass.ClassFeatureID,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(ClassFeaturePrerequisite)
		if ok && value.GetFeatureID() == featureID {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require class feature %q", featID, featureID)
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

	t.Fatalf("expected core combat feat %q to require skill %q at %d ranks", featID, skillID, minimumRanks)
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

func assertCombatHasAnyFeatPrerequisite(
	t *testing.T,
	featID FeatID,
	featIDs []FeatID,
) {
	t.Helper()

	for _, prerequisite := range mustCoreCombatFeat(t, featID).GetPrerequisites() {
		value, ok := prerequisite.(AnyFeatPrerequisite)
		if !ok {
			continue
		}

		actualFeatIDs := value.GetFeatIDs()
		if len(actualFeatIDs) != len(featIDs) {
			continue
		}

		matches := true
		for i, expected := range featIDs {
			if actualFeatIDs[i] != expected {
				matches = false
				break
			}
		}

		if matches {
			return
		}
	}

	t.Fatalf("expected core combat feat %q to require any of %v", featID, featIDs)
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
