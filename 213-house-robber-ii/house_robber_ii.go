package house_robber_ii

func rob(nums []int) int {

	switch l := len(nums); l {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
	}

	nums1 := nums[:len(nums)-1]
	nums2 := nums[1:len(nums)]
	// 切分成两个数组
	return max(doDP(nums1), doDP(nums2))
}

func doDP(nums []int) int {
	switch l := len(nums); l {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
	}

	MAX := make([]int, len(nums))

	MAX[0] = nums[0]
	MAX[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		MAX[i] = max(
			MAX[i-2]+nums[i],
			MAX[i-1],
		)
	}

	return MAX[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
