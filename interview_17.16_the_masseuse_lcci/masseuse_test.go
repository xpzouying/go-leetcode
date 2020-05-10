package masseuse

import "testing"

func TestMassage(t *testing.T) {

	ts := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 4, 5, 3, 1, 1, 3}, 12},
	}

	for _, tc := range ts {
		n := tc.nums
		got := massage(n)
		if tc.want != got {
			t.Errorf("nums=%v, want=%d, got=%d", n, tc.want, got)
		}
	}

}
