package ability

import "testing"

func TestNewSpeed_AllowsMixedMovementAndFlyManeuverability(t *testing.T) {
	speed, ok := NewSpeed(30, 10, 20, 60, 15, GoodFlyManeuverability)
	if !ok {
		t.Fatal("expected speed profile to be constructed")
	}

	land, ok := speed.GetLand()
	if !ok || land != 30 {
		t.Fatalf("expected land speed (30, true), got (%d, %t)", land, ok)
	}

	burrow, ok := speed.GetBurrow()
	if !ok || burrow != 10 {
		t.Fatalf("expected burrow speed (10, true), got (%d, %t)", burrow, ok)
	}

	climb, ok := speed.GetClimb()
	if !ok || climb != 20 {
		t.Fatalf("expected climb speed (20, true), got (%d, %t)", climb, ok)
	}

	fly, ok := speed.GetFly()
	if !ok || fly != 60 {
		t.Fatalf("expected fly speed (60, true), got (%d, %t)", fly, ok)
	}

	swim, ok := speed.GetSwim()
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

	if _, ok := speed.GetLand(); ok {
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

	if speed.HasMovement(FlyMovement) {
		t.Fatal("expected fly movement to be absent")
	}

	if _, ok := speed.GetMovement(MovementType("Teleport")); ok {
		t.Fatal("expected unknown movement type to be rejected")
	}
}

func TestSpeed_SetMovement_AllowsAddingAndRemovingNonFlyModes(t *testing.T) {
	speed, ok := NewSpeed(30, 0, 0, 0, 0, "")
	if !ok {
		t.Fatal("expected speed profile to be constructed")
	}

	if ok := speed.SetMovement(ClimbMovement, 20); !ok {
		t.Fatal("expected climb speed to be added")
	}

	climb, ok := speed.GetClimb()
	if !ok || climb != 20 {
		t.Fatalf("expected climb speed (20, true), got (%d, %t)", climb, ok)
	}

	if ok := speed.SetMovement(ClimbMovement, 0); !ok {
		t.Fatal("expected climb speed to be removable")
	}

	if _, ok := speed.GetClimb(); ok {
		t.Fatal("expected climb speed to be absent after removal")
	}
}

func TestSpeed_SetFly_RequiresExplicitManeuverability(t *testing.T) {
	speed, ok := NewSpeed(30, 0, 0, 0, 0, "")
	if !ok {
		t.Fatal("expected speed profile to be constructed")
	}

	if ok := speed.SetFly(60, AverageFlyManeuverability); !ok {
		t.Fatal("expected fly speed with maneuverability to be accepted")
	}

	fly, ok := speed.GetFly()
	if !ok || fly != 60 {
		t.Fatalf("expected fly speed (60, true), got (%d, %t)", fly, ok)
	}

	maneuverability, ok := speed.GetFlyManeuverability()
	if !ok || maneuverability != AverageFlyManeuverability {
		t.Fatalf("expected maneuverability (%q, true), got (%q, %t)", AverageFlyManeuverability, maneuverability, ok)
	}

	if ok := speed.SetFly(0, AverageFlyManeuverability); ok {
		t.Fatal("expected removing fly speed without clearing maneuverability to be rejected")
	}

	if ok := speed.SetFly(0, ""); !ok {
		t.Fatal("expected fly speed removal with cleared maneuverability to be accepted")
	}

	if _, ok := speed.GetFly(); ok {
		t.Fatal("expected fly speed to be absent after removal")
	}
}

func TestSpeed_SetMovement_RejectsInvalidInput(t *testing.T) {
	speed, ok := NewSpeed(30, 0, 0, 0, 0, "")
	if !ok {
		t.Fatal("expected speed profile to be constructed")
	}

	if ok := speed.SetMovement(MovementType("Teleport"), 30); ok {
		t.Fatal("expected invalid movement type to be rejected")
	}

	if ok := speed.SetMovement(LandMovement, -5); ok {
		t.Fatal("expected negative movement speed to be rejected")
	}

	if ok := speed.SetMovement(FlyMovement, 60); ok {
		t.Fatal("expected fly movement updates without maneuverability to be rejected")
	}
}
