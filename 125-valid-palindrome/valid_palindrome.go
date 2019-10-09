package valid_palindrome

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	// i、j分别指向字符串开头和结尾
	i := 0
	j := len(s) - 1

	for {
		if i >= j {
			break
		}

		first, ok := validChar(s[i])
		if !ok {
			i++
			continue
		}

		second, ok := validChar(s[j])
		if !ok {
			j--
			continue
		}

		if first != second {
			return false
		}

		i++
		j--
	}

	return true
}

func validChar(c byte) (lowerC byte, ok bool) {
	if c >= '0' && c <= '9' {
		return c, true
	}

	if c >= 'A' && c <= 'Z' {
		// 转换成小写字符
		return (c + 'a' - 'A'), true
	}

	if c >= 'a' && c <= 'z' {
		return c, true
	}

	return c, false
}
