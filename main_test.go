package main

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
		t.Run("", func(t *testing.T) {
			if got := proposition(tc.a, tc.b, tc.c, tc.d); got != tc.want {
				t.Errorf("proposition(%v,%v,%v,%v) = %v, want %v", tc.a, tc.b, tc.c, tc.d, got, tc.want)
			}
		})
	}
}

var benchA, benchB, benchC, benchD bool = true, false, false, true

func BenchmarkProposition(b *testing.B) {
	for n := 0; n < b.N; n++ {
		proposition(benchA, benchB, benchC, benchD)
	}
}
