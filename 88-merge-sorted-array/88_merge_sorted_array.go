package merge_orted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2[:n])
		return
	}

	doMerge(nums1, m-1, nums2, n-1, len(nums1)-1)
}

// 使用递归。比较最大
// i, j 分别标识 nums1、nums2的可用位置。
// mergedPos表示已经merge的位置。
func doMerge(nums1 []int, i int, nums2 []int, j int, mergedPos int) {
	// 如果两个数组都已经遍历完了
	if mergedPos < 0 {
		return
	}

	if i < 0 && j >= 0 {
		// 如果nums1已经遍历完了
		copy(nums1[:j+1], nums2[:j+1])

		return

	} else if j < 0 {
		// 如果nums2已经遍历完了。nums1就不需要挪动了，直接返回。
		return
	}

	n1, n2 := nums1[i], nums2[j]
	// 将大的数字放到mergedPos，表示落位到最终位置
	if n1 > n2 {
		nums1[mergedPos] = n1
		i--
	} else {
		nums1[mergedPos] = n2
		j--
	}

	mergedPos--
	doMerge(nums1, i, nums2, j, mergedPos)
}
