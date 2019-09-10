package word_break

func wordBreak(s string, wordDict []string) bool {

	/**

	思路：

	输入:
	s = "applepenapple",
	wordDict = ["apple", "pen"]


	### 使用动态规划

	1. 使用空格作为单词的分割，比如把空格插入到第i个位置，
	那么单词s就分为：

	s = s[0:i], s[i:]


	2. 此时就可以拆成子问题了，
	如果插入空格放在i的位置，是满足条件，那么：
	s[0:i] 和 s[i:] 两个子字符串也都需要满足条件。

	此时就拆成了子问题最优解。


	使用 mem := make([]bool, len(s) + 1) 记录空格插入在第i个位置时，是否满足条件。
	mem的长度一共n+1位，放到最后一位表示字符串是否可以被拆分。


	列出公式：wordBreak简写wb


	mem[len(s)] = {
		(mem[1] && wb(s[1:])) ||
		(mem[2] && wb(s[2:])) ||
		...
		(mem[len(s) - 1] && wb(s[len(s) - 1])) ||
	}
	上列条件中，满足其中一项即可


	3. 动态规划反推mem的过程：
	mem[0] = true
	mem[1] = {
		mem[0] && wb(s[0:1])
	}
	mem[2] = {
		(mem[0] && wb(s[0:2])) ||
		(mem[1] && wb(s[1:2]))
	}

	mem[i] = {
		(mem[0] && wb(s[0:i])) ||
		(mem[1] && wb(s[1:i])) ||
		(mem[2] && wb(s[2:i])) ||
		...
		(mem[i-1] && wb(s[i-1:i]))
	}

	// 说明
	// mem[i-1] && wb(s[i-1:i])
	// 查询空格是否可以放在第i个位置：
	// 只需要查询 mem[i-1] 是否满足条件，
	// 如果满足条件，那么就查询 后续子字符串s[i-1:i]是否满足条件



	// 示例：

	l e e t c o d e

	1) i == 0: ==> "", "leetcode"
	2) i == 1: ==> "l", "eetcode"

	*/

	dict := listToMap(wordDict)

	// 用来保存动态规划记录表
	// 表示切分在在第i个位置时，s[0:i]是否可以满足切分条件
	mem := make([]bool, len(s)+1)

	mem[0] = true

	// 遍历记录mem值，以mem长度
	for i := 1; i < len(s)+1; i++ {

		for j := 0; j < i; j++ {
			preMem := mem[j]

			if preMem == false {
				// 第一个查表条件未满足
				mem[i] = false
				continue
			}

			suffix := s[j:i]

			if _, ok := dict[suffix]; ok {
				// 当前切分 是可行的
				mem[i] = true
				break
			}

		}

	}

	return mem[len(s)] // 返回最后一个
}

func listToMap(wordDict []string) map[string]bool {
	m := make(map[string]bool, len(wordDict))
	for _, s := range wordDict {
		m[s] = true
	}

	return m
}
