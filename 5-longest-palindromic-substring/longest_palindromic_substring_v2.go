package longest_palindromic_substring

// 第二种解法：与 https://leetcode.cn/problems/longest-palindromic-subsequence/submissions/517291224/ 的解法类似
//
// 记录 dp， dp[i][j] 表示 s[i:j] 字符串的回文字符串长度，如果不是回文字符串的话，那么 dp[i][j] = 0。
//
// 转移方程：
// 遍历过程：从下往上，从左往右 进行 dp 设置。
// dp 初始化，对角线都是为 1。
func longestPalindromeV2(s string) string {

	count := len(s)
	if count <= 1 {
		return s
	}

	dp := make([][]int, count)
	for i := 0; i < count; i++ {
		dp[i] = make([]int, count)

		// 顺便初始化对角线
		dp[i][i] = 1
	}

	var (
		maxLen        = 1
		maxStr string = string(s[0])
	)

	// 开始状态转移方程
	for i := count - 1; i >= 0; i-- {

		for j := i + 1; j < count; j++ {

			// 这里需要处理一下 状态转移方程
			// 有两种情况：s[i] 和 s[j] ：
			// 1. 不相等的情况，那么当前的 s[i:j] 就不是回文子串，那么就直接赋值为 0；
			// 2. 相等的情况，dp[i][j] 因为是表示最长的回文子串，那么就可以从 它的子串转移过来，
			//   如果 dp[i+1][j-1] 是回文子串，那么 dp[i][j] = 2 + dp[i+1][j-1]。
			//   如果不是回文，那么 dp[i][j] = 0。
			//   如果 i,j 之间的长度为 2，那么 dp[i][j] = 2。

			if s[i] != s[j] {
				dp[i][j] = 0
				continue
			}

			// 如果长度为 2
			if j-1 == i {
				dp[i][j] = 2
			} else {
				if dp[i+1][j-1] == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = 2 + dp[i+1][j-1]
				}
			}

			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
				maxStr = s[i : j+1] // 保存最大字符串的状态
			}
		}
	}

	return maxStr
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
