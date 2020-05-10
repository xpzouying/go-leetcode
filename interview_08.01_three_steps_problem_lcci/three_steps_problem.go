package three_steps_problem

func waysToStep(n int) int {

	/** 解题思路：
	  1. 与2步走的问题类似
	  2. 走到N台阶有3种方式可获得：
	      2.1 从n-3上来
	      2.2 从n-2上来
	      2.3 从n-1上来

	  3. 转移方程为：DP[n] = DP[n-3] + DP[n-2] + DP[n-1]

	  4. 初始化：
	  DP[1] = 1
	  DP[2] = 2
	  DP[3] = 4
	*/

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
	}

	DP := make([]int, n+1)
	DP[1] = 1
	DP[2] = 2
	DP[3] = 4

	for i := 4; i <= n; i++ {
		DP[i] = (DP[i-3] + DP[i-2] + DP[i-1]) % 1000000007
	}

	return DP[n]
}
