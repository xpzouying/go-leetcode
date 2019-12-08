package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	wantList := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {

		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] == 0 {
			if nums[i+1] == 0 && nums[i+2] == 0 {
				wantList = append(wantList, []int{0, 0, 0})
				return wantList
			}
		}

		// 使用双指针的方案
		low := i + 1
		high := len(nums) - 1

		for {
			if low >= high {
				break
			}

			sum := nums[i] + nums[low] + nums[high]

			if sum == 0 {
				// 双指针变化
				wantList = append(wantList, []int{nums[i], nums[low], nums[high]})

				// 表示 和 太大，右边指针左移
				oldNum := nums[high]
				high--
				for ; low < high; high-- {
					if nums[high] != oldNum {
						break
					}
				}
			} else if sum > 0 {
				// 表示 和 太大，右边指针左移
				oldNum := nums[high]
				high--
				for ; low < high; high-- {
					if nums[high] != oldNum {
						break
					}
				}
			} else {
				// 表示 和 太小，左边指针右移
				oldNum := nums[low]
				low++
				for ; low < high; low++ {
					if nums[low] != oldNum {
						break
					}
				}
			}

		}
	}

	return wantList
}

/**

解题思路：

1. 三数之和，先对数组进行排序，假设得到[n1, n2, n3, ...]
2. for n1...nx, 从第一个数字开始遍历，
  假设n_i是当前遍历的数字，
  2.1 如果n_i大于0，则没有三元组；
  2.2 如果n_i小于0，假设为X，则在后面数组中找到 -X的和。可以使用折半查找法
  2.3 如果n_i等于0，则判断后续两个是否为0。因为要求找出不重复的三元组。

异常情况：
  - nums小于3个

*/
