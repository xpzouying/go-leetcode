package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	l := len(numbers)

	if l < 2 {
		return nil
	}

	i := 0
	j := l - 1

	for {
		if i >= j {
			return nil
		}

		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum > target {
			j--
			continue
		}

		if sum < target {
			i++
			continue
		}
	}
}
