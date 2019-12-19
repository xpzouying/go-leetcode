package maximum_subarray

import "math"

// func maxSubArray(nums []int) int {
// 	if len(nums) == 0 {
// 		return math.MinInt32
// 	}

// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	DP := make([]int, len(nums)+1)

// 	DP[0] = math.MinInt32
// 	max := math.MinInt32

// 	for i := 0; i < len(nums); i++ {
// 		n1 := DP[i] + nums[i]

// 		if n1 > nums[i] {
// 			DP[i+1] = n1
// 		} else {
// 			DP[i+1] = nums[i]
// 		}

// 		if DP[i+1] > max {
// 			max = DP[i+1]
// 		}
// 	}

// 	return max
// }

// 内存优化版本
// DP表格优化为只记录前面一个DP值
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}

	if len(nums) == 1 {
		return nums[0]
	}

	preDP := math.MinInt32
	max := math.MinInt32

	for i := 0; i < len(nums); i++ {
		n1 := preDP + nums[i]

		if n1 > nums[i] {
			// 如果累计当前元素更大
			if n1 > max {
				max = n1
			}

			preDP = n1
		} else {
			// 如果只取当前元素更大
			if nums[i] > max {
				max = nums[i]
			}
			preDP = nums[i]
		}
	}

	return max
}
