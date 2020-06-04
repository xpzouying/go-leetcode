package count_squares

func countSquares(matrix [][]int) int {

	M := len(matrix)
	if M == 0 {
		return 0
	}
	N := len(matrix[0])
	if N == 0 {
		return 0
	}

	// --- 初始化 ---

	// k为边长。正方形最长不超过某一边的长度；k+1的原因是，坐标从0开始
	K := min(M, N) + 1

	dp := make([][][]int, M)

	for m := 0; m < M; m++ {

		dp[m] = make([][]int, N)

		for n := 0; n < N; n++ {
			dp[m][n] = make([]int, K)
		}
	}
	// 此时dp已经为一个3维数组，为dp[i][j][k]

	// --- 进行动态规划 ---
	// dp[i][j][k]的定义：对于右下角为[i,j]节点时，边长为k的矩阵，新增的正方形的数量

	// 记录所有新增的次数
	count := 0

	// 初始化第0行和第0列，长度只能为1
	for i := 0; i < N; i++ {
		dp[0][i][1] = matrix[0][i]

		// 计数
		if dp[0][i][1] == 1 {
			count++
		}
	}
	for j := 1; j < M; j++ { // j从1开始，忽略重复计算[0,0]
		dp[j][0][1] = matrix[j][0]

		// 计数
		if dp[j][0][1] == 1 {
			count++
		}
	}

	for i := 1; i < M; i++ {

		for j := 1; j < N; j++ {

			// 首先判断matrix[i][j]，若[i,j]点为0，则所有边长的新增个数都为0
			if matrix[i][j] == 0 {
				continue
			}

			// 边长为1的正方形，新增数量为1
			dp[i][j][1] = 1

			// 计数：边长为1
			count++

			// 初始化其他的坐标点
			// k为边长，从边长为2开始
			for k := 2; k <= K; k++ {
				// 如果已经超出某一边的边长长度，则直接跳过该坐标
				if k > i+1 || k > j+1 {
					break
				}

				// 上、左、左上的上一次边长（即k-1)dp都新增1，
				// 那么表示新增matrix[i][j]点后，边长k的dp也可新增1
				up := dp[i-1][j][k-1]
				leftUp := dp[i-1][j-1][k-1]
				left := dp[i][j-1][k-1]

				if up == 1 && leftUp == 1 && left == 1 {
					dp[i][j][k] = 1

					// 计数：边长为k
					count++

					continue
				}

				// 否则没有增加
				// dp[i][j][k] = 0
			}
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/**

思路：

1. 求由1组成的正方形矩阵的个数总数，可以看出该题目为：动态规划。

2. 求解由1组成的正方形个数的方法：

动态规划：

- 转移方程

见图解：![](IMG_5FA3346160CE-1.jpeg)

- 初始化

初始化[0, 0]点




*/
