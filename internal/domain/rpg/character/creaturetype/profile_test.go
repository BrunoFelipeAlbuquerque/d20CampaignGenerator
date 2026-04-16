package creaturetype

import (
	"testing"

	ability "d20campaigngenerator/internal/domain/rpg/character/ability"
)

func TestNewCreatureTypeProfile_UsesRacialHDMetadata(t *testing.T) {
	profile, ok := NewCreatureTypeProfile(
		ability.D10HitDie,
		ability.BaseAttackBonusFull,
		[]ability.SavingThrowID{ability.FortitudeSave, ability.ReflexSave},
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{Darkvision60Trait, LowLightVisionTrait},
	)
	if !ok {
		t.Fatal("expected creature type profile to be constructed")
	}

	if profile.GetHitDieType() != ability.D10HitDie {
		t.Fatalf("expected hit die type %q, got %q", ability.D10HitDie, profile.GetHitDieType())
	}

	if profile.GetBABProgression() != ability.BaseAttackBonusFull {
		t.Fatalf("expected BAB progression %q, got %q", ability.BaseAttackBonusFull, profile.GetBABProgression())
	}

	if profile.GetSkillPointsPerHD() != 2 {
		t.Fatalf("expected skill points per HD 2, got %d", profile.GetSkillPointsPerHD())
	}

	if profile.GetHitPointKind() != ability.StandardHitPoints {
		t.Fatalf("expected hit point kind %q, got %q", ability.StandardHitPoints, profile.GetHitPointKind())
	}

	goodSaves := profile.GetGoodSaves()
	if len(goodSaves) != 2 || goodSaves[0] != ability.FortitudeSave || goodSaves[1] != ability.ReflexSave {
		t.Fatalf("expected good saves [Fortitude Reflex], got %v", goodSaves)
	}

	if !profile.HasTrait(Darkvision60Trait) {
		t.Fatal("expected Darkvision60 trait to be present")
	}
}

func TestNewCreatureTypeProfile_DedupesSavesAndTraits(t *testing.T) {
	profile, ok := NewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		[]ability.SavingThrowID{ability.WillSave, ability.WillSave},
		4,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{Darkvision60Trait, Darkvision60Trait, BreatheEatSleepTrait},
	)
	if !ok {
		t.Fatal("expected creature type profile to be constructed")
	}

	if len(profile.GetGoodSaves()) != 1 {
		t.Fatalf("expected deduped good saves length 1, got %d", len(profile.GetGoodSaves()))
	}

	if len(profile.GetTraitIDs()) != 2 {
		t.Fatalf("expected deduped trait ids length 2, got %d", len(profile.GetTraitIDs()))
	}
}

func TestNewCreatureTypeProfile_RejectsInvalidInputs(t *testing.T) {
	if _, ok := NewCreatureTypeProfile(
		ability.HitDieType("d20"),
		ability.BaseAttackBonusFull,
		nil,
		2,
		ability.StandardHitPoints,
		nil,
	); ok {
		t.Fatal("expected invalid hit die type to be rejected")
	}

	if _, ok := NewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusProgression("2/1"),
		nil,
		2,
		ability.StandardHitPoints,
		nil,
	); ok {
		t.Fatal("expected invalid BAB progression to be rejected")
	}

	if _, ok := NewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		[]ability.SavingThrowID{ability.SavingThrowID("Luck")},
		2,
		ability.StandardHitPoints,
		nil,
	); ok {
		t.Fatal("expected invalid save id to be rejected")
	}

	if _, ok := NewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		nil,
		-1,
		ability.StandardHitPoints,
		nil,
	); ok {
		t.Fatal("expected negative skill points per HD to be rejected")
	}

	if _, ok := NewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		nil,
		2,
		ability.HitPointKind("Vehicle"),
		nil,
	); ok {
		t.Fatal("expected invalid hit point kind to be rejected")
	}

	if _, ok := NewCreatureTypeProfile(
		ability.D8HitDie,
		ability.BaseAttackBonusThreeQuarters,
		nil,
		2,
		ability.StandardHitPoints,
		[]CreatureTypeTraitID{CreatureTypeTraitID("FastHealing")},
	); ok {
		t.Fatal("expected invalid trait id to be rejected")
	}
}

func TestGetCreatureTypeProfile_ReturnsCoreProfiles(t *testing.T) {
	undead, ok := GetCreatureTypeProfile(UndeadType)
	if !ok {
		t.Fatal("expected undead profile to exist")
	}

	if undead.GetHitPointKind() != ability.UndeadHitPoints {
		t.Fatalf("expected undead hit point kind %q, got %q", ability.UndeadHitPoints, undead.GetHitPointKind())
	}

	if undead.GetHitDieType() != ability.D8HitDie {
		t.Fatalf("expected undead hit die type %q, got %q", ability.D8HitDie, undead.GetHitDieType())
	}

	if !undead.HasTrait(NoConstitutionTrait) {
		t.Fatal("expected undead to have NoConstitution trait")
	}

	construct, ok := GetCreatureTypeProfile(ConstructType)
	if !ok {
		t.Fatal("expected construct profile to exist")
	}

	if construct.GetHitPointKind() != ability.ConstructHitPoints {
		t.Fatalf("expected construct hit point kind %q, got %q", ability.ConstructHitPoints, construct.GetHitPointKind())
	}

	humanoid, ok := GetCreatureTypeProfile(HumanoidType)
	if !ok {
		t.Fatal("expected humanoid profile to exist")
	}

	if len(humanoid.GetGoodSaves()) != 0 {
		t.Fatalf("expected humanoid fixed good saves to be empty, got %v", humanoid.GetGoodSaves())
	}

	if !humanoid.HasTrait(OneGoodSaveChoiceTrait) {
		t.Fatal("expected humanoid to carry one-good-save choice trait")
	}

	outsider, ok := GetCreatureTypeProfile(OutsiderType)
	if !ok {
		t.Fatal("expected outsider profile to exist")
	}

	if len(outsider.GetGoodSaves()) != 0 {
		t.Fatalf("expected outsider fixed good saves to be empty, got %v", outsider.GetGoodSaves())
	}

	if !outsider.HasTrait(TwoGoodSaveChoicesTrait) {
		t.Fatal("expected outsider to carry two-good-saves choice trait")
	}
}

func TestGetCreatureTypeProfile_RejectsUnknownType(t *testing.T) {
	if _, ok := GetCreatureTypeProfile(CreatureTypeID("Vehicle")); ok {
		t.Fatal("expected unknown creature type profile lookup to fail")
	}
}
