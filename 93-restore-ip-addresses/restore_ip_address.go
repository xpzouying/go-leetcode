package restore_ip_addresses

import (
	"strconv"
	"strings"
)

// 复原IP
// 25525511135
// i: [2]
// j: [5]
// k: [525511135]
func restoreIpAddresses(s string) []string {
	curr := make([]string, 0, 4)
	ans := make([]string, 0, 4)

	helper(s, curr, &ans)

	return ans
}

// 返回值：是否获得结果
// 返回 空字符串表示没有找到合适的组合
func helper(s string, curr []string, ans *[]string) {
	if len(s) == 0 {
		if len(curr) == 4 {
			*ans = append(*ans, strings.Join(curr, "."))
			return
		}

		return
	}

	if len(curr) == 4 && len(s) != 0 {
		return
	}

	// i 取s的多少位长度
	for sublen := 1; sublen <= 3; sublen++ {
		if sublen > len(s) {
			break
		}

		// 验证当前子字符串是否合法
		substr := s[:sublen]
		if len(substr) != 1 && substr[0] == '0' {
			// 排除首字母为0的子字符串，比如："01"
			return
		}

		val, _ := strconv.Atoi(substr)
		if val > 255 {
			break
		}

		curr = append(curr, substr) // 增加
		helper(s[sublen:], curr, ans)
		curr = curr[:len(curr)-1] // 回溯
	}
}
