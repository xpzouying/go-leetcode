package unique_binary_search_trees

import "testing"

func TestNumTrees(t *testing.T) {
	ts := []struct {
		Input int
		Want  int
	}{
		{
			Input: 3,
			Want:  5,
		},
	}

	for _, tc := range ts {
		got := numTrees(tc.Input)
		if tc.Want != got {
			t.Errorf("input: %d, want: %d, got: %d",
				tc.Input, tc.Want, got)
		}
	}
}
