package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {

	ts := []struct {
		M    int
		N    int
		Want int
	}{
		{
			M:    3,
			N:    2,
			Want: 3,
		},
		{
			M:    2,
			N:    2,
			Want: 2,
		},
		{
			M:    1,
			N:    2,
			Want: 1,
		},
		{
			M:    2,
			N:    1,
			Want: 1,
		},
	}

	for _, tc := range ts {

		got := uniquePaths(tc.M, tc.N)
		if tc.Want != got {
			t.Errorf("m=%d n=%d want=%d got=%d",
				tc.M, tc.N, tc.Want, got)
		}
	}

}
