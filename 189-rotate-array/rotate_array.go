package rotate_array

// 暴力法
// 解题思路：
// 每次往后挪动一位
// func rotate(nums []int, k int) {
// 	if k >= len(nums) {
// 		k = k % len(nums)
// 	}

// 	if len(nums) <= 1 || k == 0 {
// 		return
// 	}

// 	for {
// 		if k == 0 {
// 			break
// 		}

// 		n := nums[len(nums)-1]

// 		for i := len(nums) - 1; i >= 1; i-- {
// 			nums[i] = nums[i-1]
// 		}

// 		nums[0] = n
// 		k--
// 	}
// }

// 将nums拆分成两部分：
// split = len(nums) - k
// nums[0:split], nums[split:]
// 1. 翻转 nums[:split]
// 2. 翻转 nums[split:]
// 3. 翻转 nums[:]
//
// 举例
// [1, 2, 3, 4, 5], k=2
// ==> [4, 5, 1, 2, 3]

func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}

	if len(nums) <= 1 || k == 0 {
		return
	}

	split := len(nums) - k

	reverse(nums[:split])
	reverse(nums[split:])
	reverse(nums)
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for {
		if i >= j {
			return
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
