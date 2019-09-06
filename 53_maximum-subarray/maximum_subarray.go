package maximum_subarray

func maxSubArray(nums []int) int {

	/*
		SUM[i] = max{
			SUM[i-1] + a[i],
			a[i]
		}
	*/

	if len(nums) == 1 {
		return nums[0]
	}

	sum := make([]int, len(nums))
	sum[0] = nums[0]
	max := sum[0]

	for i := 1; i < len(nums); i++ {
		n := sum[i-1] + nums[i]

		if n > nums[i] {
			sum[i] = n
		} else {
			sum[i] = nums[i]
		}

		if sum[i] > max {
			max = sum[i]
		}
	}

	return max
}
