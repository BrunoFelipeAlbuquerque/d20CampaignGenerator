package ability

type movementType string
type MovementType = movementType

type flyManeuverability string
type FlyManeuverability = flyManeuverability

type speed struct {
	burrow             int
	climb              int
	fly                int
	land               int
	swim               int
	flyManeuverability flyManeuverability
}
type Speed = speed

const (
	BurrowMovement MovementType = "Burrow"
	ClimbMovement  MovementType = "Climb"
	FlyMovement    MovementType = "Fly"
	LandMovement   MovementType = "Land"
	SwimMovement   MovementType = "Swim"
)

const (
	ClumsyFlyManeuverability  FlyManeuverability = "Clumsy"
	PoorFlyManeuverability    FlyManeuverability = "Poor"
	AverageFlyManeuverability FlyManeuverability = "Average"
	GoodFlyManeuverability    FlyManeuverability = "Good"
	PerfectFlyManeuverability FlyManeuverability = "Perfect"
)

func NewSpeed(
	land int,
	burrow int,
	climb int,
	fly int,
	swim int,
	maneuverability FlyManeuverability,
) (Speed, bool) {
	profile := speed{
		burrow: burrow,
		climb:  climb,
		fly:    fly,
		land:   land,
		swim:   swim,
	}

	if !profile.isValidBaseSpeeds() || !isValidFlyConfiguration(fly, maneuverability) {
		return speed{}, false
	}

	profile.flyManeuverability = maneuverability
	return profile, true
}

func (s speed) GetMovement(kind MovementType) (int, bool) {
	switch kind {
	case BurrowMovement:
		return getOptionalSpeedValue(s.burrow)
	case ClimbMovement:
		return getOptionalSpeedValue(s.climb)
	case FlyMovement:
		return getOptionalSpeedValue(s.fly)
	case LandMovement:
		return getOptionalSpeedValue(s.land)
	case SwimMovement:
		return getOptionalSpeedValue(s.swim)
	default:
		return 0, false
	}
}

func (s speed) GetFlyManeuverability() (FlyManeuverability, bool) {
	if s.fly <= 0 || !isValidFlyManeuverability(s.flyManeuverability) {
		return "", false
	}

	return s.flyManeuverability, true
}

func (s speed) IsImmovable() bool {
	return s.land == 0 && s.burrow == 0 && s.climb == 0 && s.fly == 0 && s.swim == 0
}

func (s speed) isValidBaseSpeeds() bool {
	return s.land >= 0 && s.burrow >= 0 && s.climb >= 0 && s.fly >= 0 && s.swim >= 0
}

func getOptionalSpeedValue(value int) (int, bool) {
	if value <= 0 {
		return 0, false
	}

	return value, true
}

func isValidFlyManeuverability(value FlyManeuverability) bool {
	switch value {
	case ClumsyFlyManeuverability, PoorFlyManeuverability, AverageFlyManeuverability, GoodFlyManeuverability, PerfectFlyManeuverability:
		return true
	default:
		return false
	}
}

func isValidFlyConfiguration(fly int, maneuverability FlyManeuverability) bool {
	if fly < 0 {
		return false
	}

	if fly == 0 {
		return maneuverability == ""
	}

	return isValidFlyManeuverability(maneuverability)
}
