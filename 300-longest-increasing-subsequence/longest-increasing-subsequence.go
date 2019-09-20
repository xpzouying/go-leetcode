package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {

	l := len(nums)

	if l <= 1 {
		return l
	}

	dp := make([]int, l)
	dp[0] = 1 // 初始化

	maxRes := dp[0] // 最长递增子序列的长度，返回结果，初始化为第0个

	for i := 1; i < l; i++ {
		// 从当前下标向前搜索，找到一个比当前数更小的值，然后判断大小
		longest := 1 // 预设最小的最长递增子序列长度为1，即单个数字

		// 开始搜索i之前的子序列中，小于nums[i]的dp值，并且选择出最长的结果作为dp[i]
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] {
				longest = max(longest, dp[j]+1)
			}
		}

		dp[i] = longest

		if longest > maxRes {
			maxRes = longest
		}

	}

	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/**


解题思路：

如果输入为：

[10, 9, 2, 3, 1， 5, 4]


动态规划表：var dp []int
作用：保存到第i位时，最长递增子序列

dp[i] = max{
	dp[0,...j] + 1, j < i 且 if num[j] < num[i],
}

解释上述表达式：dp[0,...j] + 1, if num[j] < num[i]

查找 num[0, ..., i-1]数字中，所有的比num[i]小的数字，都可以与num[i]构成递增子序列，
找到对应的dp[i]最大的值，加1（即加上当前num[i]）构成最大递增子序列，
其结果与dp[i-1]进行对比，
更大的值为dp[i]的结果。

初始化：dp[0] = 1

比如：

[10, 9, 2, 3, 1， 5, 4]

dp[0] = 1
dp[1] = 1
dp[2] = 1  // 数值为2
dp[3] = 2  // 数值为3
dp[4] = 1  // 数值为1

如何计算dp[5]，

此时数值为num[i] == 5，
从0开始，找到小于5的数值，dp长度+1，
  - j = 0，num[0] == 10；10>5，跳过
  - j = 1，num[1] == 9；跳过
  - j = 2, num[2] == 2；该处为：dp[2]+1 ==> 2
  - j = 3, num[3] == 3; dp[3] = dp[2]+1 ==> 3
  - j = 4, num[4] == 1, 没有找到比1小的数，所以dp[4] = 1
  - j = 5, num[5] == 5, dp[5] = max{dp[2]+1, dp[3]+1, dp[4]+1} = dp[3]+1 = 3

*/
