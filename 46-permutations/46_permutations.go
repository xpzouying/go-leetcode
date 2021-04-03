// Package permutations - 46. 全排列
//
// 难度：`中等`
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

package permutations

var (
	res = [][]int{}
)

func permute(nums []int) [][]int {

	used := make([]bool, len(nums))

	backtrack(nums, used, 0, []int{})

	return res
}

// index - 当前的开始index
// path 表示当前的路径
func backtrack(nums []int, used []bool, index int, path []int) {

	allUsed := true
	for _, u := range used {
		if !u {
			allUsed = false
		}
	}
	if allUsed {
		res = append(res,
			append([]int{}, path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		// 已经使用过
		if used[i] {
			continue
		}

		// 1. 选择
		n := nums[i]
		path = append(path, n)
		used[i] = true

		// 2. 遍历
		backtrack(nums, used, 0, path)

		// 3. 回溯
		path = path[:len(path)-1]
		used[i] = false
	}
}

/*

思路笔记本：

1. 全排列，必然使用回溯。

2. 由于排列出来的结果有顺序之分，所以对于每次递归遍历的时候都需要从0开始。
既然从0开始的话，那么就需要记录哪些数字被遍历过了。

3.
遍历过的路径：记录下来。
可选择的路径：从0开始，使用一个used记录哪些数字已经走过。

*/
