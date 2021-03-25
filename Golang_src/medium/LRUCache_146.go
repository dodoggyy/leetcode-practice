package medium

// Use hash table + double linkedlist:
// Time Complexity: O(1)
// Space Complexity:O(capacity)
// Runtime: 124 ms, faster than 72.53%
// Memory Usage: 17.2 MB, less than 41.57%
type LRUCache struct {
	cache      map[int]*DoubleLinkedNode
	size       int
	capacity   int
	head, tail *DoubleLinkedNode
}

type DoubleLinkedNode struct {
	key, val  int
	pre, next *DoubleLinkedNode
}

func createDoubleLinkedNode(key, val int) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		key: key,
		val: val,
	}
}

func CreateLRU(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DoubleLinkedNode{},
		head:     createDoubleLinkedNode(0, 0),
		tail:     createDoubleLinkedNode(0, 0),
		capacity: capacity,
	}

	l.head.next = l.tail
	l.tail.next = l.head

	return l
}

func (this *LRUCache) Get(key int) int {
	if _, exist := this.cache[key]; !exist {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if _, exist := this.cache[key]; !exist {
		node := createDoubleLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			remove := this.removeTail()
			delete(this.cache, remove.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.val = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DoubleLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DoubleLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *DoubleLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DoubleLinkedNode {
	node := this.tail.pre
	this.removeNode(node)

	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
