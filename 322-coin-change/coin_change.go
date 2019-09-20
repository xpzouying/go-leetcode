package coin_change

import (
	"log"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if len(coins) == 0 {
		return -1
	}

	sort.Ints(coins)
	minCoin := coins[0]

	log.Printf("coins: %v, amount: %d", coins, amount)

	dp := make([]int, amount+1)

	for _, c := range coins {
		if c > amount {
			// 小技巧：抛去面值比amount还大的
			continue
		}

		dp[c] = 1
	}

	// 构建dp表
	for i := 1; i <= amount; i++ {

		currAmount := i

		if dp[currAmount] != 0 {
			// 初始化过，存在对应的硬币面值
			continue
		}

		// needCoins := make([]int, 0) // 需要多少个硬币记录，最终取其中的最小值
		minNeed := -1 // 构成i值，最少需要多少枚硬币，默认为-1，没法拼凑

		// 遍历所有的硬币
		for j := 0; j < len(coins); j++ {
			// 当前硬币的面值
			val := coins[j]

			// 已经超出构建的范围，跳出。
			// 比如：我们还在构建dp[2], 此时我们不需要统计面值为5的数量
			if val > i {
				break
			}

			// 剩余多少值需要拼凑
			leftVal := currAmount - val

			// 剩下的残值比最小硬币还大
			if leftVal < minCoin {
				continue
			}

			// 已经记录过拼凑不出来
			if dp[leftVal] == -1 {
				continue
			}

			// 可以计算
			if minNeed == -1 {
				// 如果之前拼凑不出来
				minNeed = dp[leftVal] + 1
			} else {
				// 如果之前有值，那么计算最小值
				minNeed = min(dp[leftVal]+1, minNeed)
			}
		}

		dp[i] = minNeed
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/**


解题思路：


输入: coins = [1, 2, 5], amount = 11


那么最少要多少个硬币：

假设
dp[i] 为值为i时，最少需要的硬币数量。

那么
dp[i] = min{
	dp[i-1],  // 使用面值为1的硬币
	dp[i-2],  // 使用面值为2的硬币
	...
} + 1  // 加上当前使用的这个硬币


那么示例即可拆分为：

dp[11] = min{
	dp[11-1],
	dp[11-2],
	dp[11-5],
} + 1


初始条件为：

dp[1] = dp[2] = dp[5] = 1

小于最小面值，则返回-1。

*/
