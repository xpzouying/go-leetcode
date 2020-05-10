package masseuse

func massage(nums []int) int {

	/**
	DP[i-2], DP[i-1], DP[i]

	DP[i] = MAX{
		DP[i-2] + nums[i],
		DP[i-1]
	}
	*/

	l := len(nums)

	switch l {
	case 0:
		return 0
	case 1:
		return nums[0]
	// case 2:
	// 	return max(nums[0], nums[1])
	default:
	}

	DP1, DP2 := 0, nums[0]

	for i := 1; i < l; i++ {
		n1 := DP1 + nums[i]
		n2 := DP2

		if n1 > n2 {
			DP1, DP2 = DP2, n1
		} else {
			// 如果相加都不够前一个DP大的话，
			// 则，不用叠加
			DP1, DP2 = DP2, DP2
		}
	}

	return DP2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
