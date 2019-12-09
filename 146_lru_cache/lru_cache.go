package lru_cache

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	len  int           // 当前队列的个数
	cap  int           // 上限
	head *Node         // 队首节点
	tail *Node         // 指向队尾节点
	m    map[int]*Node // hashmap，key为key，value为指向双链表中的节点
}

func Constructor(capacity int) LRUCache {
	head := &Node{Pre: nil}
	tail := &Node{Pre: head, Next: nil}
	head.Next = tail

	return LRUCache{len: 0, cap: capacity, head: head, tail: tail, m: make(map[int]*Node, capacity)}
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

	// 直接查找hash map，若不存在则直接返回
	node, ok := this.m[key]
	if !ok {
		return -1
	}

	if node.Pre == this.head {
		return node.Val
	}

	// 移动当前节点到队首
	// 1. 从队列中拿出来
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	// 2. 挂载到head后面
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next.Pre = node
	this.head.Next = node

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	// 如果存在，则直接更新节点在列表中的位置
	if ok {
		// 1. 从队列中拿出来
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre

		// 2. 挂载到head后面
		node.Next = this.head.Next
		node.Pre = this.head
		this.head.Next.Pre = node
		this.head.Next = node

		node.Val = value
		return
	}

	// 不存在的情况，需要判断是否超出容量
	if this.len == this.cap {
		// 删除队尾元素

		// 删除hash map中的对应的key
		key := this.tail.Pre.Key
		delete(this.m, key)

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
	this.m[key] = newNode // 增加hashmap
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

/**

优化思路：

如果完成“获取” / “添加” 要O(1)的时间复杂度，
需要额外借助数据结构，O(1)直接联想到的是hash map。

现在速度满的原因在于查找，
当查找元素是否存在时，需要遍历列表，由此每次插入的时候，都需要遍历列表，时间复杂度为：O(n)。

增加hash map，表示节点是否存在，如果存在的话，其value指向链表中的节点。

*/
