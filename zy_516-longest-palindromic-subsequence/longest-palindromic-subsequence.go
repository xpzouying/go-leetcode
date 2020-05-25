package longest_palindromic_subsequence

func longestPalindromeSubseq(s string) int {

	l := len(s)
	if l <= 1 {
		return l
	}

	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	max := 1
	// strLen is length of substring
	for strLen := 0; strLen < l; strLen++ {

		// begin is begin of substring
		for begin := 0; begin < l; begin++ {

			// end of substring,
			// substring is [begin, end], also is [begin, end+1)

			end := begin + strLen

			// out of range
			if end >= l {
				// break inside for scope
				break
			}

			if strLen == 0 {
				// [0,0], [1,1], ...
				dp[begin][end] = 1
				continue
			} else if strLen == 1 {
				if s[begin] == s[end] {
					dp[begin][end] = 1

					curSubstrLen := strLen + 1
					if curSubstrLen > max {
						max = curSubstrLen
					}

					continue
				}
			}

			// else, other length
			if s[begin] == s[end] && dp[begin+1][end-1] == 1 {
				dp[begin][end] = 1

				curSubstrLen := strLen + 1
				if curSubstrLen > max {
					max = curSubstrLen
				}

			}
		}
	}

	return max
}

/**

题解：

1. 判断题型为动态规划。

需要求出最长的回文子序列。

2. 找出动态规划的子最优解关系，即推导出动态规划转移方程。

字符串s[i, j]为回文字符串，那么其子字符串[i+1, j-1]也必须为回文字符串。

所以，DP[i][j] 表示子字符串[i, j]是否为回文。

DP[i][j]为回文，即为1的条件是：

- s[i] == s[j]，并且
- DP[i+1][j-1] == 1
*/
