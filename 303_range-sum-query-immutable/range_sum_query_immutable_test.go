package range_sum_query_immutable

import "testing"

func TestSumRange(t *testing.T) {

	ts := []struct {
		nums []int
		i    int
		j    int
		want int
	}{

		{
			nums: []int{1},
			i:    0,
			j:    0,
			want: 1,
		},
		{
			nums: []int{1, 1, 2, 5, 10},
			i:    0,
			j:    3,
			want: 9,
		},
		{
			nums: []int{1, 1, 2, 5, 10},
			i:    1,
			j:    3,
			want: 8,
		},
		{
			nums: []int{1, 1, 2, 5, 10},
			i:    2,
			j:    4,
			want: 17,
		},
	}

	for _, tc := range ts {
		obj := Constructor(tc.nums)
		got := obj.SumRange(tc.i, tc.j)
		if tc.want != got {

			t.Errorf("nums: %+v, range: (%d, %d), want=%d, got=%d",
				tc.nums, tc.i, tc.j, tc.want, got)
		}

	}

}
