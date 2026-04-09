package text

import "testing"

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{"empty strings", "", "", 0},
		{"same strings", "test", "test", 0},
		{"insert", "test", "tests", 1},
		{"delete", "tests", "test", 1},
		{"replace", "test", "tent", 1},
		{"completely different", "abc", "xyz", 3},
		{"kitten-sitting", "kitten", "sitting", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Levenshtein(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestMin(t *testing.T) {
	if min(1, 2, 3) != 1 {
		t.Error("min failed")
	}
	if min(3, 2, 1) != 1 {
		t.Error("min failed")
	}
	if min(2, 1, 3) != 1 {
		t.Error("min failed")
	}
}
