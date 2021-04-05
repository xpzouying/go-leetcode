package roman_to_integer

var (
	romanNums = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,

		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
)

func romanToInt(s string) int {
	return helper(s, 0)
}

func helper(s string, sum int) int {
	slen := len(s)
	if slen == 0 {
		return sum
	}

	var romanstr string

	// 先处理特殊情况
	// 1. I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	// 2. X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	// 3. C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	if slen >= 2 {

		romanstr = s[:2]

		if value, ok := romanNums[romanstr]; ok {
			sum += value
			return helper(s[2:], sum)
		}
	}

	// 如果2位数字没有取到，则取一位数字
	romanstr = s[:1]
	sum += romanNums[romanstr]
	return helper(s[1:], sum)
}

/*

执行结果：通过

执行用时： 12 ms , 在所有 Go 提交中击败了 43.02% 的用户

内存消耗： 3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
