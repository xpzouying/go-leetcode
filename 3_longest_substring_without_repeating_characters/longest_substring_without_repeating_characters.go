package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	max := 1 // 记录最大长度
	begin := 0
	end := 1

	// m 记录字符与下标的关系
	m := map[byte]int{
		s[0]: 0,
	}

	// var currLen int
	for {
		if end >= len(s) {
			break
		}

		b := s[end]

		index, ok := m[b]
		if !ok {
			// 如果是新的字符出现
			m[b] = end
			end++

			currLen := end - begin
			if currLen > max {
				max = currLen
			}
		} else {
			// 如果已经存在
			// 1. 更新m中的记录
			// 2. 挪动begin下标
			begin = index + 1
			end++

			newSubLen := end - begin
			newM := make(map[byte]int, newSubLen)
			for i := begin; i < end; i++ {
				newM[s[i]] = i
			}
			m = newM
		}
	}

	return max
}

/**

解体思路：

1. 使用两个指针指向字符串的前后端，假设记为 begin, end，可以标示一个子字符串；

2. end向后移动，会出现下列情况：
	2.1 如果下一个字符没有在子字符串中出现，那么end后移。并检查是否需要更新最长长度；
	2.2 如果下一个字符串在子字符串中出现，则end后移，同时begin移动到出现字符串的下一个下标；

3. 检查字符是否在原有字符串中出现过，并且出现过的位置，可以使用map来记录，key为字符，value为下标；

*/
