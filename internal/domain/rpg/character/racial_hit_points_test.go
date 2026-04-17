package character

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/ability"
	"d20campaigngenerator/internal/domain/rpg/character/creaturetype"
)

func TestNewRacialHitPoints_StandardCreatureWorks(t *testing.T) {
	rules := mustResolveRules(t, creaturetype.AnimalType)

	hp, ok := NewRacialHitPoints(rules, 2, 14, 0, ability.MediumSize)
	if !ok {
		t.Fatal("expected standard racial hit points to be constructed")
	}

	if hp.GetKind() != ability.StandardHitPoints {
		t.Fatalf("expected hit point kind %q, got %q", ability.StandardHitPoints, hp.GetKind())
	}

	if hp.GetTotal() != 14 {
		t.Fatalf("expected total HP 14, got %d", hp.GetTotal())
	}
}

func TestNewRacialHitPoints_UndeadWorks(t *testing.T) {
	rules := mustResolveRules(t, creaturetype.UndeadType)

	hp, ok := NewRacialHitPoints(rules, 2, 0, 16, ability.MediumSize)
	if !ok {
		t.Fatal("expected undead racial hit points to be constructed")
	}

	if hp.GetKind() != ability.UndeadHitPoints {
		t.Fatalf("expected hit point kind %q, got %q", ability.UndeadHitPoints, hp.GetKind())
	}

	if hp.GetTotal() != 16 {
		t.Fatalf("expected total HP 16, got %d", hp.GetTotal())
	}
}

func TestNewRacialHitPoints_ConstructWorks(t *testing.T) {
	rules := mustResolveRules(t, creaturetype.ConstructType)

	hp, ok := NewRacialHitPoints(rules, 2, 0, 0, ability.MediumSize)
	if !ok {
		t.Fatal("expected construct racial hit points to be constructed")
	}

	if hp.GetKind() != ability.ConstructHitPoints {
		t.Fatalf("expected hit point kind %q, got %q", ability.ConstructHitPoints, hp.GetKind())
	}

	if hp.GetTotal() != 32 {
		t.Fatalf("expected total HP 32, got %d", hp.GetTotal())
	}
}

func TestNewRacialHitPoints_HumanoidKeepsContextualFlagAndStillBuildsHP(t *testing.T) {
	rules := mustResolveRules(t, creaturetype.HumanoidType)

	if !rules.HasContextualFlag(creaturetype.HumanoidRacialHDUsesClassRulesFlag) {
		t.Fatal("expected humanoid contextual flag to remain available")
	}

	hp, ok := NewRacialHitPoints(rules, 1, 12, 0, ability.MediumSize)
	if !ok {
		t.Fatal("expected humanoid racial hit points to be constructed")
	}

	if hp.GetKind() != ability.StandardHitPoints {
		t.Fatalf("expected hit point kind %q, got %q", ability.StandardHitPoints, hp.GetKind())
	}

	hd := hp.GetHitDie()
	d8Count, ok := hd.GetDieCount(ability.D8HitDie)
	if !ok || d8Count != 1 {
		t.Fatalf("expected humanoid racial hit dice (1 d8, true), got (%d, %t)", d8Count, ok)
	}
}

func mustResolveRules(t *testing.T, baseType creaturetype.CreatureTypeID) creaturetype.ResolvedCreatureRules {
	t.Helper()

	classification, ok := creaturetype.NewCreatureClassification(baseType, nil, nil)
	if !ok {
		t.Fatal("expected classification to be constructed")
	}

	rules, ok := creaturetype.ResolveCreatureRules(classification)
	if !ok {
		t.Fatal("expected creature rules to resolve")
	}

	return rules
}
