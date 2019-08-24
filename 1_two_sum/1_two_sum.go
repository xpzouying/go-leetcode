// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package two_sum

// --- 暴力搜索 ---
// 时间复杂度为 n^2
// func twoSum(nums []int, target int) []int {
// 	if len(nums) < 2 {
// 		return []int{}
// 	}

// 	for i := 0; i < len(nums)-1; i++ {

// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	// m: 记录搜索过的数值及其下标
	// key: value in nums
	// value: index of number in nums
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		wantNum := target - nums[i]

		if v, ok := m[wantNum]; ok {
			// 如果找到之前记录的、且匹配的，返回
			return []int{v, i}
		} else {
			// 没有找到，则记录到哈希表
			m[nums[i]] = i
		}
	}

	return []int{}
}
