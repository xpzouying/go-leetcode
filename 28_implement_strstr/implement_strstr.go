package implement_strstr

func strStr(haystack string, needle string) int {
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}

	l1 := len(haystack)
	if l1 < l2 {
		return -1
	}

	for i := 0; i <= l1-l2; i++ {
		s1 := haystack[i : i+l2]

		if s1 == needle {
			return i
		}
	}

	return -1
}
