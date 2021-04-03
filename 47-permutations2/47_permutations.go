// Package permut - 47 全排列2
//
//执行结果：通过
// 显示详情
// 执行用时： 4 ms , 在所有 Go 提交中击败了 86.09% 的用户
// 内存消耗： 3.4 MB , 在所有 Go 提交中击败了 100.00% 的用户

package permut

import "sort"

var (
	res = [][]int{}
)

func permuteUnique(nums []int) [][]int {
	defer func() {
		res = res[:0]
	}()

	sort.Ints(nums)

	used := make([]bool, len(nums))
	path := make([]int, 0, len(nums))

	backtrack(nums, 0, used, path)

	return res
}

func backtrack(nums []int, index int, used []bool, path []int) {
	if index == len(nums) {
		res = append(res,
			append([]int{}, path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		// 如果数字已经被使用过了，则跳过。
		if used[i] {
			continue
		}

		// 不是首位的话，需要判断。
		if i > 0 &&
			// 判断一下是否跟前面一位相同的数字。
			// 如果是相同的数字，并且前一个相同的数字已经使用过了
			((nums[i] == nums[i-1]) && used[i-1]) {

			continue
		}

		used[i] = true
		path = append(path, nums[i])

		backtrack(nums, index+1, used, path)

		used[i] = false
		path = path[:len(path)-1]
	}
}
