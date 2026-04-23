package class

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

func TestCoreClasses_SeedsElevenCoreEntries(t *testing.T) {
	testCases := []struct {
		id                  ClassID
		hitDieType          ability.HitDieType
		baseAttackBonus     ability.BaseAttackBonusProgression
		fortitude           ability.SavingThrowProgression
		reflex              ability.SavingThrowProgression
		will                ability.SavingThrowProgression
		skillRanksPerLevel  int
		classSkill          skill.SkillID
		weaponProficiency   WeaponProficiencyID
		armorProficiency    ArmorProficiencyID
		hasArmorProficiency bool
	}{
		{BarbarianClassID, ability.D12HitDie, ability.BaseAttackBonusFull, ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowPoor, 4, knowledgeNatureSkillID, MartialWeaponsWeaponProficiencyID, MediumArmorProficiencyID, true},
		{BardClassID, ability.D8HitDie, ability.BaseAttackBonusThreeQuarters, ability.SavingThrowPoor, ability.SavingThrowGood, ability.SavingThrowGood, 6, skill.PerformSkillID, LongswordWeaponProficiencyID, ShieldArmorProficiencyID, true},
		{ClericClassID, ability.D8HitDie, ability.BaseAttackBonusThreeQuarters, ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowGood, 2, knowledgeReligionSkillID, FavoredWeaponOfDeityWeaponProficiencyID, ShieldArmorProficiencyID, true},
		{DruidClassID, ability.D8HitDie, ability.BaseAttackBonusThreeQuarters, ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowGood, 4, knowledgeNatureSkillID, WildShapeNaturalAttacksWeaponProficiencyID, ShieldArmorProficiencyID, true},
		{FighterClassID, ability.D10HitDie, ability.BaseAttackBonusFull, ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowPoor, 2, knowledgeEngineeringSkillID, MartialWeaponsWeaponProficiencyID, TowerShieldArmorProficiencyID, true},
		{MonkClassID, ability.D8HitDie, ability.BaseAttackBonusThreeQuarters, ability.SavingThrowGood, ability.SavingThrowGood, ability.SavingThrowGood, 4, knowledgeReligionSkillID, ShurikenWeaponProficiencyID, "", false},
		{PaladinClassID, ability.D10HitDie, ability.BaseAttackBonusFull, ability.SavingThrowGood, ability.SavingThrowPoor, ability.SavingThrowGood, 2, knowledgeNobilitySkillID, MartialWeaponsWeaponProficiencyID, HeavyArmorProficiencyID, true},
		{RangerClassID, ability.D10HitDie, ability.BaseAttackBonusFull, ability.SavingThrowGood, ability.SavingThrowGood, ability.SavingThrowPoor, 6, knowledgeGeographySkillID, MartialWeaponsWeaponProficiencyID, MediumArmorProficiencyID, true},
		{RogueClassID, ability.D8HitDie, ability.BaseAttackBonusThreeQuarters, ability.SavingThrowPoor, ability.SavingThrowGood, ability.SavingThrowPoor, 8, knowledgeLocalSkillID, HandCrossbowWeaponProficiencyID, LightArmorProficiencyID, true},
		{SorcererClassID, ability.D6HitDie, ability.BaseAttackBonusHalf, ability.SavingThrowPoor, ability.SavingThrowPoor, ability.SavingThrowGood, 2, knowledgeArcanaSkillID, SimpleWeaponsWeaponProficiencyID, "", false},
		{WizardClassID, ability.D6HitDie, ability.BaseAttackBonusHalf, ability.SavingThrowPoor, ability.SavingThrowPoor, ability.SavingThrowGood, 2, skill.KnowledgeSkillID, CrossbowHeavyWeaponProficiencyID, "", false},
	}

	if len(coreClasses) != len(testCases) {
		t.Fatalf("expected %d core classes, got %d", len(testCases), len(coreClasses))
	}

	if len(coreClassOrder) != len(testCases) {
		t.Fatalf("expected %d ordered core classes, got %d", len(testCases), len(coreClassOrder))
	}

	for _, tc := range testCases {
		class, ok := coreClasses[tc.id]
		if !ok {
			t.Fatalf("expected core class %q to be seeded", tc.id)
		}

		if class.GetID() != tc.id {
			t.Fatalf("expected class id %q, got %q", tc.id, class.GetID())
		}

		if class.GetHitDieType() != tc.hitDieType {
			t.Fatalf("expected class %q hit die type %q, got %q", tc.id, tc.hitDieType, class.GetHitDieType())
		}

		if class.GetBaseAttackBonusProgression() != tc.baseAttackBonus {
			t.Fatalf("expected class %q BAB progression %q, got %q", tc.id, tc.baseAttackBonus, class.GetBaseAttackBonusProgression())
		}

		if class.GetSkillRanksPerLevel() != tc.skillRanksPerLevel {
			t.Fatalf("expected class %q skill ranks per level %d, got %d", tc.id, tc.skillRanksPerLevel, class.GetSkillRanksPerLevel())
		}

		assertClassHasSkill(t, class, tc.classSkill)
		assertClassHasWeaponProficiency(t, class, tc.weaponProficiency)

		if tc.hasArmorProficiency {
			assertClassHasArmorProficiency(t, class, tc.armorProficiency)
		} else if len(class.GetArmorProficiencies()) != 0 {
			t.Fatalf("expected class %q to have no armor proficiencies, got %v", tc.id, class.GetArmorProficiencies())
		}

		saveProgressions := class.GetSaveProgressions()
		assertSaveProgression(t, tc.id, saveProgressions, ability.FortitudeSave, tc.fortitude)
		assertSaveProgression(t, tc.id, saveProgressions, ability.ReflexSave, tc.reflex)
		assertSaveProgression(t, tc.id, saveProgressions, ability.WillSave, tc.will)
	}
}

func TestCoreClasses_SeedSpellcastingMetadata(t *testing.T) {
	testCases := []struct {
		id           ClassID
		kind         SpellcastingKind
		keyAbility   ability.AbilityScoreID
		spellcasting bool
	}{
		{BarbarianClassID, NonSpellcastingKind, "", false},
		{BardClassID, ArcaneSpontaneousSpellcastingKind, ability.CharismaScore, true},
		{ClericClassID, DivinePreparedSpellcastingKind, ability.WisdomScore, true},
		{DruidClassID, DivinePreparedSpellcastingKind, ability.WisdomScore, true},
		{FighterClassID, NonSpellcastingKind, "", false},
		{MonkClassID, NonSpellcastingKind, "", false},
		{PaladinClassID, DivinePreparedSpellcastingKind, ability.CharismaScore, true},
		{RangerClassID, DivinePreparedSpellcastingKind, ability.WisdomScore, true},
		{RogueClassID, NonSpellcastingKind, "", false},
		{SorcererClassID, ArcaneSpontaneousSpellcastingKind, ability.CharismaScore, true},
		{WizardClassID, ArcanePreparedSpellcastingKind, ability.IntelligenceScore, true},
	}

	for _, tc := range testCases {
		class := coreClasses[tc.id]
		spellcasting := class.GetSpellcasting()

		if spellcasting.GetKind() != tc.kind {
			t.Fatalf("expected class %q spellcasting kind %q, got %q", tc.id, tc.kind, spellcasting.GetKind())
		}

		if spellcasting.HasSpellcasting() != tc.spellcasting {
			t.Fatalf("expected class %q spellcasting=%t, got %t", tc.id, tc.spellcasting, spellcasting.HasSpellcasting())
		}

		keyAbility, ok := spellcasting.GetKeyAbility()
		if !tc.spellcasting {
			if ok {
				t.Fatalf("expected class %q not to expose a spellcasting key ability", tc.id)
			}

			continue
		}

		if !ok || keyAbility != tc.keyAbility {
			t.Fatalf("expected class %q key ability (%q, true), got (%q, %t)", tc.id, tc.keyAbility, keyAbility, ok)
		}
	}
}

func TestNewClass_AcceptsEverySeededCoreClassID(t *testing.T) {
	for _, id := range coreClassOrder {
		seeded, ok := coreClasses[id]
		if !ok {
			t.Fatalf("expected seeded core class %q to exist", id)
		}

		class, ok := NewClass(
			seeded.GetID(),
			seeded.GetHitDieType(),
			seeded.GetBaseAttackBonusProgression(),
			seeded.GetSaveProgressions(),
			seeded.GetSkillRanksPerLevel(),
			seeded.GetClassSkills(),
			seeded.GetWeaponProficiencies(),
			seeded.GetArmorProficiencies(),
			seeded.GetSpellcasting(),
		)
		if !ok {
			t.Fatalf("expected seeded core class id %q to remain valid", id)
		}

		if class.GetID() != id {
			t.Fatalf("expected reconstructed class id %q, got %q", id, class.GetID())
		}
	}
}

func TestCoreClasses_SeedGroupedSkillSpecializations(t *testing.T) {
	testCases := []struct {
		id       ClassID
		skillIDs []skill.SkillID
	}{
		{BarbarianClassID, []skill.SkillID{knowledgeNatureSkillID}},
		{ClericClassID, []skill.SkillID{knowledgeArcanaSkillID, knowledgePlanesSkillID, knowledgeReligionSkillID}},
		{DruidClassID, []skill.SkillID{knowledgeGeographySkillID, knowledgeNatureSkillID}},
		{FighterClassID, []skill.SkillID{knowledgeDungeoneeringSkillID, knowledgeEngineeringSkillID}},
		{MonkClassID, []skill.SkillID{knowledgeHistorySkillID, knowledgeReligionSkillID}},
		{PaladinClassID, []skill.SkillID{knowledgeNobilitySkillID, knowledgeReligionSkillID}},
		{RangerClassID, []skill.SkillID{knowledgeDungeoneeringSkillID, knowledgeGeographySkillID, knowledgeNatureSkillID}},
		{RogueClassID, []skill.SkillID{knowledgeDungeoneeringSkillID, knowledgeLocalSkillID}},
		{SorcererClassID, []skill.SkillID{knowledgeArcanaSkillID}},
	}

	for _, tc := range testCases {
		class := coreClasses[tc.id]

		for _, skillID := range tc.skillIDs {
			assertClassHasSkill(t, class, skillID)
		}
	}
}

func TestGetClassByID_ReturnsSeededCoreClass(t *testing.T) {
	for _, id := range coreClassOrder {
		class, ok := GetClassByID(id)
		if !ok {
			t.Fatalf("expected class lookup for %q to succeed", id)
		}

		if class.GetID() != id {
			t.Fatalf("expected looked up class id %q, got %q", id, class.GetID())
		}
	}
}

func TestGetClassByID_RejectsUnknownClass(t *testing.T) {
	if _, ok := GetClassByID(ClassID("alchemist")); ok {
		t.Fatal("expected unknown class lookup to fail")
	}
}

func assertClassHasSkill(t *testing.T, class Class, expected skill.SkillID) {
	t.Helper()

	for _, skillID := range class.GetClassSkills() {
		if skillID == expected {
			return
		}
	}

	t.Fatalf("expected class %q to have class skill %q", class.GetID(), expected)
}

func assertClassHasWeaponProficiency(t *testing.T, class Class, expected WeaponProficiencyID) {
	t.Helper()

	for _, proficiencyID := range class.GetWeaponProficiencies() {
		if proficiencyID == expected {
			return
		}
	}

	t.Fatalf("expected class %q to have weapon proficiency %q", class.GetID(), expected)
}

func assertClassHasArmorProficiency(t *testing.T, class Class, expected ArmorProficiencyID) {
	t.Helper()

	for _, proficiencyID := range class.GetArmorProficiencies() {
		if proficiencyID == expected {
			return
		}
	}

	t.Fatalf("expected class %q to have armor proficiency %q", class.GetID(), expected)
}

func assertSaveProgression(
	t *testing.T,
	classID ClassID,
	saveProgressions SaveProgressions,
	saveID ability.SavingThrowID,
	expected ability.SavingThrowProgression,
) {
	t.Helper()

	actual, ok := saveProgressions.GetProgression(saveID)
	if !ok || actual != expected {
		t.Fatalf("expected class %q save progression for %q to be (%q, true), got (%q, %t)", classID, saveID, expected, actual, ok)
	}
}
