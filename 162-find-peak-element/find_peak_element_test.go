package find_peak_element

import "testing"

func TestFindPeakElement(t *testing.T) {
	ts := []struct {
		Input []int // 输入
		Wants []int // 可能有多个下标，结果满足其中一个即可
	}{

		{[]int{1, 2, 3, 1}, []int{2}},
		{[]int{1, 2, 1, 3, 5, 6, 4}, []int{1, 5}},
		{[]int{1, 2}, []int{1}},
		{[]int{2, 1}, []int{0}},
		{[]int{1}, []int{0}},
		{[]int{1, 2, 1}, []int{1}},
	}

	for _, tc := range ts {
		got := findPeakElement(tc.Input)
		if containsValue(tc.Wants, got) == false {
			t.Errorf("input:%v, peak:%v, got:%v", tc.Input, tc.Wants, got)
		}
	}
}

func containsValue(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}
