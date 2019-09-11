package longest_palindromic_substring

func longestPalindrome(s string) string {

	l := len(s)

	switch l {
	case 0, 1:
		return s
	case 2:
		if s[0] == s[1] {
			return s
		} else {
			return s[:len(s)-1]
		}
	default:
	}

	// 动态规划表
	DP := make([][]bool, l+1)
	for i := 0; i < l+1; i++ {
		DP[i] = make([]bool, l+1)
	}

	var maxStr string = s[0:1]

	// 从 0 开始记录DP表格
	for size := 0; size <= l; size++ {
		bound := l - size
		for begin := 0; begin <= bound; begin++ {
			end := begin + size // 字符串的结束坐标，开区间

			if size < 2 {
				DP[begin][end] = true
				continue
			}

			DP[begin][end] = DP[begin+1][end-1] && (s[begin] == s[end-1])

			if DP[begin][end] == true {
				substr := s[begin:end]

				if len(substr) > len(maxStr) {
					maxStr = substr
				}
			}
		}
	}

	return maxStr
}
