package merge_two_sorted_lists

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	ts := []struct {
		L1   []int
		L2   []int
		Want []int
	}{
		{
			L1:   []int{1, 2, 4},
			L2:   []int{1, 3, 4},
			Want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			L1:   []int{},
			L2:   []int{},
			Want: []int{},
		},
		{
			L1:   []int{1},
			L2:   []int{},
			Want: []int{1},
		},
		{
			L1:   []int{},
			L2:   []int{1},
			Want: []int{1},
		},
	}

	for _, tc := range ts {
		l1 := NewList(tc.L1)
		l2 := NewList(tc.L2)
		log.Printf("l1: %v\n", l1.Lists())
		log.Printf("l2: %v\n", l2.Lists())

		got := mergeTwoLists(l1, l2)
		log.Printf("l3: %v\n", got.Lists())

		if !equalList(got.Lists(), tc.Want) {
			t.Errorf("want: %v, got: %v", tc.Want, got.Lists())
		}
	}
}
