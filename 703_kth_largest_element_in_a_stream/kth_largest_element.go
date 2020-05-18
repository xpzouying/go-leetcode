package kth_largest_element

type KthLargest struct {
	k    int
	data []int
}

func Constructor(k int, nums []int) KthLargest {
	q := KthLargest{
		k:    k,
		data: make([]int, 0, k),
	}

	if len(nums) == 0 {
		return q
	}

	for i := 0; i < len(nums); i++ {
		q.Add(nums[i])
	}

	return q
}

func (this *KthLargest) Add(val int) int {
	if len(this.data) == 0 {
		this.data = append(this.data, val)
		return this.data[len(this.data)-1]
	}

	K := this.k

	if len(this.data) == K {
		// 如果已经队列已经满了
		last := this.data[K-1]
		if last >= val {
			// 如果不足以淘汰最小的元素
			return last
		}

		// 如果可以淘汰最后一个，则找到最合适的位置
		i := K - 1 // -2的原因为：最后一个元素（即最小的元素）淘汰，从倒数第二个元素开始
		for ; i >= 1; i-- {
			if val <= this.data[i-1] {
				break
			}

			this.data[i] = this.data[i-1]
		}
		this.data[i] = val

		return this.data[K-1]
	}

	// 如果容量未满
	if this.data[0] < val {
		// 如果比所有的数都大
		this.data = append([]int{val}, this.data...)
	} else if this.data[len(this.data)-1] > val {
		// 如果比所有的数都小
		this.data = append(this.data, val)
	} else {
		var i int
		for i = 0; i < len(this.data); i++ {
			if this.data[i] <= val {
				break
			}
		}

		this.data = append(this.data[:i], append([]int{val}, this.data[i:]...)...)
	}

	return this.data[len(this.data)-1]
}
