package climbing_stairs

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
func climbStairs(n int) int {
	/*
		count[n] = count[n-1] + count[n-2]

		count[1] = 1
		count[2] = 2
	*/

	if n <= 2 {
		return n
	}

	counter := make([]int, n+1)

	counter[1] = 1
	counter[2] = 2

	for i := 3; i < n+1; i++ {
		counter[i] = counter[i-1] + counter[i-2]
	}

	return counter[n]
}
