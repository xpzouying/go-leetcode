package sqrtx

/**

执行用时 :
12 ms , 在所有 golang 提交中击败了 22.21% 的用户

内存消耗 :
2.2 MB , 在所有 golang 提交中击败了 98.04% 的用户

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x < 4 {
		return 1
	}

	for i := 0; i < x; i++ {
		n := int(i * i)
		if n < x {
			continue
		}

		if n == x {
			return i
		}

		return i - 1
	}

	return 0
}
*/

/**
优化思路：

1. x的平方根不会超出x/2

2. 可以使用分治算法找出平方根

具体步骤：

平方根的取值范围：
[0, x/2]

*/
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x < 4 {
		return 1
	}

	left := 1
	right := x >> 1

	for {
		if left >= right {
			break
		}

		mid := (left + right + 1) >> 1

		sqrt := int(mid * mid)
		if sqrt == x {
			return mid
		}

		// 平方太大了，往低位查找
		if sqrt > x {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}
