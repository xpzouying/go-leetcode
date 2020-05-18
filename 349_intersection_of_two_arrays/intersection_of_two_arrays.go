package intersection_of_two_arrays

/* --- 空间优化法 ---
func intersection(nums1 []int, nums2 []int) []int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0)
	i := 0 // 指向nums1
	j := 0 // 指向nums2
	len1 := len(nums1)
	len2 := len(nums2)

FOR_OUTER:
	for {

		if i >= len1 || j >= len2 {
			break
		}

		// 如果不相等
		if nums1[i] > nums2[j] {
			j++
			continue
		} else if nums1[i] < nums2[j] {
			i++
			continue
		}

		res = append(res, nums1[i])

		i++
		for {
			if i >= len1 {
				break FOR_OUTER
			}

			if nums1[i] != nums1[i-1] {
				break
			}

			i++
		}

		j++
		for {
			if j >= len2 {
				break FOR_OUTER
			}

			if nums2[j] != nums2[j-1] {
				break
			}

			j++
		}
	}

	return res
}

*/

// 时间优先
// 使用map记录长度较短的表
func intersection(nums1 []int, nums2 []int) []int {

	var shortLen int
	var short []int
	var long []int
	if len(nums1) >= len(nums2) {
		short = nums2
		long = nums1
	} else {
		short = nums1
		long = nums2
	}
	shortLen = len(short)
	m := make(map[int]bool, shortLen) // m 保存short里面的元素，key为数字，value为是否已经被保存到输出结果中
	for i := 0; i < shortLen; i++ {
		m[short[i]] = false
	}

	res := make([]int, 0, shortLen)
	for i := 0; i < len(long); i++ {

		saved, ok := m[long[i]]
		if !ok {
			// 当前元素不是共有的
			continue
		}

		if saved {
			// 如果已经保存过
			continue
		}

		// 保存到输出结果中，并且记录在map中，表示已经输出
		res = append(res, long[i])
		m[long[i]] = true
	}

	return res
}
