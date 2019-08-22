package main

import (
	"log"
	"testing"
)

func TestReverse(t *testing.T) {
	ts := []struct {
		Number int
		Want   int
	}{
		{
			Number: 123,
			Want:   321,
		},
		{
			Number: 0,
			Want:   0,
		},
		{
			Number: -123,
			Want:   -321,
		},
		{
			Number: 120,
			Want:   21,
		},
		{
			Number: 901000,
			Want:   109,
		},
		{
			Number: 1534236469,
			Want:   0,
		},
	}

	for _, tc := range ts {
		got := reverse(tc.Number)

		log.Printf("number=%d, want=%d, got=%d\n",
			tc.Number, tc.Want, got)

		if got != tc.Want {
			t.Errorf("number=%d, want=%d, got=%d",
				tc.Number, tc.Want, got)
		}
	}
}
