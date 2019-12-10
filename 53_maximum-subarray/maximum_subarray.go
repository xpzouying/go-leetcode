package maximum_subarray

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	DP := make([]int, len(nums)+1)

	DP[0] = math.MinInt32
	max := math.MinInt32

	for i := 0; i < len(nums); i++ {
		n1 := DP[i] + nums[i]

		if n1 > nums[i] {
			DP[i+1] = n1
		} else {
			DP[i+1] = nums[i]
		}

		if DP[i+1] > max {
			max = DP[i+1]
		}
	}

	return max
}
