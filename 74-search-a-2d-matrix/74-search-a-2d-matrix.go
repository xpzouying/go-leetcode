package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 如果 target >= matrix[i][0] && target < matrix[i+1][0]，
	// 则在第i行。
	low := 0
	high := len(matrix) - 1
	for {
		if low > high {
			break
		}

		mid := int((low + high) / 2)

		if target == matrix[mid][0] {
			return true
		} else if target > matrix[mid][0] {

			// 如果已经到最后一行了 或者 比下一行的第一个元素小，
			// 那么直接搜索该行中的目标
			if mid == high || (target < matrix[mid+1][0]) {
				return searchInLine(matrix[mid], target)
			}

			// 否则，进行向上分治
			low = mid + 1

		} else {
			// 否则，向上分治
			high = mid - 1

		}
	}

	return false
}

func searchInLine(array []int, target int) bool {
	low := 0
	high := len(array) - 1

	for {
		if low > high {
			break
		}

		if array[low] == target || array[high] == target {
			return true
		}

		mid := (low + high) / 2

		if array[mid] == target {
			return true
		}

		if target > array[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

/*

--- 解体思路 ---

1. 由于矩阵都是有序的，所以对于每一行的数组使用二分查找；
2. 对于查找每一列，也使用二分查找；


步骤：
1. 先定位到是哪一行；
2. 再在这一行中，定位到具体的元素；


--- 执行结果 ---

执行结果： 通过

显示详情

* 执行用时： 4 ms , 在所有 Go 提交中击败了 88.33% 的用户
* 内存消耗： 2.7 MB , 在所有 Go 提交中击败了 81.32% 的用户
*/
