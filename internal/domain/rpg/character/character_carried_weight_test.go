package character

import (
	"testing"

	"d20campaigngenerator/internal/domain/rpg/character/ability"
	characterequipment "d20campaigngenerator/internal/domain/rpg/character/equipment"
)

func TestNewCharacterCarriedWeight_ComposesLightLoadFromSelectedEquipment(t *testing.T) {
	carriedWeight, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 10),
		[]CharacterEquipment{
			mustCharacterCarriedWeightEquipment(t, characterequipment.BackpackEmptyEquipmentID, 2),
			mustCharacterCarriedWeightEquipment(t, characterequipment.RopeHemp50FeetEquipmentID, 1),
			mustCharacterCarriedWeightEquipment(t, characterequipment.FlintAndSteelEquipmentID, 1),
		},
	)
	if !ok {
		t.Fatal("expected carried weight to compose through character boundary")
	}

	if carriedWeight.GetTotalOunces() != 224 {
		t.Fatalf("expected total carried weight 224 oz, got %d oz", carriedWeight.GetTotalOunces())
	}

	if carriedWeight.GetTotalPounds() != 14 {
		t.Fatalf("expected total carried weight 14 lb, got %.2f lb", carriedWeight.GetTotalPounds())
	}

	if carriedWeight.GetLoadCategory() != LightLoadCategory {
		t.Fatalf("expected light load category, got %q", carriedWeight.GetLoadCategory())
	}
}

func TestNewCharacterCarriedWeight_ClassifiesMediumHeavyAndOverMaximumLoads(t *testing.T) {
	testCases := []struct {
		name      string
		equipment []CharacterEquipment
		want      CharacterLoadCategory
	}{
		{
			name: "medium",
			equipment: []CharacterEquipment{
				mustCharacterCarriedWeightEquipment(t, characterequipment.RopeHemp50FeetEquipmentID, 3),
				mustCharacterCarriedWeightEquipment(t, characterequipment.WaterskinEquipmentID, 1),
			},
			want: MediumLoadCategory,
		},
		{
			name: "heavy",
			equipment: []CharacterEquipment{
				mustCharacterCarriedWeightEquipment(t, characterequipment.RopeHemp50FeetEquipmentID, 6),
				mustCharacterCarriedWeightEquipment(t, characterequipment.BedrollEquipmentID, 1),
				mustCharacterCarriedWeightEquipment(t, characterequipment.BackpackEmptyEquipmentID, 1),
			},
			want: HeavyLoadCategory,
		},
		{
			name: "over maximum",
			equipment: []CharacterEquipment{
				mustCharacterCarriedWeightEquipment(t, characterequipment.RopeHemp50FeetEquipmentID, 11),
			},
			want: OverMaximumLoadCategory,
		},
	}

	for _, tc := range testCases {
		carriedWeight, ok := NewCharacterCarriedWeight(mustCharacterCarriedWeightStrength(t, 10), tc.equipment)
		if !ok {
			t.Fatalf("expected %s carried weight to compose", tc.name)
		}

		if carriedWeight.GetLoadCategory() != tc.want {
			t.Fatalf("expected %s load category %q, got %q", tc.name, tc.want, carriedWeight.GetLoadCategory())
		}
	}
}

func TestNewCharacterCarriedWeight_PreservesStrengthCapacityBoundary(t *testing.T) {
	carriedWeight, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 18),
		[]CharacterEquipment{
			mustCharacterCarriedWeightEquipment(t, characterequipment.RopeHemp50FeetEquipmentID, 10),
		},
	)
	if !ok {
		t.Fatal("expected carried weight at light-load maximum to compose")
	}

	if carriedWeight.GetTotalPounds() != 100 {
		t.Fatalf("expected total carried weight 100 lb, got %.2f lb", carriedWeight.GetTotalPounds())
	}

	if carriedWeight.GetLoadCategory() != LightLoadCategory {
		t.Fatalf("expected exact light-load maximum to remain light load, got %q", carriedWeight.GetLoadCategory())
	}
}

func TestNewCharacterCarriedWeight_AllowsEmptyInventory(t *testing.T) {
	carriedWeight, ok := NewCharacterCarriedWeight(mustCharacterCarriedWeightStrength(t, 10), nil)
	if !ok {
		t.Fatal("expected empty carried inventory to compose")
	}

	if carriedWeight.GetTotalOunces() != 0 {
		t.Fatalf("expected empty carried inventory weight 0 oz, got %d oz", carriedWeight.GetTotalOunces())
	}

	if carriedWeight.GetLoadCategory() != LightLoadCategory {
		t.Fatalf("expected empty carried inventory to be light load, got %q", carriedWeight.GetLoadCategory())
	}
}

func TestNewCharacterCarriedWeight_RejectsInvalidStrength(t *testing.T) {
	if _, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightAbilityScore(t, ability.DexterityScore, 10, true),
		nil,
	); ok {
		t.Fatal("expected non-strength carrying score to be rejected")
	}

	if _, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightAbilityScore(t, ability.StrengthScore, 10, false),
		nil,
	); ok {
		t.Fatal("expected suppressed strength to be rejected")
	}
}

func TestNewCharacterCarriedWeight_RejectsInvalidCarriedEquipment(t *testing.T) {
	if _, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 10),
		[]CharacterEquipment{{id: characterequipment.EquipmentID("ten-foot-pole"), quantity: 1}},
	); ok {
		t.Fatal("expected unknown carried equipment to be rejected")
	}

	if _, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 10),
		[]CharacterEquipment{{id: characterequipment.BackpackEmptyEquipmentID, quantity: 0}},
	); ok {
		t.Fatal("expected carried equipment with zero quantity to be rejected")
	}

	if _, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 10),
		[]CharacterEquipment{{id: characterequipment.BackpackEmptyEquipmentID, quantity: -1}},
	); ok {
		t.Fatal("expected carried equipment with negative quantity to be rejected")
	}
}

func mustCharacterCarriedWeightEquipment(
	t *testing.T,
	id characterequipment.EquipmentID,
	quantity int,
) CharacterEquipment {
	t.Helper()

	selectedEquipment, ok := NewCharacterEquipment(id, quantity)
	if !ok {
		t.Fatalf("expected selected equipment %q x%d to be constructed", id, quantity)
	}

	return selectedEquipment
}

func mustCharacterCarriedWeightStrength(t *testing.T, score int) ability.AbilityScore {
	t.Helper()

	return mustCharacterCarriedWeightAbilityScore(t, ability.StrengthScore, score, true)
}

func mustCharacterCarriedWeightAbilityScore(
	t *testing.T,
	id ability.AbilityScoreID,
	score int,
	valid bool,
) ability.AbilityScore {
	t.Helper()

	value, ok := ability.NewAbilityScoreValue(score, valid)
	if !ok {
		t.Fatalf("expected ability score value %d to be constructed", score)
	}

	abilityScore, ok := ability.NewAbilityScore(id, value)
	if !ok {
		t.Fatalf("expected ability score %q to be constructed", id)
	}

	return abilityScore
}
