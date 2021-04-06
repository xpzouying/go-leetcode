package find_disappeared_numbers

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		n := nums[i]
		idxOfN := n - 1

		// 如果 下标(n-1)的元素已经是n了，那么i++，进行下一个。
		if n == nums[idxOfN] {
			i++
			continue
		}

		// 下标i的数字 和 下标(n-1)的位置，进行交换
		nums[idxOfN], nums[i] = n, nums[idxOfN]
	}

	// 遍历
	var res []int
	for i := 0; i < len(nums); i++ {

		if expect := i + 1; nums[i] != expect {
			res = append(res, expect)
		}
	}

	return res
}

/*

执行结果： 通过

显示详情

执行用时： 64 ms , 在所有 Go 提交中击败了 32.13% 的用户

内存消耗： 7.8 MB , 在所有 Go 提交中击败了 13.01% 的用户

**/
