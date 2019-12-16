package simplifypath

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0, len(path))
	stack = append(stack, "/") // 根节点

	paths := strings.Split(path[1:], "/")
	for _, path := range paths {

		switch path {
		case "", ".":
			continue
		case "..":
			if len(stack) == 1 { // 根目录
				continue
			}
			stack = stack[:len(stack)-1] // 出栈
		default:
			stack = append(stack, path) // 入栈
		}
	}

	if len(stack) == 1 {
		return "/"
	}

	return "/" + strings.Join(stack[1:], "/")
}

/**

解题思路：

1. 使用“栈”辅助
2. 使用 “/” 切分字符串
3.

*/
