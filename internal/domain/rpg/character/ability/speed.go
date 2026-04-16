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

func (s speed) GetBurrow() (int, bool) {
	return getOptionalSpeedValue(s.burrow)
}

func (s speed) GetClimb() (int, bool) {
	return getOptionalSpeedValue(s.climb)
}

func (s speed) GetFly() (int, bool) {
	return getOptionalSpeedValue(s.fly)
}

func (s speed) GetLand() (int, bool) {
	return getOptionalSpeedValue(s.land)
}

func (s speed) GetSwim() (int, bool) {
	return getOptionalSpeedValue(s.swim)
}

func (s speed) GetMovement(kind MovementType) (int, bool) {
	switch kind {
	case BurrowMovement:
		return s.GetBurrow()
	case ClimbMovement:
		return s.GetClimb()
	case FlyMovement:
		return s.GetFly()
	case LandMovement:
		return s.GetLand()
	case SwimMovement:
		return s.GetSwim()
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

func (s speed) HasMovement(kind MovementType) bool {
	_, ok := s.GetMovement(kind)
	return ok
}

func (s speed) IsImmovable() bool {
	return s.land == 0 && s.burrow == 0 && s.climb == 0 && s.fly == 0 && s.swim == 0
}

func (s *speed) SetMovement(kind MovementType, value int) bool {
	if value < 0 {
		return false
	}

	switch kind {
	case BurrowMovement:
		s.burrow = value
	case ClimbMovement:
		s.climb = value
	case FlyMovement:
		if !isValidFlyConfiguration(value, s.flyManeuverability) {
			return false
		}
		s.fly = value
	case LandMovement:
		s.land = value
	case SwimMovement:
		s.swim = value
	default:
		return false
	}

	return true
}

func (s *speed) SetFly(value int, maneuverability FlyManeuverability) bool {
	if !isValidFlyConfiguration(value, maneuverability) {
		return false
	}

	s.fly = value
	s.flyManeuverability = maneuverability
	return true
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

func isValidMovementType(value MovementType) bool {
	switch value {
	case BurrowMovement, ClimbMovement, FlyMovement, LandMovement, SwimMovement:
		return true
	default:
		return false
	}
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
