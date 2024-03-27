package generateparenttheses

// ChatGPT 的解答：https://chat.openai.com/c/bcb1bd97-28ea-4407-ab71-3540f834f534

func generateParenthesis(n int) []string {

	if n == 0 {
		return []string{}
	}

	// 动态转移方程：
	// dp[i] 表示当 len=i 时，生成的所有的括号组合。
	// 那么，dp[i] 的状态转移就是可以这么理解：
	//
	// 当需要插入一个新的括号的话，那么我们要考虑如何把括号括起来，
	// 相当于转化成：
	// 一部分是：括号括起来的所有的组合，另外一部分是：括号没有括起来的所有的组合。
	// 比如括起来的所有的组合是 dp[p]，没有括起来的所有的组合是：dp[q]，而 i = p+q+1，1的原因是新增加的这个括号的计数。
	// dp[i] = "(" + <dp[p]的所有组合> + ")" + <dp[q]的所有组合>

	dp := make([][]string, n+1)
	dp[0] = []string{""}

	for i := 1; i <= n; i++ {

		for p := 0; p < i; p++ {

			// 遍历所有的组合
			for _, inner := range dp[p] {

				q := i - 1 - p
				for _, outer := range dp[q] {

					newStr := "(" + inner + ")" + outer

					dp[i] = append(dp[i], newStr)
				}
			}
		}
	}

	return dp[n]
}
