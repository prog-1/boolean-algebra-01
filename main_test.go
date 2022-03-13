package main

import "testing"

func TestProposition(t *testing.T) {
	for _, tc := range []struct {
		name       string
		a, b, c, d bool
		want       bool
	}{

		{"test", false, false, false, false, false},
		{"test", false, false, false, true, false},
		{"test", false, false, true, false, false},
		{"test", false, false, true, true, false},
		{"test", false, true, false, false, true},
		{"test", false, true, false, true, false},
		{"test", false, true, true, false, true},
		{"test", false, true, true, true, false},
		{"test", true, false, false, false, false},
		{"test", true, false, false, true, false},
		{"test", true, false, true, false, false},
		{"test", true, false, true, true, false},
		{"test", true, true, false, false, false},
		{"test", true, true, false, true, false},
		{"test", true, true, true, false, false},
		{"test", true, true, true, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := proposition(tc.a, tc.b, tc.c, tc.d); got != tc.want {
				t.Errorf("Proposition(%v, %v, %v, %v): got = %v, want = %v", tc.a, tc.b, tc.c, tc.d, got, tc.want)
			}
		})
	}
}
