package main

import "testing"

func TestProposition(t *testing.T) {
	for _, tc := range []struct {
		row              string
		a, b, c, d, want bool
	}{
		{"1", false, false, false, false, false},
		{"2", false, false, false, true, false},
		{"3", false, false, true, false, false},
		{"4", false, false, true, true, false},
		{"5", false, true, false, false, true},
		{"6", false, true, false, true, true},
		{"7", false, true, true, false, true},
		{"8", false, true, true, true, false},
		{"9", true, false, false, false, false},
		{"10", true, false, false, true, false},
		{"11", true, false, true, false, false},
		{"12", true, false, true, true, true},
		{"13", true, true, false, false, true},
		{"14", true, true, false, true, true},
		{"15", true, true, true, false, true},
		{"16", true, true, true, true, false},
	} {
		t.Run(tc.row, func(t *testing.T) {
			if got := proposition(tc.a, tc.b, tc.c, tc.d); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
