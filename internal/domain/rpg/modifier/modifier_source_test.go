package modifier

import (
	"testing"
)

// ==============================
// VALIDATION
// ==============================

func TestValidateModifierSource(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"flanking", true},
		{"higher_ground", true},
		{"a", true},
		{"a1_b2", true},

		{"", false},
		{"Flanking", false},
		{"flanking-bonus", false},
		{"flanking bonus", false},
		{"flanking__", false},
		{"_flanking", false},
		{"flanking_", false},
		{"😏", false},
		{"§", false},
	}

	for _, tt := range tests {
		err := validateModifierSource(ModifierSource(tt.input))
		if tt.valid && err != nil {
			t.Errorf("expected valid for %q, got error: %v", tt.input, err)
		}
		if !tt.valid && err == nil {
			t.Errorf("expected error for %q, got none", tt.input)
		}
	}
}

// ==============================
// NORMALIZATION
// ==============================

func TestNormalizeModifierSource(t *testing.T) {
	tests := map[string]ModifierSource{
		"Flanking":        "flanking",
		" higher ground ": "higher_ground",
		"soft-cover":      "soft_cover",
		"Mixed Case-Test": "mixed_case_test",
	}

	for input, expected := range tests {
		got := NormalizeModifierSource(input)
		if got != expected {
			t.Errorf("normalize failed: input=%q expected=%q got=%q", input, expected, got)
		}
	}
}

// ==============================
// REGISTRY BASICS
// ==============================

func TestRegistry_DefaultEntriesExist(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	if !r.IsKnown("flanking") {
		t.Error("expected flanking to exist")
	}

	if !r.IsKnown("soft_cover") {
		t.Error("expected soft_cover to exist")
	}

	if r.IsKnown("nonexistent") {
		t.Error("did not expect nonexistent to exist")
	}
}

func TestRegistry_Get(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	info, ok := r.Get("flanking")
	if !ok {
		t.Fatal("expected flanking to be found")
	}

	if info.Description == "" {
		t.Error("expected description to be set")
	}
}

// ==============================
// REGISTER SUCCESS
// ==============================

func TestRegistry_RegisterSuccess(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	err := r.Register("new_source", "Some description")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !r.IsKnown("new_source") {
		t.Error("expected new_source to be registered")
	}
}

// ==============================
// DUPLICATE
// ==============================

func TestRegistry_RegisterDuplicate(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	err := r.Register("flanking", "Duplicate")
	if err == nil {
		t.Error("expected duplicate error")
	}
}

// ==============================
// INVALID FORMAT
// ==============================

func TestRegistry_RegisterInvalidFormat(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	invalids := []ModifierSource{
		"",
		"Flanking",
		"flanking bonus",
		"flanking-bonus",
		"😏",
	}

	for _, id := range invalids {
		err := r.Register(id, "desc")
		if err == nil {
			t.Errorf("expected error for invalid id %q", id)
		}
	}
}

// ==============================
// LEVENSHTEIN COLLISION
// ==============================

func TestRegistry_RegisterSimilar(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	// flanking already exists
	err := r.Register("flankng", "typo")
	if err == nil {
		t.Error("expected similarity error")
	}
}

// ==============================
// EDGE: CLOSE BUT VALID
// ==============================

func TestRegistry_RegisterCloseButValid(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	// distance > 1 → should pass
	err := r.Register("flank_attack", "different concept")
	if err != nil {
		t.Errorf("expected success, got error: %v", err)
	}
}

// ==============================
// INTERNAL CONSISTENCY
// ==============================

func TestRegistry_InternalMapIsolation(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	// ensure map is not nil
	if r.sources == nil {
		t.Fatal("sources map should not be nil")
	}

	// ensure no accidental overwrite
	before := len(r.sources)

	_ = r.Register("unique_source_test", "desc")

	after := len(r.sources)

	if after != before+1 {
		t.Errorf("expected map size to grow by 1, got %d -> %d", before, after)
	}
}

// ==============================
// STRESS: MANY INSERTS
// ==============================

func TestRegistry_ManyInsertions(t *testing.T) {
	r := NewDefaultCircumstanceSourceRegistry()

	for i := 0; i < 1000; i++ {
		id := ModifierSource("custom_source_" + string(rune('a'+(i%26))) + string(rune('a'+((i/26)%26))))
		_ = r.Register(id, "desc")
	}

	// just ensure no panic / corruption
	if len(r.sources) == 0 {
		t.Error("registry should not be empty")
	}
}
