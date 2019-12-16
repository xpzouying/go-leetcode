package letter_case_permutation

func letterCasePermutation(S string) []string {

	if len(S) == 0 {
		return []string{""}
	}

	res := make([]string, 0, len(S))
	// var curr string
	curr := make([]byte, 0, len(S))
	backtrack(S, curr, &res)
	return res
}

func backtrack(S string, curr []byte, res *[]string) {
	if len(S) == 0 {
		*res = append(*res, string(curr))
		return
	}

	// 每次取1个字符
	// 1. 如果是数字，则直接传入；
	// 2. 如果是字母，则传入当前字母后，继续传入大小写转化后的字母

	b := S[0]
	if b >= '0' && b <= '9' {
		curr = append(curr, b)
		backtrack(S[1:], curr, res)
	} else if b >= 'a' && b <= 'z' {
		curr = append(curr, b)
		backtrack(S[1:], curr, res)
		curr = curr[:len(curr)-1] // 回溯

		// 放入大写字母
		b = byte(int(b) + int('A') - int('a'))
		curr = append(curr, b)
		backtrack(S[1:], curr, res)
	} else if b >= 'A' && b <= 'Z' {
		curr = append(curr, b)
		backtrack(S[1:], curr, res)
		curr = curr[:len(curr)-1] // 回溯

		// 放入大写字母
		b = byte(int(b) - (int('A') - int('a')))
		curr = append(curr, b)
		backtrack(S[1:], curr, res)
	}
}
