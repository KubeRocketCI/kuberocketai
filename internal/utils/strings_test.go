package utils

import (
	"testing"
)

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
