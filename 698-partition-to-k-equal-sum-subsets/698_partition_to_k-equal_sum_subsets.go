package partition_to_k_equal_sum_subsets

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {

	if len(nums) == 0 {
		return false
	}

	if k <= 1 {
		return true
	}

	var SUM int
	for _, n := range nums {
		SUM += n
	}

	// 表示不能被整除
	if SUM%k != 0 {
		return false
	}

	AVG := SUM / k // 平均后，每一组需要达到AVG和
	if AVG == 0 {
		// [3,2,1] k=3
		// 如果每一个都不能凑到1个数值
		return false
	}

	// groups 表示k个分块，每一个元素表示分块的当前和，初始化为0， 需要到AVG总数
	groups := make([]int, k)

	sort.Ints(nums)

	return helper(nums, groups, AVG)
}

func helper(nums []int, groups []int, sum int) bool {
	if len(nums) == 0 {
		// 当nums没有元素，
		// 并且groups中的元素都满足sum的要求后，返回true；否则返回false；
		for _, g := range groups {
			if g != sum {
				return false
			}
		}

		return true
	}

	for i := len(nums) - 1; i >= 0; i-- {
		// 如果单个元素超过sum，则返回false，表示“爆掉了”
		if nums[i] > sum {
			return false
		}

		// 否则把当前元素加入到某个group中，进行递归
		for j := 0; j < len(groups); j++ {
			groupSUM := groups[j] + nums[i]

			if groupSUM > sum {
				// 如果加入到某个group中，使得当前group爆掉了，则尝试下一个group
				continue
			}

			groups[j] = groupSUM

			subRes := helper(nums[:i], groups, sum)
			if subRes == true {
				return true
			}

			// 如果加入到当前group中，没有找到解，则尝试一下group
			groups[j] -= nums[i] // 收回结果
		}

		// 表示当前元素没有加入到任何group中
		return false
	}

	return false
}

/**

### 解题思路：

1. 计算nums的和，SUM。
    1.1. 如果SUM被能被K整除，则返回false。即不能平均分成k块；
    1.2. 若可以被整除，则生成k个块，记为groups。每一个分块的和必然都为：AVG = SUM/k

2. 对nums排序处理；

3. 开始遍历处理：
    helper(nums []int, group []int, sum int) bool {}
        nums: 为输入列表参数
        group: []int, 为分块和的列表
        sum：每一组和需要拼成的和


    l1 := len(nums)
    a. 将nums中的元素，取出来，遍历加入group对应的元素后，进行递归
    b. 如果为false，则进行下一个列表；



*/
