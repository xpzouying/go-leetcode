package find_peak_element

func findPeakElement(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	begin := 0
	end := len(nums) - 1

	for {
		if begin >= end {
			return begin
		}

		mid := (begin + end) / 2

		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			begin = mid + 1
		}
	}
}
