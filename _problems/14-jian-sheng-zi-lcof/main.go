package main

func cuttingRope(n int) int {
	if n <= 2 {
		return 1
	}

	DP := make([]int, n+1)

	DP[0] = 0
	DP[1] = 1
	DP[2] = 1

	// 生成dp表
	for i := 3; i <= n; i++ {

		// j 表示要剪的长度。至少长度为2，才有收益；
		// 由于必须要剪一刀，所以需要小于i
		for j := 2; j < i; j++ {

			// 剪了一刀，j长度后，还剩下多长，为left。
			left := i - j

			// 第一个值：剪一刀的情况
			num1 := j * left

			// 第二种情况：剪一刀 + dp中的剪法
			num2 := j * DP[left]

			// 除了num1、num2比较外，还需要跟<for j>循环中的其他结果进行比较。
			DP[i] = max(
				DP[i],
				max(num1, num2),
			)
		}
	}

	return DP[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*

思路：

从题目来说，必然是动态规划。因为求的是：“最大乘积”。

1. 数字n的最大值是必然知道的，并且不会变化，所以可以保存起来。
使用数组保存：

- max: 长度为n+1，下标为n时，即为函数的返回值。


如果需要用到动态规划，那么就需要找到动态规划的状态变化。

* DP[n] 表示长度为n时，最大的乘积是什么。

* DP[n] 可以为剪一刀后，变成了2段，那么乘积为：i * (n-i)，
或者为 (i * DP[n-i])，这两个数的最大数。


2. 转换方程式：

DP[n] = Max{
	i * (n-i),
	i * DP[n-i],
}

*/
