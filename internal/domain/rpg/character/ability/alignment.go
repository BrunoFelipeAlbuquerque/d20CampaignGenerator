package ability

const (
	// ORDER AXIS
	OrderLawful  OrderAxis = "Lawful"
	OrderNeutral OrderAxis = "Neutral"
	OrderChaotic OrderAxis = "Chaotic"

	// MORALITY AXIS
	MoralityGood    MoralityAxis = "Good"
	MoralityNeutral MoralityAxis = "Neutral"
	MoralityEvil    MoralityAxis = "Evil"
)

type OrderAxis string
type MoralityAxis string

type alignment struct {
	orderAxis    OrderAxis
	moralityAxis MoralityAxis
}
type Alignment = alignment

func NewAlignment(orderAxis OrderAxis, moralityAxis MoralityAxis) (Alignment, bool) {
	value := alignment{}
	if !value.SetAlignment(orderAxis, moralityAxis) {
		return alignment{}, false
	}

	return value, true
}

func (a alignment) GetAlignmentName() string {
	if a.orderAxis == OrderNeutral && a.moralityAxis == MoralityNeutral {
		return string(OrderNeutral)
	}

	return string(a.orderAxis) + " " + string(a.moralityAxis)
}

func (a alignment) GetAlignment() (OrderAxis, MoralityAxis) {
	return a.orderAxis, a.moralityAxis
}

func (a *alignment) SetAlignment(orderAxis OrderAxis, moralityAxis MoralityAxis) bool {
	if !isValidOrderAxis(orderAxis) || !isValidMoralityAxis(moralityAxis) {
		return false
	}

	a.orderAxis = orderAxis
	a.moralityAxis = moralityAxis
	return true
}

func (a alignment) GetOrderAxis() OrderAxis {
	return a.orderAxis
}

func (a *alignment) SetOrderAxis(orderAxis OrderAxis) bool {
	if !isValidOrderAxis(orderAxis) {
		return false
	}

	a.orderAxis = orderAxis
	return true
}

func (a alignment) GetMoralityAxis() MoralityAxis {
	return a.moralityAxis
}

func (a *alignment) SetMoralityAxis(moralityAxis MoralityAxis) bool {
	if !isValidMoralityAxis(moralityAxis) {
		return false
	}

	a.moralityAxis = moralityAxis
	return true
}

func isValidOrderAxis(orderAxis OrderAxis) bool {
	switch orderAxis {
	case OrderLawful, OrderNeutral, OrderChaotic:
		return true
	default:
		return false
	}
}

func isValidMoralityAxis(moralityAxis MoralityAxis) bool {
	switch moralityAxis {
	case MoralityGood, MoralityNeutral, MoralityEvil:
		return true
	default:
		return false
	}
}
