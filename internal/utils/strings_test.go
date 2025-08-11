package utils

import "testing"

func TestContainsString(t *testing.T) {
	tests := []struct {
		name   string
		slice  []string
		target string
		want   bool
	}{
		{"present", []string{"a", "b", "c"}, "b", true},
		{"absent", []string{"a", "b", "c"}, "d", false},
		{"empty slice", nil, "a", false},
		{"empty target", []string{""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.slice, tt.target); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateStrings(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		want   []string
	}{
		{"no dups", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"with dups", []string{"b", "a", "b", "c", "a"}, []string{"a", "b", "c"}},
		{"empty", nil, nil},
		{"single", []string{"x"}, []string{"x"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeduplicateStrings(tt.values)
			if len(got) != len(tt.want) {
				t.Fatalf("DeduplicateStrings() length = %d, want %d (%v)", len(got), len(tt.want), got)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("DeduplicateStrings()[%d] = %q, want %q (%v)", i, got[i], tt.want[i], got)
				}
			}
		})
	}
}

func TestEqualFoldInSlice(t *testing.T) {
	if !EqualFoldInSlice([]string{"Hello", "World"}, "world") {
		t.Errorf("EqualFoldInSlice should be true for case-insensitive match")
	}
	if EqualFoldInSlice([]string{"Hello", "World"}, "mars") {
		t.Errorf("EqualFoldInSlice should be false for non-existing value")
	}
}
