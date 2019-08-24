package contains_duplicate

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	m := make(map[int]bool, len(nums))

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}

		m[n] = true
	}

	return false
}
