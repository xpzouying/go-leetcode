package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	/**
	使用贪心算法

	尽可能的进行多次交易，只要利润是正的，都可以买入
	*/

	if len(prices) < 2 {
		return 0
	}

	/**
	执行用时 : 8 ms , 在所有 Go 提交中击败了 75.25% 的用户
	内存消耗 : 3.4 MB , 在所有 Go 提交中击败了 5.24% 的用户
	*/

	// // 建立一个表格，记录前一天买入，第二天卖出的利润
	// profit := make([]int, len(prices)-1)

	// for i := 1; i < len(prices); i++ {
	// 	profit[i-1] = prices[i] - prices[i-1]
	// }

	// var sum int
	// for _, value := range profit {
	// 	if value > 0 {
	// 		sum += value
	// 	}
	// }

	// return sum

	// --------------

	var sum int
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]

		if profit > 0 {
			sum += profit
		}
	}

	return sum
}
