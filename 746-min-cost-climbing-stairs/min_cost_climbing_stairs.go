package min_cost_climbing_stairs

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */
func minCostClimbingStairs(cost []int) int {

	/**

	// 转移方程
	Sum[i] = Min{
		Sum[i-1]+cost[i-1],
		Sum[i-2]+cost[i-2]
	}

	// 初始化
	Sum[0] = 0
	Sum[1] = 0

	// dp_table: make([]int, len(cost)+1)

	// 边界
	last element: i == len(cost)

	---------------
	TestCase-01

	输入: cost = [10, 15, 20]
	输出: 15

	Sum[0] = 0
	Sum[1] = 0
	Sum[2] = Min{
		Sum[0] + cost[0],
		Sum[1] + cost[1],
	} = Min{ 0+10, 0+15 } = 10
	Cum[3] = Min{
		Sum[1]+cost[1],
		Sum[2]+cost[2],
	} = Min{0+15, 10+20}  = 15
	*/

	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}

	sumLen := len(cost) + 1
	SUM := make([]int, sumLen)
	SUM[0] = 0
	SUM[1] = 0

	for i := 2; i < sumLen; i++ {
		SUM[i] = Min(
			SUM[i-1]+cost[i-1],
			SUM[i-2]+cost[i-2],
		)
	}

	return SUM[sumLen-1]
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
