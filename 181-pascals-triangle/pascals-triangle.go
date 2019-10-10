package pascals_triangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	if numRows == 1 {
		return [][]int{
			[]int{1},
		}
	}

	res := make([][]int, numRows)
	res[0] = []int{1}

	// 构造每一层，从第0层开始构造
	// 每一层元素的个数为：层数+1

	for i := 1; i < numRows; i++ {

		currSize := i + 1

		currRow := make([]int, currSize)
		for j := 0; j < currSize; j++ {

			// 第一位和最后一位元素为：1
			if j == 0 || j == currSize-1 {
				currRow[j] = 1
				continue
			}

			currRow[j] = res[i-1][j-1] + res[i-1][j]
		}

		res[i] = currRow
	}

	return res
}
