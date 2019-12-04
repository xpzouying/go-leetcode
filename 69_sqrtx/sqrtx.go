package sqrtx

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x < 4 {
		return 1
	}

	for i := 0; i < x; i++ {
		n := int(i * i)
		if n < x {
			continue
		}

		if n == x {
			return i
		}

		return i - 1
	}

	return 0
}
