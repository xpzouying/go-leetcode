package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRepeatNumber(t *testing.T) {
	ts := []struct {
		nums []int

		wantContains []int // 只需要结果在其中任意一个
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, []int{2, 3}},
	}

	for _, tc := range ts {
		got := findRepeatNumber(tc.nums)

		// 校验获取的结果是否在预期之中。
		contains := false
		for _, want := range tc.wantContains {
			if want == got {
				contains = true
				break
			}
		}

		assert.Equal(t, true, contains)
	}
}
