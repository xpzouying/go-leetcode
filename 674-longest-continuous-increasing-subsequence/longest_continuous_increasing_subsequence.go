package longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	nlen := len(nums)
	if nlen <= 1 {
		return nlen
	}

	DP := 1
	max := 1

	for i := 1; i < nlen; i++ {

		if nums[i-1] < nums[i] {
			DP++
		} else {
			DP = 1
		}

		if DP > max {
			max = DP
		}
	}

	return max
}
