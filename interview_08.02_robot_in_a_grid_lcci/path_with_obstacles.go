package lost_robot

import (
	"fmt"
)

func pathWithObstacles(obstacleGrid [][]int) [][]int {

	// 这里需要计算可走的路径的可能性，所以DP保存的是走到当前路径方式
	// 由于题目的限制，那么走到当前网格的路径只能有2种：
	// 1. 从上往下走
	// 2. 从左往右走
	//
	// 计算完DP后，再通过DP反推出路径

	M := len(obstacleGrid)
	if M == 0 {
		return [][]int{}
	}

	N := len(obstacleGrid[0])
	if N == 0 {
		return [][]int{}
	}

	if obstacleGrid[0][0] == 1 {
		return [][]int{}
	}

	// DP记录从什么方向过来到当前区域块
	// 0: 不能到达该点
	// 1: 能从左边到右边
	// 2: 能从上到下
	DP := make([][]int, M)
	for i := 0; i < M; i++ {
		DP[i] = make([]int, N)
	}

	DP[0][0] = 1

	// 先初始化DP表格
	//
	// 初始化第0行
	for i := 1; i < N; i++ {
		// 如果有障碍，则不可达
		if obstacleGrid[0][i] == 1 {
			DP[0][i] = 0
			continue
		}

		// 看看左边是否可以过来
		if DP[0][i-1] == 0 {
			// 左边 --> 右边 方向不能过来
			DP[0][i] = 0
			continue
		}

		DP[0][i] = 1
	}

	// 初始化第0列
	for i := 1; i < M; i++ {
		// 如果有障碍，则不可达
		if obstacleGrid[i][0] == 1 {
			DP[i][0] = 0
			continue
		}

		// 看看上边是否可以过来
		if DP[i-1][0] == 0 {
			DP[i][0] = 0
			continue
		}

		DP[i][0] = 2
	}

	// 动态规划遍历
	for i := 1; i < M; i++ {

		for j := 1; j < N; j++ {
			// 如果有障碍
			if obstacleGrid[i][j] == 1 {
				DP[i][j] = 0 // 不可达
				continue
			}

			// 如果从左到右不可 并且从上到下不可达，
			// 则不可达
			if DP[i][j-1] == 0 && DP[i-1][j] == 0 {
				DP[i][j] = 0
				continue
			}

			// 如果从左到右可达
			if DP[i][j-1] != 0 {
				DP[i][j] = 1
			} else if DP[i-1][j] != 0 {
				DP[i][j] = 2
			}
		}
	}

	// 进行路径判断

	// 没有可达路径
	if DP[M-1][N-1] == 0 {
		return [][]int{}
	}

	paths := make([][]int, 1)
	paths[0] = []int{M - 1, N - 1}

	i := M - 1
	j := N - 1

	for {
		if i == 0 && j == 0 {
			break
		}

		var preNode []int
		if DP[i][j] == 1 {
			// 从左到右路径，则增加左边节点坐标
			preNode = []int{i, j - 1}
			j--
		} else if DP[i][j] == 2 {
			preNode = []int{i - 1, j}
			i--
		} else {
			panic(fmt.Sprintf("unknown: i=%d,j=%d", i, j))
		}

		paths = append([][]int{preNode}, paths...)
	}

	return paths
}
