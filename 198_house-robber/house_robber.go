package house_robber

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
func rob(nums []int) int {

	/*

		Max[i] = Max{(max[i-2] + a[i]), Max[i-1]}

		------------------
		Test-1: [2, 1, 1, 2]

		max[0] = 2
		max[1] = 2

		max[2] = Max{max[0] + a[2], Max[1]} = Max{2+1, 2} = 3
		max[3] = Max{max[1] + a[3], Max[2]} = Max{2+2, 3} = 4
	*/

	/*
		switch l := len(nums); l {
		case 0:
			return 0
		case 1:
			return nums[0]
		case 2:
			return Max(nums[0], nums[1])
		default:
		}

		sum := make([]int, len(nums))
		sum[0] = nums[0]
		sum[1] = nums[1]

		max := Max(nums[0], nums[1])

		for i := 2; i < len(nums); i++ {

			sum[i] = Max(
				sum[i-2]+nums[i],
				sum[i-1],
			)

			if sum[i] > max {
				max = sum[i]
			}
		}

		return max
	*/

	switch l := len(nums); l {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return Max(nums[0], nums[1])
	default:
	}

	maxList := make([]int, len(nums))
	maxList[0] = nums[0]
	maxList[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {

		maxList[i] = Max(
			maxList[i-2]+nums[i],
			maxList[i-1],
		)
	}

	return maxList[len(maxList)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
