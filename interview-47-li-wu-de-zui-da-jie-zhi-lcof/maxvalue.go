package maxvalue

func maxValue(grid [][]int) int {

	// 异常判断
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	// 构建DP表
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 进行动态规划
	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			value := grid[i][j]

			if i == 0 && j == 0 {
				// 如果是(0,0)点
				dp[i][j] = value
				continue
			}

			if i == 0 {
				// 如果是第0行，只能从左边过来
				dp[i][j] = dp[i][j-1] + value
				continue
			}

			if j == 0 {
				// 如果是第0列，只能从上到下
				dp[i][j] = dp[i-1][j] + value
				continue
			}

			// 否则，需要动态规划判断
			left := dp[i][j-1]
			up := dp[i-1][j]
			if left > up {
				dp[i][j] = left + value
			} else {
				dp[i][j] = up + value
			}
		}
	}

	return dp[m-1][n-1]
}

/**

1. 题型从“最多能拿到多少价值的礼物”中可以看出是动态规划

2. 动态规划的思路：

- 转移方程

使用DP记录；DP[i][j]为走到当前节点时，拿到的最大价值的礼物

DP[i][j] = max { DP[i-1][j], DP[i][j-1] } + NUMS[i][j]

- 初始化

*/
