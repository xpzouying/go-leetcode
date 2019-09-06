package best_time_to_buy_and_sell_stock

import (
	"log"
)

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 */
func maxProfit(prices []int) int {

	/*

		profits: [-6, 4, -2, 3, -2]


		sum[i] = max{
			sum[i-1] + A[i],
			A[i]
		}

		sum[1] = 0
		sum[2] = 4
		sum[3] = max{sum[2] + -2, -2} = 2
		sum[4] = max{sum[3] + 3, 3} = 5
		sum[5] = max{sum[4] + -2, -2} = 3

	*/

	if len(prices) < 2 {
		return 0
	}

	profits := make([]int, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		profits[i-1] = prices[i] - prices[i-1]
	}

	log.Printf("prices: %+v\n  profits: %v\n", prices, profits)

	sum := make([]int, len(profits)+1)
	sum[0] = profits[0]
	max := Max(0, sum[0])
	for i := 1; i < len(profits); i++ {
		n := sum[i-1] + profits[i]
		sum[i] = Max(n, profits[i])

		if sum[i] > max {
			max = sum[i]
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
