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

func TestNewCharacterCarriedWeight_ComposesWeaponsArmorAndShields(t *testing.T) {
	carriedWeight, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 10),
		[]CharacterEquipment{
			mustCharacterCarriedWeightCarryable(
				t,
				mustCharacterCarriedWeightWeaponRef(t, characterequipment.DaggerWeaponID),
				2,
			),
			mustCharacterCarriedWeightCarryable(
				t,
				mustCharacterCarriedWeightArmorRef(t, characterequipment.ChainShirtArmorID),
				1,
			),
			mustCharacterCarriedWeightCarryable(
				t,
				mustCharacterCarriedWeightArmorRef(t, characterequipment.ShieldHeavySteelArmorID),
				1,
			),
		},
	)
	if !ok {
		t.Fatal("expected weapons, armor, and shields to compose into carried weight")
	}

	if carriedWeight.GetTotalOunces() != 672 {
		t.Fatalf("expected total carried weight 672 oz, got %d oz", carriedWeight.GetTotalOunces())
	}

	if carriedWeight.GetTotalPounds() != 42 {
		t.Fatalf("expected total carried weight 42 lb, got %.2f lb", carriedWeight.GetTotalPounds())
	}

	if carriedWeight.GetLoadCategory() != MediumLoadCategory {
		t.Fatalf("expected medium load category, got %q", carriedWeight.GetLoadCategory())
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
	testCases := []struct {
		name string
		ref  characterequipment.CarryableItemRef
	}{
		{
			name: "equipment",
			ref:  mustCharacterCarriedWeightRef(t, characterequipment.EquipmentID("ten-foot-pole")),
		},
		{
			name: "weapon",
			ref:  mustCharacterCarriedWeightWeaponRef(t, characterequipment.WeaponID("longsword")),
		},
		{
			name: "armor",
			ref:  mustCharacterCarriedWeightArmorRef(t, characterequipment.ArmorID("breastplate")),
		},
	}

	for _, tc := range testCases {
		if _, ok := NewCharacterCarriedWeight(
			mustCharacterCarriedWeightStrength(t, 10),
			[]CharacterEquipment{{
				ref:      tc.ref,
				quantity: 1,
			}},
		); ok {
			t.Fatalf("expected unknown carried %s to be rejected", tc.name)
		}
	}

	backpackRef := mustCharacterCarriedWeightRef(t, characterequipment.BackpackEmptyEquipmentID)

	if _, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 10),
		[]CharacterEquipment{{ref: backpackRef, quantity: 0}},
	); ok {
		t.Fatal("expected carried equipment with zero quantity to be rejected")
	}

	if _, ok := NewCharacterCarriedWeight(
		mustCharacterCarriedWeightStrength(t, 10),
		[]CharacterEquipment{{ref: backpackRef, quantity: -1}},
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

	selectedEquipment, ok := NewCharacterEquipment(mustCharacterCarriedWeightRef(t, id), quantity)
	if !ok {
		t.Fatalf("expected selected equipment %q x%d to be constructed", id, quantity)
	}

	return selectedEquipment
}

func mustCharacterCarriedWeightCarryable(
	t *testing.T,
	ref characterequipment.CarryableItemRef,
	quantity int,
) CharacterEquipment {
	t.Helper()

	selectedEquipment, ok := NewCharacterEquipment(ref, quantity)
	if !ok {
		t.Fatalf("expected selected carryable item %q/%q x%d to be constructed", ref.GetKind(), ref.GetID(), quantity)
	}

	return selectedEquipment
}

func mustCharacterCarriedWeightRef(
	t *testing.T,
	id characterequipment.EquipmentID,
) characterequipment.CarryableItemRef {
	t.Helper()

	ref, ok := characterequipment.NewEquipmentCarryableItemRef(id)
	if !ok {
		t.Fatalf("expected carryable equipment ref %q to be constructed", id)
	}

	return ref
}

func mustCharacterCarriedWeightWeaponRef(
	t *testing.T,
	id characterequipment.WeaponID,
) characterequipment.CarryableItemRef {
	t.Helper()

	ref, ok := characterequipment.NewWeaponCarryableItemRef(id)
	if !ok {
		t.Fatalf("expected carryable weapon ref %q to be constructed", id)
	}

	return ref
}

func mustCharacterCarriedWeightArmorRef(
	t *testing.T,
	id characterequipment.ArmorID,
) characterequipment.CarryableItemRef {
	t.Helper()

	ref, ok := characterequipment.NewArmorCarryableItemRef(id)
	if !ok {
		t.Fatalf("expected carryable armor ref %q to be constructed", id)
	}

	return ref
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
