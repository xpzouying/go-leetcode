package plus_one


func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; {
		if digits[i] == 9 {
			digits[i] = 0

			if i != 0 {
				i--
			} else {
				digits = append([]int{0}, digits...)
				// i still is 0
			}
			continue
		} else {
			digits[i]++
			break
		}
	}

	return digits
}
