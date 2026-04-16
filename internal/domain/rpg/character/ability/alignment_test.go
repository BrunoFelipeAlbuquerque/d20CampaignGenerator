package ability

import "testing"

func TestNewAlignment_UsesValidAxes(t *testing.T) {
	alignment, ok := NewAlignment(OrderLawful, MoralityGood)
	if !ok {
		t.Fatal("expected alignment to be constructed")
	}

	orderAxis, moralityAxis := alignment.GetAlignment()
	if orderAxis != OrderLawful || moralityAxis != MoralityGood {
		t.Fatalf("expected (%q, %q), got (%q, %q)", OrderLawful, MoralityGood, orderAxis, moralityAxis)
	}
}

func TestNewAlignment_InvalidValuesAreRejected(t *testing.T) {
	if _, ok := NewAlignment(OrderAxis("Sideways"), MoralityGood); ok {
		t.Fatal("expected invalid alignment to be rejected")
	}
}

func TestAlignmentGetAlignmentName_ReturnsNeutralForTrueNeutral(t *testing.T) {
	alignment, ok := NewAlignment(OrderNeutral, MoralityNeutral)
	if !ok {
		t.Fatal("expected alignment to be constructed")
	}

	if got := alignment.GetAlignmentName(); got != "Neutral" {
		t.Fatalf("expected Neutral, got %q", got)
	}
}

func TestAlignmentGetAlignmentName_FormatsNonNeutralPairs(t *testing.T) {
	alignment, ok := NewAlignment(OrderChaotic, MoralityGood)
	if !ok {
		t.Fatal("expected alignment to be constructed")
	}

	if got := alignment.GetAlignmentName(); got != "Chaotic Good" {
		t.Fatalf("expected Chaotic Good, got %q", got)
	}
}

func TestAlignmentSetAlignment_RejectsInvalidPair(t *testing.T) {
	alignment, ok := NewAlignment(OrderLawful, MoralityGood)
	if !ok {
		t.Fatal("expected alignment to be constructed")
	}

	if ok := alignment.SetAlignment(OrderAxis("Sideways"), MoralityEvil); ok {
		t.Fatal("expected invalid alignment update to be rejected")
	}

	orderAxis, moralityAxis := alignment.GetAlignment()
	if orderAxis != OrderLawful || moralityAxis != MoralityGood {
		t.Fatalf("expected stored alignment to remain unchanged, got (%q, %q)", orderAxis, moralityAxis)
	}
}

func TestAlignmentSetAlignment_RejectsInvalidOrderAxis(t *testing.T) {
	alignment, ok := NewAlignment(OrderLawful, MoralityNeutral)
	if !ok {
		t.Fatal("expected alignment to be constructed")
	}

	if ok := alignment.SetAlignment(OrderAxis("Sideways"), MoralityNeutral); ok {
		t.Fatal("expected invalid order axis to be rejected")
	}

	orderAxis, moralityAxis := alignment.GetAlignment()
	if orderAxis != OrderLawful || moralityAxis != MoralityNeutral {
		t.Fatalf("expected alignment to remain (%q, %q), got (%q, %q)", OrderLawful, MoralityNeutral, orderAxis, moralityAxis)
	}
}

func TestAlignmentSetAlignment_RejectsInvalidMoralityAxis(t *testing.T) {
	alignment, ok := NewAlignment(OrderNeutral, MoralityGood)
	if !ok {
		t.Fatal("expected alignment to be constructed")
	}

	if ok := alignment.SetAlignment(OrderNeutral, MoralityAxis("KindOfGood")); ok {
		t.Fatal("expected invalid morality axis to be rejected")
	}

	orderAxis, moralityAxis := alignment.GetAlignment()
	if orderAxis != OrderNeutral || moralityAxis != MoralityGood {
		t.Fatalf("expected alignment to remain (%q, %q), got (%q, %q)", OrderNeutral, MoralityGood, orderAxis, moralityAxis)
	}
}
