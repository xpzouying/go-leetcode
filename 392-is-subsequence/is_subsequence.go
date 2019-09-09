package is_usbsequence

// func isSubsequence(s string, t string) bool {

// 	slen := len(s)
// 	tlen := len(t)

// 	if slen == 0 {
// 		return true
// 	}

// 	if tlen == 0 {
// 		return false
// 	}

// 	// 比较最后一个字符是否相等
// 	if s[slen-1] == t[tlen-1] {
// 		return isSubsequence(s[:slen-1], t[:tlen-1])
// 	}

// 	return isSubsequence(s, t[:tlen-1])
// }

func isSubsequence(s string, t string) bool {

	i := 0 // pointer to s
	j := 0 // pointer to t

	slen := len(s)
	tlen := len(t)

	for {
		if i == slen {
			return true
		}

		if j == tlen {
			return false
		}

		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
}
