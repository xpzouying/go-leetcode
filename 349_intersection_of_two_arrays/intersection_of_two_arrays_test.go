package intersection_of_two_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {

	ts := []struct {
		Nums1 []int
		Nums2 []int
		Want  []int
	}{
		{Nums1: []int{1, 2, 2, 1}, Nums2: []int{2, 2}, Want: []int{2}},
	}

	for _, tc := range ts {
		got := intersection(tc.Nums1, tc.Nums2)
		assert.Equal(t, tc.Want, got)
	}

}

/**

解体思路：

题目要求：

0. 求交集
1. 要求每个元素是唯一的；
2. 可以不考虑结果输出的顺序


步骤：

1. 对两个数列进行排序；
2. 使用双指针，分别指向数列A和数列B
3. 向后移动指针，
	3.1 找到数值一样的，则输出；两个指针同时移动到 下一个新的数值上
	3.2 不一样，则将指向较小值的指针向后移动；

*/
