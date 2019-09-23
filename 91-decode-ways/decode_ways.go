package decode_ways

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	if l == 1 {
		if s == "0" {
			return 0
		}

		return 1
	}

	// 0 开头，则直接返回错误
	if s[0] == '0' {
		return 0
	}

	// DP 保存到该下标（包括）的编码方法的总数
	DP := make([]int, l)
	DP[0] = 1
	switch s[0] {
	case '1': // 11, 12
		if s[1] == '0' {
			DP[1] = 1
		} else {
			DP[1] = 2
		}
	case '2':
		switch s[1] {
		case '0':
			// 20
			DP[1] = 1
		case '7', '8', '9':
			// 27, 28, 29
			DP[1] = 1
		default:
			// 21, 22
			DP[1] = 2
		}
	default:
		// 31, 99
		DP[1] = 1
	}

	for i := 2; i < len(s); i++ {

		first, second := s[i-1], s[i]

		if second == '0' {
			// 如果第二个字符是 '0'结束
			if first == '1' || first == '2' {
				DP[i] = DP[i-2] // 只有一条编码路径，就是DP[i-2]加上2个字符组成的子字符串
			} else {
				// 则无法编码
				return 0
			}
			continue
		}

		// 下面的是 second != '0'
		// 比如：01, 11, 21, 31
		if first == '0' {
			return 0
		} else if first == '1' {
			// 比如11, 12, ..., 19
			DP[i] = DP[i-1] + DP[i-2] // 可以从两条路径进行编码
		} else if first == '2' {
			// 比如 21, ..., 27, 28, 29
			if second == '7' || second == '8' || second == '9' {
				// 27, 28, 29
				DP[i] = DP[i-1]
			} else {
				DP[i] = DP[i-1] + DP[i-2]
			}
		}
	}

	return DP[len(DP)-1]
}

/**

### 解题思路

**线索：**

要求解码方式的总数，是一道动态规划的题目的线索。

状态转移：

DP记录解码方法的总数。

DP[i]为字符串到第i个字符串时的总数，

由于编码最多是变化成2为字符，所以可以假设转移方程为：

DP[i] =

分下列几种类型：

1. 当s[i] == '0': 比如： 20, 30
	1.1 s[i-1] == "1", "2"：才有解答，并且为DP[i-2]
	1.2 否则，没有解答，直接返回0
2. 当s[i] != '0'，比如：09, 19, 29
	需要判断2位字符是否为合法（大于1且小于26）
	2.1 s[i-1] == '0': 则DP[i] = DP[i-2]
	2.2 s[i-1] == '1': 则DP[i] = DP[i-2] + DP[i-1]
	2.3 s[i-1] == 其他，则无解


初始化：
取
	s[0:0] = 0，0个字符的时候，为0
	s[0:1] = 1，1个字符的时候，为1

*/
