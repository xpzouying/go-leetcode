package minimum_path_sum

func minPathSum(grid [][]int) int {

	/**

	如果需要求的最小值，也即到达右下角的和为最小值。

	到达右下角，有两条路径：
	1. 由上到下；
	2. 由左到右；

	SUM[m][n] = Min{
		SUM[m-1][n] + grid[m][n],
		SUM[m][n-1] + grid[m][n],
	}

	初始化：
	SUM[i][0] = grid[i][0]
	SUM[0][j] = grid[0][j]
	*/

	// 排除一些小的特殊的矩阵
	if len(grid) == 0 {
		return 0
	}

	if len(grid) == 1 {
		g0 := grid[0]
		if len(g0) == 1 {
			return g0[0]
		}
	}

	// 初始化
	m := len(grid)    // 行数
	n := len(grid[0]) // 列数

	SUM := make([][]int, m)
	for i := 0; i < m; i++ {
		SUM[i] = make([]int, n)
		// SUM[i][0] = grid[i][0]
	}

	// 初始化左上角（即起点）
	SUM[0][0] = grid[0][0]

	// 初始化第一行
	for j := 1; j < n; j++ {
		SUM[0][j] = SUM[0][j-1] + grid[0][j]
	}

	// 初始化第一列
	for i := 1; i < m; i++ {
		SUM[i][0] = SUM[i-1][0] + grid[i][0]
	}

	// DP
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			n := grid[i][j]
			SUM[i][j] = min(
				SUM[i-1][j]+n,
				SUM[i][j-1]+n,
			)
		}
	}

	return SUM[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
