package maximum_product_subarray

func maxProduct(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	DP1 := make([]int, l) // DP1 正数DP表
	DP2 := make([]int, l) // DP2 负数DP表
	DP1[0] = nums[0]
	DP2[0] = nums[0]
	max := nums[0]

	for i := 1; i < l; i++ {
		n := nums[i]

		if n == 0 {
			DP1[i] = 0
			DP2[i] = 0
		} else if n > 0 {
			// 最大正数子序列
			// 计算最大正数
			v1 := DP1[i-1] * n
			if v1 >= n {
				DP1[i] = v1
			} else {
				DP1[i] = nums[i]
			}

			// 最大负数子序列
			// 计算最大负数
			DP2[i] = DP2[i-1] * n
		} else { // 否则是负数
			// 计算最大正数
			DP1[i] = DP2[i-1] * n

			// 计算最大负数
			v2 := DP1[i-1] * n
			if v2 <= n {
				DP2[i] = v2
			} else {
				DP2[i] = n
			}
		}

		if DP1[i] > max {
			max = DP1[i]
		}
	}

	return max
}

/**

解题思路：

从题目中的“乘积最大”，很容易得出来是一道动态规划题目。

假设：DP为最大乘积。

如何进行拆分子任务。


### 思路：

DP[i] 和 nums[i]的关系，

`DP[i] = DP[i-1] * nums[i]`

这个是一种状态（不能满足全部）。

由于在这里子序列是需要连续的，所以就会出现下列两种子问题：

1. DP[i] = DP_2 * nums[i]  // 前面的子字符串有效
2. DP[i] = nums[i]  // 前面的子字符串在某种情况下无效，抛弃

那么可以拆分成哪些情况，

1.1 nums[i]是正数，那么需要从前面最大的正数子序列中获取
1.2 nums[i]是负数，那么需要从前面最大的负数子序列中获取
2.1 如果乘积都不如nums[i]，比如前面都是0，那么DP[i]为nums[i]

所以我们需要两张DP表，分别记录最大正数的子序列和最大负数的子序列。

*/
