package number_of_1_bits

func hammingWeight(num uint32) int {

	if num == 0 {
		return 0
	}

	count := 0

	for {
		if num == 0 {
			break
		}

		if (num & 1) == 1 {
			count++
		}

		num = num >> 1
	}

	return count
}
