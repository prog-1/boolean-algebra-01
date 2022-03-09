ppackage main

import "testing"

func TestProposition(t *testing.T) {
	for _, tc := range []struct {
		a, b, c, d bool
		want       bool
	}{

		{false, false, false, false, false},
		{false, false, false, true, false},
		{false, false, true, false, false},
		{false, false, true, true, false},
		{false, true, false, false, true},
		{false, true, false, true, true},
		{false, true, true, false, true},
		{false, true, true, true, false},
		{true, false, false, false, false},
		{true, false, false, true, false},
		{true, false, true, false, false},
		{true, false, true, true, true},
		{true, true, false, false, true},
		{true, true, false, true, true},
		{true, true, true, false, false},
		{true, true, true, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := proposition(tc.a, tc.b, tc.c, tc.d); got != tc.want {
				t.Errorf("Proposition(%v, %v, %v, %v): got = %v, want = %v", tc.a, tc.b, tc.c, tc.d, got, tc.want)
			}
		})
	}
}
