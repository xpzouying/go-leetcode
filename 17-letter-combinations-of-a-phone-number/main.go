// Package main
//
// 执行结果：通过
//
// 显示详情
//
// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB , 在所有 Go 提交中击败了 94.45% 的用户

package main

var (
	res = []string{}

	keyboard = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
)

func letterCombinations(digits string) []string {
	defer func() {
		res = res[:0]
	}()

	digitsBacktrack(digits, 0, "")

	return res
}

// digitsBacktrack 对数字进行回溯。这里不表示字母的组合.
// 由于可以接收任意顺序返回。比如按“23”和按“32”返回的都是一样的，那么就可以传入index，表示当前遍历的开始位置。
//
// 所以，在这里，digits + index 表示可以选择的路径。paths 表示已经做的选择。
func digitsBacktrack(digits string, index int, path string) {
	if index == len(digits) {
		if len(path) != 0 {
			// 输入完毕，记录所有的结果
			res = append(res, path)
		}

		return
	}

	currNum := byte(digits[index])

	for _, char := range keyboard[currNum] {
		digitsBacktrack(digits, index+1, path+string(char))
	}
}
