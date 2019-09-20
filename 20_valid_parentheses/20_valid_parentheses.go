package valid_parentheses

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (39.30%)
 * Likes:    989
 * Dislikes: 0
 * Total Accepted:    109.6K
 * Total Submissions: 279K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
func isValid(s string) bool {
	switch len(s) {
	case 0:
		return true
	case 1:
		return false
	default:
	}

	stack := NewStack(len(s))

	for _, c := range s {

		switch c {
		case '{', '[', '(':
			stack.push(byte(c))

		case ']':

			if stack.pop() != '[' {
				return false
			}

		case '}':
			if stack.pop() != '{' {
				return false
			}

		case ')':
			if stack.pop() != '(' {
				return false
			}

		default:
			return false
		}
	}

	if stack.IsEmpty() {
		return true
	}

	return false
}

type Stack struct {
	data []byte
}

func NewStack(len int) *Stack {
	return &Stack{data: make([]byte, 0, len)}
}

func (s *Stack) IsEmpty() bool {

	if len(s.data) == 0 {
		return true
	}

	return false
}

func (s *Stack) String() string {
	return string(s.data)
}

func (s *Stack) push(c byte) {
	s.data = append(s.data, c)
}

func (s *Stack) pop() byte {
	if s.IsEmpty() {
		return byte('x')
	}

	l := len(s.data)

	c := s.data[l-1]
	s.data = s.data[:l-1]

	return c
}
