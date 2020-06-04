package multiply_strings

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num1) == 0 {
		return ""
	}

	// a*b
	var a string
	var b string
	if len(num1) > len(num2) {
		a = num1
		b = num2
	} else {
		a = num2
		b = num1
	}

	sum := "0"
	for i := 0; i < len(b); i++ {
		res := strMultiple(a, b[len(b)-1-i], i)
		sum = strPlus(sum, res)
	}

	return sum
}

// 字符串乘法
// factor：补"0"的因子
func strMultiple(n1 string, n2 byte, factor int) string {
	// res := make([]byte, 0, len(n1)+1+factor)

	v2 := int(n2) - int('0')

	sum := "0"
	needZeros := 0
	var carry int
	for i := len(n1) - 1; i >= 0; i-- {
		b1 := int(n1[i] - '0') // 转化成整形

		// 计算当前值
		n := b1*v2 + carry

		str1 := strconv.Itoa(n)
		for j := 0; j < needZeros; j++ {
			str1 = str1 + "0"
		}

		sum = strPlus(sum, str1)

		carry = n % 10
	}

	return sum
}

func strPlus(s1, s2 string) string {
	var res string

	i := len(s1) - 1
	j := len(s2) - 1
	carry := 0
	for {
		if i < 0 {
			res = s2[:j+1] + res
			break
		}

		if j < 0 {
			res = s1[:i+1] + res
			break
		}

		v1 := int(s1[i]) - '0'
		v2 := int(s2[i]) - '0'

		n := v1 + v2 + carry
		if n < 10 {
			carry = 0
		} else {
			carry = 1
			n -= 10
		}
		// res = fmt.Sprintf("%d%s", n, string(res))

		res = fmt.Sprintf("%s%d", string(res), n)

		i--
		j--
	}

	return res
}
