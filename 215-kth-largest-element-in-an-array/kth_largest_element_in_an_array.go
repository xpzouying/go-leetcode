package kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	klist := NewKList(k)

	for _, n := range nums {
		klist.Insert(n)
	}

	return klist.Kth()
}

type KList struct {
	// data 保存k个数
	// 从小到大排序
	data []int
	k    int
}

func NewKList(k int) *KList {
	return &KList{data: make([]int, 0, k), k: k}
}

// Kth 返回第k大的数
// 如果klist为空，则返回错误
func (kl *KList) Kth() int {
	return kl.data[0]
}

func (kl *KList) Full() bool {
	return len(kl.data) == kl.k
}

// Insert 插入数字n
func (kl *KList) Insert(n int) {

	// 使用替换
	if kl.Full() {
		kl.doReplace(n)
	} else {
		kl.doInsert(n)
	}
}

func (kl *KList) doReplace(n int) {
	if n < kl.data[0] {
		// 太小，不足以替换
		return
	}

	i := 0
	for i = 0; i < len(kl.data); i++ {

		if kl.data[i] > n {
			// 小于i的都小于n
			break
		}
	}

	if i == 1 {
		// 替换第一个数字
		kl.data[0] = n
		return
	} else if i == len(kl.data) {
		// 增加到最后一位
		kl.data = append(kl.data[1:], n)
		return
	}

	// 插入在中间
	kl.data = append(
		kl.data[1:i],
		append([]int{n}, kl.data[i:]...)...,
	)
}

// 还有空间，插入元素
func (kl *KList) doInsert(n int) {

	i := 0
	for i = 0; i < len(kl.data); i++ {
		if kl.data[i] > n {
			break
		}
	}

	kl.data = append(
		kl.data[:i],
		append([]int{n}, kl.data[i:]...)...,
	)

}

/**

解题思路：


使用数组保存，数组排列升序排列，数组长度为k。

数组下标0，即为第k大的数。

优化点：

1. 使用小顶堆保存，更优化。
2. 使用折半查找法搜索插入的位置i，搜索目标位置的时间复杂度从N降到logN

*/
