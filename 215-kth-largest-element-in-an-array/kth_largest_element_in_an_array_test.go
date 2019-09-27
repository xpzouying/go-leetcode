package kth_largest_element_in_an_array

import "testing"

func TestFindKthLargest(t *testing.T) {

	ts := []struct {
		input []int
		k     int
		want  int
	}{
		{input: []int{3, 2, 1, 5, 6, 4}, k: 2, want: 5},
		{input: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, want: 4},
		{input: []int{1}, k: 1, want: 1},
		{input: []int{2, 1}, k: 2, want: 1},
		{input: []int{-1, 2, 0}, k: 3, want: -1},
		{input: []int{3, 1, 2, 4}, k: 2, want: 3},
		{input: []int{-1, 2, 0}, k: 2, want: 0},
		{input: []int{5, 2, 4, 1, 3, 6, 0}, k: 4, want: 3},
	}

	for _, tc := range ts {
		got := findKthLargest(tc.input, tc.k)
		if tc.want != got {
			t.Errorf("input:%v, k:%d, want:%d, got:%d",
				tc.input, tc.k, tc.want, got)
		}
	}
}
