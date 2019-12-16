package reverse_words

// 1. 反转整个字符串
// 2. 反转每一个单词
// 3. 去除多余的空格
func reverseWords(s string) string {
	lens := len(s)

	b := []byte(s)

	// 1. 反转整个字符串
	i := 0
	j := lens - 1
	for {
		if i >= j {
			break
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	// 2. 反转单词
	begin := 0
FOR_OUT:
	for {
		for {
			if begin >= lens-1 {
				// 达到最后一个
				break FOR_OUT
			}
			if b[begin] != ' ' {
				break
			}
			begin++
		}

		end := begin + 1
		for {
			if end >= lens {
				break
			}

			if b[end] == ' ' {
				// 找到空格
				break
			}
			end++
		}

		i := begin
		j := end - 1
		for {
			if i > j {
				break
			}
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}

		begin = end
	}

	// 3. 清除空格
	// 3.1 trim left
	for i = 0; i < lens; i++ {
		if b[i] != ' ' {
			break
		}
	}
	b = b[i:]

	// 3.2 trim right
	for j = len(b) - 1; j >= 0; j-- {
		if b[j] != ' ' {
			break
		}
	}
	b = b[:j+1]

	// 3.3 trim middle space
	i = 0
	for {
		if i >= len(b)-1 {
			break
		}

		if b[i] == ' ' && b[i+1] == ' ' {
			b = append(b[:i], b[i+1:]...)
			continue
		}
		i++
	}

	return string(b)
}

// func reverseWords(s string) string {
// 	s = strings.TrimSpace(s)

// 	ss := strings.Split(s, " ")

// 	res := make([]string, 0, len(ss))
// 	for _, substr := range ss {
// 		if len(substr) != 0 {
// 			res = append(res, substr)
// 		}
// 	}

// 	// 反转
// 	l := len(res)
// 	for i := 0; i < l/2; i++ {
// 		res[i], res[l-1-i] = res[l-1-i], res[i]
// 	}

// 	log.Printf("%+v", res)

// 	return strings.Join(res, " ")
// }

// func reverse(s []byte) string {
// 	i := 0
// 	j := len(s) - 1
// 	for {
// 		if i >= j {
// 			break
// 		}

// 		s[i], s[j] = s[j], s[i]
// 		i++
// 		j--
// 	}

// 	return string(s)
// }

/**

解题思路：
反转两次，

第1次，反转整个字符串；在此过程中，删除前后端的空格。
第2次，反转每一个单独的单词；

*/
