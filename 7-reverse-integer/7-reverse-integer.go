package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
func reverse(x int) int {
	if x == 0 {
		return 0
	}

	var result int

	for {
		res := x / 10       // 除法结果
		remainder := x % 10 // 余数
		x = res

		result = result * 10
		result += remainder

		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}

		if x == 0 {
			break
		}
	}

	if x < 0 {
		result = -result
	}

	return result
}
