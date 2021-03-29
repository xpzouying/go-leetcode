package main

import "sort"

// 方法1 - 使用hash map缓存记录
// func findRepeatNumber(nums []int) int {
// 	m := make(map[int]struct{}, len(nums))

// 	for _, num := range nums {

// 		if _, ok := m[num]; ok {
// 			return num
// 		} else {
// 			m[num] = struct{}{}
// 		}
// 	}

// 	return -1
// }

/*

--- 解体思路 --

使用map保存遇到的数。

如果map中已经存在该数字，则直接返回。


--- 运行结果 ---
执行用时： 44 ms , 在所有 Go 提交中击败了 56.50% 的用户

内存消耗： 10.1 MB , 在所有 Go 提交中击败了 6.10% 的用户
*/

// -----
// 方法2 - 先使用快排进行排序，然后再查看连续的2个数字是否有相同的。
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return -1
}

/* 方法2的运行结果

执行用时： 52 ms , 在所有 Go 提交中击败了 18.30% 的用户

内存消耗： 8 MB , 在所有 Go 提交中击败了 96.88% 的用户
*/

/* 方法3

感觉可以省去上面最后一步的for循环遍历；
直接使用快排的过程中，在比较的过程中，直接判断是否有相同的元素。

*/
