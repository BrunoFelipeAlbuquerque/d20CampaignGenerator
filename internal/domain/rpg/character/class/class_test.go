package class

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
	"d20campaigngenerator/internal/domain/rpg/character/skill"
)

func TestNewClass_ConstructsValidatedClassChassis(t *testing.T) {
	saveProgressions := mustSaveProgressionsForTest(
		t,
		ability.SavingThrowPoor,
		ability.SavingThrowPoor,
		ability.SavingThrowGood,
	)
	spellcasting := mustSpellcastingProfileForTest(
		t,
		ArcanePreparedSpellcastingKind,
		ability.IntelligenceScore,
	)

	class, ok := NewClass(
		ClassID("wizard"),
		ability.D6HitDie,
		ability.BaseAttackBonusHalf,
		saveProgressions,
		2,
		[]skill.SkillID{skill.KnowledgeSkillID, skill.SpellcraftSkillID},
		[]WeaponProficiencyID{
			ClubWeaponProficiencyID,
			DaggerWeaponProficiencyID,
			CrossbowHeavyWeaponProficiencyID,
			CrossbowLightWeaponProficiencyID,
			QuarterstaffWeaponProficiencyID,
		},
		nil,
		spellcasting,
	)
	if !ok {
		t.Fatal("expected class chassis to be constructed")
	}

	if class.GetID() != ClassID("wizard") {
		t.Fatalf("expected class id %q, got %q", ClassID("wizard"), class.GetID())
	}

	if class.GetHitDieType() != ability.D6HitDie {
		t.Fatalf("expected hit die type %q, got %q", ability.D6HitDie, class.GetHitDieType())
	}

	if class.GetBaseAttackBonusProgression() != ability.BaseAttackBonusHalf {
		t.Fatalf(
			"expected BAB progression %q, got %q",
			ability.BaseAttackBonusHalf,
			class.GetBaseAttackBonusProgression(),
		)
	}

	if class.GetSkillRanksPerLevel() != 2 {
		t.Fatalf("expected skill ranks per level 2, got %d", class.GetSkillRanksPerLevel())
	}

	classSkills := class.GetClassSkills()
	if len(classSkills) != 2 ||
		classSkills[0] != skill.KnowledgeSkillID ||
		classSkills[1] != skill.SpellcraftSkillID {
		t.Fatalf("expected class skills [Knowledge Spellcraft], got %v", classSkills)
	}

	weaponProficiencies := class.GetWeaponProficiencies()
	if len(weaponProficiencies) != 5 {
		t.Fatalf("expected 5 weapon proficiencies, got %d", len(weaponProficiencies))
	}

	if len(class.GetArmorProficiencies()) != 0 {
		t.Fatalf("expected no armor proficiencies, got %v", class.GetArmorProficiencies())
	}

	fortitude, ok := class.GetSaveProgressions().GetProgression(ability.FortitudeSave)
	if !ok || fortitude != ability.SavingThrowPoor {
		t.Fatalf("expected fortitude progression (%q, true), got (%q, %t)", ability.SavingThrowPoor, fortitude, ok)
	}

	will, ok := class.GetSaveProgressions().GetProgression(ability.WillSave)
	if !ok || will != ability.SavingThrowGood {
		t.Fatalf("expected will progression (%q, true), got (%q, %t)", ability.SavingThrowGood, will, ok)
	}

	if !class.GetSpellcasting().HasSpellcasting() {
		t.Fatal("expected wizard to expose spellcasting metadata")
	}

	if class.GetSpellcasting().GetKind() != ArcanePreparedSpellcastingKind {
		t.Fatalf(
			"expected spellcasting kind %q, got %q",
			ArcanePreparedSpellcastingKind,
			class.GetSpellcasting().GetKind(),
		)
	}

	keyAbility, ok := class.GetSpellcasting().GetKeyAbility()
	if !ok || keyAbility != ability.IntelligenceScore {
		t.Fatalf("expected key ability (%q, true), got (%q, %t)", ability.IntelligenceScore, keyAbility, ok)
	}
}

func TestNewClass_StoresExplicitNonSpellcastingMetadata(t *testing.T) {
	saveProgressions := mustSaveProgressionsForTest(
		t,
		ability.SavingThrowGood,
		ability.SavingThrowPoor,
		ability.SavingThrowPoor,
	)

	class, ok := NewClass(
		ClassID("fighter"),
		ability.D10HitDie,
		ability.BaseAttackBonusFull,
		saveProgressions,
		2,
		[]skill.SkillID{skill.ClimbSkillID, skill.SwimSkillID},
		[]WeaponProficiencyID{
			SimpleWeaponsWeaponProficiencyID,
			MartialWeaponsWeaponProficiencyID,
		},
		[]ArmorProficiencyID{
			LightArmorProficiencyID,
			MediumArmorProficiencyID,
			HeavyArmorProficiencyID,
			ShieldArmorProficiencyID,
			TowerShieldArmorProficiencyID,
		},
		NewNonSpellcastingProfile(),
	)
	if !ok {
		t.Fatal("expected non-spellcasting class chassis to be constructed")
	}

	if class.GetSpellcasting().HasSpellcasting() {
		t.Fatal("expected fighter not to expose spellcasting metadata")
	}

	if _, ok := class.GetSpellcasting().GetKeyAbility(); ok {
		t.Fatal("expected non-spellcasting class not to expose key ability metadata")
	}
}

func TestNewClass_DedupesSkillsAndProficiencies(t *testing.T) {
	saveProgressions := mustSaveProgressionsForTest(
		t,
		ability.SavingThrowPoor,
		ability.SavingThrowGood,
		ability.SavingThrowPoor,
	)

	class, ok := NewClass(
		ClassID("rogue"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		saveProgressions,
		8,
		[]skill.SkillID{
			skill.BluffSkillID,
			skill.BluffSkillID,
			skill.DisableDeviceSkillID,
		},
		[]WeaponProficiencyID{
			SimpleWeaponsWeaponProficiencyID,
			SimpleWeaponsWeaponProficiencyID,
			HandCrossbowWeaponProficiencyID,
		},
		[]ArmorProficiencyID{
			LightArmorProficiencyID,
			LightArmorProficiencyID,
		},
		NewNonSpellcastingProfile(),
	)
	if !ok {
		t.Fatal("expected class chassis with duplicated metadata to be constructed")
	}

	if len(class.GetClassSkills()) != 2 {
		t.Fatalf("expected deduped class skills length 2, got %d", len(class.GetClassSkills()))
	}

	if len(class.GetWeaponProficiencies()) != 2 {
		t.Fatalf("expected deduped weapon proficiencies length 2, got %d", len(class.GetWeaponProficiencies()))
	}

	if len(class.GetArmorProficiencies()) != 1 {
		t.Fatalf("expected deduped armor proficiencies length 1, got %d", len(class.GetArmorProficiencies()))
	}
}

func TestNewClass_AcceptsSpecializedGroupedClassSkills(t *testing.T) {
	saveProgressions := mustSaveProgressionsForTest(
		t,
		ability.SavingThrowGood,
		ability.SavingThrowPoor,
		ability.SavingThrowPoor,
	)

	class, ok := NewClass(
		ClassID("barbarian"),
		ability.D12HitDie,
		ability.BaseAttackBonusFull,
		saveProgressions,
		4,
		[]skill.SkillID{skill.CraftSkillID, skill.SkillID("Knowledge (nature)")},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID, MartialWeaponsWeaponProficiencyID},
		[]ArmorProficiencyID{LightArmorProficiencyID},
		NewNonSpellcastingProfile(),
	)
	if !ok {
		t.Fatal("expected specialized grouped class skill ids to be accepted")
	}

	if got := class.GetClassSkills()[1]; got != skill.SkillID("Knowledge (nature)") {
		t.Fatalf("expected specialized grouped class skill %q, got %q", skill.SkillID("Knowledge (nature)"), got)
	}
}

func TestNewClass_RejectsInvalidInputs(t *testing.T) {
	validSaves := mustSaveProgressionsForTest(
		t,
		ability.SavingThrowGood,
		ability.SavingThrowPoor,
		ability.SavingThrowPoor,
	)
	validSpellcasting := mustSpellcastingProfileForTest(
		t,
		DivinePreparedSpellcastingKind,
		ability.WisdomScore,
	)

	if _, ok := NewSaveProgressions(ability.SavingThrowProgression("bad"), ability.SavingThrowPoor, ability.SavingThrowPoor); ok {
		t.Fatal("expected invalid save progression metadata to be rejected")
	}

	if _, ok := NewSpellcastingProfile(NonSpellcastingKind, ability.IntelligenceScore); ok {
		t.Fatal("expected non-spellcasting kind to require the explicit non-spellcasting constructor")
	}

	if _, ok := NewSpellcastingProfile(ArcanePreparedSpellcastingKind, ability.AbilityScoreID("LCK")); ok {
		t.Fatal("expected invalid spellcasting key ability to be rejected")
	}

	if _, ok := NewClass(
		ClassID(" "),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		NewNonSpellcastingProfile(),
	); ok {
		t.Fatal("expected blank class id to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.HitDieType("d20"),
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected invalid hit die type to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusProgression("2/3"),
		validSaves,
		6,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected invalid BAB progression to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		SaveProgressions{},
		6,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected zero-valued save progression metadata to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		0,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected non-positive skill ranks per level to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		nil,
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected missing class skills to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		[]skill.SkillID{skill.SkillID("Jump")},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected unknown class skill to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		[]skill.SkillID{skill.BluffSkillID},
		nil,
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected missing weapon proficiency metadata to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{WeaponProficiencyID("Laser Rifle")},
		nil,
		validSpellcasting,
	); ok {
		t.Fatal("expected unknown weapon proficiency to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		[]ArmorProficiencyID{ArmorProficiencyID("Power Armor")},
		validSpellcasting,
	); ok {
		t.Fatal("expected unknown armor proficiency to be rejected")
	}

	if _, ok := NewClass(
		ClassID("bard"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		validSaves,
		6,
		[]skill.SkillID{skill.BluffSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		SpellcastingProfile{},
	); ok {
		t.Fatal("expected zero-valued spellcasting metadata to be rejected")
	}

	if _, ok := NewClass(
		ClassID("fighter"),
		ability.D10HitDie,
		ability.BaseAttackBonusFull,
		validSaves,
		2,
		[]skill.SkillID{skill.ClimbSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID},
		nil,
		SpellcastingProfile{
			kind:       NonSpellcastingKind,
			keyAbility: ability.IntelligenceScore,
		},
	); ok {
		t.Fatal("expected non-spellcasting profile with key ability metadata to be rejected")
	}
}

func TestClass_GettersReturnDefensiveCopies(t *testing.T) {
	saveProgressions := mustSaveProgressionsForTest(
		t,
		ability.SavingThrowPoor,
		ability.SavingThrowGood,
		ability.SavingThrowPoor,
	)

	class, ok := NewClass(
		ClassID("rogue"),
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		saveProgressions,
		8,
		[]skill.SkillID{skill.BluffSkillID, skill.DisableDeviceSkillID},
		[]WeaponProficiencyID{SimpleWeaponsWeaponProficiencyID, HandCrossbowWeaponProficiencyID},
		[]ArmorProficiencyID{LightArmorProficiencyID},
		NewNonSpellcastingProfile(),
	)
	if !ok {
		t.Fatal("expected class chassis to be constructed")
	}

	classSkills := class.GetClassSkills()
	weaponProficiencies := class.GetWeaponProficiencies()
	armorProficiencies := class.GetArmorProficiencies()

	classSkills[0] = skill.SkillID("Jump")
	weaponProficiencies[0] = WeaponProficiencyID("Laser Rifle")
	armorProficiencies[0] = ArmorProficiencyID("Power Armor")

	if class.GetClassSkills()[0] != skill.BluffSkillID {
		t.Fatalf("expected defensive class skill copy to preserve %q", skill.BluffSkillID)
	}

	if class.GetWeaponProficiencies()[0] != SimpleWeaponsWeaponProficiencyID {
		t.Fatalf("expected defensive weapon proficiency copy to preserve %q", SimpleWeaponsWeaponProficiencyID)
	}

	if class.GetArmorProficiencies()[0] != LightArmorProficiencyID {
		t.Fatalf("expected defensive armor proficiency copy to preserve %q", LightArmorProficiencyID)
	}
}

func TestSaveProgressions_GetProgression_RejectsUnknownSaveID(t *testing.T) {
	saveProgressions := mustSaveProgressionsForTest(
		t,
		ability.SavingThrowGood,
		ability.SavingThrowPoor,
		ability.SavingThrowPoor,
	)

	if _, ok := saveProgressions.GetProgression(ability.SavingThrowID("Initiative")); ok {
		t.Fatal("expected unknown save id lookup to fail")
	}
}

func mustSaveProgressionsForTest(
	t *testing.T,
	fortitude ability.SavingThrowProgression,
	reflex ability.SavingThrowProgression,
	will ability.SavingThrowProgression,
) SaveProgressions {
	t.Helper()

	value, ok := NewSaveProgressions(fortitude, reflex, will)
	if !ok {
		t.Fatal("expected save progression metadata to be constructed")
	}

	return value
}

func mustSpellcastingProfileForTest(
	t *testing.T,
	kind SpellcastingKind,
	keyAbility ability.AbilityScoreID,
) SpellcastingProfile {
	t.Helper()

	value, ok := NewSpellcastingProfile(kind, keyAbility)
	if !ok {
		t.Fatal("expected spellcasting metadata to be constructed")
	}

	return value
}
