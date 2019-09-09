package unique_paths

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
func uniquePaths(m int, n int) int {
	/**

	Path[0][0] = 0
	Path[0][j] = 1
	Path[i][0] = 1

	Path[i][j] = Path[i-1][j] + Path[i][j-1]
	*/

	// 1. 构建矩阵 m*n
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}

	// 2. 初始化
	// 2.1 第一列为1
	for i := 0; i < m; i++ {
		path[i][0] = 1
	}
	// 2.2 第一行为1
	for i := 0; i < n; i++ {
		path[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}

	return path[m-1][n-1]
}
