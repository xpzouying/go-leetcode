package power_of_two

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for {
		if n == 1 {
			return true
		}

		a := n / 2
		b := n % 2

		if b != 0 {
			return false
		}

		n = a
	}
}
