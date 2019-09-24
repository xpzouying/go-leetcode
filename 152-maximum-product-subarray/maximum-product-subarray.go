package maximum_product_subarray

import (
	"log"
	"math"
)

func maxProduct(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	DP := make([]int, l)
	DP[0] = nums[0]

	max := math.MinInt64

	for i := 1; i < l; i++ {

		if DP[i-1] == 0 {
			DP[i] = nums[i]
		} else {
			DP[i] = DP[i-1] * nums[i]
		}

		if DP[i] > max {
			max = DP[i]
		}
	}

	log.Printf("DP: %v", DP)

	return max
}

/**

解题思路：

从题目中的“乘积最大”，很容易得出来是一道动态规划题目。

假设：DP为最大乘积。

如何进行拆分子任务。


### 思路：

DP[i] 和 nums[i]的关系，

如果要nums[i]可以加入进来，也即可以构成更大值，那么 DP[i-1]*nums[i] 要可以累计起来，即DP[i-1]不能为0。

所以问题就被拆成：

1. DP[i-1] == 0，那么DP[i] = nums[i]
2. DP[i-1] != 0，那么DP[i] = DP[i-1]*nums[i]

*/
