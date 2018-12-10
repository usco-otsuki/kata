package main

import (
	"testing"
)

func TestOverlap(t *testing.T) {
	tests := []struct {
		f1, f2 fabric
		want   bool
	}{
		{
			fabric{1, 0, 0, 2, 2},
			fabric{2, 1, 1, 3, 3},
			true,
		},
		{
			fabric{1, 0, 0, 2, 2},
			fabric{2, 3, 3, 3, 3},
			false,
		},
		{
			fabric{1, 1, 3, 4, 4},
			fabric{2, 5, 5, 2, 2},
			false,
		},
	}

	for _, test := range tests {
		if result := overlap(test.f1, test.f2); result != test.want {
			t.Errorf("overlap(%v, %v) = %t, want %t\n", test.f1, test.f2, result, test.want)
		}

	}
}
