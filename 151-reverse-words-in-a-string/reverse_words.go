package reverse_words

import (
	"log"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	ss := strings.Split(s, " ")

	res := make([]string, 0, len(ss))
	for _, substr := range ss {
		if len(substr) != 0 {
			res = append(res, substr)
		}
	}

	// 反转
	l := len(res)
	for i := 0; i < l/2; i++ {
		res[i], res[l-1-i] = res[l-1-i], res[i]
	}

	log.Printf("%+v", res)

	return strings.Join(res, " ")
}

func reverse(s []byte) string {
	i := 0
	j := len(s) - 1
	for {
		if i >= j {
			break
		}

		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return string(s)
}

/**

解题思路：
反转两次，

第1次，反转整个字符串；在此过程中，删除前后端的空格。
第2次，反转每一个单独的单词；

*/
