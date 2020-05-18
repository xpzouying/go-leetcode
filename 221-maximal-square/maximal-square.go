package maximal_square

func maximalSquare(matrix [][]byte) int {
	DP := make([][]int, len(matrix))

	// 最大边长
	var maxSquare int = 0

	// 初始化空间和第一行和第一列的值
	for i := 0; i < len(matrix); i++ {
		DP[i] = make([]int, len(matrix[i]))

		for j := 0; j < len(matrix[i]); j++ {

			// 初始化第一行和第一列
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					DP[i][j] = 1
					maxSquare = 1
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {

		for j := 1; j < len(matrix[i]); j++ {

			// 只有当matrix[i][j] 为 1时，才开始计算
			if matrix[i][j] == '0' {
				continue
			}

			// 比较各个方向的边长
			DP[i][j] = Min(
				Min(DP[i-1][j], DP[i][j-1]),
				DP[i-1][j-1]) + 1

			if DP[i][j] > maxSquare {
				maxSquare = DP[i][j]
			}
		}
	}

	return maxSquare * maxSquare
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/**

思路：

求最大正方形面积，判断题型为动态规划。

转换方程：

DP[i, j]保存终点加上M[i,j]点后的最大边长。

如果M[i,j] == 0，则为0，因为包括(0,0)点为0。

如果M[i,j] == 1，则：

	DP[i,j] 可以从3个方向扩展，

	1. D[i,j-1] --> D[i,j]
	2. D[i-1,j] --> D[i,j]
	3. D[i-1,j-1] --> D[i,j]

在上述情况下，边长可以为：

	1. 竖方向，如果M[i,j]为1，那么竖方向的边长可以加1，则 DP[i,j-1]+1
	2. 横方向，如果M[i,j]为1，那么横方向的边长可以加1，则 DP[i-1,j]+1
	3. 斜方向，如果M[i,j]为1，那么需要保证起点为1，DP[i-1,j-1]+1

由于正边形的长和宽相等，所以取上述3种情况的最小值，形成一个正方形的边长。


*/
