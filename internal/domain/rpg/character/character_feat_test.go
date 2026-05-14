package character

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/ability"
	characterclass "d20campaigngenerator/internal/domain/rpg/character/class"
	characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"
	characterfeat "d20campaigngenerator/internal/domain/rpg/character/feat"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

func TestNewCharacterFeat_ComposesAbilityScoreAndBaseAttackBonusPrerequisites(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		[]CharacterAbilityScore{mustNewCharacterAbilityScoreForTest(t, ability.StrengthScore, 13)},
		1,
		nil,
		nil,
		nil,
		nil,
	)

	selectedFeat, ok := NewCharacterFeat(characterfeat.PowerAttackFeatID, state)
	if !ok {
		t.Fatal("expected power attack prerequisites to compose")
	}

	if selectedFeat.GetFeatID() != characterfeat.PowerAttackFeatID {
		t.Fatalf("expected selected feat id %q, got %q", characterfeat.PowerAttackFeatID, selectedFeat.GetFeatID())
	}

	feat, ok := selectedFeat.GetFeat()
	if !ok {
		t.Fatal("expected selected feat to resolve")
	}

	if feat.GetCategory() != characterfeat.CombatFeatCategory {
		t.Fatalf("expected selected feat category %q, got %q", characterfeat.CombatFeatCategory, feat.GetCategory())
	}

	lowStrengthState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		[]CharacterAbilityScore{mustNewCharacterAbilityScoreForTest(t, ability.StrengthScore, 12)},
		1,
		nil,
		nil,
		nil,
		nil,
	)
	if _, ok := NewCharacterFeat(characterfeat.PowerAttackFeatID, lowStrengthState); ok {
		t.Fatal("expected power attack to reject a low strength score")
	}

	lowBaseAttackState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		[]CharacterAbilityScore{mustNewCharacterAbilityScoreForTest(t, ability.StrengthScore, 13)},
		0,
		nil,
		nil,
		nil,
		nil,
	)
	if _, ok := NewCharacterFeat(characterfeat.PowerAttackFeatID, lowBaseAttackState); ok {
		t.Fatal("expected power attack to reject a low base attack bonus")
	}
}

func TestNewCharacterFeat_ComposesSkillRanksAndFeatPrerequisites(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		nil,
		nil,
		[]CharacterSkillRanks{mustNewCharacterSkillRanksForTest(t, skill.RideSkillID, 1)},
		[]characterfeat.FeatID{characterfeat.MountedCombatFeatID},
	)

	if _, ok := NewCharacterFeat(characterfeat.MountedArcheryFeatID, state); !ok {
		t.Fatal("expected mounted archery prerequisites to compose")
	}

	missingFeatState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		nil,
		nil,
		[]CharacterSkillRanks{mustNewCharacterSkillRanksForTest(t, skill.RideSkillID, 1)},
		nil,
	)
	if _, ok := NewCharacterFeat(characterfeat.MountedArcheryFeatID, missingFeatState); ok {
		t.Fatal("expected mounted archery to reject missing mounted combat")
	}

	missingSkillState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		nil,
		nil,
		nil,
		[]characterfeat.FeatID{characterfeat.MountedCombatFeatID},
	)
	if _, ok := NewCharacterFeat(characterfeat.MountedArcheryFeatID, missingSkillState); ok {
		t.Fatal("expected mounted archery to reject missing ride ranks")
	}
}

func TestNewCharacterFeat_ComposesClassFeaturePrerequisites(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		nil,
		[]characterclass.ClassFeatureID{characterclass.RageClassFeatureID},
		nil,
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.ExtraRageFeatID, state); !ok {
		t.Fatal("expected extra rage prerequisites to compose")
	}

	missingFeatureState := mustNewCharacterFeatPrerequisiteStateForTest(t, nil, 0, nil, nil, nil, nil)
	if _, ok := NewCharacterFeat(characterfeat.ExtraRageFeatID, missingFeatureState); ok {
		t.Fatal("expected extra rage to reject missing rage feature")
	}
}

func TestNewCharacterFeat_ComposesClassLevelAndAnyFeatPrerequisites(t *testing.T) {
	anyFeatState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		8,
		nil,
		nil,
		nil,
		[]characterfeat.FeatID{characterfeat.CatchOffGuardFeatID},
	)

	if _, ok := NewCharacterFeat(characterfeat.ImprovisedWeaponMasteryFeatID, anyFeatState); !ok {
		t.Fatal("expected any-feat prerequisite to compose")
	}

	classLevelState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		1,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 8)},
		nil,
		nil,
		[]characterfeat.FeatID{
			characterfeat.ShieldFocusFeatID,
			characterfeat.ShieldProficiencyFeatID,
		},
	)

	if _, ok := NewCharacterFeat(characterfeat.GreaterShieldFocusFeatID, classLevelState); !ok {
		t.Fatal("expected class-level feat prerequisites to compose")
	}

	lowClassLevelState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		1,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 7)},
		nil,
		nil,
		[]characterfeat.FeatID{
			characterfeat.ShieldFocusFeatID,
			characterfeat.ShieldProficiencyFeatID,
		},
	)

	if _, ok := NewCharacterFeat(characterfeat.GreaterShieldFocusFeatID, lowClassLevelState); ok {
		t.Fatal("expected class-level feat prerequisites to reject a low fighter level")
	}
}

func TestNewCharacterFeat_ComposesCharacterLevelAndSpellcastingPrerequisites(t *testing.T) {
	characterLevelState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		[]CharacterClassLevel{
			mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 5),
			mustNewCharacterClassLevelForTest(t, characterclass.RogueClassID, 2),
		},
		nil,
		nil,
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.LeadershipFeatID, characterLevelState); !ok {
		t.Fatal("expected character-level prerequisite to compose from total class levels")
	}

	spellcastingState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1)},
		nil,
		nil,
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.ArcaneStrikeFeatID, spellcastingState); !ok {
		t.Fatal("expected arcane spellcasting prerequisite to compose from selected class levels")
	}
}

func TestNewCharacterFeat_ComposesCasterLevelPrerequisites(t *testing.T) {
	firstLevelCasterState := mustNewCharacterFeatPrerequisiteStateWithCasterLevelsForTest(
		t,
		nil,
		0,
		[]CharacterCasterLevel{mustNewCharacterCasterLevelForTest(t, ability.ArcaneCasterSource, 1)},
		nil,
		nil,
		nil,
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.ScribeScrollFeatID, firstLevelCasterState); !ok {
		t.Fatal("expected scribe scroll caster-level prerequisite to compose")
	}

	if _, ok := NewCharacterFeat(characterfeat.BrewPotionFeatID, firstLevelCasterState); ok {
		t.Fatal("expected brew potion to reject a low caster level")
	}

	thirdLevelCasterState := mustNewCharacterFeatPrerequisiteStateWithCasterLevelsForTest(
		t,
		nil,
		0,
		[]CharacterCasterLevel{mustNewCharacterCasterLevelForTest(t, ability.DivineCasterSource, 3)},
		nil,
		nil,
		nil,
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.BrewPotionFeatID, thirdLevelCasterState); !ok {
		t.Fatal("expected brew potion caster-level prerequisite to compose from any caster source")
	}
}

func TestNewCharacterFeat_CasterLevelPrerequisiteRequiresExplicitCasterLevelFact(t *testing.T) {
	classOnlyState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 5)},
		nil,
		nil,
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.BrewPotionFeatID, classOnlyState); ok {
		t.Fatal("expected caster-level prerequisites to reject class levels without a caster-level fact")
	}
}

func TestNewCharacterFeat_ComposesSelectedWeaponProficiencyPrerequisiteFromCategory(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateWithSelectedWeaponForTest(
		t,
		nil,
		1,
		nil,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)},
		nil,
		nil,
		mustNewCharacterSelectedWeaponForTest(t, characterequipment.DaggerWeaponID),
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.WeaponFocusFeatID, state); !ok {
		t.Fatal("expected weapon focus to compose from selected simple weapon and fighter category proficiency")
	}
}

func TestNewCharacterFeat_ComposesSelectedWeaponProficiencyPrerequisiteFromIndividualWeapon(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateWithSelectedWeaponForTest(
		t,
		nil,
		1,
		nil,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1)},
		nil,
		nil,
		mustNewCharacterSelectedWeaponForTest(t, characterequipment.CrossbowHeavyWeaponID),
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.WeaponFocusFeatID, state); !ok {
		t.Fatal("expected weapon focus to compose from selected heavy crossbow and wizard individual proficiency")
	}
}

func TestNewCharacterFeat_RejectsSelectedWeaponProficiencyWithoutMatchingClassProficiency(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateWithSelectedWeaponForTest(
		t,
		nil,
		1,
		nil,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.WizardClassID, 1)},
		nil,
		nil,
		mustNewCharacterSelectedWeaponForTest(t, characterequipment.SlingWeaponID),
		nil,
	)

	if _, ok := NewCharacterFeat(characterfeat.WeaponFocusFeatID, state); ok {
		t.Fatal("expected weapon focus to reject selected sling without wizard proficiency")
	}
}

func TestNewCharacterFeat_RejectsSelectedWeaponProficiencyWithoutSelectedWeaponContext(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateForTest(t, nil, 1, nil, nil, nil, nil)

	if _, ok := NewCharacterFeat(characterfeat.WeaponFocusFeatID, state); ok {
		t.Fatal("expected selected weapon prerequisite to reject without selected weapon context")
	}
}

func TestNewCharacterFeat_RejectsRemainingUnsupportedSelectionPrerequisites(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateForTest(t, nil, 1, nil, nil, nil, nil)

	if _, ok := NewCharacterFeat(characterfeat.GreaterSpellFocusFeatID, state); ok {
		t.Fatal("expected same-selection prerequisite to reject without selection context")
	}

	spellSchoolState := mustNewCharacterFeatPrerequisiteStateForTest(
		t,
		nil,
		0,
		nil,
		nil,
		nil,
		[]characterfeat.FeatID{characterfeat.SpellFocusFeatID},
	)
	if _, ok := NewCharacterFeat(characterfeat.AugmentSummoningFeatID, spellSchoolState); ok {
		t.Fatal("expected spell-school prerequisite to reject without spell school context")
	}
}

func TestCharacterFeatPrerequisiteState_RejectsMalformedSelectedWeaponFacts(t *testing.T) {
	if _, ok := NewCharacterFeatPrerequisiteStateWithSelectedWeapon(
		nil,
		1,
		nil,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)},
		nil,
		nil,
		characterSelectedWeapon{
			id:                  characterequipment.WeaponID(" dagger"),
			proficiencyCategory: characterequipment.SimpleWeaponProficiencyCategory,
			valid:               true,
		},
		nil,
	); ok {
		t.Fatal("expected malformed selected weapon id to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteStateWithSelectedWeapon(
		nil,
		1,
		nil,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)},
		nil,
		nil,
		characterSelectedWeapon{
			id:                  characterequipment.WeaponID("longsword"),
			proficiencyCategory: characterequipment.MartialWeaponProficiencyCategory,
			valid:               true,
		},
		nil,
	); ok {
		t.Fatal("expected unknown selected weapon id to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteStateWithSelectedWeapon(
		nil,
		1,
		nil,
		[]CharacterClassLevel{mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)},
		nil,
		nil,
		characterSelectedWeapon{
			id:                  characterequipment.DaggerWeaponID,
			proficiencyCategory: characterequipment.ExoticWeaponProficiencyCategory,
			valid:               true,
		},
		nil,
	); ok {
		t.Fatal("expected selected weapon fact with mismatched proficiency category to be rejected")
	}
}

func TestCharacterFeatPrerequisiteState_UnsupportedSelectedWeaponMappingFailsClosed(t *testing.T) {
	state := characterFeatPrerequisiteState{
		valid: true,
		classLevels: map[characterclass.ClassID]int{
			characterclass.FighterClassID: 1,
		},
		selectedWeapon: characterSelectedWeapon{
			id:                  characterequipment.WeaponID("net"),
			proficiencyCategory: characterequipment.ExoticWeaponProficiencyCategory,
			valid:               true,
		},
	}

	if state.SatisfiesPrerequisite(characterfeat.NewSelectedWeaponProficiencyPrerequisite()) {
		t.Fatal("expected unsupported selected weapon proficiency mapping to fail closed")
	}
}

func TestNewCharacterFeatPrerequisiteState_RejectsInvalidEntries(t *testing.T) {
	if _, ok := NewCharacterFeatPrerequisiteState(nil, -1, nil, nil, nil, nil, nil); ok {
		t.Fatal("expected negative base attack bonus to be rejected")
	}

	if _, ok := NewCharacterCasterLevel(ability.CasterSource("Mystic"), 1); ok {
		t.Fatal("expected unknown caster source to be rejected")
	}

	if _, ok := NewCharacterCasterLevel(ability.ArcaneCasterSource, 0); ok {
		t.Fatal("expected zero caster level to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteState(
		[]CharacterAbilityScore{{id: ability.AbilityScoreID("LUCK"), score: 10}},
		0,
		nil,
		nil,
		nil,
		nil,
		nil,
	); ok {
		t.Fatal("expected invalid ability score to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		[]CharacterCasterLevel{{source: ability.CasterSource("Mystic"), level: 1}},
		nil,
		nil,
		nil,
		nil,
	); ok {
		t.Fatal("expected invalid caster level to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		nil,
		[]CharacterClassLevel{{classID: characterclass.ClassID("alchemist"), level: 1}},
		nil,
		nil,
		nil,
	); ok {
		t.Fatal("expected invalid class level to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		nil,
		nil,
		[]characterclass.ClassFeatureID{characterclass.ClassFeatureID("alchemy")},
		nil,
		nil,
	); ok {
		t.Fatal("expected invalid class feature to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		nil,
		nil,
		nil,
		[]CharacterSkillRanks{{skillID: skill.SkillID("Sailing"), ranks: 1}},
		nil,
	); ok {
		t.Fatal("expected invalid skill ranks to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		nil,
		nil,
		nil,
		nil,
		[]characterfeat.FeatID{characterfeat.FeatID("Extra Alchemy")},
	); ok {
		t.Fatal("expected invalid known-feat reference to be rejected")
	}
}

func TestNewCharacterFeatPrerequisiteState_RejectsDuplicateEntries(t *testing.T) {
	strength := mustNewCharacterAbilityScoreForTest(t, ability.StrengthScore, 13)
	if _, ok := NewCharacterFeatPrerequisiteState(
		[]CharacterAbilityScore{strength, strength},
		0,
		nil,
		nil,
		nil,
		nil,
		nil,
	); ok {
		t.Fatal("expected duplicate ability scores to be rejected")
	}

	casterLevel := mustNewCharacterCasterLevelForTest(t, ability.ArcaneCasterSource, 1)
	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		[]CharacterCasterLevel{casterLevel, casterLevel},
		nil,
		nil,
		nil,
		nil,
	); ok {
		t.Fatal("expected duplicate caster levels to be rejected")
	}

	fighterLevel := mustNewCharacterClassLevelForTest(t, characterclass.FighterClassID, 1)
	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		nil,
		[]CharacterClassLevel{fighterLevel, fighterLevel},
		nil,
		nil,
		nil,
	); ok {
		t.Fatal("expected duplicate class levels to be rejected")
	}

	rideRanks := mustNewCharacterSkillRanksForTest(t, skill.RideSkillID, 1)
	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		nil,
		nil,
		nil,
		[]CharacterSkillRanks{rideRanks, rideRanks},
		nil,
	); ok {
		t.Fatal("expected duplicate skill ranks to be rejected")
	}

	if _, ok := NewCharacterFeatPrerequisiteState(
		nil,
		0,
		nil,
		nil,
		nil,
		nil,
		[]characterfeat.FeatID{characterfeat.EnduranceFeatID, characterfeat.EnduranceFeatID},
	); ok {
		t.Fatal("expected duplicate feats to be rejected")
	}
}

func TestCharacterFeatPrerequisiteState_ZeroValueDoesNotSatisfyFeat(t *testing.T) {
	var state CharacterFeatPrerequisiteState

	feat, ok := characterfeat.GetFeatByID(characterfeat.AcrobaticFeatID)
	if !ok {
		t.Fatal("expected acrobatic feat seed to resolve")
	}

	if state.SatisfiesFeat(feat) {
		t.Fatal("expected zero-value prerequisite state not to satisfy feats")
	}
}

func TestCharacterFeatPrerequisiteState_DoesNotSatisfyZeroValueFeat(t *testing.T) {
	state := mustNewCharacterFeatPrerequisiteStateForTest(t, nil, 0, nil, nil, nil, nil)
	var feat characterfeat.Feat

	if state.SatisfiesFeat(feat) {
		t.Fatal("expected prerequisite state not to satisfy zero-value feat")
	}
}

func TestCharacterFeat_ZeroValueDoesNotResolve(t *testing.T) {
	var selectedFeat CharacterFeat

	if _, ok := selectedFeat.GetFeat(); ok {
		t.Fatal("expected zero-value character feat not to resolve")
	}
}

func mustNewCharacterAbilityScoreForTest(
	t *testing.T,
	id ability.AbilityScoreID,
	score int,
) CharacterAbilityScore {
	t.Helper()

	value, ok := NewCharacterAbilityScore(id, score)
	if !ok {
		t.Fatalf("expected ability score %q %d to compose", id, score)
	}

	return value
}

func mustNewCharacterClassLevelForTest(
	t *testing.T,
	id characterclass.ClassID,
	level int,
) CharacterClassLevel {
	t.Helper()

	value, ok := NewCharacterClassLevel(id, level)
	if !ok {
		t.Fatalf("expected class level %q %d to compose", id, level)
	}

	return value
}

func mustNewCharacterCasterLevelForTest(
	t *testing.T,
	source ability.CasterSource,
	level int,
) CharacterCasterLevel {
	t.Helper()

	value, ok := NewCharacterCasterLevel(source, level)
	if !ok {
		t.Fatalf("expected caster level %q %d to compose", source, level)
	}

	return value
}

func mustNewCharacterSkillRanksForTest(
	t *testing.T,
	id skill.SkillID,
	ranks int,
) CharacterSkillRanks {
	t.Helper()

	value, ok := NewCharacterSkillRanks(id, ranks)
	if !ok {
		t.Fatalf("expected skill ranks %q %d to compose", id, ranks)
	}

	return value
}

func mustNewCharacterFeatPrerequisiteStateForTest(
	t *testing.T,
	abilityScores []CharacterAbilityScore,
	baseAttackBonus int,
	classLevels []CharacterClassLevel,
	classFeatures []characterclass.ClassFeatureID,
	skillRanks []CharacterSkillRanks,
	feats []characterfeat.FeatID,
) CharacterFeatPrerequisiteState {
	t.Helper()

	return mustNewCharacterFeatPrerequisiteStateWithCasterLevelsForTest(
		t,
		abilityScores,
		baseAttackBonus,
		nil,
		classLevels,
		classFeatures,
		skillRanks,
		feats,
	)
}

func mustNewCharacterFeatPrerequisiteStateWithCasterLevelsForTest(
	t *testing.T,
	abilityScores []CharacterAbilityScore,
	baseAttackBonus int,
	casterLevels []CharacterCasterLevel,
	classLevels []CharacterClassLevel,
	classFeatures []characterclass.ClassFeatureID,
	skillRanks []CharacterSkillRanks,
	feats []characterfeat.FeatID,
) CharacterFeatPrerequisiteState {
	t.Helper()

	return mustNewCharacterFeatPrerequisiteStateWithSelectedWeaponForTest(
		t,
		abilityScores,
		baseAttackBonus,
		casterLevels,
		classLevels,
		classFeatures,
		skillRanks,
		CharacterSelectedWeapon{},
		feats,
	)
}

func mustNewCharacterFeatPrerequisiteStateWithSelectedWeaponForTest(
	t *testing.T,
	abilityScores []CharacterAbilityScore,
	baseAttackBonus int,
	casterLevels []CharacterCasterLevel,
	classLevels []CharacterClassLevel,
	classFeatures []characterclass.ClassFeatureID,
	skillRanks []CharacterSkillRanks,
	selectedWeapon CharacterSelectedWeapon,
	feats []characterfeat.FeatID,
) CharacterFeatPrerequisiteState {
	t.Helper()

	state, ok := NewCharacterFeatPrerequisiteStateWithSelectedWeapon(
		abilityScores,
		baseAttackBonus,
		casterLevels,
		classLevels,
		classFeatures,
		skillRanks,
		selectedWeapon,
		feats,
	)
	if !ok {
		t.Fatal("expected feat prerequisite state to compose")
	}

	return state
}
