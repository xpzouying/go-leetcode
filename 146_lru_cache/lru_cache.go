package lru_cache

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	len  int   // 当前队列的个数
	cap  int   // 上限
	head *Node // 队首节点
	tail *Node // 指向队尾节点
}

func Constructor(capacity int) LRUCache {
	head := &Node{Pre: nil}
	tail := &Node{Pre: head, Next: nil}
	head.Next = tail

	return LRUCache{len: 0, cap: capacity, head: head, tail: tail}
}

func (this *LRUCache) Println() []int {

	if this.len == 0 {
		return []int{}
	}

	l := make([]int, 0, this.len)
	p := this.head.Next
	for {
		if p == this.tail {
			break
		}

		l = append(l, p.Val)

		p = p.Next
	}

	return l
}

func (this *LRUCache) Get(key int) int {
	if this.len == 0 {
		return -1
	}
	p := this.head.Next
	for {
		if p == this.tail {
			break
		}

		if p.Key == key {
			// 找到当前节点
			if p.Pre == this.head {
				// 如果已经是队首，则直接返回
				return p.Val
			}
			// 移动当前节点到队首

			// 1. 从队列中拿出来
			p.Pre.Next = p.Next
			p.Next.Pre = p.Pre

			// 2. 挂载到head后面
			p.Next = this.head.Next
			p.Pre = this.head
			this.head.Next.Pre = p
			this.head.Next = p

			return p.Val
		}

		p = p.Next
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	p := this.head.Next

	for {
		if p == this.tail {
			// 没有找到节点
			break
		}

		if p.Key == key {
			// 找到节点
			break
		}

		p = p.Next
	}

	// 如果没有找到
	if p == this.tail {
		// 如果容量已满，则淘汰后，再增加
		if this.len == this.cap {
			// 删除队尾元素
			this.tail.Pre.Pre.Next = this.tail
			this.tail.Pre = this.tail.Pre.Pre
		} else {
			// 如果容量未满
			this.len++
		}

		newNode := &Node{
			Key: key,
			Val: value,
			Pre: this.head,
		}
		newNode.Next = this.head.Next
		newNode.Next.Pre = newNode
		this.head.Next = newNode
		return
	}

	// 如果找到了
	// 移动当前节点到队首

	// 1. 从队列中拿出来
	p.Pre.Next = p.Next
	p.Next.Pre = p.Pre

	// 2. 挂载到head后面
	p.Next = this.head.Next
	p.Pre = this.head
	this.head.Next.Pre = p
	this.head.Next = p

	p.Val = value
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**

解题思路：

最简单的方案：使用双向列表

1. 使用两个节点head和tail，分别指向队首和队尾；

2. 判断当前节点是否已经存在；
	2.1 如果存在，则更新该节点到队首；
	2.2 如果不存在，则：
		2.2.1 如LRU长度未满，则添加到队首；
		2.2.2 如LRU长度已满，则删除队尾节点，添加

*/
