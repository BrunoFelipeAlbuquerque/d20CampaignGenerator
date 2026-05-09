package equipment

import (
	"math"
	"testing"
)

func TestEquipmentWeight_StoresImperialAndMetricUnits(t *testing.T) {
	weight, ok := NewEquipmentWeightOunces(32)
	if !ok {
		t.Fatal("expected equipment weight to be constructed")
	}

	if weight.GetOunces() != 32 {
		t.Fatalf("expected weight 32 oz, got %d oz", weight.GetOunces())
	}

	assertFloatAlmostEqual(t, weight.GetKilograms(), 0.90718474)
}

func TestEquipmentWeight_ConstructsFromMetricUnits(t *testing.T) {
	weight, ok := NewEquipmentWeightKilograms(0.90718474)
	if !ok {
		t.Fatal("expected equipment weight to be constructed from kilograms")
	}

	if weight.GetOunces() != 32 {
		t.Fatalf("expected metric weight to preserve equivalent 32 oz, got %d oz", weight.GetOunces())
	}

	assertFloatAlmostEqual(t, weight.GetKilograms(), 0.90718474)
}

func TestWeaponRangeIncrement_StoresImperialAndMetricUnits(t *testing.T) {
	rangeIncrement, ok := NewWeaponRangeIncrementFeet(80)
	if !ok {
		t.Fatal("expected weapon range increment to be constructed")
	}

	if rangeIncrement.GetFeet() != 80 {
		t.Fatalf("expected range increment 80 ft, got %d ft", rangeIncrement.GetFeet())
	}

	assertFloatAlmostEqual(t, rangeIncrement.GetMeters(), 24.384)
}

func TestWeaponRangeIncrement_ConstructsFromMetricUnits(t *testing.T) {
	rangeIncrement, ok := NewWeaponRangeIncrementMeters(24.384)
	if !ok {
		t.Fatal("expected weapon range increment to be constructed from meters")
	}

	if rangeIncrement.GetFeet() != 80 {
		t.Fatalf("expected metric range increment to preserve equivalent 80 ft, got %d ft", rangeIncrement.GetFeet())
	}

	assertFloatAlmostEqual(t, rangeIncrement.GetMeters(), 24.384)
}

func TestWeaponRangeIncrement_NoRangeHasNoMetricLength(t *testing.T) {
	if meters := NewNoWeaponRangeIncrement().GetMeters(); meters != 0 {
		t.Fatalf("expected no-range weapon increment to report 0 meters, got %.6f", meters)
	}
}

func TestArmorSpeedImpact_StoresImperialAndMetricUnits(t *testing.T) {
	speedImpact, ok := NewArmorSpeedImpact(20, 15, true)
	if !ok {
		t.Fatal("expected armor speed impact to be constructed")
	}

	if speedImpact.GetSpeedFor30FeetBase() != 20 {
		t.Fatalf("expected 30-foot base speed impact 20 ft, got %d ft", speedImpact.GetSpeedFor30FeetBase())
	}

	if speedImpact.GetSpeedFor20FeetBase() != 15 {
		t.Fatalf("expected 20-foot base speed impact 15 ft, got %d ft", speedImpact.GetSpeedFor20FeetBase())
	}

	assertFloatAlmostEqual(t, speedImpact.GetSpeedFor30FeetBaseMeters(), 6.096)
	assertFloatAlmostEqual(t, speedImpact.GetSpeedFor20FeetBaseMeters(), 4.572)
}

func TestArmorSpeedImpact_ConstructsFromMetricUnits(t *testing.T) {
	speedImpact, ok := NewArmorSpeedImpactMeters(6.096, 4.572, true)
	if !ok {
		t.Fatal("expected armor speed impact to be constructed from meters")
	}

	if speedImpact.GetSpeedFor30FeetBase() != 20 {
		t.Fatalf("expected metric 30-foot base speed impact to preserve equivalent 20 ft, got %d ft", speedImpact.GetSpeedFor30FeetBase())
	}

	if speedImpact.GetSpeedFor20FeetBase() != 15 {
		t.Fatalf("expected metric 20-foot base speed impact to preserve equivalent 15 ft, got %d ft", speedImpact.GetSpeedFor20FeetBase())
	}

	assertFloatAlmostEqual(t, speedImpact.GetSpeedFor30FeetBaseMeters(), 6.096)
	assertFloatAlmostEqual(t, speedImpact.GetSpeedFor20FeetBaseMeters(), 4.572)
}

func TestArmorSpeedImpact_NoImpactHasNoMetricLength(t *testing.T) {
	speedImpact := NewNoArmorSpeedImpact()

	if meters := speedImpact.GetSpeedFor30FeetBaseMeters(); meters != 0 {
		t.Fatalf("expected no-impact 30-foot base speed to report 0 meters, got %.6f", meters)
	}

	if meters := speedImpact.GetSpeedFor20FeetBaseMeters(); meters != 0 {
		t.Fatalf("expected no-impact 20-foot base speed to report 0 meters, got %.6f", meters)
	}
}

func TestCoreEquipmentSeeds_StoreImperialAndMetricWeight(t *testing.T) {
	equipment := coreEquipment[BackpackEmptyEquipmentID]

	if equipment.weight.ounces != 32 {
		t.Fatalf("expected backpack seed to store 32 oz, got %d oz", equipment.weight.ounces)
	}

	assertFloatAlmostEqual(t, equipment.weight.kilograms, 0.90718474)
}

func TestCoreWeaponSeeds_StoreImperialAndMetricRangeAndWeight(t *testing.T) {
	weapon := coreSimpleWeapons[DaggerWeaponID]

	if weapon.rangeIncrement.feet != 10 {
		t.Fatalf("expected dagger seed to store 10 ft range increment, got %d ft", weapon.rangeIncrement.feet)
	}

	assertFloatAlmostEqual(t, weapon.rangeIncrement.meters, 3.048)

	if weapon.weight.ounces != 16 {
		t.Fatalf("expected dagger seed to store 16 oz, got %d oz", weapon.weight.ounces)
	}

	assertFloatAlmostEqual(t, weapon.weight.kilograms, 0.45359237)
}

func TestCoreArmorSeeds_StoreImperialAndMetricWeight(t *testing.T) {
	armor := coreArmor[ChainShirtArmorID]

	if armor.weight.ounces != 400 {
		t.Fatalf("expected chain shirt seed to store 400 oz, got %d oz", armor.weight.ounces)
	}

	assertFloatAlmostEqual(t, armor.weight.kilograms, 11.33980925)
}

func assertFloatAlmostEqual(t *testing.T, got float64, expected float64) {
	t.Helper()

	if math.Abs(got-expected) > 0.000001 {
		t.Fatalf("expected %.6f, got %.6f", expected, got)
	}
}
