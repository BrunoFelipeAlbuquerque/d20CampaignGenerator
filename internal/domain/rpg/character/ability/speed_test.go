package ability

import "testing"

func TestNewSpeed_AllowsMixedMovementAndFlyManeuverability(t *testing.T) {
	speed, ok := NewSpeed(30, 10, 20, 60, 15, GoodFlyManeuverability)
	if !ok {
		t.Fatal("expected speed profile to be constructed")
	}

	land, ok := speed.GetMovement(LandMovement)
	if !ok || land != 30 {
		t.Fatalf("expected land speed (30, true), got (%d, %t)", land, ok)
	}

	burrow, ok := speed.GetMovement(BurrowMovement)
	if !ok || burrow != 10 {
		t.Fatalf("expected burrow speed (10, true), got (%d, %t)", burrow, ok)
	}

	climb, ok := speed.GetMovement(ClimbMovement)
	if !ok || climb != 20 {
		t.Fatalf("expected climb speed (20, true), got (%d, %t)", climb, ok)
	}

	fly, ok := speed.GetMovement(FlyMovement)
	if !ok || fly != 60 {
		t.Fatalf("expected fly speed (60, true), got (%d, %t)", fly, ok)
	}

	swim, ok := speed.GetMovement(SwimMovement)
	if !ok || swim != 15 {
		t.Fatalf("expected swim speed (15, true), got (%d, %t)", swim, ok)
	}

	maneuverability, ok := speed.GetFlyManeuverability()
	if !ok || maneuverability != GoodFlyManeuverability {
		t.Fatalf("expected fly maneuverability (%q, true), got (%q, %t)", GoodFlyManeuverability, maneuverability, ok)
	}
}

func TestNewSpeed_AllowsImmovableProfile(t *testing.T) {
	speed, ok := NewSpeed(0, 0, 0, 0, 0, "")
	if !ok {
		t.Fatal("expected immovable speed profile to be constructed")
	}

	if !speed.IsImmovable() {
		t.Fatal("expected all-null movement profile to be immovable")
	}

	if _, ok := speed.GetMovement(LandMovement); ok {
		t.Fatal("expected immovable profile to have no land speed")
	}

	if _, ok := speed.GetFlyManeuverability(); ok {
		t.Fatal("expected immovable profile to have no fly maneuverability")
	}
}

func TestNewSpeed_RejectsInvalidFlyConfiguration(t *testing.T) {
	if _, ok := NewSpeed(30, 0, 0, 60, 0, ""); ok {
		t.Fatal("expected fly speed without maneuverability to be rejected")
	}

	if _, ok := NewSpeed(30, 0, 0, 0, 0, AverageFlyManeuverability); ok {
		t.Fatal("expected maneuverability without fly speed to be rejected")
	}
}

func TestSpeed_GetMovement_UsesMovementTypeEnum(t *testing.T) {
	speed, ok := NewSpeed(30, 10, 0, 0, 15, "")
	if !ok {
		t.Fatal("expected speed profile to be constructed")
	}

	land, ok := speed.GetMovement(LandMovement)
	if !ok || land != 30 {
		t.Fatalf("expected land movement (30, true), got (%d, %t)", land, ok)
	}

	if _, ok := speed.GetMovement(FlyMovement); ok {
		t.Fatal("expected fly movement to be absent")
	}

	if _, ok := speed.GetMovement(MovementType("Teleport")); ok {
		t.Fatal("expected unknown movement type to be rejected")
	}
}

func TestNewSpeed_RejectsNegativeMovementValues(t *testing.T) {
	if _, ok := NewSpeed(-5, 0, 0, 0, 0, ""); ok {
		t.Fatal("expected negative land speed to be rejected")
	}

	if _, ok := NewSpeed(30, -10, 0, 0, 0, ""); ok {
		t.Fatal("expected negative burrow speed to be rejected")
	}
}
