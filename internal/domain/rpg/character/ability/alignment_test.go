package ability

import "testing"

func TestNewAlignment_UsesValidAxes(t *testing.T) {
	alignment := NewAlignment(OrderLawful, MoralityGood)

	orderAxis, moralityAxis := alignment.GetAlignment()
	if orderAxis != OrderLawful || moralityAxis != MoralityGood {
		t.Fatalf("expected (%q, %q), got (%q, %q)", OrderLawful, MoralityGood, orderAxis, moralityAxis)
	}
}

func TestNewAlignment_InvalidValuesProduceZeroAlignment(t *testing.T) {
	alignment := NewAlignment(OrderAxis("Sideways"), MoralityGood)

	orderAxis, moralityAxis := alignment.GetAlignment()
	if orderAxis != "" || moralityAxis != "" {
		t.Fatalf("expected zero alignment, got (%q, %q)", orderAxis, moralityAxis)
	}
}

func TestAlignmentGetAlignmentName_ReturnsNeutralForTrueNeutral(t *testing.T) {
	alignment := NewAlignment(OrderNeutral, MoralityNeutral)

	if got := alignment.GetAlignmentName(); got != "Neutral" {
		t.Fatalf("expected Neutral, got %q", got)
	}
}

func TestAlignmentGetAlignmentName_FormatsNonNeutralPairs(t *testing.T) {
	alignment := NewAlignment(OrderChaotic, MoralityGood)

	if got := alignment.GetAlignmentName(); got != "Chaotic Good" {
		t.Fatalf("expected Chaotic Good, got %q", got)
	}
}

func TestAlignmentSetAlignment_RejectsInvalidPair(t *testing.T) {
	alignment := NewAlignment(OrderLawful, MoralityGood)

	if ok := alignment.SetAlignment(OrderAxis("Sideways"), MoralityEvil); ok {
		t.Fatal("expected invalid alignment update to be rejected")
	}

	orderAxis, moralityAxis := alignment.GetAlignment()
	if orderAxis != OrderLawful || moralityAxis != MoralityGood {
		t.Fatalf("expected stored alignment to remain unchanged, got (%q, %q)", orderAxis, moralityAxis)
	}
}

func TestAlignmentSetOrderAxis_RejectsInvalidValue(t *testing.T) {
	alignment := NewAlignment(OrderLawful, MoralityNeutral)

	if ok := alignment.SetOrderAxis(OrderAxis("Sideways")); ok {
		t.Fatal("expected invalid order axis to be rejected")
	}

	if alignment.GetOrderAxis() != OrderLawful {
		t.Fatalf("expected order axis to remain %q, got %q", OrderLawful, alignment.GetOrderAxis())
	}
}

func TestAlignmentSetMoralityAxis_RejectsInvalidValue(t *testing.T) {
	alignment := NewAlignment(OrderNeutral, MoralityGood)

	if ok := alignment.SetMoralityAxis(MoralityAxis("KindOfGood")); ok {
		t.Fatal("expected invalid morality axis to be rejected")
	}

	if alignment.GetMoralityAxis() != MoralityGood {
		t.Fatalf("expected morality axis to remain %q, got %q", MoralityGood, alignment.GetMoralityAxis())
	}
}
