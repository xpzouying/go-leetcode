package power_of_two

func isPowerOfTwo(n int) bool {

	// --- 解法1 ---
	// if n == 0 {
	// 	return false
	// }

	// if n == 1 {
	// 	return true
	// }

	// for {
	// 	if n == 1 {
	// 		return true
	// 	}

	// 	a := n / 2
	// 	b := n % 2

	// 	if b != 0 {
	// 		return false
	// 	}

	// 	n = a
	// }

	// --- 解法2 ---

	return (n > 0) && ((n & (n - 1)) == 0)
}
