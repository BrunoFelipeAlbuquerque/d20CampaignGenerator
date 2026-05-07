package feat

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
	"d20campaigngenerator/internal/domain/rpg/character/spell"
)

func TestNewPrerequisiteList_AcceptsCorePrerequisiteShapes(t *testing.T) {
	strengthPrerequisite := mustNewAbilityScorePrerequisite(t, ability.StrengthScore, 13)
	baseAttackPrerequisite := mustNewBaseAttackBonusPrerequisite(t, 1)
	skillRanksPrerequisite := mustNewSkillRanksPrerequisite(t, skill.RideSkillID, 1)
	anySkillRanksPrerequisite := mustNewAnySkillRanksPrerequisite(
		t,
		[]skill.SkillID{skill.CraftSkillID, skill.ProfessionSkillID},
		5,
	)
	spellcastingPrerequisite := mustNewSpellcastingPrerequisite(t, ArcaneSpellcastingAccess)
	casterLevelPrerequisite := mustNewCasterLevelPrerequisite(t, 3)
	characterLevelPrerequisite := mustNewCharacterLevelPrerequisite(t, 7)
	classLevelPrerequisite := mustNewClassLevelPrerequisite(t, characterclass.FighterClassID, 4)
	classFeaturePrerequisite := mustNewClassFeaturePrerequisite(t, characterclass.ChannelEnergyClassFeatureID)
	weaponProficiencyPrerequisite := NewSelectedWeaponProficiencyPrerequisite()
	selectedFamiliarEligibilityPrerequisite := NewSelectedFamiliarEligibilityPrerequisite()
	featPrerequisite := mustNewFeatPrerequisite(t, FeatID("Power Attack"))
	sameSelectionFeatPrerequisite := mustNewSameSelectionFeatPrerequisite(t, FeatID("Weapon Focus"))
	spellSchoolFeatPrerequisite := mustNewSpellSchoolFeatPrerequisite(t, FeatID("Spell Focus"), spell.ConjurationSchoolID)

	prerequisites, ok := NewPrerequisiteList([]Prerequisite{
		strengthPrerequisite,
		baseAttackPrerequisite,
		skillRanksPrerequisite,
		anySkillRanksPrerequisite,
		spellcastingPrerequisite,
		casterLevelPrerequisite,
		characterLevelPrerequisite,
		classLevelPrerequisite,
		classFeaturePrerequisite,
		weaponProficiencyPrerequisite,
		selectedFamiliarEligibilityPrerequisite,
		featPrerequisite,
		sameSelectionFeatPrerequisite,
		spellSchoolFeatPrerequisite,
	})
	if !ok {
		t.Fatal("expected core prerequisite shapes to be accepted")
	}

	got := prerequisites.GetPrerequisites()
	if len(got) != 14 {
		t.Fatalf("expected 14 prerequisites, got %d", len(got))
	}

	if got[0].GetKind() != AbilityScorePrerequisiteKind {
		t.Fatalf("expected first prerequisite kind %q, got %q", AbilityScorePrerequisiteKind, got[0].GetKind())
	}

	if strengthPrerequisite.GetAbilityScoreID() != ability.StrengthScore ||
		strengthPrerequisite.GetMinimumScore() != 13 {
		t.Fatal("expected ability-score prerequisite to preserve score and minimum")
	}

	if classLevelPrerequisite.GetClassID() != characterclass.FighterClassID ||
		classLevelPrerequisite.GetMinimumLevel() != 4 {
		t.Fatal("expected class-level prerequisite to preserve class and minimum level")
	}

	anySkillIDs := anySkillRanksPrerequisite.GetSkillIDs()
	if len(anySkillIDs) != 2 ||
		anySkillIDs[0] != skill.CraftSkillID ||
		anySkillIDs[1] != skill.ProfessionSkillID ||
		anySkillRanksPrerequisite.GetMinimumRanks() != 5 {
		t.Fatal("expected any-skill ranks prerequisite to preserve skill choices and minimum ranks")
	}

	if spellSchoolFeatPrerequisite.GetFeatID() != FeatID("Spell Focus") ||
		spellSchoolFeatPrerequisite.GetSchoolID() != spell.ConjurationSchoolID {
		t.Fatal("expected spell-school feat prerequisite to preserve feat and school")
	}
}

func TestNewPrerequisiteList_RejectsNilAndZeroValuePrerequisites(t *testing.T) {
	if _, ok := NewPrerequisiteList([]Prerequisite{nil}); ok {
		t.Fatal("expected nil prerequisite to be rejected")
	}

	var zero AbilityScorePrerequisite
	if _, ok := NewPrerequisiteList([]Prerequisite{zero}); ok {
		t.Fatal("expected zero-value prerequisite to be rejected")
	}

	var zeroSelectedWeapon SelectedWeaponProficiencyPrerequisite
	if _, ok := NewPrerequisiteList([]Prerequisite{zeroSelectedWeapon}); ok {
		t.Fatal("expected zero-value selected weapon prerequisite to be rejected")
	}

	var zeroSelectedFamiliarEligibility SelectedFamiliarEligibilityPrerequisite
	if _, ok := NewPrerequisiteList([]Prerequisite{zeroSelectedFamiliarEligibility}); ok {
		t.Fatal("expected zero-value selected familiar eligibility prerequisite to be rejected")
	}
}

func TestNewPrerequisiteList_ReturnsDefensiveCopy(t *testing.T) {
	strengthPrerequisite := mustNewAbilityScorePrerequisite(t, ability.StrengthScore, 13)
	dexterityPrerequisite := mustNewAbilityScorePrerequisite(t, ability.DexterityScore, 15)

	prerequisites, ok := NewPrerequisiteList([]Prerequisite{strengthPrerequisite})
	if !ok {
		t.Fatal("expected prerequisite list to be created")
	}

	first := prerequisites.GetPrerequisites()
	first[0] = dexterityPrerequisite

	second := prerequisites.GetPrerequisites()
	if second[0].(AbilityScorePrerequisite).GetAbilityScoreID() != ability.StrengthScore {
		t.Fatal("expected prerequisite list getter to return a defensive copy")
	}

	anySkillRanksPrerequisite := mustNewAnySkillRanksPrerequisite(
		t,
		[]skill.SkillID{skill.CraftSkillID, skill.ProfessionSkillID},
		5,
	)
	skillIDs := anySkillRanksPrerequisite.GetSkillIDs()
	skillIDs[0] = skill.RideSkillID

	if anySkillRanksPrerequisite.GetSkillIDs()[0] != skill.CraftSkillID {
		t.Fatal("expected any-skill ranks prerequisite getter to return a defensive copy")
	}
}

func TestPrerequisiteConstructors_RejectInvalidInputs(t *testing.T) {
	if _, ok := NewAbilityScorePrerequisite(ability.AbilityScoreID("BAD"), 13); ok {
		t.Fatal("expected unknown ability score prerequisite to be rejected")
	}

	if _, ok := NewAbilityScorePrerequisite(ability.StrengthScore, 0); ok {
		t.Fatal("expected zero ability score prerequisite to be rejected")
	}

	if _, ok := NewBaseAttackBonusPrerequisite(0); ok {
		t.Fatal("expected zero base attack bonus prerequisite to be rejected")
	}

	if _, ok := NewSkillRanksPrerequisite(skill.SkillID("Ride "), 1); ok {
		t.Fatal("expected invalid skill prerequisite to be rejected")
	}

	if _, ok := NewSkillRanksPrerequisite(skill.RideSkillID, 0); ok {
		t.Fatal("expected zero skill ranks prerequisite to be rejected")
	}

	if _, ok := NewAnySkillRanksPrerequisite(nil, 5); ok {
		t.Fatal("expected empty any-skill ranks prerequisite to be rejected")
	}

	if _, ok := NewAnySkillRanksPrerequisite([]skill.SkillID{skill.SkillID("Jump")}, 5); ok {
		t.Fatal("expected unknown any-skill ranks prerequisite to be rejected")
	}

	if _, ok := NewAnySkillRanksPrerequisite([]skill.SkillID{skill.CraftSkillID}, 0); ok {
		t.Fatal("expected zero any-skill ranks prerequisite to be rejected")
	}

	if _, ok := NewAnySkillRanksPrerequisite([]skill.SkillID{skill.CraftSkillID, skill.CraftSkillID}, 5); ok {
		t.Fatal("expected duplicate any-skill ranks prerequisite to be rejected")
	}

	if _, ok := NewSpellcastingPrerequisite(SpellcastingAccess("Psionic")); ok {
		t.Fatal("expected unsupported spellcasting access prerequisite to be rejected")
	}

	if _, ok := NewCasterLevelPrerequisite(0); ok {
		t.Fatal("expected zero caster level prerequisite to be rejected")
	}

	if _, ok := NewCharacterLevelPrerequisite(0); ok {
		t.Fatal("expected zero character level prerequisite to be rejected")
	}

	if _, ok := NewClassLevelPrerequisite(characterclass.ClassID("adept"), 1); ok {
		t.Fatal("expected non-core class-level prerequisite to be rejected")
	}

	if _, ok := NewClassLevelPrerequisite(characterclass.FighterClassID, 0); ok {
		t.Fatal("expected zero class level prerequisite to be rejected")
	}

	if _, ok := NewClassFeaturePrerequisite(characterclass.ClassFeatureID("Sneak Attack")); ok {
		t.Fatal("expected unsupported class feature prerequisite to be rejected")
	}

	if _, ok := NewFeatPrerequisite(FeatID(" Power Attack")); ok {
		t.Fatal("expected unnormalized feat prerequisite to be rejected")
	}

	if _, ok := NewSameSelectionFeatPrerequisite(FeatID(" Weapon Focus")); ok {
		t.Fatal("expected unnormalized same-selection feat prerequisite to be rejected")
	}

	if _, ok := NewSpellSchoolFeatPrerequisite(FeatID("Spell Focus"), spell.SchoolID("Void")); ok {
		t.Fatal("expected unsupported spell-school feat prerequisite to be rejected")
	}
}

func mustNewAbilityScorePrerequisite(
	t *testing.T,
	id ability.AbilityScoreID,
	minimumScore int,
) AbilityScorePrerequisite {
	t.Helper()

	prerequisite, ok := NewAbilityScorePrerequisite(id, minimumScore)
	if !ok {
		t.Fatalf("expected ability score prerequisite %q %d to be valid", id, minimumScore)
	}

	return prerequisite
}

func mustNewBaseAttackBonusPrerequisite(t *testing.T, minimumBonus int) BaseAttackBonusPrerequisite {
	t.Helper()

	prerequisite, ok := NewBaseAttackBonusPrerequisite(minimumBonus)
	if !ok {
		t.Fatalf("expected base attack bonus prerequisite %d to be valid", minimumBonus)
	}

	return prerequisite
}

func mustNewSkillRanksPrerequisite(
	t *testing.T,
	id skill.SkillID,
	minimumRanks int,
) SkillRanksPrerequisite {
	t.Helper()

	prerequisite, ok := NewSkillRanksPrerequisite(id, minimumRanks)
	if !ok {
		t.Fatalf("expected skill ranks prerequisite %q %d to be valid", id, minimumRanks)
	}

	return prerequisite
}

func mustNewAnySkillRanksPrerequisite(
	t *testing.T,
	ids []skill.SkillID,
	minimumRanks int,
) AnySkillRanksPrerequisite {
	t.Helper()

	prerequisite, ok := NewAnySkillRanksPrerequisite(ids, minimumRanks)
	if !ok {
		t.Fatalf("expected any-skill ranks prerequisite %v %d to be valid", ids, minimumRanks)
	}

	return prerequisite
}

func mustNewSpellcastingPrerequisite(
	t *testing.T,
	access SpellcastingAccess,
) SpellcastingPrerequisite {
	t.Helper()

	prerequisite, ok := NewSpellcastingPrerequisite(access)
	if !ok {
		t.Fatalf("expected spellcasting prerequisite %q to be valid", access)
	}

	return prerequisite
}

func mustNewCasterLevelPrerequisite(t *testing.T, minimumLevel int) CasterLevelPrerequisite {
	t.Helper()

	prerequisite, ok := NewCasterLevelPrerequisite(minimumLevel)
	if !ok {
		t.Fatalf("expected caster level prerequisite %d to be valid", minimumLevel)
	}

	return prerequisite
}

func mustNewCharacterLevelPrerequisite(t *testing.T, minimumLevel int) CharacterLevelPrerequisite {
	t.Helper()

	prerequisite, ok := NewCharacterLevelPrerequisite(minimumLevel)
	if !ok {
		t.Fatalf("expected character level prerequisite %d to be valid", minimumLevel)
	}

	return prerequisite
}

func mustNewClassLevelPrerequisite(
	t *testing.T,
	id characterclass.ClassID,
	minimumLevel int,
) ClassLevelPrerequisite {
	t.Helper()

	prerequisite, ok := NewClassLevelPrerequisite(id, minimumLevel)
	if !ok {
		t.Fatalf("expected class level prerequisite %q %d to be valid", id, minimumLevel)
	}

	return prerequisite
}

func mustNewClassFeaturePrerequisite(
	t *testing.T,
	id characterclass.ClassFeatureID,
) ClassFeaturePrerequisite {
	t.Helper()

	prerequisite, ok := NewClassFeaturePrerequisite(id)
	if !ok {
		t.Fatalf("expected class feature prerequisite %q to be valid", id)
	}

	return prerequisite
}

func mustNewFeatPrerequisite(t *testing.T, id FeatID) FeatPrerequisite {
	t.Helper()

	prerequisite, ok := NewFeatPrerequisite(id)
	if !ok {
		t.Fatalf("expected feat prerequisite %q to be valid", id)
	}

	return prerequisite
}

func mustNewSameSelectionFeatPrerequisite(t *testing.T, id FeatID) SameSelectionFeatPrerequisite {
	t.Helper()

	prerequisite, ok := NewSameSelectionFeatPrerequisite(id)
	if !ok {
		t.Fatalf("expected same-selection feat prerequisite %q to be valid", id)
	}

	return prerequisite
}

func mustNewSpellSchoolFeatPrerequisite(
	t *testing.T,
	id FeatID,
	schoolID spell.SchoolID,
) SpellSchoolFeatPrerequisite {
	t.Helper()

	prerequisite, ok := NewSpellSchoolFeatPrerequisite(id, schoolID)
	if !ok {
		t.Fatalf("expected spell-school feat prerequisite %q %q to be valid", id, schoolID)
	}

	return prerequisite
}
