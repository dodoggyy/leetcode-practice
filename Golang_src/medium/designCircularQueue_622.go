package medium

// Use array:
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 28 ms, faster than 7.61%
// Memory Usage: 7.2 MB, less than 31.52%
type MyCircularQueue struct {
	rear, front int
	size        int
	arr         []int
}

func ConstructorMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		size: k + 1,
		arr:  make([]int, k+1),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr[this.rear] = value
	this.rear = (this.rear + 1) % this.size

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.size

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[(this.rear-1+this.size)%this.size]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.rear == this.front
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.rear+1)%this.size == this.front
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
