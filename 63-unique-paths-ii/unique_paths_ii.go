package unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	// 出口即是障碍，出不去
	// if obstacleGrid[0][1] == 1 && obstacleGrid[1][0] == 1 {
	//     return 0
	// }

	lines := len(obstacleGrid)
	cols := len(obstacleGrid[0])

	DP := make([][]int, lines)
	for i := 0; i < lines; i++ {
		DP[i] = make([]int, cols)
	}

	// 初始化
	DP[0][0] = 1 - obstacleGrid[0][0]
	// 初始化第一行
	for i := 1; i < cols; i++ {
		DP[0][i] = DP[0][i-1] * (1 - obstacleGrid[0][i])
	}
	// 初始化第一列
	for i := 1; i < lines; i++ {
		DP[i][0] = DP[i-1][0] * (1 - obstacleGrid[i][0])
	}

	// 初始化数组的边缘
	for i := 1; i < lines; i++ {

		for j := 1; j < cols; j++ {
			count := 0
			// 先算横向
			if DP[i-1][j] != 0 && obstacleGrid[i][j] == 0 {
				// 如果横向没有被封堵死了
				count += DP[i-1][j]
			}
			// count += DP[i-1][j] * (1-obstacleGrid[i][j])

			// 计算纵向
			if DP[i][j-1] != 0 && obstacleGrid[i][j] == 0 {
				count += DP[i][j-1]
			}
			// count += DP[i][j-1] * (1-obstacleGrid[i][j])

			DP[i][j] = count
		}
	}

	return DP[lines-1][cols-1]
}

/**
解题思路：

有多少条不同的路径，可以看出是动态规划。

假设DP[i][j]为走到当前节点有多少条路径，

DP[i][j]的值可以从：DP[i][j-1]或者DP[i-1][j]中推算获得。

*/
