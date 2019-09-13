package majority_element

import "testing"

func TestMajorityElement(t *testing.T) {

	ts := []struct {
		input []int
		want  int
	}{
		{
			input: []int{3, 2, 3},
			want:  3,
		},
		{
			input: []int{2, 2, 1, 1, 1, 2, 2},
			want:  2,
		},
	}

	for _, tc := range ts {
		got := majorityElement(tc.input)
		if tc.want != got {
			t.Errorf("input: %v, want: %d, got: %d",
				tc.input, tc.want, got)
		}
	}
}
