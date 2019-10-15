package rotate_array

// 暴力法
// 解题思路：
// 每次往后挪动一位
func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}

	for {
		if k == 0 {
			break
		}

		n := nums[len(nums)-1]

		for i := len(nums) - 1; i >= 1; i-- {
			nums[i] = nums[i-1]
		}

		nums[0] = n
		k--
	}
}
