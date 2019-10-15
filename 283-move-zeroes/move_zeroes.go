package move_zeroes

func moveZeroes(nums []int) {

	if len(nums) <= 1 {
		return
	}

	i, j := 0, 0

	for {
		if i == len(nums) || j == len(nums) {
			break
		}

		if nums[i] != 0 {
			i++
			continue
		}

		if i > j {
			j = i
		}

		if nums[j] == 0 {
			j++
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j++
	}
}

/**

解体思路：

设定i，j分别指向数组的前端，一起往后移动。

1. 当i指向一个0时，停止；

2. 当j指向一个非0时，停止；

3. 交换i、j指向的数值；

4. 移动i、j；

*/
