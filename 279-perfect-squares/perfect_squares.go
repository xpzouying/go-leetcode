package perfect_squares

import "math"

func numSquares(n int) int {
	if n <= 3 {
		return n
	}

	DP := make([]int, n+1) // 从1到n的下标，抛弃下标0
	for i := 0; i <= 3; i++ {
		DP[i] = i
	}

	for i := 4; i <= n; i++ {

		// 计算n==i时，最少需要多少个
		min := math.MaxInt64

		j := 1
		for {
			sqr := j * j
			if sqr > i {
				break
			}

			// 需要计算i的结果
			need := DP[i-sqr] + 1
			if need < min {
				min = need
			}

			j++
		}

		DP[i] = min
	}

	return DP[n]
}

/**

思路：

由于题目中包括：要求最少个数，那么可以使用动态规划推断。

使用DP记录最少需要多少个数拼接。DP[i]表示，值为i时，最少需要的个数。

那么考虑如何拆解成子问题，也即状态转移方程是什么？

	> 示例：
	>
	>数字X=12，可以由多种情况构成：
	>
	>1. 全部用1构成，那么需要12个
	>2. 全部由2构成，那么是3个
	>3. 由9+3，而3又有1+1+1构成，一共4个
	>4. 等等

拆解：

DP[X=12] = MIN{
	DP[12-9],
	DP[12-4],
	DP[12-1],
} + 1，尝试所有小于X的完全平方数




*/
