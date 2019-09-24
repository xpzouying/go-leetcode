package maximum_product_subarray

import "testing"

func TestMaxProduct(t *testing.T) {

	ts := []struct {
		input []int
		want  int
	}{
		{
			input: []int{2, 3, -2, 4},
			want:  6,
		},
		{
			input: []int{-2, 0, -1},
			want:  0,
		},
		{
			input: []int{-1, 1, -2},
			want:  2,
		},
		{
			input: []int{-1, 2, -2, -3, 0},
			want:  12,
		},
	}

	for _, tc := range ts {
		got := maxProduct(tc.input)
		if tc.want != got {
			t.Errorf("input:%+v, want:%d, got:%d", tc.input, tc.want, got)
		}
	}

}
