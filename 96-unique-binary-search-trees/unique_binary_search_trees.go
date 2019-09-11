package unique_binary_search_trees

func numTrees(n int) int {
	if n < 2 {
		return n
	}

	DP := make([]int, n+1)
	DP[0] = 1
	DP[1] = 1

	for i := 2; i <= n; i++ {

		count := 0 // 计算形态总和

		// j 表示根节点的值
		for j := 1; j <= i; j++ {
			// 左边节点：1, 2, ..., i-1
			left := j - 1

			// 右边节点：i+1, i+2, ..., n
			right := i - j

			leftCount := DP[left]
			rightCount := DP[right]
			count += (leftCount * rightCount)
		}

		DP[i] = count
	}

	return DP[n]
}
