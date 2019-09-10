package house_robber_ii

func rob(nums []int) int {

	/**

	Test1:

	输入: [2,3,2]
	输出: 3

	A = [2, 3, 2]

	SUM[i] 记录到第i个坐标时，最大值
	SUM[i] = Max{
		SUM[i-2],
		SUM[i-1],
	}
	if i != len(A)-1 {
		// 如果不是最后一个元素
		SUM[i] += A[i]
	}
	else {
		// 如果是最后一个元素，
		// 需要判断是否应该添加最后一个
	}


	如果i是最后一个元素，需要特殊判断。
	- 如果i（下标）为奇数，也即元素总长为偶数，则可以把A[i] 累计进入SUM[i]；
	- 如果i（下标）为偶数，也即元素总长度为奇数，则不可以把A[i]累计进入SUM[i]；
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

	SUM := make([]int, len(nums))

	SUM[0] = nums[0]
	SUM[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		// 如果不是最后一个元素 或者 i 下标为奇数，可以把nums[i]累计入SUM[i]中
		if i != len(nums)-1 || (i%2 != 0) {
			SUM[i] = Max(
				SUM[i-2]+nums[i],
				SUM[i-1],
			)
			continue
		}

		// 否则，不可以将最后一个元素累计入SUM[i]中，
		// 此时需要考虑是选择第0个元素还是最后一个元素

		// 对比SUM[0]和SUM[1]，
		// 如果SUM[0] > SUM[1]，
		// 那么SUM[1]取的是SUM[0]，
		// 否则取的是SUM[1]，此时不影响SUM[i]

		// 如果第0个元素，不起作用，那么最后一个元素直接累计
		if nums[0] <= nums[1] {
			SUM[i] = Max(
				SUM[i-2]+nums[i],
				SUM[i-1],
			)
			continue
		}

		// 如果第0个元素，已经累计到最后一个元素中，则：
		// 此时，需要考虑是将最后一个元素放入SUM[i]，
		// 还是第0个元素放入SUM[i]，
		first := SUM[i-2]
		second := SUM[i-1]
		// 求第一个元素和最后一个元素的差值
		third := SUM[i-2] - nums[0] + nums[len(nums)-1]

		SUM[i] = Max(
			Max(first, third), // 决定进入SUM的是第0个元素还是最后一个元素
			second,
		)
	}

	return SUM[len(nums)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
