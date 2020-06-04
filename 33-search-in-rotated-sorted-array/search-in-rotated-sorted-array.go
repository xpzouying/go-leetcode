package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return doSearch(nums, 0, len(nums)-1, target)
}

func doSearch(nums []int, low, high int, target int) int {
	if low > high {
		return -1
	}

	if nums[low] == target {
		return low
	}

	if nums[high] == target {
		return high
	}

	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	}

	// 检查左半部分
	findLeft := -1
	// 如果没有转序
	if nums[low] < nums[mid] {
		if nums[low] < target && target < nums[mid] {
			findLeft = doSearch(nums, low+1, mid-1, target)
		}
	} else {
		// 由于有转序，则继续递归查找
		findLeft = doSearch(nums, low+1, mid-1, target)
	}

	// 在左侧找到，则不需要找右侧
	if findLeft != -1 {
		return findLeft
	}

	// 检查右半部分
	findRight := -1
	// 如果右侧没有转序
	if nums[mid] < nums[high] {
		if nums[mid] < target && target < nums[high] {
			findRight = doSearch(nums, mid+1, high-1, target)
		}
	} else {
		// 由于转序，则继续递归查找
		findRight = doSearch(nums, mid+1, high-1, target)
	}

	return findRight
}

/**

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4


因为需要

*/
